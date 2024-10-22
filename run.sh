#!/bin/bash

set -e

SHELL_FOLDER=$(dirname "$(readlink -f "$0")")
datadir=${SHELL_FOLDER}/beacon_explorer
compose_file=${datadir}/docker-compose.yaml
config_file=beacon_explorer/config.yaml

sudo rm -rf ${datadir}
mkdir -p ${datadir}

touch ${compose_file}
cat >${compose_file} <<EOL
version: '3'
x-service: &default-service
  image: golang:1.22
  working_dir: /app
  network_mode: "host"
  volumes:
    - ${SHELL_FOLDER}:/app
    - /tmp/go-cache:/go
    - /tmp/go-build-cache:/root/.cache/go-build
services:
  build-once:
    <<: *default-service
    profiles:
      - build-once
    command: /bin/bash -c "trap stopit SIGINT; stopit() { echo 'trapped SIGINT'; exit; }; git config --global --add safe.directory '*' && make -B all"
    container_name: build-once
  indexer:
    <<: *default-service
    command: go run ./cmd/explorer -config /app/${config_file}
    container_name: indexer
    environment:
      - INDEXER_ENABLED=true
  eth1indexer:
    <<: *default-service
    command: go run ./cmd/eth1indexer -config /app/${config_file} -blocks.concurrency 1 -blocks.tracemode 'geth' -data.concurrency 1 --balances.enabled
    container_name: eth1indexer
  rewards-exporter:
    <<: *default-service
    command: go run ./cmd/rewards-exporter -config /app/${config_file}
    container_name: rewards-exporter
  statistics:
    <<: *default-service
    command: go run ./cmd/statistics -config /app/${config_file} --charts.enabled --graffiti.enabled -validators.enabled
    container_name: statistics
  frontend-data-updater:
    <<: *default-service
    command: go run ./cmd/frontend-data-updater -config /app/${config_file}
    container_name: frontend-data-updater
  frontend:
    <<: *default-service
    command: go run ./cmd/explorer -config /app/${config_file}
    container_name: frontend
    environment:
      - FRONTEND_ENABLED=true
  ratelimits-updater:
    <<: *default-service
    command: go run ./cmd/misc -config /app/${config_file} -command=update-ratelimits
    container_name: ratelimits-updater
  misc:
    <<: *default-service
    command: /bin/bash -c "trap stopit SIGINT; stopit() { echo 'trapped SIGINT'; exit; }; while true; do date; sleep 1; done"
    container_name: misc
  redis-sessions:
    image: redis:7
    container_name: redis-sessions
    volumes:
      - redis-sessions:/data
    ports:
      - "16379:6379"

volumes:
  redis-sessions:

EOL

PROJECT="explorer"
INSTANCE="explorer"
ELCONFIG_FILE=${datadir}/elconfig.json
touch ${ELCONFIG_FILE}
cat >${ELCONFIG_FILE} <<EOL
{
    "byzantiumBlock": 0,
    "constantinopleBlock": 0
}
EOL

