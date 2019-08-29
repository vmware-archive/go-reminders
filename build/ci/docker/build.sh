#!/bin/bash
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#
# This build script assumes there are two important build packages available
# in ~/bin :
#
#    linux-amd64: this is the helm (latest) download; and
#    kubectl: this is the (latest) kubectl download.
#
# Note that in both cases, the executables (e.g., ~/bin/linux-amd64/bin/helm)
# already have the executable mode set (i.e. chmod 755 ~/bin/kubectl).
#

BUILDIMAGE=${BUILDIMAGE:=reminders-build:1.0.0}

# Grab the latest executables
cp -a ${HOME}/bin/linux-amd64 .
cp -a ${HOME}/bin/kubectl .

# Build and push the container
docker build --rm -t ${BUILDIMAGE} .

if [ -z "$1" ]; then
    docker push ${BUILDIMAGE}
fi
