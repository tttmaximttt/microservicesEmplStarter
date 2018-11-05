build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli

runLocal:
	go build . && ./consignment-cli ${file} ${token}

