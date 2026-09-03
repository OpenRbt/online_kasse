#!/bin/bash
export TAG=v0.0.8
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t reg.registry.open-rbt.com/lcw-kasse:$TAG \
  --push .

