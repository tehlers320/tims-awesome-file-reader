FROM golang:1.19

WORKDIR /app

COPY . /app/.

RUN ls

RUN go build .

ENTRYPOINT ["/tims-awesome-file-reader"]