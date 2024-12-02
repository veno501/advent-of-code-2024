#!/bin/bash
# Builds and runs all the Go modules and their tests

# run tests

go test ./...


# build and run 1_compare_difference

CHALLENGE_NAME=1_compare_difference

go build -o ./bin/$CHALLENGE_NAME/ ./cmd/$CHALLENGE_NAME
./bin/${CHALLENGE_NAME}/${CHALLENGE_NAME}.exe


# build and run 2_similarity_score

CHALLENGE_NAME=2_similarity_score

go build -o ./bin/$CHALLENGE_NAME/ ./cmd/$CHALLENGE_NAME
./bin/${CHALLENGE_NAME}/${CHALLENGE_NAME}.exe