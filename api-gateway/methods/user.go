package methods

import (
	"api-gateway/proto/genuser"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectUser() genuser.UserServiceClient {
	log.Println("Ulanib bildimi api surov yuborishga")
	conn, err := grpc.NewClient("user_service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect user micro...", err)
	}

	client := genuser.NewUserServiceClient(conn)
	return client
}
