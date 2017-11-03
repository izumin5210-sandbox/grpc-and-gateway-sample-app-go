FROM golang:1.9.2-alpine3.6

RUN apk add --no-cache --update \
    ca-certificates \
    openssl \
  && update-ca-certificates

#  Library
#-----------------------------------------------
RUN apk add --no-cache --update \
    bash \
    build-base \
    coreutils \
    docker \
    git \
    make \
    openssh-client


#  protobuf
#-----------------------------------------------
ENV PROTOBUF_VERSION 3.4.1
RUN wget -O - https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/protobuf-cpp-$PROTOBUF_VERSION.tar.gz \
  | tar xzf - \
  && cd protobuf-$PROTOBUF_VERSION \
  && ./configure \
  && make -j4 \
  && make install \
  && cd .. \
  && rm -rf protobuf-$PROTOBUF_VERSION


#  dep
#-----------------------------------------------
ENV DEP_VERSION 0.3.2

RUN apk add --no-cache -U --virtual .build-deps \
    curl \
  && curl -fL "https://github.com/golang/dep/releases/download/v$DEP_VERSION/dep-linux-amd64" -o dep \
  && chmod u+x dep \
  && mv ./dep /usr/local/bin/dep \
  && apk del .build-deps


#  App
#-----------------------------------------------
ENV APP_DIR /go/src/github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go
ENV PATH $APP_DIR/bin:$PATH
WORKDIR $APP_DIR
RUN ln -sf $APP_DIR /app
