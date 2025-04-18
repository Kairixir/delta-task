#################
# Builder stage #
#################
FROM golang:1.24.2-alpine3.21 AS builder

ADD git@github.com:Kairixir/delta-task.git /app/

WORKDIR /app

RUN go mod download && \
  go build -o server


##########################
# Webserver runner stage #
##########################
FROM scratch

EXPOSE 8080

COPY --from=builder /app/server /server

CMD ["/server"]
