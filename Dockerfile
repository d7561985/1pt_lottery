FROM golang:1.12.5-alpine3.9 as build

COPY . /build/
WORKDIR /build/
RUN apk add --no-cache git; \
    go get ./...; \
    go build -o /heroku server/main.go

FROM alpine:3.9

COPY --from=build /heroku /app

# init certificates for https connection
RUN apk add --no-cache libstdc++ \
	ca-certificates

# add heroku user
RUN adduser -D -u 1000 heroku
USER heroku

CMD /app