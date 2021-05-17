#!/usr/bin/env bash

set -eu


REPO=github.com/openshift/cloud-provider-azure
REPO_ROOT="$(git rev-parse --show-toplevel)"
WHAT=${1:-cloud-controller-manager}
GLDFLAGS=${GLDFLAGS:-}

eval $(go env | grep -e "GOHOSTOS" -e "GOHOSTARCH")

: "${GOOS:=${GOHOSTOS}}"
: "${GOARCH:=${GOHOSTARCH}}"


cd "$REPO_ROOT"
if [ -z ${VERSION_OVERRIDE+a} ]; then
	echo "Using version from git..."
	VERSION_OVERRIDE=$(git describe --abbrev=8 --dirty --always)
fi

VERSION_PKG=sigs.k8s.io/cloud-provider-azure/pkg/version

GLDFLAGS="-extldflags '-static'"
GLDFLAGS="$GLDFLAGS -X $VERSION_PKG.gitVersion=${REPO}"
GLDFLAGS="$GLDFLAGS -X $VERSION_PKG.gitCommit=${VERSION_OVERRIDE}"
GLDFLAGS="$GLDFLAGS -X $VERSION_PKG.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

eval $(go env)

echo "Building ${REPO}/cmd/${WHAT} (${VERSION_OVERRIDE})"
GO111MODULE=${GO111MODULE} CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${GOFLAGS} -ldflags "${GLDFLAGS}" -o bin/${WHAT} ${REPO_ROOT}/cmd/${WHAT}
