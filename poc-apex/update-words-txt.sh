#!/bin/bash

WORDS_GOBUSTER="./functions/gobuster/words.txt"
WORDS_LIBGOBUSTER="./functions/libgobuster/words.txt"
DEFAULT_SAMPLE=10
DEFAULT_BUCKET="gopherblazer-test-discovery"

aws s3 ls s3://$DEFAULT_BUCKET | awk '{print $4}' | head -n $DEFAULT_SAMPLE > "$WORDS_GOBUSTER"
aws s3 ls s3://$DEFAULT_BUCKET | awk '{print $4}' | head -n $DEFAULT_SAMPLE > "$WORDS_LIBGOBUSTER"
