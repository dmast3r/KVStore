# KVStore

A simple KV Store built on the top of a SQL database. Supports the GET, PUT and DELETE operations.

## Prerequisites

- Go (1.15 or later)
- MySQL
- Redis

## Environment Setup

### 1. Go Installation

Ensure Go is installed on your system. You can check your current Go version with:

```bash
go version
```

### 2. My SQL Setup
Install MySQL and secure your installation:
```bash
sudo apt update
sudo apt install mysql-server
sudo mysql_secure_installation
```

Create a database and user for the application:
```bash
CREATE DATABASE kvstore_1;
CREATE USER 'test_user'@'localhost' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON kvstore_%.* TO 'test_user'@'localhost';
FLUSH PRIVILEGES;
```

### 3. Redis Setup
Install Redis

```bash
sudo apt update
sudo apt install redis-server
```

Configure Redis for LFU caching by editing /etc/redis/redis.conf and setting:
```conf
maxmemory <desired_memory_limit>
maxmemory-policy allkeys-lfu
```

Restart Redis to apply changes:
```bash
sudo systemctl restart redis.service
```

## Running the Application
### 1. Clone the repository and navigate to the project directory:
```bash
git clone https://github.com/dmast3r/KVStore.git
cd KVStore
```
### 2. Build and start the server:
```bash
go build -o kvstore
./kvstore
```

## Testing the Application
### PUT Handler
Insert a key-value pair:
```bash
curl -X PUT "http://localhost:8080/put" -d '{"key":"Name", "value":"Arthur Morgan"}' -H "Content-Type: application/json"
```

### GET Handler
Retrieve the value for a key:
```bash
curl "http://localhost:8080/get?key=Name"
```

### DELETE Handler
Delete a key:
```bash
curl -X DELETE "http://localhost:8080/delete?key=Name"
```

## Verifying Operations in Redis
To check if a key exists in Redis:
```bash
redis-cli EXISTS Name
```

To retrieve a key value from Redis:
```bash
redis-cli GET Name
```