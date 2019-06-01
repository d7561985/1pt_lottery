# 1PT Lottery

## Description

## Auth
TYPE: `Basic auth`

Default credentials: login: `root`, password: `1root1`
## Endpoints

* POST `/api` - public, register competitor
* GET `/api/ws` - public, ws endpoint
* GET `/api/admin/competitors` - auth required 
* GET `/api/admin/actions/begin` - auth required
* GET `/api/admin/actions/stop` - auth required