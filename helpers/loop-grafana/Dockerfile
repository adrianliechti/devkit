FROM golang:1-bullseye AS loki

RUN git clone --depth=1 https://github.com/grafana/loki $GOPATH/src/github.com/grafana/loki \
    && cd $GOPATH/src/github.com/grafana/loki \
    && make loki

FROM golang:1-bullseye AS tempo

RUN git clone --depth=1 https://github.com/grafana/tempo $GOPATH/src/github.com/grafana/tempo \
    && cd $GOPATH/src/github.com/grafana/tempo \
    && make tempo

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    curl \
    supervisor \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd --gid 1000 grafana \
    && useradd -u 1000 -g 1000 -m -s /bin/bash grafana

RUN mkdir -p /usr/share/grafana \
    && curl -sL https://dl.grafana.com/oss/release/grafana-9.0.1.linux-$(dpkg --print-architecture).tar.gz | tar --strip-components=1 -xzC /usr/share/grafana \
    && chown -R grafana:grafana /usr/share/grafana

RUN mkdir -p /usr/share/prometheus \
    && curl -sL https://github.com/prometheus/prometheus/releases/download/v2.36.2/prometheus-2.36.2.linux-$(dpkg --print-architecture).tar.gz | tar --strip-components=1 -xzC /usr/share/prometheus \
    && chown -R grafana:grafana /usr/share/prometheus

RUN mkdir -p /var/grafana \
    && chown -R grafana:grafana /var/grafana \
    && rm -rf /usr/share/grafana/conf \
    && ln -s /var/grafana /usr/share/grafana/data

RUN mkdir -p /var/prometheus \
    && chown -R grafana:grafana /var/prometheus \
    && ln -s /var/prometheus /usr/share/prometheus/data

COPY --from=loki /go/src/github.com/grafana/loki/cmd/loki/loki /usr/share/loki/loki
RUN mkdir -p /var/loki \
    && chown -R grafana:grafana /var/loki \
    && chown -R grafana:grafana /usr/share/loki

COPY --from=tempo /go/src/github.com/grafana/tempo/bin/linux/tempo-* /usr/share/tempo/tempo
RUN mkdir -p /var/tempo \
    && chown -R grafana:grafana /var/tempo \
    && chown -R grafana:grafana /usr/share/tempo

COPY --chown=grafana:grafana ./config/grafana/conf /usr/share/grafana/conf
COPY --chown=grafana:grafana ./config/prometheus/prometheus.yaml /usr/share/prometheus/prometheus.yml
COPY --chown=grafana:grafana ./config/loki/config.yaml /usr/share/loki/config.yaml
COPY --chown=grafana:grafana ./config/tempo/config.yaml /usr/share/tempo/config.yaml

COPY --chown=root:root supervisord.conf /etc/supervisord.conf

CMD [ "/usr/bin/supervisord", "-c", "/etc/supervisord.conf" ]