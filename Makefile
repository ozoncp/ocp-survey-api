run:
	go run cmd/ocp-survey-api/main.go

lint:
	golangci-lint run

test:
	go test -v ./...


.PHONY: build
build: vendor-proto .generate .build

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-survey-api cmd/ocp-survey-api/main.go


.PHONY: generate
generate: vendor-proto .generate

PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-survey-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-survey-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-survey-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-survey-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-survey-api/ocp-survey-api.proto
		mv pkg/ocp-survey-api/github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api/* pkg/ocp-survey-api/
		rm -rf pkg/ocp-survey-api/github.com
		mkdir -p cmd/ocp-survey-api


.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/ocp-survey-api/main.go


.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-survey-api
		cp api/ocp-survey-api/ocp-survey-api.proto vendor.protogen/api/ocp-survey-api/ocp-survey-api.proto
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init github.com/ozoncp/ocp-survey-api
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
