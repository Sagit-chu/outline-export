#!/bin/sh
if [ $(ls -l /backup/*.zip | wc -l) -gt $b ]
then
    rm -r $(ls -rt /backup/*.zip | head -n1)
    echo "delete old backup success"
fi