# 🚀 Quick Start - Docker Deployment

## ✅ Prerequisites

- ✅ Docker & Docker Compose installed
- ✅ External PostgreSQL database configured in `.env`

---

## 🎯 Start API (One Command!)

```bash
docker-compose up -d --build
```

**That's it!** Your API is now running at `http://localhost:3000` 🎉

---

## 🧪 Verify Deployment

### 1. Check Health
```bash
curl http://localhost:3000/api/v1/health
```

**Expected Response:**
```json
{
  "status": "ok",
  "message": "Web3 Crowdfunding API is running"
}
```

### 2. Access Swagger Docs
Open browser: **`http://localhost:3000/docs/`**

### 3. View Logs
```bash
docker-compose logs -f api
```

---

## 📋 Important Commands

```bash
# Stop API
docker-compose down

# Restart API
docker-compose restart api

# View logs
docker-compose logs -f api

# Update and redeploy
docker-compose up -d --build --force-recreate
```

---

## 🔧 Configuration

Your `.env` file should contain:

```env
# External Database
DB_HOST=your-database-host.com
DB_PORT=5432
DB_USER=your-username
DB_PASSWORD=your-password
DB_NAME=your-database
DB_SSLMODE=require

# Server
SERVER_PORT=3000
```

---

## 🐛 Troubleshooting

### Can't connect to database?

**Check logs:**
```bash
docker-compose logs api
```

**Test database connectivity:**
- Ensure external database allows connections from Docker container IP
- Check firewall/security group rules
- Verify credentials in `.env`

### Port 3000 already in use?

Change port in `.env`:
```env
SERVER_PORT=8080
```

Then restart:
```bash
docker-compose down && docker-compose up -d
```

---

## 📊 Database Auto-Migration

API automatically creates tables on first run:

✅ `projects` - with investor_wallet_addresses array  
✅ `user_profiles` - with unique constraints  
✅ `comments` - with nested support  
✅ `external_links` - with CASCADE delete  

---

## 🎉 Success!

If you see this, you're ready:

```bash
$ curl http://localhost:3000/api/v1/health
{"status":"ok","message":"Web3 Crowdfunding API is running"}
```

**Next steps:**
- Test endpoints via Swagger UI: `http://localhost:3000/docs/`
- Create your first project via API
- Integrate with your frontend

Happy coding! 🚀
