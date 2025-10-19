#!/bin/bash

# Script untuk create database di PostgreSQL external
# Run sekali sebelum docker-compose up

echo "=========================================="
echo "Creating Database on External PostgreSQL"
echo "=========================================="

# Load environment variables from .env
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
else
    echo "Error: .env file not found!"
    exit 1
fi

echo "Host: $DB_HOST"
echo "Port: $DB_PORT"
echo "User: $DB_USER"
echo "Database: $DB_NAME"
echo ""

# Create database
echo "Creating database '$DB_NAME'..."

PGPASSWORD=$DB_PASSWORD psql \
    -h $DB_HOST \
    -p $DB_PORT \
    -U $DB_USER \
    -d postgres \
    -c "CREATE DATABASE $DB_NAME;"

if [ $? -eq 0 ]; then
    echo "✓ Database '$DB_NAME' created successfully!"
else
    echo "Note: Database might already exist or there was an error."
    echo "Checking if database exists..."
    
    PGPASSWORD=$DB_PASSWORD psql \
        -h $DB_HOST \
        -p $DB_PORT \
        -U $DB_USER \
        -d postgres \
        -c "SELECT 1 FROM pg_database WHERE datname='$DB_NAME';" | grep -q 1
    
    if [ $? -eq 0 ]; then
        echo "✓ Database '$DB_NAME' already exists!"
    else
        echo "✗ Failed to create database. Please create manually:"
        echo ""
        echo "PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c \"CREATE DATABASE $DB_NAME;\""
        exit 1
    fi
fi

echo ""
echo "=========================================="
echo "Database setup complete!"
echo "Now you can run: docker-compose up -d"
echo "=========================================="
