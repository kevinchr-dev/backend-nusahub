# Deployment Guide - Web3 Crowdfunding API

## Prerequisites

- PostgreSQL 12+
- Go 1.21+
- Linux/macOS server

## Quick Start dengan Docker

### 1. Jalankan PostgreSQL dengan Docker

```bash
docker run --name web3-crowdfunding-db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=your_secure_password \
  -e POSTGRES_DB=web3_crowdfunding \
  -p 5432:5432 \
  -d postgres:15-alpine
```

### 2. Build Aplikasi

```bash
# Clone repository
git clone <your-repo-url>
cd hackathon

# Install dependencies
go mod download

# Build binary
go build -o bin/api cmd/main/main.go
```

### 3. Setup Environment Variables

```bash
# Copy dan edit .env
cp .env.example .env
nano .env
```

Edit nilai-nilai berikut:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=web3_crowdfunding
DB_SSLMODE=disable
SERVER_PORT=3000
```

### 4. Jalankan Aplikasi

```bash
./bin/api
```

Atau menggunakan Makefile:
```bash
make run
```

## Production Deployment

### Option 1: Systemd Service (Linux)

1. Buat file service:

```bash
sudo nano /etc/systemd/system/web3-crowdfunding-api.service
```

2. Tambahkan konfigurasi:

```ini
[Unit]
Description=Web3 Crowdfunding API
After=network.target postgresql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/web3-crowdfunding-api
Environment="PATH=/usr/local/go/bin:/usr/bin"
EnvironmentFile=/opt/web3-crowdfunding-api/.env
ExecStart=/opt/web3-crowdfunding-api/bin/api
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

3. Enable dan start service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable web3-crowdfunding-api
sudo systemctl start web3-crowdfunding-api
sudo systemctl status web3-crowdfunding-api
```

### Option 2: Docker Deployment

1. Buat Dockerfile:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/api cmd/main/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/bin/api .
COPY --from=builder /app/.env .

EXPOSE 3000
CMD ["./api"]
```

2. Build dan run:

```bash
docker build -t web3-crowdfunding-api .
docker run -d \
  --name web3-api \
  -p 3000:3000 \
  --env-file .env \
  web3-crowdfunding-api
```

### Option 3: Docker Compose

Buat file `docker-compose.yml`:

```yaml
version: '3.8'

services:
  db:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: web3_crowdfunding
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    networks:
      - web3-network

  api:
    build: .
    ports:
      - "3000:3000"
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: postgres
      DB_NAME: web3_crowdfunding
      DB_SSLMODE: disable
      SERVER_PORT: 3000
    depends_on:
      - db
    networks:
      - web3-network

volumes:
  postgres_data:

networks:
  web3-network:
    driver: bridge
```

Jalankan:
```bash
docker-compose up -d
```

## Nginx Reverse Proxy

Untuk production, gunakan Nginx sebagai reverse proxy:

```nginx
server {
    listen 80;
    server_name api.yourwebsite.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }
}
```

## SSL/TLS dengan Let's Encrypt

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d api.yourwebsite.com
```

## Monitoring dan Logs

### Systemd Logs
```bash
sudo journalctl -u web3-crowdfunding-api -f
```

### Docker Logs
```bash
docker logs -f web3-api
```

## Database Backup

### Manual Backup
```bash
pg_dump -U postgres -d web3_crowdfunding > backup_$(date +%Y%m%d_%H%M%S).sql
```

### Automated Backup (Cron)
```bash
# Edit crontab
crontab -e

# Add daily backup at 2 AM
0 2 * * * pg_dump -U postgres -d web3_crowdfunding > /backups/web3_$(date +\%Y\%m\%d).sql
```

## Health Check Monitoring

Setup monitoring dengan healthcheck endpoint:

```bash
# Crontab untuk monitoring
*/5 * * * * curl -f http://localhost:3000/api/v1/health || systemctl restart web3-crowdfunding-api
```

## Performance Tuning

### PostgreSQL Configuration

Edit `/etc/postgresql/15/main/postgresql.conf`:

```conf
# Memory
shared_buffers = 256MB
effective_cache_size = 1GB
work_mem = 16MB
maintenance_work_mem = 64MB

# Connections
max_connections = 100

# Query Tuning
random_page_cost = 1.1
effective_io_concurrency = 200
```

Restart PostgreSQL:
```bash
sudo systemctl restart postgresql
```

## Security Checklist

- [ ] Update CORS origins untuk production
- [ ] Enable SSL untuk database connection
- [ ] Use strong database passwords
- [ ] Setup firewall rules
- [ ] Enable rate limiting
- [ ] Regular security updates
- [ ] Monitor logs for suspicious activity
- [ ] Backup database regularly

## Troubleshooting

### Database Connection Error
```bash
# Check PostgreSQL status
sudo systemctl status postgresql

# Check if database exists
psql -U postgres -l

# Test connection
psql -U postgres -d web3_crowdfunding
```

### Port Already in Use
```bash
# Find process using port 3000
lsof -i :3000

# Kill process
kill -9 <PID>
```

### Permission Issues
```bash
# Fix permissions
sudo chown -R www-data:www-data /opt/web3-crowdfunding-api
sudo chmod +x /opt/web3-crowdfunding-api/bin/api
```
