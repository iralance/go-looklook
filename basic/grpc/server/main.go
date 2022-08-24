package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	trippb "github.com/iralance/go-looklook/basic/grpc/proto/gen/go"
	"github.com/iralance/go-looklook/basic/grpc/tripservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	ln, err := net.Listen("tcp", ":8600")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &tripservice.ServiceServer{})
	log.Printf("server listening at %v", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
				UseEnumNumbers:  true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		//mux: multiplexer
		mux,
		":8600",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	log.Println(err)
	if err != nil {
		log.Fatalf("cannot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8601", mux)
	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}
}
