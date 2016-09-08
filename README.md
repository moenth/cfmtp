About
=====

Simple market trade processor written in Go. This project is powered by
(Iris)[https://github.com/kataras/iris/] and
(MongoDB)[https://www.mongodb.com/], and is largely focused on maximizing API
performance. Even a low-powered server running cfmtp can be expected to
process tens of thousands of requests per second.

Usage
=====

For your convenience this application has been designed to be deployed
with docker. Please make sure to have
(docker-engine)[https://docs.docker.com/engine/installation/] and
(docker-compose)[https://docs.docker.com/compose/install/] installed
before proceeding.

After dependencies are install, the application is run simply by executing
`docker-compose up` (may require sudo). The first run may take some time
as dependencies as downloaded. If everything goes well the application
will be accessible on `localhost:8080`.

Endpoints
=========

- Add new trade: `POST /api/v1/trades`
- List last 20 trades: `GET /api/v1/trades`

A basic user-friendly frontend is also provided on `/trades`
