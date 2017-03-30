#!/bin/bash

java -jar selenium-server-standalone-3.3.1.jar -role hub &
sleep 5
java -jar selenium-server-standalone-3.3.1.jar -role node -hub http://localhost:4444/grid/register -port 6666 &



# java -jar selenium-server-standalone-3.3.1.jar -role hub
# phantomjs --webdriver=8080 --webdriver-selenium-grid-hub=http://127.0.0.1:4444
