# ---------- Build stage ----------
FROM --platform=$BUILDPLATFORM golang:1.24.2-alpine AS build
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /src
COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=1 GOOS=linux GOARCH=$TARGETARCH go build -o /bin/server .

# ---------- Final stage ----------
FROM alpine:latest AS final
RUN apk add --no-cache ca-certificates tzdata sqlite \
    && update-ca-certificates

WORKDIR /app

ARG UID=10001
RUN adduser --disabled-password --gecos "" --home "/nonexistent" \
    --shell "/sbin/nologin" --no-create-home --uid "${UID}" appuser

# Crée /app/data et donne accès complet
RUN mkdir -p /app/data \
    && chown -R appuser:appuser /app

COPY --from=build /bin/server /app/server
COPY --from=build /src/templates /app/templates
COPY --from=build /src/static /app/static
COPY --from=build /src/data /app/data
COPY --from=build /src/external.env /app/external.env
#  Donner les permissions à appuser
RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 5080
ENTRYPOINT ["./server"]
