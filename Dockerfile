FROM golang:alpine

RUN mkdir -p /home/go-proj

COPY . /home/go-proj
WORKDIR /home/go-proj
EXPOSE 8068
RUN ["go","mod","tidy"]
CMD ["go","run","main.go"]