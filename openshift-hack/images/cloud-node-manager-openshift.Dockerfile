FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.16-openshift-4.8 AS builder
WORKDIR /go/src/github.com/openshift/cloud-provider-azure
COPY . .

RUN make azure-cloud-node-manager

FROM registry.ci.openshift.org/ocp/4.8:base
COPY --from=builder /go/src/github.com/openshift/cloud-provider-azure/bin/cloud-node-manager /bin/azure-cloud-node-manager

LABEL description="Azure Cloud Node Manager"

ENTRYPOINT ["/bin/azure-cloud-node-manager"]
