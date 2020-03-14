#!/bin/bash

cd cmd/web
go build
cd ..

cd drivers/web_kasse
./build.sh
cp kasse.exe ../../cmd/web
cd ../..

