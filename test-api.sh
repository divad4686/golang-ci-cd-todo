#!/bin/bash
set -e

host="http://cicdexample.com/staging/todoapi"
# host="localhost:8080"

curl $host/ping
echo ""
curl $host/todos \
  -H  "accept: application/json" \
  -H  "Content-Type: application/json" \
  -d '{
      "Title":"comprar",
      "Text":"comprar cosas",
      "Completed":false
    }'