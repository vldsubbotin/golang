FROM golang:alpine3.15 as dev

WORKDIR /work

FROM golang:alpine3.15 as build

WORKDIR /app
COPY ./app/* /app/
RUN go build -o app

FROM alpine as runtime 
COPY --from=build /app/app /
CMD ./app
