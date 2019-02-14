#!/bin/bash
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -eu


if [ $# -ne 1 ]; then
  echo "Usage: $0 [create|destroy]"
  exit -1
fi

SUBCMD=$1

case "$SUBCMD" in

  create)
    docker-compose up -d
    sleep 10
    docker run -d --link goreminders_db_1:db --link goreminders_etcdsrv_1:etcdsrv \
      --name go-reminders -p 8080:8080 greent/go-reminders \
      -cfgurl etcdsrv:4001 -host db
    ;;

  destroy)
    docker-compose stop
    docker-compose rm -f
    docker stop go-reminders
    docker rm go-reminders
    ;;
  *)
    echo "Invalid option $1"

esac
