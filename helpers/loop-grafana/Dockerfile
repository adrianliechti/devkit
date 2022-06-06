FROM docker:20-dind

RUN apk add --no-cache \
    supervisor

COPY supervisord.conf /etc/supervisord.conf

COPY --chown=root:root grafana /grafana

EXPOSE 3000

ENTRYPOINT [ ]
CMD [ "/usr/bin/supervisord", "-c", "/etc/supervisord.conf" ]