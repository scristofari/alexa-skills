protoc:
	@cd alexa/ && protoc *.proto --go_out=plugins=grpc:.

deploy-aws-lambda:
	@cd deploy && GOOS=linux go build ../cmd/lambda/main.go
	@cd deploy && zip deployment.zip main