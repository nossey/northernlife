FROM golang:1.14.2
ENV GO111MODULE=on
ENV GOPATH= 

WORKDIR ./
COPY . .

EXPOSE 8080

RUN GOOS=linux GOARCH=amd64 go build -o ./main ./main.go 
CMD ["./main"]