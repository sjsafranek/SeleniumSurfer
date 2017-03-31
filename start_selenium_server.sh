#!/bin/bash

java -jar selenium-server-standalone-3.3.1.jar -role hub &

sleep 5

java -jar selenium-server-standalone-3.3.1.jar -role node -hub http://localhost:4444/grid/register -port 6666 &

# ./phantomjs --webdriver=8080 --webdriver-selenium-grid-hub=http://localhost:4444

#sleep 5
#java -jar selenium-server-standalone-3.3.1.jar
#java -jar ios-server-0.6.5-jar-with-dependencies.jar \
#    -hub http://localhost:4444/grid/register \
#    -proxy org.uiautomation.ios.grid.IOSRemoteProxy \
#    -port 3333 &
