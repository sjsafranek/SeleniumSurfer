#!/bin/sh
set -e

# Start Xvfb, Chrome, and Selenium in the background
#export DISPLAY=:10
#cd /vagrant

echo "Starting Xvfb ..."
Xvfb :10 -screen 0 1366x768x24 -ac &

echo "Starting Google Chrome ..."
google-chrome --remote-debugging-port=9222 &
