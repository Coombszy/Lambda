#!/bin/sh

docker-compose -f docker/docker-compose.yml down && docker volume prune -f 