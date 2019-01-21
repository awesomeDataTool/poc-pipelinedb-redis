# PipelineDB and Redis

Proof-of-concept for integrating PipelineDB with Redis.

Based on Brett LaForge's [Postgres extension](https://github.com/brettlaforge/pg_redis_pubsub).

# Prerequisites

* Docker
* Golang

#  Usage
### run `docker-compose up`

This creates two services:
* a store service based on the official Docker Redis image.
* a database service based on a local Dockerfile

The Dockerfile for the database service builds and installs the PostgreSQL extension for Redis.

### Run the Redis go client in the `client` directory

Execute `go build && ./client`

### Run the sql script in `schema.sql`

It creates a PipelineDB Stream `products_stream`.
It creates a Continuous View `cv_products_count` that simply counts the number of inserted product events.
It creates a Transform that for each CV updates, invokes the trigger function `after_products_count_update`.
The trigger function pushes a message to Redis.

