#!/bin/bash

APPDIR=/tmp/SCR-NETWORK-temp
SCR-NETWORK_RPC_PORT=29587

rm -rf "${APPDIR}"

SCR-NETWORK --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${SCR-NETWORK_RPC_PORT}" --profile=6061 &
SCR-NETWORK_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${SCR-NETWORK_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $SCR-NETWORK_PID

wait $SCR-NETWORK_PID
SCR-NETWORK_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "SCR-NETWORK exit code: $SCR-NETWORK_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SCR-NETWORK_EXIT_CODE -eq 0 ]; then
APPDIR=/tmp/SCR_Network-temp
SCR_Network_RPC_PORT=29587

rm -rf "${APPDIR}"

SCR_Network --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${SCR_Network_RPC_PORT}" --profile=6061 &
SCR_Network_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${SCR_Network_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $SCR_Network_PID

wait $SCR_Network_PID
SCR_Network_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "SCR_Network exit code: $SCR_Network_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SCR_Network_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
