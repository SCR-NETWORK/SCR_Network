#!/bin/bash
SCR-NETWORKsanity --command-list-file ./commands-list --profile=7000
SCR_Networksanity --command-list-file ./commands-list --profile=7000
TEST_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ]; then
  echo "SCR-NETWORKsanity test: PASSED"
  exit 0
fi
echo "SCR-NETWORKsanity test: FAILED"
  echo "SCR_Networksanity test: PASSED"
  exit 0
fi
echo "SCR_Networksanity test: FAILED"
exit 1
