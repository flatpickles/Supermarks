# Supermarks

The goal of this project is to allow you to easily publish the contents of a subfolder within your chrome bookmarks. The `supermarks.go` program parses Chrome's bookmark JSON format into a templated HTML page, and I've written a couple shell scripts to assist with watching Chrome's bookmark file for changes, and uploading the HTML file to your destination of choice after generation.

To use this:
- Checkout this repo, if you haven't already.
- Edit `supermarks.cfg` to fit your needs:
  - `chrome_bookmarks_path` is the location of your Chrome bookmarks (default is probably correct)
  - `output_file_name` is what you want your HTML output to be named
  - `root_folder` is the string name of the subfolder you want to upload
  - `scp_destination` may be left undefined, but set this if you want to configure automatic upload
- If you want automatic upload to work, we'll need to be able to SCP to your server. Make sure your SSH keys are set up appropriately. Here are some [helpful tips](https://www.digitalocean.com/community/tutorials/how-to-set-up-ssh-keys--2).
- Run `supermarks.sh`! I've compiled `supermarks.go` and included the build in this repo, but if that doesn't work on your machine, you'll need to build your own. To do this, install Go, delete my `supermarks` build, and try running `supermarks.sh` again.
- If you don't need to enter a password for your SSH config, we can automatically upload new files in the background whenever your Chrome bookmarks file is updated. Use the `watch.sh` script to this end, with `nohup` if you want it to run indefinitely as a background task. This script uses [fswatch](https://github.com/emcrisostomo/fswatch), which you'll need to install first if you want to set this up.

This project is very much a work in progress! The generated bookmarks file looks like garbage for the time being, but before too long I'll be able to get my hands dirty with some stylish CSS.

Check out this [example output](http://man1.biz/supermarks/) on my site – my bookmark collection is also in progress. Stay tuned!