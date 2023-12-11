#!/bin/bash
# post-create.sh

# Run a SQL script in the db service
docker-compose exec db psql -U postgres -f /docker-entrypoint-initdb.d/init.sql

# Rest of your script...