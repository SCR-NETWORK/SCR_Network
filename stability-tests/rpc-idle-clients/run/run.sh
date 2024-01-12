#!/bin/bash
rm -rf /tmp/SCR-NETWORK-temp

NUM_CLIENTS=128
SCR-NETWORK --devnet --appdir=/tmp/SCR-NETWORK-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SCR-NETWORK_PID=$!
SCR-NETWORK_KILLED=0
function killSCR-NETWORKIfNotKilled() {
  if [ $SCR-NETWORK_KILLED -eq 0 ]; then
    kill $SCR-NETWORK_PID
  fi
}
trap "killSCR-NETWORKIfNotKilled" EXIT
rm -rf /tmp/SCR_Network-temp

NUM_CLIENTS=128
SCR_Network --devnet --appdir=/tmp/SCR_Network-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SCR_Network_PID=$!
SCR_Network_KILLED=0
function killSCR_NetworkIfNotKilled() {
  if [ $SCR_Network_KILLED -eq 0 ]; then
    kill $SCR_Network_PID
  fi
}
trap "killSCR_NetworkIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $SCR-NETWORK_PID

wait $SCR-NETWORK_PID
SCR-NETWORK_EXIT_CODE=$?
SCR-NETWORK_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "SCR-NETWORK exit code: $SCR-NETWORK_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SCR-NETWORK_EXIT_CODE -eq 0 ]; then
kill $SCR_Network_PID

wait $SCR_Network_PID
SCR_Network_EXIT_CODE=$?
SCR_Network_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "SCR_Network exit code: $SCR_Network_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SCR_Network_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
