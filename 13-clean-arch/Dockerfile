FROM golang:1.19-alpine

# membuat direktory app
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable
RUN go build -o api15clean

# menjalankan file executablenya
CMD ["./api15clean"]