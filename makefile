.PHONY: gen build push apply all run
gen:
	### go grpc gen
	rm -rf ./pkg/api
	mkdir -p ./pkg/api
	protoc \
    	-I=./api \
    	--proto_path=./api \
    	--go_opt=paths=source_relative \
    	--go_out=./pkg/api \
    	--go-grpc_opt=paths=source_relative \
    	--go-grpc_out=./pkg/api \
    	$(shell find ./api -iname "*.proto") 2>&1 > /dev/null
	
	### js grpc gen
	rm -rf ./src/api
	mkdir -p ./src/api

	protoc -I=./api \
  		--js_out=import_style=commonjs,binary:./src/api \
  		--grpc-web_out=import_style=typescript,mode=grpcweb:./src/api \
		$(shell find ./api -iname "*.proto") 2>&1 > /dev/null
	
	
	# build deps
	go mod tidy
	yarn

build:
	docker build -t njpowell/testo .

push: 
	docker push njpowell/testo

apply: 
	-kubectl delete -k kustomize
	kubectl apply -k kustomize

all: generate build push apply

run:
	go run cmd/main.go

emissary-update:
	helm repo add datawire https://app.getambassador.io
	helm repo update

emissary: emissary-update
	kubectl apply -f https://app.getambassador.io/yaml/emissary/3.1.0/emissary-crds.yaml
	kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system
	helm install -n emissary --create-namespace \
    	emissary-ingress datawire/emissary-ingress && \
 		kubectl rollout status  -n emissary deployment/emissary-ingress -w
