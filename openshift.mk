## --------------------------------------
## Openshift specific make targets,
## intended to be included in root Makefile in this repository
## --------------------------------------

azure-cloud-controller-manager:
	openshift-hack/build-go.sh cloud-controller-manager
.PHONY: azure-cloud-controller-manager

azure-cloud-node-manager:
	openshift-hack/build-go.sh cloud-node-manager
.PHONY: azure-cloud-node-manager

binaries: azure-cloud-controller-manager azure-cloud-node-manager
.PHONY: binaries

verify-history:
	openshift-hack/verify-history.sh
.PHONY: verify-history

verify: test-lint test-gofmt test-govet
.PHONY: verify

test-unit-ci:
	openshift-hack/test-unit-ci.sh
.PHONY: test-unit-ci
