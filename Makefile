migrateup:
	migrate -path ./migrations -database "postgres://postgres:123456@localhost:5432/golang?sslmode=disable" up

migratedown:
	migrate -path ./migrations -database "postgres://postgres:123456@localhost:5432/golang?sslmode=disable" down
