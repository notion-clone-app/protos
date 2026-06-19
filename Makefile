PLUGIN_GO=$(HOME)/go/bin/protoc-gen-go
PLUGIN_GRPC=$(HOME)/go/bin/protoc-gen-go-grpc
PLUGIN_GATEWAY=$(HOME)/go/bin/protoc-gen-grpc-gateway
# Добавляем плагин для OpenAPI спецификации
PLUGIN_OPENAPI=$(HOME)/go/bin/protoc-gen-openapiv2

OUT_DIR=./gen/go
TMP_SWAGGER_DIR=./gen/swagger

PATH_TO_PROTO ?= $(shell find . -name "*.proto" -not -path "*/google/*")

.PHONY: all clean generate

all: clean generate

clean:
	@echo "Cleaning up old generated files..."
	rm -rf $(OUT_DIR)
	rm -rf $(TMP_SWAGGER_DIR)

generate:
	@echo "Creating output directories..."
	mkdir -p $(OUT_DIR)
	mkdir -p $(TMP_SWAGGER_DIR)
	
	@echo "Running protoc compiler..."
	protoc -I . \
		--plugin=protoc-gen-go=$(PLUGIN_GO) \
		--plugin=protoc-gen-go-grpc=$(PLUGIN_GRPC) \
		--plugin=protoc-gen-grpc-gateway=$(PLUGIN_GATEWAY) \
		--plugin=protoc-gen-openapiv2=$(PLUGIN_OPENAPI) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUT_DIR) --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=$(TMP_SWAGGER_DIR) --openapiv2_opt=logtostderr=true,json_names_for_fields=true \
		$(PATH_TO_PROTO)

	@echo "Packing Swagger JSON into Go source code..."
	@# Берем сгенерированный json, экранируем его в сырую строку и пишем в файл рядом с sso.pb.go
	@echo "package ssov1\n\nconst SwaggerJSON = \`$$(cat $(TMP_SWAGGER_DIR)/proto/sso/sso.swagger.json)\`" > $(OUT_DIR)/proto/sso/swagger.pb.go
	
	@# Удаляем временную папку с "чистым" JSON, так как он нам больше не нужен на диске
	rm -rf $(TMP_SWAGGER_DIR)
	@echo "Generation successful! All files packed into $(OUT_DIR)"
