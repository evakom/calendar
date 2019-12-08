gen:
	protoc --go_out=plugins=grpc:internal/grpc api/*.proto

build: gen
	go build -o grpc_client cmd/grpc_client/main.go
	go build -o grpc_server cmd/grpc_server/main.go
	go build -o http_server cmd/http_server/main.go
	go build -o alert_publisher cmd/notifications/publisher/*.go
	go build -o alert_sender cmd/notifications/sender/*.go

test: build
	go test ./...

docker-http-server:
	go build -o http_server cmd/http_server/main.go
	docker build -t calendar-http-server -f ./deployments/docker/http_server/Dockerfile .
	docker run --rm -d -p 8080:8080 --name calendar-http-server --network calendar_calendar calendar-http-server