FROM golang:1.8 AS build-env
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -v -t -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine

COPY --from=build-env /app/main /main
CMD '/main'
