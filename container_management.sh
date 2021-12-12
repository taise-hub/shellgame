#!/bin/bash
# cronで15分おきくらいに実行するように設定する。
docker ps --filter "ancestor=shellgame" | awk '{if($4>30 && NR!=1) print $1}' | xargs -r docker container rm -f