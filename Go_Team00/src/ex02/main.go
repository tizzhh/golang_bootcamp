package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
	pb "transmitterGRPCWriteToDB/transmitter"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	MIN_MEAN         float64 = -10
	MAX_MEAN         float64 = 10
	MIN_SD           float64 = 0.3
	MAX_SD           float64 = 1.5
	DB_USER_ENV      string  = "DB_USER"
	DB_PASS_ENV      string  = "DB_PASS"
	DB_HOST_ENV      string  = "DB_HOST"
	DB_NAME_ENV      string  = "DB_NAME"
	DB_PORT_ENV      string  = "DB_PORT"
	DSN_URL_TEMPLATE string  = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

type Frequencies struct {
	SessionId string    `json:"UUID"`
	Frequency float64   `json:"Frequency"`
	Timestamp time.Time `json:"Timestamp"`
}

type generatedVals struct {
	mean, sd float64
	uuid     string
}

type Server struct {
	pb.UnimplementedTransmitServer
	db *gorm.DB
}

func (s *Server) TransmitVals(em *pb.EmptyBody, stream pb.Transmit_TransmitValsServer) error {
	vals := generateRandomVals()
	log.Printf("UUID: %s, SD: %f, mean: %f\n", vals.uuid, vals.sd, vals.mean)

	for {
		freq := rand.NormFloat64()*vals.sd + vals.mean
		timestmp := timestamppb.Now()
		resp := &pb.Reponses{
			SessionId: vals.uuid,
			Frequency: freq,
			Timestamp: timestmp,
		}

		result := s.db.Create(Frequencies{SessionId: vals.uuid, Frequency: freq, Timestamp: timestmp.AsTime()})
		if result.Error != nil {
			return result.Error
		}

		err := stream.Send(resp)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func generateRandomVals() generatedVals {
	return generatedVals{
		uuid: uuid.New().String(),
		mean: MIN_MEAN + rand.Float64()*(MAX_MEAN-MIN_MEAN),
		sd:   MIN_SD + rand.Float64()*(MAX_SD-MIN_SD),
	}

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
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listed: %v\n", err)
	}

	dbConn, err := getDbConn()
	if err != nil {
		log.Fatalf("error openning a db connection: %v\n", err)
	}
	dbConn.AutoMigrate(&Frequencies{})

	server := grpc.NewServer()
	myTransmitterServer := &Server{db: dbConn}
	pb.RegisterTransmitServer(server, myTransmitterServer)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
