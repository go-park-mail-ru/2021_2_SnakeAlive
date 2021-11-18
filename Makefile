protogen-api-with-validator:
	protoc 								\
    		-I. 								\
    		-I./third_party 								\
    		-I./third_party/googleapis 								\
    		-I./third_party/grpc-gateway/v2 								\
    		--go_out=plugins=grpc,paths=source_relative:. 								\
    		--validate_out=lang=go,paths=source_relative:. \
    		$(path)


protogen-api-auth-service:
	make protogen-api-with-validator path=pkg/services/auth/api.proto

protogen-all-service:
	make protogen-api-auth-service