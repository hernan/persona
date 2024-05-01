#! /bin/bash

source .env

export GOOSE_DRIVER=mysql
export GOOSE_DBSTRING="$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME?parseTime=true"
export GOOSE_MIGRATION_DIR=./migrations

~/go/bin/goose $@