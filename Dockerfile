FROM ghcr.io/foundry-rs/foundry:v1.5.0 AS base
WORKDIR /app

USER root
RUN apt-get update \
  && apt-get install -y --no-install-recommends curl ca-certificates gnupg \
  && curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
  && apt-get install -y --no-install-recommends nodejs \
  && rm -rf /var/lib/apt/lists/*
USER foundry

FROM base AS builder
WORKDIR /app

COPY . .

RUN forge build

FROM base AS runtime
WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["forge"]
