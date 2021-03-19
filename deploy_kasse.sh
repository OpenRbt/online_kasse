#!/bin/bash

function GetDIR {
  echo "$(pwd)" | sed 's/\//\\\//g'
}

CUR_DIR=$(GetDIR)
echo "CUR_DIR is $CUR_DIR"

sed -i "s/{DIR}/$CUR_DIR/g" kasse.service
sed -i "s/{USER}/$USER/g" kasse.service
sudo cp kasse.service /etc/systemd/system/
sudo systemctl enable kasse
