#!/bin/bash
echo "SAMPLE BUILD"
docker buildx build --platform=linux/amd64 -t romanfedyashov/lcw-kasse:v0.0.1 .
