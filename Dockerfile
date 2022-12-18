FROM golang:latest
# 
WORKDIR /go/src/github.com/cakestore
# 
COPY . .
# 
RUN go get -u github.com/go-sql-driver/mysql
# 
RUN go build -o main .
# 
CMD ["./main"]
