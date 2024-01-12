#!/bin/sh

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until nc -z -w 1 "$host" "$port"; do
  echo "Waiting for MySQL to be ready..."
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd
