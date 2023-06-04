#!/bin/bash

# NOTE: Execute with `source` keyword
# source ./setEnvs.sh

filename='./.env'
n=1
while read -r line; do
  # reading each line
  export $line
  n=$((n+1))
done < $filename