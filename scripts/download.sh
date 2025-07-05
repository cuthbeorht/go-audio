#!/bin/bash

if [ ! -f ./tmp/media.zip ]; then
    curl --output ./tmp/media.zip https://s3.amazonaws.com/citizen-dj-assets.labs.loc.gov/samplepacks/loc.gov_fma_mp3.zip

    echo "Unzipping media"
    unzip ./tmp/media.zip -d ./tmp
fi