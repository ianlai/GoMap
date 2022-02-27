#!/bin/sh
export DBHOST=`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' gomap_db`
echo Set '$DBHOST' env variable to "$DBHOST"