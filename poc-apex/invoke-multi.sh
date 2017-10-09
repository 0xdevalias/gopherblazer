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
  time jsonTemplate $1 $2 | apex invoke libgobuster | jq && echo "This timing is for slice $1-$2" && echo &
}

function invokeMulti {
  invokeApexSlice 1 5
  invokeApexSlice 6 10
  invokeApexSlice 11 17

  wait
}

# Count words
echo "$DICT contains $(wc -l ./functions/libgobuster/$DICT | awk '{print $1}') lines."

time invokeMulti && echo "This timing is for the whole job"
