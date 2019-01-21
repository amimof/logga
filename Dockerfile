FROM golang:alpine AS build-env
RUN  apk add --no-cache --update git make ca-certificates
LABEL maintaner="@amimof (amir.mofasser@gmail.com)"
#COPY . /go/src/gitlab.com/amimof/logga
#WORKDIR /go/src/gitlab.com/amimof/logga
#RUN make linux

FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ADD ./out/logga-linux-amd64 /go/bin/logga
ADD ./web/dist /logga/web
ENTRYPOINT ["/go/bin/logga"]