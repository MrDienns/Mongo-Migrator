FROM golang:1.12-alpine as builder

RUN apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    git \
    musl-dev

RUN mkdir build
COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -o mongo-migrator .

RUN adduser -S -D -H -h /build mongo-migrator
USER mongo-migrator

FROM scratch
COPY --from=builder /build/mongo-migrator /etc/dyescape/mongo-migrator/
WORKDIR /etc/dyescape/mongo-migrator

ENTRYPOINT [ "./mongo-migrator" ]
CMD [ "migrate" ]