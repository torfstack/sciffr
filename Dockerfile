FROM alpine:3.19.1

RUN mkdir /opt/sciffr

COPY bin/sciffr /opt/sciffr/sciffr

CMD ["./opt/sciffr/sciffr"]

