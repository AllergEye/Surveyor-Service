mocks:
	mockgen -source=./pkg/surveyor/restaurant/router.go -destination=./pkg/surveyor/mocks/restaurant/router.go
	mockgen -source=./pkg/surveyor/restaurant/controller.go -destination=./pkg/surveyor/mocks/restaurant/controller.go
	mockgen -source=./pkg/surveyor/restaurant/service.go -destination=./pkg/surveyor/mocks/restaurant/service.go
	mockgen -source=./internal/database/restaurant_repository.go -destination=./pkg/surveyor/mocks/database/restaurant_repository.go
	mockgen -source=./internal/lib/helpers.go -destination=./pkg/surveyor/mocks/lib/helpers.go

test:
	go test ./... -cover