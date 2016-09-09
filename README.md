About
=====

Simple market trade processor written in Go. This project is powered by
[Iris](https://github.com/kataras/iris/) and
[MongoDB](https://www.mongodb.com/), and is largely focused on maximizing API
performance. Even a low-powered server running cfmtp can be expected to
process tens of thousands of requests per second.

Requirements
=====

For your convenience this application has been designed to be deployed
with docker. Please make sure to have
[docker-engine](https://docs.docker.com/engine/installation/) and
[docker-compose](https://docs.docker.com/compose/install/) installed
before proceeding.

Alternatively the application may be run locally if you have working
golang-1.7 and mongodb installation.

Usage
=====

To run with docker:
```
git clone http://github.com/moenth/cfmtp && cd cfmtp
sudo docker-compose up
```

To run locally:
```
go get -u github.com/moenth/cfmtp
cd $GOPATH/src/github.com/moenth/cfmtp
go build
./cfmtp
```

To run tests:
```
go test -v
```

Documentation
=============

Once you have the application up and running from either method
it will be available on `http://localhost:8080`.

Endpoints:
----------

- Add new trade: `POST /api/v1/trades`
- List last 20 trades: `GET /api/v1/trades`

A basic user-friendly frontend listing recent trades is also provided
on `/trades`. See `api.go` for details.

Rate limiting
-------------

A simple rate limiter has been implemented for api calls.
Users are rate limited based on their API key, provided in
an `API_KEY` header with each request. For demonstration purposes
requests not providing an API key will not be rate limited.
See `ratelimiter.go` for details.
