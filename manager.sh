#!/usr/bin/env bash

echo 'msg_rabbit'
# локальный запуск
cd "/Users/anya/IdeaProjects/messages"
docker-compose up --force-recreate msg_rabbit_local

echo 'msg_postgres'
# локальный запуск
cd "/Users/anya/IdeaProjects/messages"
docker-compose up --force-recreate msg_postgres_local

echo "acceptor"
# сборка бинарника acceptor
cd '/Users/anya/IdeaProjects/messages/acceptor'
GOOS=linux GOARCH=amd64 go build -o deploy/app_acceptor
# создать образ
docker build -f deploy/Dockerfile . -t siannarom/msg_acceptor
rm deploy/app_acceptor
# локальный запуск
cd "/Users/anya/IdeaProjects/messages"
docker-compose up --force-recreate msg_acceptor


echo 'msg_worker'
# сборка бинарника msg_worker
cd '/Users/anya/IdeaProjects/messages/worker'
GOOS=linux GOARCH=amd64 go build -o deploy/app_msg_worker
# создать образ
docker build -f deploy/Dockerfile . -t siannarom/msg_worker
rm deploy/app_msg_worker
# локальный запуск
cd "/Users/anya/IdeaProjects/messages"
docker-compose up --force-recreate msg_worker_local