build:
	protoc -I. --go_out=plugins=micro:. ./proto/consignment/consignment.proto
	docker build -t consignment-service:latest .

run:
	docker run -p 50051:50051 --net="host" \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e DISABLE_AUTH=true \go
	consignment-service

