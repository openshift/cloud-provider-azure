FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.21-openshift-4.16 AS builder
WORKDIR /go/src/github.com/openshift/cloud-provider-azure
COPY . .

RUN make azure-cloud-node-manager ARCH=$(go env GOARCH)

FROM registry.ci.openshift.org/ocp/4.16:base-rhel9
COPY --from=builder /go/src/github.com/openshift/cloud-provider-azure/bin/cloud-node-manager /bin/azure-cloud-node-manager

LABEL description="Azure Cloud Node Manager"

ENTRYPOINT ["/bin/azure-cloud-node-manager"]
