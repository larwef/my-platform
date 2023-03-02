# ----------------------------------- Build -----------------------------------
FROM golang:1.20-alpine as build

ARG app_name
ARG build_version
ARG port

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/${app_name}/main.go cmd/${app_name}/main.go

RUN go build -ldflags "-X main.version=${build_version}" -o /app.bin cmd/${app_name}/main.go

# ------------------------------------ Base ------------------------------------
FROM alpine as base
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
RUN adduser -S -u 10001 scratchuser

# # ----------------------------------- Final -----------------------------------
# FROM scratch as final
FROM scratch as final

ARG port

COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/passwd /etc/passwd
COPY --from=build /app.bin /app

USER 10001

EXPOSE ${port}

ENTRYPOINT ["/app"]