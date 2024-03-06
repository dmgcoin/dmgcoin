#!/bin/bash
rm -rf /tmp/dmgcoin-temp

dmgcoin --devnet --appdir=/tmp/dmgcoin-temp --profile=6061 --loglevel=debug &
DMGCOIN_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $DMGCOIN_PID

wait $DMGCOIN_PID
DMGCOIN_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dmgcoin exit code: $DMGCOIN_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DMGCOIN_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
