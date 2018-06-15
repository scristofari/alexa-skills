protoc:
	@cd alexa/ && protoc *.proto --go_out=plugins=grpc:.