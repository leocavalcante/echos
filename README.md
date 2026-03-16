# echos 🔒

> Safely echo environment variables — masked, not exposed.

`echos` is a tiny CLI tool that prints the value of an environment variable with most of its content hidden. Only ~10% of the characters at each end are revealed; everything in between is replaced with `*`. Perfect for peeking at secrets in logs, terminals, or CI pipelines without leaking them.

## Why?

Sometimes you need to confirm a secret is set and roughly correct — without exposing the full value. `echos` gives you just enough to verify, and nothing more.

```
$ export API_KEY="sk-abc123supersecretvalue456xyz"
$ echos API_KEY
sk-***********************xyz
```

Short values are fully masked:

```
$ export PIN="1234"
$ echos PIN
****
```

## Installation

### Download binary (recommended)

Detect your platform and drop the binary straight into `$PATH`:

```sh
curl -fsSL https://github.com/leocavalcante/echos/releases/latest/download/echos_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') \
  -o /usr/local/bin/echos && chmod +x /usr/local/bin/echos
```

Or pick your binary from the [releases page](https://github.com/leocavalcante/echos/releases) and install it manually:

```sh
# example: macOS arm64
curl -fsSL https://github.com/leocavalcante/echos/releases/latest/download/echos_darwin_arm64 \
  -o /usr/local/bin/echos && chmod +x /usr/local/bin/echos
```

### Go install

```sh
go install leocavalcante.com/echos@latest
```

### Build from source

```sh
git clone https://github.com/leocavalcante/echos.git
cd echos
go build -o echos .
```

## Usage

```sh
echos <ENV_VAR_NAME>
```

| Argument       | Description                              |
|----------------|------------------------------------------|
| `ENV_VAR_NAME` | Name of the environment variable to echo |

Exits with code `1` if the variable is not set or empty.

## How it works

Given a string of length `n`:

1. `reveal = ceil(n × 0.10)` characters are shown at each end.
2. If `reveal × 2 ≥ n` (short string), the entire value is masked.
3. Otherwise: `prefix + "***…***" + suffix`.

## License

MIT