echo 'please input postgres url'
printf '\teg: postgresql://postgres:123456@172.17.0.1:5432/postgres\n'
read postgresql
echo
postgresql=$(echo $postgresql | awk '$1=$1')
array_postgresql1=(${postgresql//:/ })
array_postgresql2=(${array_postgresql1[2]//@/ })
array_postgresql3=(${array_postgresql1[3]//\// })

postgres_user=$(echo ${array_postgresql1[1]} | sed 's/\///g')
postgres_password=${array_postgresql2[0]}
postgres_host=${array_postgresql2[1]}
postgres_port=${array_postgresql3[0]}
postgres_db=${array_postgresql3[1]}

echo 'please input web3 rpc url'
printf '\teg: http://172.17.0.1:8545\n'
read web3_rpc_url
echo
web3_rpc_url=$(echo $web3_rpc_url | awk '$1=$1')
array_web3_rpc_url=(${web3_rpc_url//:/ })
web3_rpc_host=$(echo ${array_web3_rpc_url[1]} | sed 's/\///g')
web3_rpc_port=$(echo ${array_web3_rpc_url[2]} | sed 's/\///g')

echo 'please input bigtable endpoint'
printf '\teg: 172.17.0.1:9000\n'
read bigtable_endpoint
echo
bigtable_endpoint=$(echo $bigtable_endpoint | awk '$1=$1')
HOST="${bigtable_endpoint}"
array_bigtable_endpoint=(${bigtable_endpoint//:/ })
bigtable_port=${array_bigtable_endpoint[1]}

echo 'please input redia cache endpoint'
printf '\teg: 172.17.0.1:6379\n'
read redia_cache_endpoint
echo
redia_cache_endpoint=$(echo $redia_cache_endpoint | awk '$1=$1')

echo 'please input redia session store endpoint'
printf '\teg: 172.17.0.1:16379\n'
read redia_session_store_endpoint
echo
redia_session_store_endpoint=$(echo $redia_session_store_endpoint | awk '$1=$1')

echo 'please input site domain'
printf '\teg: 127.0.0.1:8080\n'
read site_domain
echo
site_domain=$(echo $site_domain | awk '$1=$1')
array_site_domain=(${site_domain//:/ })
site_port=$(echo ${array_site_domain[1]} | sed 's/\///g')

echo 'please input site name'
printf '\teg: Open Source NBnet Explorer\n'
read site_name
echo
site_name=$(echo $site_name | awk '$1=$1')

echo 'please input bn url'
printf '\teg: 172.17.0.1:5052\n'
read bn_url
echo
bn_url=$(echo $bn_url | awk '$1=$1')
array_bn_url=(${bn_url//:/ })
bn_host=$(echo ${array_bn_url[0]} | sed 's/\///g')
bn_port=$(echo ${array_bn_url[1]} | sed 's/\///g')

touch ${SHELL_FOLDER}/${config_file}
cat >${SHELL_FOLDER}/${config_file} <<EOL
chain:
  clConfigPath: 'node'
  elConfigPath: 'beacon_explorer/elconfig.json'
readerDatabase:
  name: ${postgres_db}
  host: ${postgres_host}
  port: ${postgres_port}
  user: ${postgres_user}
  password: ${postgres_password}
writerDatabase:
  name: ${postgres_db}
  host: ${postgres_host}
  port: ${postgres_port}
  user: ${postgres_user}
  password: ${postgres_password}
bigtable:
  project: ${PROJECT}
  instance: ${INSTANCE}
  emulator: true
  emulatorPort: ${bigtable_port}
eth1ErigonEndpoint: '${web3_rpc_url}'
eth1GethEndpoint: '${web3_rpc_url}'
redisCacheEndpoint: '${redia_cache_endpoint}'
redisSessionStoreEndpoint: '${redia_session_store_endpoint}'
tieredCacheProvider: 'redis'
frontend:
  siteDomain: "${site_domain}"
  siteName: "${site_name}"
  siteSubtitle: "Showing a local testnet."
  server:
    host: '0.0.0.0' # Address to listen on
    port: '${site_port}' # Port to listen on
  readerDatabase:
    name: ${postgres_db}
    host: ${postgres_host}
    port: ${postgres_port}
    user: ${postgres_user}
    password: ${postgres_password}
  writerDatabase:
    name: ${postgres_db}
    host: ${postgres_host}
    port: ${postgres_port}
    user: ${postgres_user}
    password: ${postgres_password}
  sessionSecret: "11111111111111111111111111111111"
  jwtSigningSecret: "1111111111111111111111111111111111111111111111111111111111111111"
  jwtIssuer: "localhost"
  jwtValidityInMinutes: 30
  maxMailsPerEmailPerDay: 10
  mail:
    mailgun:
      sender: no-reply@localhost
      domain: mg.localhost
      privateKey: "key-11111111111111111111111111111111"
  csrfAuthKey: '1111111111111111111111111111111111111111111111111111111111111111'
  legal:
    termsOfServiceUrl: "tos.pdf"
    privacyPolicyUrl: "privacy.pdf"
    imprintTemplate: '{{ define "js" }}{{ end }}{{ define "css" }}{{ end }}{{ define "content" }}Imprint{{ end }}'
  stripe:
    sapphire: price_sapphire
    emerald: price_emerald
    diamond: price_diamond
  ratelimitUpdateInterval: 1s

indexer:
  #fullIndexOnStartup: false # Perform a one time full db index on startup
  #indexMissingEpochsOnStartup: true # Check for missing epochs and export them after startup
  node:
    host: ${bn_host}
    port: ${bn_port}
    type: lighthouse
  eth1DepositContractFirstBlock: 0
EOL

go run ./cmd/misc/main.go -config ${SHELL_FOLDER}/${config_file} -command initBigtableSchema
go run ./cmd/misc/main.go -config ${SHELL_FOLDER}/${config_file} -command applyDbSchema

docker-compose -f ${compose_file} --profile=build-once run -T build-once
docker-compose -f ${compose_file} up -d