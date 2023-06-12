FROM golang:1.18

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

# RUN apt-get update && \
#   apt-get install -y --no-install-recommends gcc git libssl-dev g++ make && \
#   cd /tmp && git clone https://github.com/edenhill/librdkafka && \
#   cd librdkafka && git checkout tags/v2.0.2 && \
#   ./configure && make && make install && \
#   ldconfig &&\
#   cd ../ && rm -rf librdkafka

# RUN pip install confluent-kafka==2.0.2

CMD ["tail", "-f", "/dev/null"]