#!/bin/bash

echo "Rolling back last migration..."

source .env

migrate -path migrations -database "$DB_URL" down 1

echo "Rollback completed."
