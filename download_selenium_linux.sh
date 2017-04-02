#!/bin/bash

# Get selenium server
if [ ! -f "selenium-server-standalone-3.3.1.jar" ]; then
    wget http://selenium-release.storage.googleapis.com/3.3/selenium-server-standalone-3.3.1.jar
fi

# Get browser drivers
if [ ! -f "chromedriver" ]; then
    wget https://chromedriver.storage.googleapis.com/2.28/chromedriver_linux64.zip
    unzip chromedriver*.zip
    rm chromedriver*.zip
fi

if [ ! -f "geckodriver" ]; then
    wget https://github.com/mozilla/geckodriver/releases/download/v0.15.0/geckodriver-v0.15.0-linux64.tar.gz
    tar -xvf geckodriver*.tar.gz
    rm geckodriver*.tar.gz
fi

# download selenium server and drivers
#wget http://selenium-release.storage.googleapis.com/3.3/selenium-server-standalone-3.3.1.jar
#wget https://github.com/mozilla/geckodriver/releases/download/v0.15.0/geckodriver-v0.15.0-linux64.tar.gz
#wget https://chromedriver.storage.googleapis.com/2.28/chromedriver_linux64.zip
#wget https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2
#wget http://central.maven.org/maven2/org/seleniumhq/selenium/selenium-htmlunit-driver/2.52.0/selenium-htmlunit-driver-2.52.0.jar

# unpack files
#unzip chromedriver*.zip
#tar -xvf geckodriver*.tar.gz
#tar -xvf phantomjs-2.1.1-linux-x86_64.tar.bz2

#cp phantomjs-2.1.1-linux-x86_64/bin/phantomjs .

# clean up
# rm chromedriver*.zip
# rm geckodriver*.tar.gz
#rm phantomjs*.tar.bz2
#rm -rf phantomjs-2.1.1-linux-x86_64
