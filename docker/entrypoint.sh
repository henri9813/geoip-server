#!/usr/bin/env bash
set -e

getent group "${UNIX_GROUP}" || \
 groupadd "${UNIX_GROUP}" \
    --gid "${UNIX_GID}"

id --user "${UNIX_USER}" || \
 useradd "${UNIX_USERNAME}" \
    --uid "${UNIX_UID}" \
    --gid "${UNIX_GROUP}" \
    --create-home \
    --home-dir "/home/${UNIX_USERNAME}"

exec "$@"
