.PHONY: gen-grpc
gen-grpc:
	rm -rf ./src/api
	mkdir -p ./src/api
	rm -rf ./pkg/api
	mkdir -p ./pkg/api
	protoc \
    	-I=./api \
    	--proto_path=./api \
    	--go_opt=paths=source_relative \
    	--go_out=./pkg/api \
    	--go-grpc_opt=paths=source_relative \
    	--go-grpc_out=./pkg/api \
		--js_out=import_style=commonjs,binary:./src/api \
  		--grpc-web_out=import_style=typescript,mode=grpcweb:./src/api \
    	$(shell find ./api -iname "*.proto") 2>&1 > /dev/null

.PHONY: gen-go
gen-go:
	go mod tidy
	go generate ./...

.PHONY: gen-ts
gen-ts:	
	yarn

.PHONY: gen
gen: gen-grpc gen-go gen-ts

.PHONY: build
build:
	docker build -t njpowell/parthenon .

.PHONY: push
push: 
	docker push njpowell/parthenon

.PHONY: apply
apply:
	-kubectl delete -k kustomize
	kubectl apply -k kustomize

.PHONY: all
all: build push apply

# runs main.go
.PHONY: run
run:
	go run cmd/main.go

.PHONY: install
install:
	go build -o ~/go/bin/parth cmd/main.go

# starts the ts server
.PHONY: start
start:
	yarn
	yarn start

.PHONY: emissary-update
emissary-update:
	helm repo add datawire https://app.getambassador.io
	helm repo update

.PHONY: emissary
emissary: emissary-update
	kubectl apply -f https://app.getambassador.io/yaml/emissary/3.2.0/emissary-crds.yaml
	kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system
	helm install -n emissary --create-namespace \
    	emissary-ingress datawire/emissary-ingress && \
 		kubectl rollout status -n emissary deployment/emissary-ingress -w

.PHONY: db
db:
	-docker rm -f mysql
	docker run --name mysql -p 3306:3306 --rm -d -it -e MYSQL_ROOT_PASSWORD=password mysql

.PHONY: dbcli
dbcli:
	kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword

.PHONY: test-ts
test-ts: gen-ts
	yarn test --watchAll=false

.PHONY: test-go
test-go: lint-go
	go test -run TestIndex ./...

.PHONY: lint-go
lint-go:
	errcheck -ignoretests ./...

.PHONY: test
test: test-go test-ts

.PHONY: tools
tools: tools-kubectl tools-protobuf tools-grpc-web


.PHONY: tools-kubectl
tools-kubectl:
	brew install kubectl

# pins protoc at v3.20 because v3.21 is broken
# TODO remove overwrite when v3.22 is released
.PHONY: tools-protobuf
tools-protobuf:
	brew install protobuf@3
	brew link --overwrite protobuf@3

# downloads amd binary pinned at 1.4.2
# TODO get latest instead?
.PHONY: tools-grpc-web
tools-grpc-web:
	curl https://github.com/grpc/grpc-web/releases/download/1.4.2/protoc-gen-grpc-web-1.4.2-darwin-x86_64 \
		-Lo /usr/local/bin/protoc-gen-grpc-web
	sudo chmod +x /usr/local/bin/protoc-gen-grpc-web
