# Restaurant Order System

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
