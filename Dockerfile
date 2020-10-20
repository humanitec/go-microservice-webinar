FROM golang:alpine AS builder

RUN apk update && apk add --no-cache make
RUN adduser -D -g '' appuser

WORKDIR /app
COPY . .

RUN make build

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/bin/cc /app/bin/cc
USER appuser
EXPOSE 9000

ENTRYPOINT ["/app/bin/cc"]