#!/bin/bash
set -e

solved_count=${solved_count:-0}
repo_url=${GITHUB_REPO_URL:-"https://github.com/orgball2608/leet-code/blob/main"}

# Hàm tạo danh sách bài giải từ thư mục
generate_section() {
  level=$1
  solved_count=$(find "$level" -type f -name "*.go" | wc -l)

  level_capitalized=$(echo "$level" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')

  echo "## $level_capitalized ($solved_count solved)" >> README.md
  echo "" >> README.md
  IFS=$'\n'
  for file in $(find "$level" -type f -name "*.go" | sort -V -t'.' -k1,1); do
    filename=$(basename "$file" .go)

    problem_name=$(echo "$filename" | cut -d'.' -f2-)
    problem_name=$(echo "$problem_name" | tr '[:upper:]' '[:lower:]')
    problem_name=${problem_name// /-}
    problem_name=${problem_name#-}

    encoded_filename=${filename// /%20}
    file_link="$repo_url/$level/$encoded_filename.go"

    leetcode_link="https://leetcode.com/problems/$problem_name/"
    echo "- [$filename]($file_link) - [LeetCode Problem]($leetcode_link)" >> README.md
  done
  echo "" >> README.md
}

# Tạo README.md
cat << EOF > README.md
# LeetCode Solutions

Practice and improve your LeetCode skills every day.

[LeetCode profile](https://leetcode.com/u/orgball2608/)

## Total Problems Solved

**$solved_count** problems solved.

EOF

for level in easy medium hard; do
  if [ -d "$level" ]; then
    generate_section "$level"
  fi
done
