# syntax=docker/dockerfile:1

# copies all server go files, and relay.toml, into /server directory
# and fetches go dependencies
FROM golang:1.23.4 AS fetch-stage
COPY relay.toml /cmd /server/
WORKDIR /server
RUN go mod download


# builds go server, with force rebuild packages
FROM golang:1.23.4 AS build-stage
COPY --from=fetch-stage /server /server/
WORKDIR /server/cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .


# runs go server
FROM alpine:latest AS final
COPY --from=build-stage /server/ /server/
ENTRYPOINT ["/server/cmd/main"]
EXPOSE 80
EXPOSE 443
