#!/bin/bash

DEFAULT_DICT="./functions/libgobuster/words.txt"
DEFAULT_SAMPLE=3
DEFAULT_BUCKET="gopherblazer-test-discovery"

DICT=${1:-$DEFAULT_DICT}
SAMPLE=${2:-$DEFAULT_SAMPLE}
BUCKET=${3:-$DEFAULT_BUCKET}

rm -rf ./discoveryFiles/
mkdir -p ./discoveryFiles/

./samplen.py "$SAMPLE" "$DICT" | (while read line; do
  line=${line%$'\r'}
  # echo $line
  echo "This is a test file for discovery" > ./discoveryFiles/$line
  # touch ./discoveryFiles/$line
done)

aws s3 sync ./discoveryFiles/ s3://$BUCKET --exclude "*.rsl*" --delete

# echo -e "\nThe following files exist for discovery:"
# aws s3 ls s3://$BUCKET | awk '{print $4}'

echo -e "\n$SAMPLE files created for discovery."

rm -rf ./discoveryFiles/
