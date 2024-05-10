package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	pb "anomalyDetection/transmitter"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	SERVER_PORT                  string = ":3333"
	NUMBER_OF_NORM_DISTR_ELEMS   int    = 50
	LOG_PERIODICITY_BY_RESPONSES int    = 10
	ANOMALY_DETECTED             bool   = false
	DB_USER_ENV                  string = "DB_USER"
	DB_PASS_ENV                  string = "DB_PASS"
	DB_HOST_ENV                  string = "DB_HOST"
	DB_NAME_ENV                  string = "DB_NAME"
	DB_PORT_ENV                  string = "DB_PORT"
	DSN_URL_TEMPLATE             string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

type Frequencies struct {
	SessionId string    `json:"UUID"`
	Frequency float64   `json:"Frequency"`
	Timestamp time.Time `json:"Timestamp"`
}

func checkFreqAnomaly(freq, sd, coeff, mean float64) bool {
	return freq-mean > sd*coeff
}

func calcMean(sum float64) float64 {
	return sum / float64(NUMBER_OF_NORM_DISTR_ELEMS)
}

func calcSD(prevSD, mean float64, normSamples []float64) float64 {
	sumSquareDiff := prevSD * float64(NUMBER_OF_NORM_DISTR_ELEMS)
	for _, val := range normSamples {
		diff := val - mean
		sumSquareDiff += diff * diff
	}
	return sumSquareDiff / float64(NUMBER_OF_NORM_DISTR_ELEMS)
}

func putNormInPool(sl *[]float64, pool *sync.Pool) {
	if cap(*sl) <= NUMBER_OF_NORM_DISTR_ELEMS {
		*sl = (*sl)[:0]
		pool.Put(sl)
	}
}

func getFloatForPool(pool *sync.Pool) *[]float64 {
	ifc := pool.Get()
	var sl *[]float64
	if ifc != nil {
		sl = ifc.(*[]float64)
	}
	return sl
}

func getDbConn() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dsn := fmt.Sprintf(DSN_URL_TEMPLATE, os.Getenv(DB_HOST_ENV), os.Getenv(DB_USER_ENV), os.Getenv(DB_PASS_ENV), os.Getenv(DB_NAME_ENV), os.Getenv(DB_PORT_ENV))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error openning a database connection: %w", err)
	}

	return db, nil
}

func main() {
	sdCoeff := flag.Float64("k", 0, "SD coefficient")
	flag.Parse()

	conn, err := grpc.Dial(SERVER_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect with server: %v\n", err)
	}

	client := pb.NewTransmitClient(conn)
	stream, err := client.TransmitVals(context.Background(), &pb.EmptyBody{})
	if err != nil {
		log.Fatalf("stream openning error: %v\n", err)
	}

	dbConn, err := getDbConn()

	if err != nil {
		log.Fatalf("error openning a db connection: %v\n", err)
	}
	dbConn.AutoMigrate(&Frequencies{})

	var normPool = sync.Pool{
		New: func() interface{} { return &[]float64{} },
	}

	normSamples := getFloatForPool(&normPool)
	defer putNormInPool(normSamples, &normPool)
	var countIncVals int
	var mean, sd float64
	var normSum float64
	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			normSum += resp.Frequency
			countIncVals++

			*normSamples = append(*normSamples, resp.Frequency)

			if checkFreqAnomaly(resp.Frequency, sd, *sdCoeff, mean) {
				fmt.Println("Anomaly found: ", resp.Frequency)
				dbConn.Create(Frequencies{SessionId: resp.SessionId, Frequency: resp.Frequency, Timestamp: timestamppb.Now().AsTime()})
			} else {
				fmt.Println("Normal frequency: ", resp.Frequency)
			}

			if countIncVals%LOG_PERIODICITY_BY_RESPONSES == 0 {
				mean = calcMean(normSum)
				sd = calcSD(sd, mean, *normSamples)
				log.Printf("Values processed: %d, approx mean: %f, approx sd: %f\n", countIncVals, mean, sd)
				putNormInPool(normSamples, &normPool)
				normSamples = getFloatForPool(&normPool)
			}

			if err == io.EOF || countIncVals == NUMBER_OF_NORM_DISTR_ELEMS {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("cannot recieve: %v\n", err)
			}
		}
	}()
	<-done
}
