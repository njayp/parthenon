.PHONY: gen build push apply all run emissary emissary-update db dbcli test-go test-ts test tools tools-kubectl tools-protobuf tools-grpc-web lint lint-ts lint-go

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

gen-go:
	go mod tidy
	go generate ./...

gen-ts:	
	yarn

gen: gen-grpc gen-go gen-ts

build:
	docker build -t njpowell/parthenon .

push: 
	docker push njpowell/parthenon

apply:
	-kubectl delete -k kustomize
	kubectl apply -k kustomize

all: build push apply

# runs main.go
run:
	go run cmd/main.go

# starts the ts server
start:
	yarn
	yarn start

emissary-update:
	helm repo add datawire https://app.getambassador.io
	helm repo update

emissary: emissary-update
	kubectl apply -f https://app.getambassador.io/yaml/emissary/3.2.0/emissary-crds.yaml
	kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system
	helm install -n emissary --create-namespace \
    	emissary-ingress datawire/emissary-ingress && \
 		kubectl rollout status -n emissary deployment/emissary-ingress -w

db:
	-docker rm -f mysql
	docker run --name mysql -p 3306:3306 --rm -d -it -e MYSQL_ROOT_PASSWORD=password mysql

dbcli:
	kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword

test-ts:
	yarn test --watchAll=false
	
test-go: lint-go
	go test -run \S*TestIndex$  ./...

test: test-go test-ts

tools: tools-kubectl tools-protobuf tools-grpc-web

tools-kubectl:
	brew install kubectl

# pins protoc at v3.20 because v3.21 is broken
# TODO remove overwrite when v3.22 is released
tools-protobuf:
	brew install protobuf@3
	brew link --overwrite protobuf@3

# downloads amd binary pinned at 1.4.2
# TODO get latest instead?
tools-grpc-web:
	curl https://github.com/grpc/grpc-web/releases/download/1.4.2/protoc-gen-grpc-web-1.4.2-darwin-x86_64 \
		-Lo /usr/local/bin/protoc-gen-grpc-web
	sudo chmod +x /usr/local/bin/protoc-gen-grpc-web

lint-go:
	errcheck -ignoretests ./...