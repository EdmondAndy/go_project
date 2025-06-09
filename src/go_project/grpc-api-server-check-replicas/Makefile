KUBECTL := kubectl
KUBECONFIG ?= $(HOME)/.kube/config
DEPLOY_YAML := test/testdata/deployment.yaml
TEST_NS := default

.PHONY: test deploy clean

deploy:
	@echo "Deploying test deployment..."
	$(KUBECTL) apply -f $(DEPLOY_YAML)

test: deploy
	@echo "Running Go integration tests..."
	go test ./test -v

clean:
	@echo "Cleaning up..."
	$(KUBECTL) delete -f $(DEPLOY_YAML) --ignore-not-found
