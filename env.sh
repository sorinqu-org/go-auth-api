#!/usr/bin/env sh
set -e

set -a
. ./.env
set +a

./bin/app
