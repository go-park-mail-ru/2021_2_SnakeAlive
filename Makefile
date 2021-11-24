protogen-api-with-validator:
	protoc 								\
    		-I. 								\
    		-I./third_party 								\
    		-I./third_party/googleapis 								\
    		--go_out=. --go_opt=paths=source_relative 				\
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    		--validate_out=lang=go,paths=source_relative:. \
    		$(path)


protogen-api-auth-service:
	make protogen-api-with-validator path=pkg/services/auth/api.proto

protogen-api-trip-service:
	make protogen-api-with-validator path=pkg/services/trip/api.proto

protogen-api-sight-service:
	make protogen-api-with-validator path=pkg/services/sight/api.proto

protogen-api-review-service:
	make protogen-api-with-validator path=pkg/services/review/api.proto

protogen-all-services:
	make protogen-api-auth-service && \
	make protogen-api-trip-service && \
	make protogen-api-sight-service

prepare-auth_service-env:
	export USER_DB_URL="postgres://tripadvisor:12345@localhost:5432/tripadvisor" && \
			export USER_GRPC_PORT="10123" && export USER_PREFIX_LEN="0"


prepare-trip_service-env:
	export TRIP_DB_URL="postgres://tripadvisor:12345@localhost:5432/tripadvisor" && \
			export TRIP_GRPC_PORT="6666"

prepare-gateway-env:
	export GATEWAY_HTTP_PORT=":8080" && export GATEWAY_AUTH_ENDPOINT="localhost:10123" && export GATEWAY_TRIP_ENDPOINT="localhost:6666"

run-auth:
	make prepare-auth_service-env && go run cmd/auth_service/main.go
run-trip:
	make prepare-trip_service-env && go run cmd/trip_service/main.go
run-gateway:
	make prepare-gateway-env && go run cmd/gateway/main.go
