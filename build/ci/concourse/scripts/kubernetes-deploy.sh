#!/bin/bash
# go-winery kubernetes-deploy.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

# Save current directory
TOP="$(pwd)"

# install kubectl
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.8.3/bin/linux/amd64/kubectl
chmod +x kubectl

image=$(cat $image_name)
tag=$(cat $image_tag)

echo "$cluster_ca" | base64 -d > ca.pem

# if using a bearer token, admin keys are useless
if [ -z "$admin_token" ]; then
    echo "$admin_key" | base64 -d > key.pem
    echo "$admin_cert" | base64 -d > cert.pem
fi

# list the directory now for debugging purposes
ls -lat

ret=0

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
if [ -z "${container_name}" ]; then
    container_name=$resource_name
fi
if [ -z "${image}" ]; then
    echo "ERROR: image_name not supplied. Aborting!"
    ret=1
fi
if [ -z "${tag}" ]; then
    echo "ERROR: image_tag not supplied. Aborting!"
    ret=1
fi

if [ $ret -ne 0 ]; then
    exit $ret
fi

# build kubectl command line
KUBECTL="./kubectl --server=$cluster_url --namespace=$namespace --certificate-authority=ca.pem"
if [ -z "$admin_token" ]; then
    KUBECTL="$KUBECTL --client-key=key.pem --client-certificate=cert.pem"
else
    KUBECTL="$KUBECTL --token=${admin_token}"
fi

$KUBECTL set image deployment/$resource_name $container_name=$image:$tag

echo ""
