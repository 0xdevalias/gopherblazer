#!/bin/bash

DICT="./functions/libgobuster/words.txt"
SAMPLE=3
BUCKET="gopherblazer-test-discovery"

rm -rf ./discoveryFiles/
mkdir -p ./discoveryFiles/

./samplen.py "$SAMPLE" "$DICT" | (while read line; do
  line=${line%$'\r'}
  # echo $line
  echo "This is a test file for discovery" > ./discoveryFiles/$line
  # touch ./discoveryFiles/$line
done)

aws s3 sync ./discoveryFiles/ s3://$BUCKET --exclude "*.rsl*" --delete

echo -e "\nThe following files exist for discovery:"
aws s3 ls s3://$BUCKET | awk '{print $4}'

rm -rf ./discoveryFiles/
