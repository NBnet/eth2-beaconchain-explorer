#!/bin/bash

set -e

SHELL_FOLDER=$(dirname "$(readlink -f "$0")")
datadir=${SHELL_FOLDER}/database
compose_file=${datadir}/docker-compose.yaml

sudo rm -rf ${datadir}
mkdir -p ${datadir}

touch ${compose_file}
cat >${compose_file} <<EOL
version: '3'
services:
  postgres:
    image: postgres:15
    container_name: postgres
    command: postgres -c 'max_connections=5000'
    shm_size: 128mb
    environment:
      POSTGRES_PASSWORD: cwbqhqnryd2416
    volumes:
      - ${datadir}/postgres/data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  redis:
    image: redis:7
    container_name: redis
    restart: always
    ports:
      - '6379:6379'
    volumes:
      - ${datadir}/redis/data:/data
      #- ${datadir}/redis/redis.conf:/usr/local/etc/redis/redis.conf
      - ${datadir}/redis/logs:/logs
    #command: redis-server /usr/local/etc/redis/redis.conf

  little_bigtable:
    image: gobitfly/little_bigtable
    container_name: lbt
    restart: always
    ports:
      - '19000:9000'
EOL

docker-compose -f ${compose_file} up -d