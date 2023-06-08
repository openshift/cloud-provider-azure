FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.20-openshift-4.14 AS builder
WORKDIR /go/src/github.com/openshift/cloud-provider-azure
COPY . .

RUN make azure-cloud-controller-manager ARCH=$(go env GOARCH)

FROM registry.ci.openshift.org/ocp/4.14:base
COPY --from=builder /go/src/github.com/openshift/cloud-provider-azure/bin/cloud-controller-manager /bin/azure-cloud-controller-manager

LABEL description="Azure Cloud Controller Manager"

ENTRYPOINT ["/bin/azure-cloud-controller-manager"]
