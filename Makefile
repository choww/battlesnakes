ECR_URL = $(AWS_ID).dkr.ecr.us-west-2.amazonaws.com/core/battlesnake
SNAKE_NAME ?=carmen-go
# Version is done like this because the docker images are tagged by the commit hash, so if you just pass it v0.1.2 it won't find an image
# IWe have to search the repo for what commit a particular tag refers to.
VERSION := $(if $(VERSION),$(shell git rev-parse --short $(VERSION) || echo $(VERSION)),$(shell git rev-parse --short HEAD))
COMMON_MANIFESTS := deployment service ingress

.PHONY: deploy
deploy: ## Deploy Kubernetes resources
	$(call check_variable, AWS_REGION, You must specify the region to use (eg. 'us-east-1'))
	$(call check_variable, ENVIRONMENT, You must specify the environment to use (eg. 'sandbox'))
	@$(MAKE) -f Makefile $(foreach manifest,$(COMMON_MANIFESTS),deploy-manifest-$(manifest))
	kubectl rollout status --namespace battlesnake deployment/$(SNAKE_NAME)

.PHONY: deploy-manifest-%
deploy-manifest-%: ## Deploy a Kubernetes manifest
	@ECR_URL=$(ECR_URL) \
	VERSION=$(VERSION) \
		iidy render \
			--region us-west-2 \
			--environment integration \
			infrastructure/kube/$*.yaml \
		| kubectl apply -f -

.PHONY: cleanup
cleanup: ## Remove all Kubernetes resources
	$(call check_variable, AWS_REGION, You must specify the region to use (eg. 'us-east-1'))
	$(call check_variable, ENVIRONMENT, You must specify the environment to use (eg. 'sandbox'))
	ECR_URL=$(ECR_URL) \
	VERSION=$(VERSION) \
		iidy render \
			--region us-west-2 \
			--environment integration \
			infrastructure/kube/ \
		| kubectl delete -f -

.PHONY: debug
debug: ## Show status of all Kubernetes resources
	kubectl describe pods --namespace battlesnake

ecr-login: # Login to ECR
	aws ecr get-login-password --region us-west-2 | docker login --password-stdin --username AWS $(AWS_ID).dkr.ecr.us-west-2.amazonaws.com

.PHONY: release
release: ecr-login ## Push Docker container to ECR
	docker push $(ECR_URL):$(SNAKE_NAME)-$(VERSION)

.PHONY: run
run: ## Run application in Docker container
	docker run --rm -v $(shell pwd):/opt/app $(SNAKE_NAME):latest

.PHONY: build
build: ## Build Docker container
	docker build -t $(ECR_URL):$(SNAKE_NAME)-$(VERSION) -t $(SNAKE_NAME):latest .

.PHONY: help
help: ## Display this message
	@grep --no-filename -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort | awk 'BEGIN {FS = ":.*?## "} \
		{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'	
