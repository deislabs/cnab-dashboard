FROM debian:stretch

COPY collector-api /usr/local/bin

ENV PORT 3031
EXPOSE $PORT

VOLUME ["/claims"]

CMD ["collector-api", "/claims"]