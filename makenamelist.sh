#!/bin/sh
unzip -l /usr/local/go/lib/time/zoneinfo.zip| sed -n -e '1,3d' -e '/^[ ][ ]*[0-9]/s,[^A-Za-z/]*,,' -e '/\/$/d' -e 'p'
