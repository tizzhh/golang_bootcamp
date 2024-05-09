package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"sync"

	pb "anomalyDetection/transmitter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SERVER_PORT                  string = ":3333"
	NUMBER_OF_NORM_DISTR_ELEMS   int    = 50
	LOG_PERIODICITY_BY_RESPONSES int    = 10
	ANOMALY_DETECTED             bool   = false
)

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
			fmt.Println(resp.Frequency)
			normSum += resp.Frequency
			countIncVals++

			*normSamples = append(*normSamples, resp.Frequency)

			if checkFreqAnomaly(resp.Frequency, sd, *sdCoeff, mean) {
				done <- ANOMALY_DETECTED
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

	res := <-done
	if res {
		fmt.Println("Anomalies are not present")
	} else {
		fmt.Println("Anomaly detected!")
	}
}
