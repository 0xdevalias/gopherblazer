#!/bin/bash

DEFAULT_SLICES=2
DEFAULT_URL="http://test-discovery.gopherblazer.devalias.net"
DEFAULT_DICT="words.txt"
DEFAULT_THREADS=10

SLICE_CNT=${1:-$DEFAULT_SLICES}
URL=${2:-$DEFAULT_URL}
DICT=${3:-$DEFAULT_DICT}
THREADS=${4:-$DEFAULT_THREADS}

function jsonTemplate {
  echo "{\"url\": \"$URL\", \"wordlist\":\"$DICT\", \"threads\": $THREADS, \"sliceStart\": $1, \"sliceEnd\": $2}"
}

function invokeApexSlice {
  jsonTemplate $1 $2 | apex invoke libgobuster | jq &
}

function invokeMulti {
  # Count words
  DICT_SIZE=$(wc -l ./functions/libgobuster/$DICT | awk '{print $1}')
  SLICE_SIZE=$(($DICT_SIZE / $SLICE_CNT))
  SLICE_EXCESS=$(($DICT_SIZE % $SLICE_CNT))

  echo "$DICT contains $DICT_SIZE lines."
  echo "  $SLICE_CNT slices will be $SLICE_SIZE lines each (evenly) with $SLICE_EXCESS leftover"
  echo -e "  (Remember 0 index counts as a line..)\n"

  TOTAL_SLICES=0
  while [ $TOTAL_SLICES -lt $SLICE_CNT ]; do
    if [ $TOTAL_SLICES -eq 0 ]; then
      SLICE_FROM=0
      SLICE_TO=$(($SLICE_FROM+$SLICE_SIZE+$SLICE_EXCESS-1))
    else
      SLICE_FROM=$((SLICE_TO+1))
      SLICE_TO=$(($SLICE_FROM+$SLICE_SIZE-1))
    fi

    if [ $SLICE_FROM -lt $DICT_SIZE ]; then
      echo "Slice #$((TOTAL_SLICES+1)): $SLICE_FROM to $SLICE_TO ($(($SLICE_TO-$SLICE_FROM+1)) lines) "
      invokeApexSlice $SLICE_FROM $SLICE_TO
    else
      echo "Slice #$((TOTAL_SLICES+1)) : Skipping slice, outside dictionary range" #$SLICE_FROM to $SLICE_TO ($(($SLICE_TO-$SLICE_FROM+1)) lines)"
    fi

    let TOTAL_SLICES=$(($TOTAL_SLICES+1))
  done

  wait
}

time invokeMulti && echo "This timing is for the whole job"
