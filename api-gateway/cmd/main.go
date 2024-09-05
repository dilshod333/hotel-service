package main

import (
	"api-gateway/api"
	"api-gateway/jwt"
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	server := api.Conn()

	r := gin.Default()

	r.POST("api/users/register", func(c *gin.Context) {
		server.RegisterUser(c)
	})
	r.POST("api/users/verify", func(ctx *gin.Context) {
		server.VerifyUser(ctx)
	})

	r.POST("api/users/login", func(ctx *gin.Context) {
		server.LoginUser(ctx)
	})

	r.DELETE("api/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.IndentedJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		valid, email, err := jwt.VerifyToken(token)
		if err != nil || !valid {
			ctx.JSON(401, gin.H{"error": "Unauthorized: Invalid or expired token"})
			return
		}
		log.Println(">>>>> token email", email)
		server.DeleteUserr(ctx, id)
	})

	r.GET("api/users/:id", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.IndentedJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		valid, email, err := jwt.VerifyToken(token)
		if err != nil || !valid {
			c.JSON(401, gin.H{"error": "Unauthorized: Invalid or expired token"})
			return
		}
		log.Println(">>>>> token email", email)
		id := c.Param("id")
		server.GetUserr(c, id)
	})

	r.GET("api/hotels", func(ctx *gin.Context) {
		server.GetHotelss(ctx)
	})

	r.GET("api/hotels/:id", func(ctx *gin.Context) {
		a := ctx.Param("id")
		server.GetHotelID(ctx, a)
	})

	r.POST("api/hotels/create", func(ctx *gin.Context) {
		server.CreateHotelll(ctx)
	})

	r.GET("api/hotels/rooms/availabilty/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		server.CheckRoom(ctx, idStr)
	})

	r.POST("api/booking/", func(ctx *gin.Context) {
		server.CreateBookingg(ctx)
	})

	r.GET("api/booking/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error on id", err)
			ctx.IndentedJSON(500, gin.H{"error": "error on convert id"})
			return
		}
		server.GetBookingIdd(ctx, intId)
	})

	r.PUT("api/booking/update/:id", func(ctx *gin.Context) {
		server.UpdateBokkingID(ctx)

	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("API Gateway running on :%s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server foreeedd to shutdown:", err)
	}

	log.Println("Server stopped")
}
