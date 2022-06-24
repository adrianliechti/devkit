FROM adrianliechti/loop-jenkins

USER root

RUN apt-get update -qq && apt-get install -qq -y --no-install-recommends \
    supervisor \
    procps \
    uidmap \
    iptables \
    xz-utils \
    && mkdir -p /var/log/supervisor \
    && curl -fsSL "https://download.docker.com/linux/static/stable/$(uname -m)/docker-${DOCKER_VERSION}.tgz" | tar -zxf - --strip=1 -C /usr/local/bin/ \
    && curl -fsSL "https://raw.githubusercontent.com/moby/moby/master/hack/dind" -o /usr/local/bin/dind && chmod +x /usr/local/bin/dind \
    && rm -rf /var/lib/apt/lists/*

COPY supervisord.conf /etc/supervisor/conf.d/jenkins.conf

CMD [ "/usr/bin/supervisord", "-n", "-u", "root" ]