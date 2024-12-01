#!/bin/bash
# Builds and runs all the Go modules and their tests


# 1_compare_lists

CHALLENGE_NAME=1_compare_lists

go build -o ./bin/$CHALLENGE_NAME/ ./cmd/$CHALLENGE_NAME
./bin/${CHALLENGE_NAME}/${CHALLENGE_NAME}.exe
