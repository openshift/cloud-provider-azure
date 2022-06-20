FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.17-openshift-4.10 AS builder
WORKDIR /go/src/github.com/openshift/cloud-provider-azure
COPY . .

RUN make azure-cloud-controller-manager

FROM registry.ci.openshift.org/ocp/4.10:base
COPY --from=builder /go/src/github.com/openshift/cloud-provider-azure/bin/cloud-controller-manager /bin/azure-cloud-controller-manager

LABEL description="Azure Cloud Controller Manager"

ENTRYPOINT ["/bin/azure-cloud-controller-manager"]
