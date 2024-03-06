#!/bin/bash
rm -rf /tmp/dmgcoin-temp

dmgcoin --devnet --appdir=/tmp/dmgcoin-temp --profile=6061 &
DMGCOIN_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $DMGCOIN_PID

wait $DMGCOIN_PID
DMGCOIN_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dmgcoin exit code: $DMGCOIN_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DMGCOIN_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
