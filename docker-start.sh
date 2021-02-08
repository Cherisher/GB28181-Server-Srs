#!/bin/sh
envsubst '$API_SERVER' < nginx.conf.tpl > /etc/nginx/nginx.conf

#envsubst '$API_SERVER' < nginx.conf.tpl > nginx.conf

env nginx

./gb28181-proxy-srs