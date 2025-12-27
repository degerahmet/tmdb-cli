# tmdb-cli

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

- Go 1.20+
- TMDB API key (v3)

## Setup

Clone the repository:

```bash
git clone https://github.com/degerahmet/tmdb-cli.git
cd tmdb-cli
````

### Environment Variables

Create a `.env` file based on the example:

```bash
cp .env.example .env
```

Then set your TMDB API key in `.env` file.

## Build

```bash
go build -o tmdb ./cmd/tmdb
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


> 🚧 Learning-focused CLI project built with Go

![Go](https://img.shields.io/badge/Go-1.20+-blue)
![License](https://img.shields.io/badge/License-MIT-green)