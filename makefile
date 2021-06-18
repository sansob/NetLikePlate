include .env

gen:
	protoc --gofast_out=plugins=grpc:./models/proto/*.proto ./models/proto/*

gen-cov:
	go test main_test.go -coverprofile=coverage.out