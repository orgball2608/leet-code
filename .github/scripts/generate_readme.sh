#!/bin/bash
set -e

# Get the absolute path of the repository root
REPO_ROOT=$(git rev-parse --show-toplevel)

TEMPLATE_FILE="$REPO_ROOT/README.md.template"
README_FILE="$REPO_ROOT/README.md"

if [ ! -f "$TEMPLATE_FILE" ]; then
    echo "Template file not found at: $TEMPLATE_FILE"
    exit 1
fi

cp "$TEMPLATE_FILE" "$README_FILE"
sed -i "s|{{GITHUB_REPOSITORY}}|$GITHUB_REPOSITORY|g" "$README_FILE"

# Function to generate the problem list for a given difficulty
generate_problem_list() {
    local difficulty=$1
    find "$REPO_ROOT/$difficulty" -type f -name "*.go" | sort -V | while IFS= read -r file; do
        local problem_title
        problem_title=$(basename "$file" .go)
        
        # Extract the problem name for the LeetCode URL, e.g., "two-sum" from "1. Two Sum"
        local problem_slug
        problem_slug=$(echo "$problem_title" | cut -d'.' -f2- | tr '[:upper:]' '[:lower:]' | sed 's/ /-/g' | sed 's/^-//')

        local leetcode_link="[Problem](https://leetcode.com/problems/$problem_slug/)"
        
        local relative_file
        relative_file=${file#"$REPO_ROOT/"}
        
        local encoded_file
        encoded_file=$(printf '%s' "$relative_file" | jq -s -R -r @uri)
        
        local solution_link="[Solution]($GITHUB_SERVER_URL/$GITHUB_REPOSITORY/blob/main/$encoded_file)"
        
        echo "| $problem_title | $leetcode_link | $solution_link |"
    done
}

EASY_SOLVED=$(find "$REPO_ROOT/easy" -type f -name "*.go" | wc -l)
MEDIUM_SOLVED=$(find "$REPO_ROOT/medium" -type f -name "*.go" | wc -l)
HARD_SOLVED=$(find "$REPO_ROOT/hard" -type f -name "*.go" | wc -l)
ALL_SOLVED=$((EASY_SOLVED + MEDIUM_SOLVED + HARD_SOLVED))

sed -i "s/{{EASY_SOLVED}}/$EASY_SOLVED/g" "$README_FILE"
sed -i "s/{{MEDIUM_SOLVED}}/$MEDIUM_SOLVED/g" "$README_FILE"
sed -i "s/{{HARD_SOLVED}}/$HARD_SOLVED/g" "$README_FILE"
sed -i "s/{{ALL_SOLVED}}/$ALL_SOLVED/g" "$README_FILE"

generate_problem_list "easy" > easy_list.tmp
generate_problem_list "medium" > medium_list.tmp
generate_problem_list "hard" > hard_list.tmp

sed -i -e '/{{EASY_LIST}}/r easy_list.tmp' -e '/{{EASY_LIST}}/d' "$README_FILE"
sed -i -e '/{{MEDIUM_LIST}}/r medium_list.tmp' -e '/{{MEDIUM_LIST}}/d' "$README_FILE"
sed -i -e '/{{HARD_LIST}}/r hard_list.tmp' -e '/{{HARD_LIST}}/d' "$README_FILE"

rm easy_list.tmp medium_list.tmp hard_list.tmp

echo "README.md generated successfully."
