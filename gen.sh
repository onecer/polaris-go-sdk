#!/usr/bin/env bash
wget -O swagger.json http://127.0.0.1:8090/apidocs.json
openapi-generator generate -i swagger.json -g go --package-name polaris -o ./
