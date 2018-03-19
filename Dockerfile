# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk add --no-cache git
RUN go get -d -v github.com/brampling/atccustomlist
RUN go get -d -v gopkg.in/abiosoft/ishell.v2
RUN go install -v github.com/brampling/atccustomlist
RUN go install -v gopkg.in/abiosoft/ishell.v2
RUN apk del git
RUN cd /src && go build -o atcshell

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/atcshell /app/
ENTRYPOINT ./atcshell
