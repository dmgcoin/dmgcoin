#!/bin/bash
rm -rf /tmp/dmgcoin-temp

dmgcoin --devnet --appdir=/tmp/dmgcoin-temp --profile=6061 --loglevel=debug &
DMGCOIN_PID=$!
DMGCOIN_KILLED=0
function killDmgcoinIfNotKilled() {
    if [ $DMGCOIN_KILLED -eq 0 ]; then
      kill $DMGCOIN_PID
    fi
}
trap "killDmgcoinIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $DMGCOIN_PID

wait $DMGCOIN_PID
DMGCOIN_KILLED=1
DMGCOIN_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dmgcoin exit code: $DMGCOIN_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DMGCOIN_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
