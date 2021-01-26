TAG?=latest
NS?=openfaas

.PHONY: build-gateway
build-gateway:
	(buildctl build --opt build-arg:http_proxy=$http_proxy --opt build-arg:https_proxy=$https_proxy --frontend=dockerfile.v0 --local context=./gateway/ --local dockerfile=./gateway/ --output type=image,name=docker.io/lwyan/gateway:latest)

# .PHONY: test-ci
# test-ci:
# 	./contrib/ci.sh
