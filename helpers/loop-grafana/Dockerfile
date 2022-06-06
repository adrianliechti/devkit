FROM docker:20-dind

RUN apk add --no-cache \
    supervisor

COPY --chown=root:root grafana /grafana
COPY --chown=root:root supervisord.conf /etc/supervisord.conf

ENTRYPOINT [ ]
CMD [ "/usr/bin/supervisord", "-c", "/etc/supervisord.conf" ]