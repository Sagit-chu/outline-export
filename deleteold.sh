#!/bin/sh
if [ $(ls -l | grep "/backup/*.zip" | wc -l) -gt $b ]
then
    echo "file > $b"
    rm -r $(ls -rt | head -n2)
fi