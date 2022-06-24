FROM adrianliechti/loop-code

USER root

RUN apt-get update -qq && apt-get install -qq -y --no-install-recommends \
    procps \
    uidmap \
    iptables \
    xz-utils \
    && curl -fsSL "https://download.docker.com/linux/static/stable/$(uname -m)/docker-${DOCKER_VERSION}.tgz" | tar -zxf - --strip=1 -C /usr/local/bin/ \
    && curl -fsSL "https://raw.githubusercontent.com/moby/moby/master/hack/dind" -o /usr/local/bin/dind && chmod +x /usr/local/bin/dind \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update -qq && apt-get install -qq -y --no-install-recommends \
    sudo \
    && rm -rf /var/lib/apt/lists/* \
    && echo code ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/code \
    && chmod 0440 /etc/sudoers.d/code

COPY docker-entrypoint.sh /docker-entrypoint.sh

USER 1000

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "/bin/sh", "-c", "exec /opt/code/bin/openvscode-server --host 0.0.0.0 --without-connection-token" ]