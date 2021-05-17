#!/bin/bash
# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Verifies git commits starting from predefined prefix.

# default to checking most recent commit only if not run by CI pipeline
check_base="${PULL_BASE_SHA:-HEAD^}"
check_sha="${PULL_PULL_SHA:-HEAD}"

read -d '' help_message << EOF
commit messages should look like one of:
UPSTREAM: <carry>: message  (commits that should be carried indefinitely)
UPSTREAM: <drop>: message   (commits that should be dropped on the next upstream rebase)
UPSTREAM: 1234: message     (commits that should be carried until an upstream rebase includes upstream PR 1234)
EOF

prefix='UPSTREAM: ([0-9]+|<(carry|drop)>): '

echo "examining commits between $check_base and $check_sha"
echo

while read -r message; do
  if ! [[ "$message" =~ ^$prefix ]]; then
    echo "Git history in this PR doesn't conform to set commit message standards. Offending commit message is:"
    echo "$message"
    echo
    echo "$help_message"
    exit 1
  fi
  echo "$message"
done < <(git log "$check_base".."$check_sha" --pretty=%s --no-merges --ancestry-path)

if [[ "$CI" == "true" ]]; then
    merge_base=$(git merge-base "$PULL_BASE_SHA" "$PULL_PULL_SHA")
    echo "Comparing the base branch (${PULL_BASE_REF} - commit: ${PULL_BASE_SHA}) and the merge base (${merge_base})."
    if [[ "$PULL_BASE_SHA" != "$merge_base" ]]; then
        echo "ERROR: This pull request is not based on the latest commit of the base branch."
        echo "To resolve this issue, please rebase your pull request on top of the latest base branch commit."
        echo
        echo "Additional commits on the base branch (${PULL_BASE_REF}) that are not included in this pull request:"
        git log --oneline "$PULL_BASE_SHA" ^"$PULL_PULL_SHA"
        exit 1
    fi
    echo "SUCCESS: This pull request includes all commits from the base branch."
fi

echo
echo "All looks good"
