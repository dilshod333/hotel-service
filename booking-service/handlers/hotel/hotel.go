// // package hotel

// // import (
// // 	"booking-service/models"
// // 	"booking-service/proto/genhotel"
// // 	"context"
// // 	"log"

// // 	"google.golang.org/grpc"
// // 	"google.golang.org/grpc/credentials/insecure"
// // )

// // func HandleHotel(id string) (models.RoomHotel, error) {
// // 	log.Println("Handlehotel ichida")
// // 	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
// // 	if err != nil {
// // 		log.Println("Error on connection handhotel", err)
// // 		return models.RoomHotel{}, err
// // 	}

// // 	client := genhotel.NewHotelServiceClient(conn)

// // 	hotel, err := client.GetHotelDetails(context.Background(), &genhotel.GetHotelDetailsRequest{HotelId: id})
// // 	log.Println("Hotel>>>", hotel)
// // 	if err != nil {
// // 		log.Println("Error on gethotel", err)
// // 		return models.RoomHotel{}, err
// // 	}
// // 	var roomHotel models.RoomHotel
// // 	var m models.Only
// // 	log.Println("M>>>.", m)
// // 	for _, room := range hotel.Hotel.Rooms {
// // 		if m.RoomType == room.RoomType{
// // 			roomHotel = models.RoomHotel{
// // 				RoomType:     room.RoomType,
// // 				PricePerNight: float32(room.PricePerNight),
// // 			}
// // 		}

// // 		log.Printf("Room Type: %s, Price Per Night: %.2f", roomHotel.RoomType, roomHotel.PricePerNight)
// // 	}
// // 	roomHotel.HotelId = int32(hotel.Hotel.HotelId)

// // 	return roomHotel, nil
// // }

// package hotel

// import (
// 	"booking-service/models"
// 	"booking-service/proto/genhotel"
// 	"context"
// 	"fmt"
// 	"log"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func HandleHotel(id string) (models.RoomHotel, error) {
// 	log.Println("HandleHotel function started")
// 	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Println("Error connecting to gRPC server:", err)
// 		return models.RoomHotel{}, err
// 	}
// 	defer conn.Close()

// 	client := genhotel.NewHotelServiceClient(conn)

// 	hotel, err := client.GetHotelDetails(context.Background(), &genhotel.GetHotelDetailsRequest{HotelId: id})
// 	if err != nil {
// 		log.Println("Error getting hotel details:", err)
// 		return models.RoomHotel{}, err
// 	}

// 	// Initialize a variable to keep track of the roomHotel details
// 	var roomHotel models.RoomHotel
// 	found := false

// 	// Iterate through the rooms to find the required room type
// 	for _, room := range hotel.Hotel.Rooms {
// 		if !found {
// 			roomHotel = models.RoomHotel{
// 				RoomType:      room.RoomType,
// 				PricePerNight: float32(room.PricePerNight),
// 			}
// 			found = true
// 		}

// 		log.Printf("Room Type: %s, Price Per Night: %.2f", roomHotel.RoomType, roomHotel.PricePerNight)
// 	}

// 	// Set the hotel ID in the roomHotel struct
// 	roomHotel.HotelId = int32(hotel.Hotel.HotelId)

// 	if !found {
// 		return models.RoomHotel{}, fmt.Errorf("no rooms found for hotel with ID %s", id)
// 	}

// 	return roomHotel, nil
// }



package hotel

import (
	"booking-service/models"
	"booking-service/proto/genhotel"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HandleHotel(id string) (models.RoomHotel, error) {
	log.Println("HandleHotel function started")
	conn, err := grpc.NewClient("hotel_service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error connecting to gRPC server:", err)
		return models.RoomHotel{}, err
	}
	defer conn.Close()

	client := genhotel.NewHotelServiceClient(conn)

	hotelResponse, err := client.GetHotelDetails(context.Background(), &genhotel.GetHotelDetailsRequest{HotelId: id})
	if err != nil {
		log.Println("Error getting hotel details:", err)
		return models.RoomHotel{}, err
	}

	
	hotel := models.RoomHotel{
		HotelId: int32(hotelResponse.Hotel.HotelId),
		Rooms:   make([]models.Room, len(hotelResponse.Hotel.Rooms)),
	}

	for i, room := range hotelResponse.Hotel.Rooms {
		hotel.Rooms[i] = models.Room{
			RoomType:     room.RoomType,
			PricePerNight: float32(room.PricePerNight),
		}
	}

	return hotel, nil
}
