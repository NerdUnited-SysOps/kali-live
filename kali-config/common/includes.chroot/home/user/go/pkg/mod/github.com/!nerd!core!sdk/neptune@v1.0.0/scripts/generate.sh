#!/usr/bin/env sh
#

cat "$1" | python3 -c 'import json,sys;obj=json.load(sys.stdin);print(json.dumps(obj["abi"]))' > "$2"
cat "$1" | python3 -c 'import json,sys;obj=json.load(sys.stdin);print(obj["bytecode"])' > "$3"
