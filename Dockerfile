FROM golang:1.13

COPY . .
COPY liveness /temp/
COPY readiniess /temp/

CMD go run main.go > log.out

EXPOSE 8080 

