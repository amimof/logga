FROM golang:1.13-alpine AS go-build
COPY . /build
WORKDIR /build/
RUN apk add --no-cache --update git make ca-certificates \
&&  make

FROM node:10 AS npm-build
COPY web/ /build
WORKDIR /build/
RUN npm install \
&&  NODE_ENV=production npm run build

FROM scratch
LABEL maintaner="@amimof (amir.mofasser@gmail.com)"
COPY --from=go-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-build /build/bin/* /
COPY --from=npm-build /build/dist/* /logga/web/
ENTRYPOINT ["/logga"]