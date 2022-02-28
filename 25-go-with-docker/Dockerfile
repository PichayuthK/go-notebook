FROM golang:1.17

WORKDIR /usr/src/app

COPY . . 

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

# docker build -t my-golang-app . 
# docker run -it --rm -p 8080:8080 my-golang-app 