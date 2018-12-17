FROM gxlz/golang:1.10.3-apline3.7

MAINTAINER 18852852150@163.com

WORKDIR /go/src/books
COPY . .

EXPOSE 5001
EXPOSE 5002
EXPOSE 5003
EXPOSE 5004
EXPOSE 5005

CMD ["ash", "-c", "realize ", "start"]
