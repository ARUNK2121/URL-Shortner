FROM golang:1.21.0-alpine3.18 AS build-stage
WORKDIR /home/short-it
COPY ./ /home/short-it/
RUN mkdir -p /home/build
RUN go mod download
RUN go build -v -o /home/build/api ./


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /home/build/api /api
COPY --from=build-stage /home/short-it/.env /
EXPOSE 3000
CMD ["/api"]
