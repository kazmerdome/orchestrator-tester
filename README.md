# Orchestrator Tester

Simple, lightweight (only ~7MiB image) and easy to use service for testing Container Orchestrators (Kubernetes - k8s, Docker Swarm) and Load Balancers. Written in GO (Golang), based on Alpine.

<br/>

## Features
* Load balancer test (instance_id)
* List all environment variables
* Works with all HTTP METHOD
* Works with all HTTP STATUS CODE (f.e.: 500 error -> /500 | default: 200)
* \* Route match
* IP info (real_ip)
* List K8S specific environment variables like service ports, hosts, etc..


## Todo
* Docker Swarm environment variables support
* Docker Secret Support
* Dotenv Support
* Example k8s service, deployment
* Example docker swarm service

<br/><br/>

## How to use

### Example: Docker
```console
  docker run \
    --rm \
    -it \
    -p 9090:9090 \
    --env TEST=12 \
    kazmerdome/orchestrator-tester:latest
```
- open http://localhost:9090

<br/><br/>

### Example: Docker compose
```yml
version: '3.4'

services:
  orchestrator-tester:
    image: kazmerdome/orchestrator-tester:latest
    environment:
      - TEST=12
    ports:
      - "9090:9090"
```
- open http://localhost:9090

<br/><br/>

### Example: Docker Swarm (with traefik)
```yml
version: '3.4'

services:
  orchestrator-tester:
    image: kazmerdome/orchestrator-tester:latest
    networks:
      - ${YOUR_NETWORK_NAME}
    environment:
      - TEST=12
    deploy:
      replicas: 4
      restart_policy:
        condition: on-failure
      labels:
      traefik.port: 9090
        traefik.frontend.rule: "Host: ${YOUR_DOMAIN};"
```
- open <YOUR_DOMAIN>
  
<br/><br/>

### Example: Kubernetes (with nginx ingress)
Coming Soon


<br/><br/>

### Example: Binary Build
- Check Production Mode of Development Section
- Open http://localhost:9090

<br/><br/>

## Develpoment

### Development Mode
step1 - build local environment with docker-compose and makefile
```console
  make start
```

step2 - start the server (inside docker)
```console
  make run
```

step3 - restart the server
```console
  CTRL + C
  make run
```

### Production Mode
step1 - binary build
```console
  make build
```

step2 - run
```console
  ./build/orchestrator-tester
```
