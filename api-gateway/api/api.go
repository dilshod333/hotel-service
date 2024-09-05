package api

import (
	"api-gateway/methods"
	"api-gateway/proto/genbooking"
	"api-gateway/proto/genhotel"
	"api-gateway/proto/genuser"
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Serverr struct {
	userr    genuser.UserServiceClient
	hotell   genhotel.HotelServiceClient
	bookingg genbooking.BookingServiceClient
}

func Conn() *Serverr {
	log.Println("Logger is working on Conn")
	user := methods.ConnectUser()
	hotel := methods.ConnectHotel()
	book := methods.CoonectBooking()
	return &Serverr{userr: user, hotell: hotel, bookingg: book}
}

func (u *Serverr) RegisterUser(c *gin.Context) {
	log.Println("Registeruser  ichiga kirdi apida... name")
	var req genuser.RegisterReq

	var err error
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := u.userr.RegisterUser(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(200, resp)
}

func (s *Serverr) VerifyUser(c *gin.Context) {
	log.Println("Verify userrr")

	var req genuser.VerifyReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	resp, err := s.userr.VerifyUser(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Error on verify check email or password"})
		return
	}

	c.JSON(200, resp)

}

func (s *Serverr) LoginUser(c *gin.Context) {
	var req genuser.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
	}

	resp, err := s.userr.LoginUser(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error on login user"})
		return
	}

	c.JSON(200, resp)

}

func (s *Serverr) DeleteUserr(c *gin.Context, id string) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	req := &genuser.DeleteUserReq{
		UserId: int32(idInt),
	}

	resp, err := s.userr.DeleteUser(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, resp)
}

func (s *Serverr) GetUserr(c *gin.Context, id string) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	req := &genuser.GetUserReq{
		UserId: int32(idInt),
	}

	resp, err := s.userr.GetUser(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to Getuser"})
		return
	}

	c.JSON(200, resp)

}

func (s *Serverr) GetHotelss(c *gin.Context) {

	req := &genhotel.GetHotelsRequest{}

	resp, err := s.hotell.GetHotels(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": "did not  get hotels"})
		return
	}

	c.JSON(200, resp)
}

func (s *Serverr) GetHotelID(c *gin.Context, id string) {
	req := &genhotel.GetHotelDetailsRequest{HotelId: id}

	resp, err := s.hotell.GetHotelDetails(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": "did not  get hotel details"})
		return
	}

	c.JSON(200, resp)
}

func (s *Serverr) CreateHotelll(c *gin.Context) {
	var req genhotel.HotelReq
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Error on shouldbind", err)
		return
	}

	resp, err := s.hotell.AddHotel(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "error on create hotel"})
		return
	}
	c.JSON(200, resp)

}

func (s *Serverr) CheckRoom(c *gin.Context, id string) {
	req := &genhotel.CheckRoomAvailabilityRequest{HotelId: id}

	resp, err := s.hotell.CheckRoomAvailability(c, req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Not found or error on checkroomavailabilty"})
		return
	}

	c.IndentedJSON(200, resp)
}

func (s *Serverr) CreateBookingg(c *gin.Context) {
	var req genbooking.BookingRequest

	if err := c.ShouldBind(&req); err != nil {
		c.IndentedJSON(500, gin.H{"error": "error on shouldbind"})
		return
	}

	resp, err := s.bookingg.CreateBooking(c, &req)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Error on createbook check it again"})
		return
	}

	c.IndentedJSON(200, resp)
}

func (s *Serverr) GetBookingIdd(c *gin.Context, id int) {
	req := &genbooking.BookingIdReq{BookingId: int32(id)}

	resp, err := s.bookingg.GetBooking(c, req)

	if err != nil {
		c.JSON(500, gin.H{"error": "error on getbookingid"})
		return
	}
	c.JSON(200, resp)
}

func (s *Serverr) UpdateBokkingID(c *gin.Context) {

	id := c.Param("id")


	bookingID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid booking ID"})
		return
	}

	
	var req genbooking.UpdateBookIdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error on ShouldBindJSON:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	
	req.BookingId = int32(bookingID)


	resp, err := s.bookingg.UpdateBooking(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating booking"})
		return
	}


	c.JSON(200, resp)
}
