build:
	docker build -t user-cli .

run:
	docker run -e MICRO_REGISTRY=mdns user-cli

runLocal:
	go build . && ./user-cli
