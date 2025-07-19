#!/bin/bash
set -e

# Get the absolute path of the repository root
REPO_ROOT=$(git rev-parse --show-toplevel)

# First, try to fetch the solved count from the LeetCode API
# The LEETCODE_USERNAME should be set as a secret in the GitHub repository settings
solved_count=$(curl -s "https://leetcode-stats-api.herokuapp.com/${LEETCODE_USERNAME}" | jq '.totalSolved // 0')

# If the API fails or returns 0, fall back to counting files locally
if [ -z "$solved_count" ] || [ "$solved_count" -eq "0" ]; then
  echo "API call failed or returned 0. Counting files locally as a fallback."
  # Count .go files in the difficulty directories relative to the repo root
  solved_count=$(find "$REPO_ROOT/easy" "$REPO_ROOT/medium" "$REPO_ROOT/hard" -type f -name "*.go" | wc -l)
fi

# Export the final count to the GitHub Actions environment
echo "solved_count=$solved_count" >> "$GITHUB_ENV"
echo "Final solved count: $solved_count"
