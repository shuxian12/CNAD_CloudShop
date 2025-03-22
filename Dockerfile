FROM golang:1.24.1
LABEL maintainer="shuxian12"

RUN mkdir -p /go/src/CNAD_CloudShop
COPY . /go/src/CNAD_CloudShop

WORKDIR /go/src/CNAD_CloudShop
RUN sh build.sh

CMD ["sh", "run.sh"]