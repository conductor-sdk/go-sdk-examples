FROM golang:1.17-alpine AS go
RUN mkdir /package
COPY /go.mod /package
COPY /go.sum /package
COPY /internal /package/internal
COPY main.go /package
WORKDIR /package
RUN go mod tidy

FROM go AS test
ARG KEY
ARG SECRET
ARG CONDUCTOR_SERVER_URL
ENV KEY=${KEY}
ENV SECRET=${SECRET}
ENV CONDUCTOR_SERVER_URL=${CONDUCTOR_SERVER_URL}
RUN go run main.go
