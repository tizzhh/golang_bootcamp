package main

import (
	"log"
	"math/rand"
	"net"
	"time"
	pb "transmitterGRPC/transmitter"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	MIN_MEAN float64 = -10
	MAX_MEAN float64 = 10
	MIN_SD   float64 = 0.3
	MAX_SD   float64 = 1.5
)

type generatedVals struct {
	mean, sd float64
	uuid     string
}

type Server struct {
	pb.UnimplementedTransmitServer
}

func (s *Server) TransmitVals(em *pb.EmptyBody, stream pb.Transmit_TransmitValsServer) error {
	vals := generateRandomVals()
	log.Printf("UUID: %s, SD: %f, mean: %f\n", vals.uuid, vals.sd, vals.mean)

	for {
		resp := &pb.Reponses{
			SessionId: vals.uuid,
			Frequency: rand.NormFloat64()*vals.sd + vals.mean,
			Timestamp: timestamppb.Now(),
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

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listed: %v\n", err)
	}
	server := grpc.NewServer()
	myTransmitterServer := &Server{}
	pb.RegisterTransmitServer(server, myTransmitterServer)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
