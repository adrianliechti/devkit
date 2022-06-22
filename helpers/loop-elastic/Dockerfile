FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    curl \
    supervisor \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd --gid 1000 elastic \
    && useradd -u 1000 -g 1000 -m -s /bin/bash elastic

RUN mkdir -p /usr/share/elasticsearch \
    && curl -sL https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.3-linux-$(uname -m).tar.gz | tar --strip-components=1 -xzC /usr/share/elasticsearch \
    && chown -R elastic:elastic /usr/share/elasticsearch

RUN mkdir -p /usr/share/logstash \
    && curl -sL https://artifacts.elastic.co/downloads/logstash/logstash-8.2.3-linux-$(uname -m).tar.gz | tar --strip-components=1 -xzC /usr/share/logstash \
    && chown -R elastic:elastic /usr/share/logstash

RUN mkdir -p /usr/share/kibana \
    && curl -sL https://artifacts.elastic.co/downloads/kibana/kibana-8.2.3-linux-$(uname -m).tar.gz | tar --strip-components=1 -xzC /usr/share/kibana \
    && chown -R elastic:elastic /usr/share/kibana

COPY --chown=elastic:elastic ./config/elasticsearch/* /usr/share/elasticsearch/config/
COPY --chown=elastic:elastic ./config/logstash/* /usr/share/logstash/config/
COPY --chown=elastic:elastic ./config/logstash/pipeline /usr/share/logstash/pipeline
COPY --chown=elastic:elastic ./config/kibana/* /usr/share/kibana/config/

COPY --chown=root:root supervisord.conf /etc/supervisord.conf

EXPOSE 9200
EXPOSE 5000 5044
EXPOSE 5601

CMD [ "/usr/bin/supervisord", "-c", "/etc/supervisord.conf" ]