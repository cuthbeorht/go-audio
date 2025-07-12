#!/bin/bash

my_string=$(ffmpeg -v debug -i /Users/davidsciacchettano/src/go-audio/sample.mp3 2>&1 | grep len)

# Loop through the words in the string
for word in $my_string; do

  if [[ $word == *"len"* ]]; then
    delimiter=":"
    IFS="$delimiter"

    read -ra id3_length <<< "$word"

    echo "${id3_length[1]}" >&1
  fi
done