#!/bin/sh
set -e
ssm-get /nginx/config /etc/nginx/conf.d/default.conf
# For testing purposes, print the file to logs
cat /etc/nginx/conf.d/default.conf