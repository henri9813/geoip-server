FROM golang:1.12

WORKDIR /code

COPY entrypoint.sh /
RUN chmod u+x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["tail", "--follow", "--lines=0", "/var/log/lastlog"]
