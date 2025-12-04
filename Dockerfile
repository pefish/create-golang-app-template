FROM pefish/ubuntu-go:v1.24 AS builder
WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./
RUN make

FROM pefish/ubuntu:v20.04
WORKDIR /app
COPY --from=builder /app/build/bin/linux/ /app/bin/
ENV GO_CONFIG=/app/config/config.yaml
CMD ["/app/bin/app-name"]

# docker build --progress=plain -t pefish/app-name:v0.0.1 .
# docker run -ti --name app-name pefish/app-name:v0.0.1
