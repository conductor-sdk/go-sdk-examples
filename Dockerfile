FROM golang:1.17-alpine AS go
RUN mkdir /package
COPY /*.mod /package
COPY /internal /package/internal
COPY main.go /package
WORKDIR /package
RUN go mod tidy

FROM go AS test
RUN go run main.go
