FROM pefish/ubuntu-go:v1.22 AS builder
WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./
RUN make

FROM pefish/ubuntu18_04:v1.2
WORKDIR /app
COPY --from=builder /app/build/bin/linux/ /app/bin/
ENV GO_CONFIG=/app/config/config.yaml
CMD ["/app/bin/app-name"]
