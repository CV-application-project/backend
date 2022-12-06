gen-api-proto:
	protoc $(folder_path)/api/*.proto --go_out=$(folder_path)/api --grpc-gateway_out=$(folder_path)/api --go-grpc_out=$(folder_path)/api --go_opt=paths=import --go-grpc_opt=paths=import --validate_out=paths=source_relative,lang=go:./
update:
	go mod tidy
	go mod vendor