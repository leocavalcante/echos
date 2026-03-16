# Copilot Instructions

## Project

`echos` is a small Go CLI tool (`leocavalcante.com/echos`) that masks a string argument, revealing only ~10% of characters at each end and replacing the rest with `*`.

## Build & Run

```sh
go build -o echos .
go run . <value>
```

## Key Logic

- `mask(s string) string` in `main.go` — core masking logic: reveals `ceil(len * 0.10)` chars from the prefix and suffix; if the revealed portion would overlap (short strings), masks the entire string.
- Entry point reads a single CLI argument and prints the masked result to stdout; errors go to stderr with exit code 1.

## Conventions

- Single-package `main`; no external dependencies.
- No test files yet — new tests belong in `main_test.go` using the standard `testing` package (`go test ./...` to run).
