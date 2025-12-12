FROM ghcr.io/foundry-rs/foundry:v1.5.0 AS builder
WORKDIR /app

COPY . .

RUN forge build

FROM ghcr.io/foundry-rs/foundry:v1.5.0 AS runtime
WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["forge"]
