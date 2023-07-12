FROM golang:1.20 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
EXPOSE "8080:8080"
WORKDIR /opt/app/api
CMD ["air"]