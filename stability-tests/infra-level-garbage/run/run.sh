#!/bin/bash
rm -rf /tmp/SCR-NETWORK-temp

SCR-NETWORK --devnet --appdir=/tmp/SCR-NETWORK-temp --profile=6061 &
SCR-NETWORK_PID=$!
rm -rf /tmp/SCR_Network-temp

SCR_Network --devnet --appdir=/tmp/SCR_Network-temp --profile=6061 &
SCR_Network_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SCR-NETWORK_PID

wait $SCR-NETWORK_PID
SCR-NETWORK_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "SCR-NETWORK exit code: $SCR-NETWORK_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SCR-NETWORK_EXIT_CODE -eq 0 ]; then
kill $SCR_Network_PID

wait $SCR_Network_PID
SCR_Network_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "SCR_Network exit code: $SCR_Network_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SCR_Network_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
