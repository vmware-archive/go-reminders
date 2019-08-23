#!/bin/bash
# go-reminders kubernetes-deploy.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -x

# Save current directory
TOP="$(pwd)"

# install kubectl
LATEST="$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)"
if [ $? -ne 0 ]; then
    echo "Failed to obtain latest kubectl release. Aborting!"
    exit 1
fi

curl -LO https://storage.googleapis.com/kubernetes-release/release/${LATEST}/bin/linux/amd64/kubectl
if [ $? -ne 0 ]; then
    echo "Failed to download kubectl. Aborting!"
    exit 1
fi
chmod +x kubectl

# install kubectl
LATEST=${helm_ver}
SHA="$(curl -s https://storage.googleapis.com/kubernetes-helm/helm-${LATEST}-linux-amd64.tar.gz.sha256)"
if [ $? -ne 0 ]; then
    echo "Failed to obtain helm checksum. Aborting!"
    exit 1
fi

curl -LO https://storage.googleapis.com/kubernetes-helm/helm-${LATEST}-linux-amd64.tar.gz
if [ $? -ne 0 ]; then
    echo "Failed to download helm tarball. Aborting!"
    exit 1
fi

which sha256 >/dev/null 2>&1
if [ $? -eq 0 ]; then
    dsha=$(sha256 helm-${LATEST}-linux-amd64.tar.gz)
    if [ ! "{SHA}" == "${dsha}" ]; then
        echo "Checksum for helm download is incorrect."
        exit 1
    fi
fi
tar xvzf helm-${LATEST}-linux-amd64.tar.gz
mv linux-amd64/helm ${TOP}/
mv linux-amd64/tiller ${TOP}/
chmod +x ${TOP}/helm
chmod +x ${TOP}/tiller

# Expand the path to include k8s deploy tooling
export PATH=${PATH}:${TOP}

# get the tag for the docker container
tag=$(cat version/version)

# validate parameters
ret=0
if [ -z "${cluster_ca}" ]; then
    echo "ERROR: cluster_ca not supplied. Aborting!"
    ret=1
fi
if [ -z "${cluster_url}" ]; then
    echo "ERROR: cluster_url not supplied. Aborting!"
    ret=1
fi
if [ -z "${namespace}" ]; then
    echo "ERROR: namespace not supplied. Aborting!"
    ret=1
fi
if [ -z "${resource_type}" ]; then
    echo "ERROR: resource_type not supplied. Aborting!"
    ret=1
fi
if [ -z "${resource_name}" ]; then
    echo "ERROR: resource_name not supplied. Aborting!"
    ret=1
fi
if [ -z "${container}" ]; then
    echo "ERROR: container not supplied. Aborting!"
    ret=1
fi
if [ -z "${admin_key}" ]; then
    echo "ERROR: admin_key not supplied. Aborting!"
    ret=1
fi
if [ -z "${admin_cert}" ]; then
    echo "ERROR: admin_cert not supplied. Aborting!"
    ret=1
fi
if [ -z "${admin_token}" ]; then
    echo "WARN: admin_token not supplied. Using Certs Only!"
    ret=1
fi
if [ -z "${helm_ver}" ]; then
    echo "ERROR: helm version not supplied. Aborting!"
    ret=1
fi
if [ -z "${deployenv}" ]; then

    echo "ERROR: helm version not supplied. Aborting!"
    ret=1
fi
if [ -z "${tag}" ]; then
    echo "ERROR: tag (version) not supplied. Aborting!"
    ret=1
fi
if [ $ret -ne 0 ]; then
    exit $ret
fi

echo "build credentials"

echo "$cluster_ca" | base64 -d > ca.pem

# if using a bearer token or using minikube (or kind), admin keys are useless
if [ -z "$admin_token" -o "${admin_token}" == "MINIKUBE" ]; then
    echo "$admin_key" | base64 -d > key.pem
    echo "$admin_cert" | base64 -d > cert.pem
fi

# list the directory now for debugging purposes
ls -lat

# setup kube config
kubectl config set-cluster go-reminders --server=${cluster_url} --certificate-authority=${TOP}/ca.pem


# set kube user  
kubectl config set-credentials go-reminders --client-key=${TOP}/key.pem --client-certificate=${TOP}/cert.pem

# enable the context
kubectl config set-context go-reminders --user=go-reminders --cluster=go-reminders
kubectl config use-context go-reminders

# check kubectl for validity
kubectl get all --all-namespaces
if [ $? -ne 0 ]; then
    echo "kubectl failed to connect to master."
    exit 1
fi

# Kustomize the manifestst and deploy
kubectl kustomize kubernetes/overlays/${deployenv} | kubectl -n ${namespace} apply -f -
