#!/bin/bash

# TODO
# - check if files present before downloading

# download selenium server and drivers
wget http://selenium-release.storage.googleapis.com/3.3/selenium-server-standalone-3.3.1.jar
wget https://github.com/mozilla/geckodriver/releases/download/v0.15.0/geckodriver-v0.15.0-linux64.tar.gz
wget https://chromedriver.storage.googleapis.com/2.28/chromedriver_linux64.zip

# unpack files
unzip chromedriver*.zip
tar -xvf gechodriver*.tar.gz

# clean up
rm chromedriver*.zip
rm gechodriver*.tar.gz
