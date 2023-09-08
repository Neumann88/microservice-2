package main

import (
	"net/http"

	pb "github.com/Neum/microservice-2/internal/gen/proto/services/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	microserviceOneConn, err := grpc.Dial("localhost:5010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer microserviceOneConn.Close()

	microServiceOneClient := pb.NewMicroserviceOneExampleClient(microserviceOneConn)

	mux := http.NewServeMux()
	mux.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		msg := &pb.ExampleMessageRequest{Id: 1}
		resp, err := microServiceOneClient.GetExampleMessage(r.Context(), msg)
		if err != nil {
			return
		}

		w.Write([]byte(resp.Value))
	})

	http.ListenAndServe(":8080", mux)
}
