#!/bin/bash
set -e

solved_count=$(curl -s "https://leetcode-stats-api.herokuapp.com/orgball2608" | jq '.totalSolved // 0')

if [ "$solved_count" -eq "0" ]; then
  solved_count=$(find easy medium hard -type f -name "*.go" | wc -l)
fi

echo "solved_count=$solved_count" >> "$GITHUB_ENV"
echo "Solved count: $solved_count"
