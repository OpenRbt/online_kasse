#!/bin/bash

cd cmd/web
go build
cd ..

./deploy_kasse.sh
