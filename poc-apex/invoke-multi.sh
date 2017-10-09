#!/bin/bash

DEFAULT_URL="http://gopherblazer-test-discovery.s3-ap-southeast-2.amazonaws.com/"
DEFAULT_DICT="words.txt"
DEFAULT_THREADS=10

URL=${1:-$DEFAULT_URL}
DICT=${2:-$DEFAULT_DICT}
THREADS=${3:-$DEFAULT_THREADS}

function jsonTemplate {
  echo "{\"url\": \"$URL\", \"wordlist\":\"$DICT\", \"threads\": $THREADS, \"sliceStart\": $1, \"sliceEnd\": $2}"
}

function invokeApexSlice {
  jsonTemplate $1 $2 | apex invoke libgobuster | jq &
}

function invokeMulti {
  # Example for words.txt
  invokeApexSlice 1 5
  invokeApexSlice 6 10
  invokeApexSlice 11 17

  # Example for big.txt
  # invokeApexSlice 1     2000
  # invokeApexSlice 2001  4000
  # invokeApexSlice 4001  6000
  # invokeApexSlice 6001  8000
  # invokeApexSlice 8001  10000
  # invokeApexSlice 10001 12000
  # invokeApexSlice 12001 14000
  # invokeApexSlice 14001 16000
  # invokeApexSlice 16001 18000
  # invokeApexSlice 18001 20000
  # invokeApexSlice 20001 20469

  # Example 2 for big.txt
  # invokeApexSlice 1     1000
  # invokeApexSlice 1001  2000
  # invokeApexSlice 2001  3000
  # invokeApexSlice 3001  4000
  # invokeApexSlice 4001  5000
  # invokeApexSlice 5001  6000
  # invokeApexSlice 6001  7000
  # invokeApexSlice 7001  8000
  # invokeApexSlice 8001  9000
  # invokeApexSlice 9001  10000
  # invokeApexSlice 10001 11000
  # invokeApexSlice 11001 12000
  # invokeApexSlice 12001 13000
  # invokeApexSlice 13001 14000
  # invokeApexSlice 14001 15000
  # invokeApexSlice 15001 16000
  # invokeApexSlice 16001 17000
  # invokeApexSlice 17001 18000
  # invokeApexSlice 18001 19000
  # invokeApexSlice 19001 20000
  # invokeApexSlice 20001 20469

  wait
}

# Count words
echo "$DICT contains $(wc -l ./functions/libgobuster/$DICT | awk '{print $1}') lines."

time invokeMulti && echo "This timing is for the whole job"
