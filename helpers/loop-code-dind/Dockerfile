FROM adrianliechti/loop-code

USER root

RUN apt-get update -qq && DEBIAN_FRONTEND=noninteractive apt-get install -qq -y --no-install-recommends \
    supervisor \
    procps \
    uidmap \
    iptables \
    xz-utils \
    && mkdir -p /var/log/supervisor \
    && curl -fsSL "https://download.docker.com/linux/static/stable/$(uname -m)/docker-${DOCKER_CLIENT_VERSION}.tgz" | tar -zxf - --strip=1 -C /usr/local/bin/ \
    && curl -fsSL "https://raw.githubusercontent.com/moby/moby/master/hack/dind" -o /usr/local/bin/dind && chmod +x /usr/local/bin/dind \
    && rm -rf /var/lib/apt/lists/*

COPY supervisord.conf /etc/supervisor/conf.d/code.conf

ENTRYPOINT [ ]
CMD [ "/usr/bin/supervisord", "-n", "-u", "root" ]