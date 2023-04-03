watch:
	reflex -s -r '\.go$$' make run
	
proto:
	protoc -I ./protos/ ./protos/*.proto --go_out=./ --go-grpc_out=./