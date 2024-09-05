run:
	@go run ./cmd/main.go


tidy:
	@go mod tidy 

migrate:
	migrate create -dir ./storage/db -ext sql db

migrate-up:
	migrate -path ./storage/db -database "postgres://postgres:dilshod@localhost:5432/booking?sslmode=disable" up

migrate-down:
	migrate -path ./storage/db -database "postgres://postgres:dilshod@localhost:5432/booking?sslmode=disable" down

migrate-force:
	migrate -path ./storage/db -database "postgres://postgres:dilshod@localhost:5432/booking?sslmode=disable" force 
