#!/usr/bin/env bash

node --version
npm --version
npm install -g prettier

ls -l -a

PRETTIER_CHECK_OUTPUT="$(prettier --config /default-prettier-config.json --check .)"
PRETTIER_CHECK_EXIT_CODE=$?

echo 'check_output<<EOF' >> $GITHUB_OUTPUT
echo "$PRETTIER_CHECK_OUTPUT" >> $GITHUB_OUTPUT
echo 'EOF' >> $GITHUB_OUTPUT
echo "check_exit_code=$PRETTIER_CHECK_EXIT_CODE" >> $GITHUB_OUTPUT