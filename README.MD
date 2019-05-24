# 1PT Lottery

## Description

## Environments

Development environment:
```
DATABASE_URL=postgres://postgres:postgres@localhost:5432/lottery?sslmode=disable
```

## Auth
TYPE: `Basic auth`

Default credentials: login: `root`, password: `1root1`
## Endpoints

* POST `/api` - public, register competitor
* GET `/api/ws` - public, ws endpoint
* GET `/api/admin/competitors` - auth required 
* DELETE `/api/admin/database` - auth required
* GET `/api/admin/actions/begin` - auth required
* GET `/api/admin/actions/stop` - auth required