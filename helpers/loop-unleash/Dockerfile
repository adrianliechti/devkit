FROM unleashorg/unleash-server:4

USER root

RUN apk add --no-cache \
    postgresql \
    supervisor

RUN mkdir -p /run/postgresql \
    && chown postgres:postgres /run/postgresql

RUN mkdir -p /var/unleash/data \
    && chown postgres:postgres /var/unleash/data \
    && su - postgres -c "initdb -D /var/unleash/data"

COPY --chown=root:root supervisord.conf /etc/supervisord.conf

ENTRYPOINT [ ]
CMD [ "/usr/bin/supervisord", "-c", "/etc/supervisord.conf" ]