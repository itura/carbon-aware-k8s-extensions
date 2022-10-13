 #!/usr/bin/env bash

 # start self-built image for CAAPI

PORT=${1:-8080}

docker run \
  -it --rm \
  -p $PORT:80 \
  -e CarbonAwareVars__CarbonIntensityDataSource=JSON \
  ca-api

