# Пути к плагинам
PLUGIN_GO=$(HOME)/go/bin/protoc-gen-go
PLUGIN_GRPC=$(HOME)/go/bin/protoc-gen-go-grpc
PLUGIN_GATEWAY=$(HOME)/go/bin/protoc-gen-grpc-gateway

# Папка, куда складываем весь сгенерированный код
OUT_DIR=./gen/go

# Переменная по умолчанию. Если пользователь не передал путь, 
# ищем вообще все .proto файлы в репозитории (игнорируя папку google с зависимостями)
PATH_TO_PROTO ?= $(shell find . -name "*.proto" -not -path "*/google/*")

.PHONY: all clean generate

all: clean generate

clean:
	@echo "Cleaning up old generated files..."
	rm -rf $(OUT_DIR)

generate:
	@echo "Creating output directory..."
	mkdir -p $(OUT_DIR)
	@echo "Running protoc compiler for: $(PATH_TO_PROTO)"
	protoc -I . \
		--plugin=protoc-gen-go=$(PLUGIN_GO) \
		--plugin=protoc-gen-go-grpc=$(PLUGIN_GRPC) \
		--plugin=protoc-gen-grpc-gateway=$(PLUGIN_GATEWAY) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUT_DIR) --grpc-gateway_opt=paths=source_relative \
		$(PATH_TO_PROTO)
	@echo "Generation successful!"
