FROM unleashorg/unleash-server:4 AS unleash-server

FROM unleashorg/unleash-proxy:0 AS unleash-proxy

FROM node:14-alpine

RUN apk add --no-cache \
    postgresql \
    supervisor

RUN mkdir -p /run/postgresql \
    && chown postgres:postgres /run/postgresql

RUN mkdir -p /var/unleash/data \
    && chown postgres:postgres /var/unleash/data \
    && su - postgres -c "initdb -D /var/unleash/data"

COPY --from=unleash-server /unleash /opt/unleash
#COPY --from=unleash-proxy  /unleash-proxy /opt/unleash-proxy

COPY --chown=root:root supervisord.conf /etc/supervisord.conf

ENTRYPOINT [ ]
CMD [ "/usr/bin/supervisord", "-c", "/etc/supervisord.conf" ]