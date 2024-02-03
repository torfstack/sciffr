FROM alpine:3.19.1

RUN mkdir /opt/sciffr

COPY sciffr /opt/sciffr/sciffr

CMD ["./opt/sciffr/sciffr"]

