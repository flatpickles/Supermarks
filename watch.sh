#!/bin/bash

source supermarks.cfg

if ! type "fswatch" > /dev/null; then
  echo "Please install fswatch: https://github.com/emcrisostomo/fswatch"
fi

echo "Watching for changes in $chrome_bookmarks_path"
fswatch "$chrome_bookmarks_path" | (while read; do ./supermarks.sh; done)