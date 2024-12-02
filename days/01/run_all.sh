#!/bin/bash
# Builds and runs all the Go modules and their tests


# 1_compare_difference

CHALLENGE_NAME=1_compare_difference

go build -o ./bin/$CHALLENGE_NAME/ ./cmd/$CHALLENGE_NAME
./bin/${CHALLENGE_NAME}/${CHALLENGE_NAME}.exe

# 2_similarity_score

CHALLENGE_NAME=2_similarity_score

go build -o ./bin/$CHALLENGE_NAME/ ./cmd/$CHALLENGE_NAME
./bin/${CHALLENGE_NAME}/${CHALLENGE_NAME}.exe