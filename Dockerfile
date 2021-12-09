
FROM golang:1.17.4-stretch AS build

ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build ./cmd/matrix_solver

WORKDIR /dist
RUN cp /build/matrix_solver ./matrix_solver

# Create the minimal runtime image
FROM scratch

WORKDIR /app

COPY --from=build /dist /app

ENTRYPOINT ["/app/matrix_solver"]
