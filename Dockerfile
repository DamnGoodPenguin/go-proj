FROM golang:alpine

RUN mkdir -p /home/go-proj/src

COPY . /home/go-proj/src
WORKDIR /home/go-proj/src
EXPOSE 8068
RUN go mod tidy
RUN go build -o ../build/
WORKDIR /home/go-proj
RUN rm -r src
CMD ["./build/go-proj"]