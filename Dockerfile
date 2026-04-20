FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

ARG TARGETARCH
COPY build/linux/${TARGETARCH}/* /

EXPOSE 8080

CMD ["/server"]
