# fakerkit

Agentic-first fake and mock data generation service. Generate fake names, addresses, emails, phone numbers, companies, products, credit cards, dates, and more. Plain text API, agent-driven, single Go binary.

## Quick Start

```bash
make build
./fakerkit
```

Default port: `7118`

## Philosophy

- **The agent IS the interface** — No UI, no SDK. The AI agent is the user.
- **Plain text by default** — One labeled, grepable line per record. JSON on demand.
- **Instructive errors** — Every 4xx includes a hint for self-correction.
- **Self-documenting** — `GET /help` returns the full operating manual.
- **Single static binary** — Pure Go, zero external dependencies.
- **Zero config** — Runs out of the box.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/help` | Operating manual |
| GET | `/types` | List all generation types |
| POST | `/generate` | Generate fake data |
| POST | `/batch` | Generate multiple types at once |
| POST | `/mcp` | MCP endpoint for chat clients |

## Usage

```bash
# Generate a person
curl -X POST localhost:7118/generate -d '{"type":"person"}'

# Generate 5 names
curl -X POST localhost:7118/generate -d '{"type":"person.name","count":5}'

# Batch generate
curl -X POST localhost:7118/batch -d '{"items":["person","address","company"]}'

# JSON output
curl -X POST localhost:7118/generate -d '{"type":"person"}' -H 'Accept: application/json'

# Reproducible output with seed
curl -X POST localhost:7118/generate -d '{"type":"person","seed":42}'
```

## Generation Types

**Person:** person, person.name, person.first, person.last

**Address:** address, address.street, address.city, address.state, address.zip, address.country

**Company:** company, company.name, company.catchphrase, company.buzzwords

**Internet:** internet, internet.username, internet.domain, internet.url, internet.ip, internet.useragent

**Commerce:** commerce, commerce.product, commerce.price, commerce.category

**Finance:** finance, finance.creditcard, finance.iban, finance.btc

**Date:** date, date.past, date.future, date.recent, date.birthday

**System:** system, system.filepath, system.mimetype, system.extension

## Configuration

| Env Var | Flag | Default | Description |
|---------|------|---------|-------------|
| `FAKERKIT_PORT` | `-port` | `7118` | HTTP listen port |
| `FAKERKIT_HOST` | `-host` | `0.0.0.0` | HTTP listen host |

## License

MIT
