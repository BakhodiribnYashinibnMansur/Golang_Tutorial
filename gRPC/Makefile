genp:
	# protoc --proto_path=proto proto/*.proto  --go_out=:gen/ --go-grpc_out=:gen/
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:gen

rm:
	rm gen/proto/*.go

server:
	go run command/server/*.go --port 8080

client:
	go run command/client/*.go --address 0.0.0.0:8080
