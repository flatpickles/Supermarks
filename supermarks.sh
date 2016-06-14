#!/bin/bash

source supermarks.cfg

if [ ! -f supermarks ]; then
  echo "Building executable from Go source..."
  go build supermarks.go
fi

echo "Generating HTML page..."
./supermarks -bookmarks "$chrome_bookmarks_path" -output $output_file_name -root "$root_folder"

if [ ! -z "$scp_destination" ]; then
  echo "Uploading..."
  scp $output_file_name $scp_destination
  scp -r resources $scp_destination
else
  echo "Upload disabled."
fi

echo "Done!"