#!/bin/bash

docker compose exec db psql -h localhost -U tam4 -c "INSERT INTO api_keys VALUES ('4N807YC0D9LGC0N4', 'DevMaster')"
