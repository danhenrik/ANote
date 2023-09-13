FROM golang:1.21.1

COPY . /

WORKDIR /

RUN echo "Running"