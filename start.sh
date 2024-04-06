#!/bin/sh
set -e
echo "run db migration"
#/app/migrate -path /app/migration -database "$DATA_SOURCE" -verbose up

echo "start the app"

###exec all the parameters
exec "$@"