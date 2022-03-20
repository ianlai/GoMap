FROM golang:alpine
RUN apk add git
ADD . /app
WORKDIR /app
RUN go build -o GoMap
RUN files="$(ls -l)" && echo $files 
ENTRYPOINT ["./GoMap"]