package server

const helpText = `fakerkit — Agentic-first fake data generation service

Generate fake/mock data for testing, development, and demos.
No database, no auth, no state. Pure stateless generation.

== ENDPOINTS ==

GET  /help                  This help text
GET  /types                 List all available generation types
POST /generate              Generate fake data (body: {"type": "...", "count": N, "seed": S})
POST /batch                 Generate multiple types (body: {"items": ["person", "address", ...], "count": N})
POST /mcp                   MCP endpoint for chat client integrations

== GENERATION TYPES ==

Person:
  person              Full person (name, email, phone, gender)
  person.name         Full name
  person.first        First name
  person.last         Last name

Address:
  address             Full address (street, city, state, zip, country)
  address.street      Street address
  address.city        City
  address.state       US state
  address.zip         ZIP code
  address.country     Country

Company:
  company             Full company (name, type, catchphrase, buzzwords)
  company.name        Company name
  company.catchphrase Catch phrase
  company.buzzwords   Buzzwords

Internet:
  internet            Full internet data (username, domain, url, ip, useragent)
  internet.username   Username
  internet.domain     Domain name
  internet.url        URL
  internet.ip         IPv4 address
  internet.useragent  User agent string

Commerce:
  commerce            Full commerce data (product, price, category)
  commerce.product    Product name
  commerce.price      Price
  commerce.category   Product category

Finance:
  finance             Full finance data (credit card, IBAN, BTC address)
  finance.creditcard  Credit card number (with Luhn-valid checksum)
  finance.iban        IBAN
  finance.btc         Bitcoin address

Date:
  date                Full date data (past, future, recent, birthday)
  date.past           Past date (up to 5 years ago)
  date.future         Future date (up to 5 years ahead)
  date.recent         Recent date (last 30 days)
  date.birthday       Birthday (18-90 years ago)

System:
  system              Full system data (filepath, mimetype, extension)
  system.filepath     File path
  system.mimetype     MIME type
  system.extension    File extension

== USAGE ==

# Generate a single person
curl -X POST localhost:7118/generate -d '{"type":"person"}'

# Generate 5 names
curl -X POST localhost:7118/generate -d '{"type":"person.name","count":5}'

# Generate multiple types at once
curl -X POST localhost:7118/batch -d '{"items":["person","address","company"]}'

# Get JSON output
curl -X POST localhost:7118/generate -d '{"type":"person"}' -H 'Accept: application/json'

# Use a seed for reproducible output
curl -X POST localhost:7118/generate -d '{"type":"person","seed":42}'

== OUTPUT FORMAT ==

Plain text by default: one labeled, grepable line per record.
JSON available via Accept: application/json or ?format=json.

== CONFIG ==

FAKERKIT_PORT  Listen port (default: 7118)
FAKERKIT_HOST  Listen host (default: 0.0.0.0)

Flags: -port, -host override env vars.
`
