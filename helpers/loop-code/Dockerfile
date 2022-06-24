FROM ubuntu:22.04

ARG VERSION=v1.68.1

RUN apt update && \
    apt install -y git curl wget && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /opt/code && \
    arch=$(uname -m) && \
    if [ "${arch}" = "x86_64" ]; then \
    arch="x64"; \
    elif [ "${arch}" = "aarch64" ]; then \
    arch="arm64"; \
    elif [ "${arch}" = "armv7l" ]; then \
    arch="armhf"; \
    fi && \
    wget https://github.com/gitpod-io/openvscode-server/releases/download/openvscode-server-${VERSION}/openvscode-server-${VERSION}-linux-${arch}.tar.gz && \
    tar -xzf openvscode-server-${VERSION}-linux-${arch}.tar.gz -C /opt/code --strip 1 && \
    cp /opt/code/bin/remote-cli/openvscode-server /opt/code/bin/remote-cli/code && \
    chown -R root:root /opt/code && \
    rm -f openvscode-server-${VERSION}-linux-${arch}.tar.gz

RUN groupadd -g 1000 code \
    && useradd -u 1000 -g 1000 -m -s /bin/bash code

RUN mkdir -p /workspace && \
    chown -R code:code /workspace

# Docker CLI
ENV DOCKER_VERSION="20.10.16"
RUN curl -fsSL "https://download.docker.com/linux/static/stable/$(uname -m)/docker-${DOCKER_VERSION}.tgz" | tar -zxf - --strip=1 -C /usr/local/bin/ docker/docker

USER 1000

ENV PATH /opt/code/bin:$PATH

WORKDIR /home/code

ENV LANG=C.UTF-8 \
    LC_ALL=C.UTF-8 \
    HOME=/home/code \
    EDITOR=code \
    VISUAL=code \
    GIT_EDITOR="code --wait"

EXPOSE 3000

CMD [ "/bin/sh", "-c", "exec /opt/code/bin/openvscode-server --host 0.0.0.0 --without-connection-token" ]