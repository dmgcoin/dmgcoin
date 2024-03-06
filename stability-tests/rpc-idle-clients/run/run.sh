#!/bin/bash
rm -rf /tmp/dmgcoin-temp

NUM_CLIENTS=128
dmgcoin --devnet --appdir=/tmp/dmgcoin-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
DMGCOIN_PID=$!
DMGCOIN_KILLED=0
function killDmgcoinIfNotKilled() {
  if [ $DMGCOIN_KILLED -eq 0 ]; then
    kill $DMGCOIN_PID
  fi
}
trap "killDmgcoinIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $DMGCOIN_PID

wait $DMGCOIN_PID
DMGCOIN_EXIT_CODE=$?
DMGCOIN_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Dmgcoin exit code: $DMGCOIN_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DMGCOIN_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
