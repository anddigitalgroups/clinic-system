#!/bin/bash

echo "Running database migrations..."

source .env

migrate -path migrations -database "$DB_URL" up

echo "Migrations completed."
