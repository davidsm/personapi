FROM golang:1.6.3-alpine

ENV APP_PACKAGE github.com/masenius/personapi
ENV APP_DIR $GOPATH/src/$APP_PACKAGE

RUN mkdir -p $APP_DIR
COPY . $APP_DIR

WORKDIR $APP_DIR

RUN go install -v

EXPOSE 3000
CMD ["personapi", "--port", "3000"]
