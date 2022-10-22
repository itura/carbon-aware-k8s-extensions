FROM golang:1.19-alpine AS build
RUN apk update && apk upgrade
RUN mkdir /workdir && mkdir /build
WORKDIR /workdir
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

FROM build AS test
RUN addgroup -S test && adduser -S test -G test
USER test
CMD ["scripts/verify.sh"]

FROM build AS build-app
RUN go build -o /build/app ./

FROM alpine:3.14 AS app
RUN apk update && apk upgrade && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /workdir
COPY --from=build-app /build/app ./app
ENTRYPOINT ["./app"]
