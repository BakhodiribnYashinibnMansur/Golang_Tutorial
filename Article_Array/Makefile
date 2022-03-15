swag-init:
	swag init -g api/api.go

vendor-update:
	go get -u ./...
	go mod vendor

run:
	go run command/main.go

startDocker:
	sudo docker start article

install:
	swag init -g api/api.go
	go mod download
	go mod vendor
	go run command/main.go