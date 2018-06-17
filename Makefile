protoc:
	@cd alexa/ && protoc *.proto --go_out=plugins=grpc:.

deploy-create-aws-lambda:
	@cd deploy && GOOS=linux go build ../cmd/lambda/main.go
	@cd deploy && zip deployment.zip main
	@aws lambda create-function \
	--region eu-west-1 \
	--function-name HelloFunction \
	--zip-file fileb://deploy/deployment.zip \
	--runtime go1.x \
	--tracing-config Mode=Active \
	--role arn:aws:iam::<account-id>:role/<role-id> \
	--handler main

deploy-update-aws-lambda:
	@cd deploy && GOOS=linux go build ../cmd/lambda/main.go
	@cd deploy && zip deployment.zip main
	@aws lambda update-function-code \
	--region eu-west-1 \
	--function-name HelloFunction \
	--zip-file fileb://deploy/deployment.zip