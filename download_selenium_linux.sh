#!/bin/bash

# TODO
# - check if files present before downloading

# download selenium server and drivers
wget http://selenium-release.storage.googleapis.com/3.3/selenium-server-standalone-3.3.1.jar
wget https://github.com/mozilla/geckodriver/releases/download/v0.15.0/geckodriver-v0.15.0-linux64.tar.gz
wget https://chromedriver.storage.googleapis.com/2.28/chromedriver_linux64.zip
wget https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2
wget http://central.maven.org/maven2/org/seleniumhq/selenium/selenium-htmlunit-driver/2.52.0/selenium-htmlunit-driver-2.52.0.jar
wget https://github.com/ios-driver/ios-driver/releases/download/0.6.5/ios-server-0.6.5-jar-with-dependencies.jar
#wget http://ios-driver-ci.ebaystratus.com/userContent/ios-server-standalone-0.6.6-SNAPSHOT.jar

# unpack files
unzip chromedriver*.zip
tar -xvf geckodriver*.tar.gz
tar -xvf phantomjs-2.1.1-linux-x86_64.tar.bz2

cp phantomjs-2.1.1-linux-x86_64/bin/phantomjs .

# clean up
rm chromedriver*.zip
rm geckodriver*.tar.gz
rm phantomjs*.tar.bz2
rm -rf phantomjs-2.1.1-linux-x86_64
