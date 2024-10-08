#!/bin/sh
set -e

STASH_NAME=pre-commit-$(date +%s)

echo "Stashing unstaged changes"
git stash save --quiet --keep-index --include-untracked $STASH_NAME

function restoreStashedIfAny {
  STASH_NUM=$(git stash list | grep $STASH_NAME | sed -re 's/stash@\{(.*)\}.*/\1/')
  if [ -n "$STASH_NUM" ]; then
    echo "Restoring stashed changes"
    git stash pop --quiet stash@{$STASH_NUM}
  fi
}

trap restoreStashedIfAny EXIT
trap restoreStashedIfAny ERR

go mod tidy

make test

make prettier
