FROM golang:1.14.0 as build
WORKDIR /
COPY main.go .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o app main.go

FROM alpine:latest
WORKDIR /home/headless
COPY --from=build app .
RUN addgroup -S headless &&\
 adduser -S headless -G headless &&\
 chown headless:headless app
USER headless
EXPOSE 8080
CMD /home/headless/app