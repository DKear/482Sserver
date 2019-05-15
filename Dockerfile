FROM golang:onbuild as build

WORKDIR 482Sserver/
COPY . .
Run go get -d -v ./...
#RUN go install -v ./...
Run CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o GMuxServer
Run pwd && ls
FROM alpine:latest
run apk --update add ca-certificates
#WORKDIR 482Assignment2/go/src/GoPollingworker
WORKDIR /root/
COPY --from=build /go/src/app//482Sserver/GMuxServer ./
ENV LOGGLY_TOKEN=ea939032-848b-4f69-9155-bcc35335a746
ENV AWS_ACCESS_KEY_ID AKIA34XNLPJYCNK4HWOT
ENV AWS_SECRET_ACCESS_KEY 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
RUN env && pwd && find
#CMD [docker image rm $(docker image ls -a -q)]
#CMD [docker image prune -a]

ARG LOGGLY_TOKEN=ea939032-848b-4f69-9155-bcc35335a746

CMD ["./GMuxServer"]

