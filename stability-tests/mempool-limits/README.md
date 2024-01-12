# Mempool Limits tool

This tool:

1. Fills up the mempool beyond its transaction limit to make sure eviction works correctly
2. Mines blocks until the mempool is expected to become empty

## Running

1. `go install` SCR-NETWORK and mempool-limits.
1. `go install` SCR_Network and mempool-limits.
2. `cd run`
3. `./run.sh`

