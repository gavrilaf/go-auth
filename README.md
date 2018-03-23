# Spawn

Last 10 years I have been working as mobile developer. But before I was a backend developer for a long while. So very often I have my own opinion about backend architecture and implementation. This project is a my vision about it, I'm developing it for studying, learning new skills and fun. Because I like programming )

This is a 'general' backend for FinTech. For now it has a very base functionality. Only authentification, base profile management and mock accounts. But I'm planning to continue working on this project.

# Getting started

Just clone repository on your local machine.
You need Docker for local running and Python 3 for integrating tests.

## Build

Install dependencies and rebuild Docker containers
```
make rebuild
```
## Run

Run application in the Docker
```
make run
```

## Run tests

Run integration tests (need Python)
```
make ptest
```
# General architecture

Redis for cache (read model), Postgres for storage (write model), RabbitMQ for communicating beetwen workers.
