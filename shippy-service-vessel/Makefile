build:
	protoc --proto_path=proto/vessel --go_out=proto/vessel --micro_out=proto/vessel proto/vessel/vessel.proto
	docker build -t shippy-service-vessel .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns shippy-service-vessel