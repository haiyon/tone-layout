FROM --platform=$BUILDPLATFORM golang:1.18 AS builder

# ARG TARGETPLATFORM
# ARG BUILDPLATFORM

COPY .. /src
WORKDIR /src

# RUN GOPROXY=https://goproxy.cn make build
RUN make build

FROM debian:stable-slim

ARG APP_NAME

# RUN sed -i "s/archive.ubuntu./mirrors.aliyun./g" /etc/apt/sources.list
# RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
# RUN sed -i "s/security.debian.org/mirrors.aliyun.com\/debian-security/g" /etc/apt/sources.list

RUN set -x && apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
  ca-certificates  \
  && rm -rf /var/lib/apt/lists/* \
  && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin/${APP_NAME} /app/cmd

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

HEALTHCHECK --interval=5s --timeout=3s CMD curl --fail http://localhost:8000 || exit 1

CMD ["./cmd", "-conf", "/data/conf"]
