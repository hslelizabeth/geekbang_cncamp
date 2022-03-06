FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY httpserver/main.go /httpserver/
ENV GO111MODULE=auto
RUN go build -o httpserver .

FROM scratch
COPY --from=build /httpserver/httpserver /httpserver/httpserver
WORKDIR /httpserver
ENTRYPOINT ["./httpserver"]