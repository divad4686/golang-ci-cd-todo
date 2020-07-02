#!/bin/bash
set -e

host="http://cicdexample.com/staging/todoapi"
# host="localhost:8080"

curl $host/ping

echo ""
result=$(curl $host/todos \
  -H  "accept: application/json" \
  -H  "Content-Type: application/json" \
  -d '{
      "Title":"comprar",
      "Text":"comprar cosas",
      "Completed":false
    }')

echo $result

url=$(jq -r '.url' <<< $result)

curl $url
