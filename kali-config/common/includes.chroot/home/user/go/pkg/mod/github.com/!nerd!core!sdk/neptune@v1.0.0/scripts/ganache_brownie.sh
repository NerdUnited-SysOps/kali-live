#!/usr/bin/env sh

ganache-cli \
	--account="0x81cc613611ad69443c10f81af20b8871e73c5c3fec154475a60535eadb29c950,100000000000000000000000000"  \
	--account="0xeedc66620ad82801e1a1da1c2b44eb17a385d83e58901330e82bf194180bab65,100000000000000000000"        \
	--account="0x36313aa3a7aa6dcf08a847e9fb2d6b7717a21b8766b4c82873bc4b31fdaab2e6,100000000000000000000"        \
	--account="0xd26cc15a624d00ccccc931e18cf98df50a36befc5319044f3326b0f4b39b2b77,100000000000000000000"        \
	--account="0xbba2c6fb324a0b47b26ec438436879b0e2702f88fb597ab7d2135a0fe38d67f1,100000000000000000000"        \
	--account="0xca3e7242c60e5bf94714cd580ffa265ef95042a719ea59026ddec6589986ee41,100000000000000000000"        \
	--account="0x364ec5b699f850695a0fedcc122e1bcc347605a6712e6831cf8e1af9601398a9,100000000000000000000"        \
	--account="0x2bc46d76ad516ed8d226550467095918a6ce25c7921d7709e32fca9932902c64,100000000000000000000"        \
	--account="0x1fa57f8c879df00058dc54556523b4b6672e7b0573c3028d8dc6577b4087a794,100000000000000000000"        \
	--account="0x18525ddf45cb75ba93c7eeea126e44800f046c41a685c152de564dc2cd5ffa45,100000000000000000000"        \
	--account="0xe8641d7d7966e3c84348c984a86d47a7b1d856a3fce7c1476fc598015c612a59,100000000000000000000"        \
	--account="0xb4d50084329b953415f6ba692fbf11394eb72fd127184effeac99b7343404eab,100000000000000000000"        \
	--account="0xdd9cbe9720dbb9cdba41cddf09e1f58998ed69999e577a31fa55a045e78ad170,100000000000000000000"        \
	--account="0x1cc317c51d105b0fd6a5a13a3e345c7f6954d6911e957a9e73d551ff6b6ad94d,100000000000000000000"        \
	--account="0x0e73e050657345c479c920906eb0b8ffee3b44aca6285c8fe276e8087461f0b2,100000000000000000000"        \
	--account="0x48d163ebc7d17edd4f54092e2e668f9dd39e6c16d27cb6addf0d008debd8f9b8,100000000000000000000"        \
    > /dev/null 2>&1 &

export GANACHE_PID=$!

status=""
printf "Starting Ganache"
for var in 0 1 2 3 4; do
    status=$(curl -X POST --write-out '%{http_code}' --silent --output /dev/null --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}' localhost:8545)
    if [ "$status" == "200" ]; then
        break
    fi
    sleep 1
    printf "."
done
printf "\n"

if [ "$status" != "200" ]; then
    kill -9 "$GANACHE_PID"
    printf "\nFailed to start Ganache in a timely manner\n"
    exit 1
fi

trap shutdown 1 2 3 6
shutdown() {
    kill -9 $GANACHE_PID
}
eval "brownie $@"
ret_code=$?
shutdown
exit "$ret_code"
