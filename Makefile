
mockgen:
		mockgen -source=./internal/storage/storage.go -destination=./mock/storage_mock.go -package=mock

run:
		go run cmd/main.go
