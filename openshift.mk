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

# due to the way the linter binary packages are included in builds for OpenShift,
# we only want to run the linting target if the binary is defined.
have_linter = $(LINTER)
ifneq ($(have_linter),)
	linter_target = lint
else
	linter_target =
endif

verify: $(linter_target)
.PHONY: verify

test-unit-ci:
	openshift-hack/test-unit-ci.sh
.PHONY: test-unit-ci
