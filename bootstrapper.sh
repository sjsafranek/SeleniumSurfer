#!/bin/bash

sudo aptitude -yy install openjdk-8-jre-headless 
sudo aptitude -yy install openjdk-8-jdk
sudo /usr/sbin/update-java-alternatives -s java-1.8.0-openjdk-amd64
