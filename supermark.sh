#!/bin/bash

source supermark.cfg

if [ ! -f supermark ]; then
  echo "Building executable from Go source..."
  go build supermark.go
fi

echo "Generating HTML page..."
./supermark -bookmarks "$chrome_bookmarks_path" -output $output_file_name -root "$root_folder"

if [ ! -z "$scp_destination" ]; then
  echo "Uploading..."
  scp $output_file_name $scp_destination
else
  echo "Upload disabled."
fi

echo "Done!"