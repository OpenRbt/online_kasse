#!/bin/bash

function GetDIR {
  echo "$(pwd)" | sed 's/\//\\\//g'
}

cd ..
CUR_DIR=$(GetDIR)
echo "CUR_DIR is $CUR_DIR"
cd install

sed -i "s/{DIR}/$CUR_DIR/g" kasse.service
sudo cp kasse.service /etc/systemd/system/
sudo systemctl enable kasse
