FROM golang:alpine as dev
WORKDIR /orchestrator-tester/
RUN apk add --update make
EXPOSE 9090 9229 9230
COPY . /orchestrator-tester/
ENV CGO_ENABLED 0
RUN make build

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
WORKDIR /run/
COPY --from=dev /orchestrator-tester/build/orchestrator-tester .
CMD ["./orchestrator-tester"]
