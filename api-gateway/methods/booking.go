package methods

import (
	"api-gateway/proto/genbooking"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CoonectBooking() genbooking.BookingServiceClient {
	log.Println("Ulanib bildimi api surov yuborishga")
	conn, err := grpc.NewClient("booking_service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect user micro...", err)
	}

	client := genbooking.NewBookingServiceClient(conn)
	return client
}
