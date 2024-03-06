#!/bin/bash

APPDIR=/tmp/dmgcoin-temp
DMGCOIN_RPC_PORT=29587

rm -rf "${APPDIR}"

dmgcoin --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${DMGCOIN_RPC_PORT}" --profile=6061 &
DMGCOIN_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${DMGCOIN_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $DMGCOIN_PID

wait $DMGCOIN_PID
DMGCOIN_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dmgcoin exit code: $DMGCOIN_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DMGCOIN_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
