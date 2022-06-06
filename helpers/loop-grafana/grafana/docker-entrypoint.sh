#! /bin/sh

while (! docker stats --no-stream 1> /dev/null 2>&1 ); do
  sleep 1
done

docker-compose up
