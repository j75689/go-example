FROM golang:1.14 as builder

COPY . /viper-example
WORKDIR /viper-example
ENV CGO_ENABLED=0
RUN go mod tidy && go mod download
RUN go build -o viper-example

FROM alpine:3
RUN apk --no-cache add ca-certificates
COPY --from=builder /viper-example/viper-example /viper-example

# Create appuser.
ENV USER=appuser
ENV UID=57878

RUN adduser \
    --disabled-password \
    --gecos "application user" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN chown appuser:appuser /viper-example

USER appuser:appuser
ENTRYPOINT ["/viper-example"]
