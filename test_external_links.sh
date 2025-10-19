#!/bin/bash

# Test script untuk External Links API
# Pastikan server sudah berjalan di http://localhost:3000

BASE_URL="http://localhost:3000/api/v1"

echo "=========================================="
echo "Testing External Links API"
echo "=========================================="

# 1. Buat project baru
echo -e "\n1. Creating new project..."
PROJECT_RESPONSE=$(curl -s -X POST $BASE_URL/projects \
  -H "Content-Type: application/json" \
  -d '{
    "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "title": "My Awesome Game",
    "description": "A great blockchain game",
    "developer_name": "GameDev Studio",
    "genre": "RPG",
    "game_type": "web3"
  }')

PROJECT_ID=$(echo $PROJECT_RESPONSE | jq -r '.id')
echo "Project created with ID: $PROJECT_ID"

# 2. Tambahkan external link - Instagram
echo -e "\n2. Adding Instagram link..."
INSTAGRAM_RESPONSE=$(curl -s -X POST $BASE_URL/projects/$PROJECT_ID/links \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Instagram",
    "url": "https://instagram.com/mygame"
  }')

INSTAGRAM_ID=$(echo $INSTAGRAM_RESPONSE | jq -r '.id')
echo "Instagram link added with ID: $INSTAGRAM_ID"
echo $INSTAGRAM_RESPONSE | jq .

# 3. Tambahkan external link - Twitter
echo -e "\n3. Adding Twitter link..."
TWITTER_RESPONSE=$(curl -s -X POST $BASE_URL/projects/$PROJECT_ID/links \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Twitter",
    "url": "https://twitter.com/mygame"
  }')

TWITTER_ID=$(echo $TWITTER_RESPONSE | jq -r '.id')
echo "Twitter link added with ID: $TWITTER_ID"
echo $TWITTER_RESPONSE | jq .

# 4. Tambahkan external link - Website
echo -e "\n4. Adding Website link..."
WEBSITE_RESPONSE=$(curl -s -X POST $BASE_URL/projects/$PROJECT_ID/links \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Website",
    "url": "https://mygame.com"
  }')

WEBSITE_ID=$(echo $WEBSITE_RESPONSE | jq -r '.id')
echo "Website link added with ID: $WEBSITE_ID"
echo $WEBSITE_RESPONSE | jq .

# 5. Get semua external links
echo -e "\n5. Getting all external links for project..."
ALL_LINKS=$(curl -s $BASE_URL/projects/$PROJECT_ID/links)
echo $ALL_LINKS | jq .

# 6. Update Instagram link
echo -e "\n6. Updating Instagram link..."
UPDATE_RESPONSE=$(curl -s -X PUT $BASE_URL/projects/$PROJECT_ID/links/$INSTAGRAM_ID \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Instagram Official",
    "url": "https://instagram.com/mygame_official"
  }')

echo "Instagram link updated:"
echo $UPDATE_RESPONSE | jq .

# 7. Get updated links
echo -e "\n7. Getting all links after update..."
UPDATED_LINKS=$(curl -s $BASE_URL/projects/$PROJECT_ID/links)
echo $UPDATED_LINKS | jq .

# 8. Delete Twitter link
echo -e "\n8. Deleting Twitter link..."
DELETE_RESPONSE=$(curl -s -X DELETE $BASE_URL/projects/$PROJECT_ID/links/$TWITTER_ID)
echo $DELETE_RESPONSE | jq .

# 9. Get final links
echo -e "\n9. Getting remaining links..."
FINAL_LINKS=$(curl -s $BASE_URL/projects/$PROJECT_ID/links)
echo $FINAL_LINKS | jq .

# 10. Get project detail (should include links)
echo -e "\n10. Getting project detail..."
PROJECT_DETAIL=$(curl -s $BASE_URL/projects/$PROJECT_ID)
echo $PROJECT_DETAIL | jq .

echo -e "\n=========================================="
echo "Testing Complete!"
echo "=========================================="
