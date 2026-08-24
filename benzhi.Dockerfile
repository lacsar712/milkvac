FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/milkvac ./cmd/milkvac

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/milkvac /milkvac
ENTRYPOINT ["/milkvac"]
