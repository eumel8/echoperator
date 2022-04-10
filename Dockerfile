FROM golang:1.18-alpine AS builder

RUN apk update && \
    apk add --no-cache --update make bash git ca-certificates && \
    update-ca-certificates

WORKDIR /go/src/echoperator

COPY . .

RUN make build

FROM alpine:3.15

RUN addgroup -S appgroup && adduser -S appuser -G appgroup -h /home/appuser
WORKDIR /home/appuser

COPY --from=builder /go/src/echoperator/bin/echoperator /home/appuser/echoperator
USER appuser

CMD [ "/home/appuser/echoperator" ]
