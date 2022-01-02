#!/bin/sh
crond # starts the cron daemon in the background so we can run ingest every few hours
echo "Waiting 1 minute for DB initialization"
sleep 1m
/serverapi ingest # performs an ingest
/serverapi serve # Runs the server