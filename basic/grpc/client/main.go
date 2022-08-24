package main

import (
	"context"
	"log"

	trippb "github.com/iralance/go-looklook/basic/grpc/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8600", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := trippb.NewTripServiceClient(conn)
	mode(err, client)
	//modeIn(err, client)
	//modeOut(err, client)
	//modeIo(err, client)
}

func mode(err error, client trippb.TripServiceClient) {
	res, err := client.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "abc456"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", res)
}
