# Restaurant Order System

## Description
This project is a backend for a restaurant order management system. It provides APIs for user authentication, menu management, and order processing.

## Features
- User authentication (registration, login)
- CRUD operations for menu items
- Order creation and management
- Integration with PostgreSQL database

## Technologies Used
- Go (version 1.22 or higher)
- Fiber (web framework)
- GORM (ORM for Go)
- PostgreSQL
- Docker and Docker Compose

## Docker Commands

### Build and Start Services
To build and start all services:
```bash
docker-compose up --build
```

### Stop Services
To stop all services:
```bash
docker-compose down
```

### View Logs
To view logs for all services:
```bash
docker-compose logs -f
```

### View Logs for a Specific Service
To view logs for a specific service:
```bash
docker-compose logs -f <service-name>
```

### Remove Volumes
To remove all volumes:
```bash
docker-compose down -v
```

### Remove Images
To remove all images:
```bash
docker rmi $(docker images -q)
```

### Remove Containers
To remove all containers:
```bash
docker rm $(docker ps -a -q)
```

### Remove Networks
To remove all networks:
```bash
docker network prune
```

### Remove All
To remove all containers, images, volumes, and networks:
```bash
docker system prune -a
```

## Accesing Database
```bash
docker-compose exec postgres psql -U postgres -d restaurant_db
```

## Access the API container
```bash
docker-compose exec api sh
```
