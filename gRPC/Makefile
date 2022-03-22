create:
	protoc --proto_path=proto proto/*.proto  --go_out=:gen/ --go-grpc_out=:gen/



clean:
	rm gen/proto/*.go

go:
	go run server/*.go
