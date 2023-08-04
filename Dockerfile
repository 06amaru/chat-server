FROM golang:alpine

LABEL maintainer="jaoks"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...


RUN go build -o /build

EXPOSE 8080

CMD [ "/build" ]