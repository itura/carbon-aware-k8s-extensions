 #!/usr/bin/env bash

 # Run swagger UI on openapi spec for Carbon Aware SDK WebApi

 PORT=${1:-80}

 docker run \
  -p $PORT:8080 \
  -v /"$(PWD)/caapi/api":/spec \
  -e SWAGGER_JSON=/spec/openapi.yaml \
  swaggerapi/swagger-ui