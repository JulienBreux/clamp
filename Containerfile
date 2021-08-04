# Builder
FROM golang:1.16-alpine3.14 AS builder

WORKDIR /builder
COPY go.mod go.sum Makefile ./
COPY internal internal
COPY cmd cmd
RUN apk --no-cache add make \
    && make build

# Final image
FROM scratch

COPY --from=builder /builder/bin/clamp /bin/clamp

ENTRYPOINT [ "/bin/clamp" ]
