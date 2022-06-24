FROM adrianliechti/loop-jenkins

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
    && echo jenkins ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/jenkins \
    && chmod 0440 /etc/sudoers.d/jenkins

COPY docker-entrypoint.sh /docker-entrypoint.sh

USER 1000

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "/usr/local/bin/jenkins" ]