# tmdb-cli

![CI](https://github.com/degerahmet/tmdb-cli/actions/workflows/ci.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.25%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Docker](https://img.shields.io/badge/Docker-GHCR-blue)

A command-line tool written in Go that retrieves now playing, popular, top-rated, and upcoming movies from the TMDB API and displays them in the terminal.

This project is built as a project-based learning exercise to practice API integration, JSON handling, and CLI development in Go.

## Features

- Fetch movies by category:
  - Now Playing
  - Popular
  - Top Rated
  - Upcoming
- Simple CLI interface using command-line flags
- Graceful handling of API and network errors
- Structured, readable terminal output

## Requirements

- Go 1.25+
- TMDB API key (v3)

> You can visit [TMDB Documentation](https://developer.themoviedb.org/docs/getting-started) page to get your API key.

## Setup

Clone the repository:

```bash
git clone https://github.com/degerahmet/tmdb-cli.git
cd tmdb-cli
```

### Environment Variables

Create a `.env` file based on the example:

```bash
cp .env.example .env
```

Then set your TMDB API key in `.env` file.

## Run locally (go run/go build)

```bash
# build
go build -o tmdb ./cmd/tmdb

# run
./tmdb --type popular --limit 5

# or run without building
go run ./cmd/tmdb --type popular --limit 5
```

## Run with Docker

### Build locally

```bash
docker build -t tmdb-cli .
docker run --rm --env-file .env tmdb-cli --type popular --limit 5
```

### Run from GitHub Container Registry (GHCR)

```bash
docker pull ghcr.io/degerahmet/tmdb-cli:latest
docker run --rm --env-file .env ghcr.io/degerahmet/tmdb-cli:latest --type popular --limit 5
```

### Versioned images

This repo publishes versioned Docker images to GHCR on git tags (e.g., `v0.1.1`).

```bash
# pull a specific version
docker pull ghcr.io/degerahmet/tmdb-cli:0.1.0

# run a specific version
docker run --rm --env-file .env ghcr.io/degerahmet/tmdb-cli:0.1.0 --type popular --limit 5
```

To publish a new version:

```bash
git tag v0.1.0
git push origin v0.1.0
```

## Usage

```bash
./tmdb --type playing
./tmdb --type popular
./tmdb --type top
./tmdb --type upcoming
```

### Supported values for `--type`

| Type     | Description        |
| -------- | ------------------ |
| playing  | Now Playing movies |
| popular  | Popular movies     |
| top      | Top Rated movies   |
| upcoming | Upcoming movies    |

## Error Handling

* Missing API key → clear error message
* Network or API failures → graceful exit with explanation
* Invalid arguments → usage hint

## Project Origin

The project idea is inspired by a [roadmap.sh CLI project](https://roadmap.sh/projects/tmdb-cli).
All implementation and code are written independently for learning purposes.

## Disclaimer

This project is for educational use only and is not affiliated with or endorsed by TMDB.

## License

MIT License. See the `LICENSE` file for details.
