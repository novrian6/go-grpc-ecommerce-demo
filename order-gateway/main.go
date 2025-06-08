package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "example.com/go_grpc/ecommerce-demo/proto-user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("gRPC dial error:", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")
		if userID == "" {
			userID = "1"
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := client.GetUser(ctx, &pb.GetUserRequest{Id: userID})
		if err != nil {
			http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "JSON encode error: "+err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Order Gateway running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
