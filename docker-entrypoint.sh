#!/bin/sh

if [ -n "${MYSQL_HOST}" ]; then
    OPTS="${OPTS} -host ${MYSQL_HOST}"
fi
if [ -n "${MYSQL_PORT}" ]; then
    OPTS="${OPTS} -port ${MYSQL_PORT}"
fi
if [ -n "${ADDRESS}" ]; then
    OPTS="${OPTS} -address ${ADDRESS}"
fi
/go/src/app/main -user ${MYSQL_USER} -password ${MYSQL_PASSWORD} -database ${MYSQL_DATABASE} -query "${MYSQL_QUERY}" ${OPTS}
