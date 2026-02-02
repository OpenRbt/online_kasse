#!/bin/bash
export TAG=v0.0.3
docker build -t reg.registry.open-rbt.com/lcw-kasse:$TAG .
docker push reg.registry.open-rbt.com/lcw-kasse:$TAG
