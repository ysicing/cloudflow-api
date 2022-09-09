#!/usr/bin/env bash

[ -d "" ]

go mod tidy
go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
    exit $retVal
fi

set -ex
TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api
cp -r ./{hack,vendor,go.mod} "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api

git clone --depth 1 https://github.com/ysicing/cloudflow.git "${TMP_DIR}"/src/github.com/ysicing/cloudflow

cp -r "${TMP_DIR}"/src/github.com/ysicing/cloudflow/apis/apps "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api

(cd "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api; \
    GOPATH=${TMP_DIR} GO111MODULE=off /bin/bash vendor/k8s.io/code-generator/generate-groups.sh all \
    github.com/ysicing/cloudflow-api/client github.com/ysicing/cloudflow-api "apps:v1beta1" -h ./hack/boilerplate.go.txt ;
    )

rm -rf ./client/{clientset,informers,listers} ./apps
tree "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api/client/
mv "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api/client .
mv "${TMP_DIR}"/src/github.com/ysicing/cloudflow-api/apps .
rm -rf ${TMP_DIR}
rm -rf vendor
