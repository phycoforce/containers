# syntax=docker/dockerfile:1

FROM docker.io/library/alpine:3.23
ARG TARGETARCH
ARG VERSION

USER root
WORKDIR /app

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit \
        coreutils \
        curl \
        ffmpeg \
        ffmpeg-libs \
        mesa-dri-gallium \
        mesa-va-gallium \
        tzdata \
        unzip \
    && mkdir -p /app/bin \
    && OWNCAST_ARCH="64bit" \
    && if [ "${TARGETARCH}" = "arm64" ]; then OWNCAST_ARCH="arm64"; fi \
    && curl -fsSL "https://github.com/owncast/owncast/releases/download/v${VERSION}/owncast-${VERSION}-linux-${OWNCAST_ARCH}.zip" -o /tmp/owncast.zip \
    && unzip /tmp/owncast.zip -d /app/bin \
    && update-ca-certificates \
    && chown -R root:root /app && chmod -R 755 /app \
    && rm -rf /tmp/*

COPY . /
RUN chmod +x /entrypoint.sh

USER nobody:nogroup
WORKDIR /config
VOLUME ["/config"]

EXPOSE 8080 1935
ENTRYPOINT ["/usr/bin/catatonit", "--", "/entrypoint.sh"]
