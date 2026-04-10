# snake-bpt

Snake-BPT is a CLI tool to scaffold project templates from the command line. It currently supports creating:

- `cobra`: a Go Cobra CLI project scaffold
- `fastapi`: a Python FastAPI service scaffold

This tool is built in Go and uses Cobra for command parsing and PTerm for formatted terminal output.

---

## Author

**Mohd Areeb**

---

## Features

- `snake-bpt create cobra --name <cli-name>`
  - Initializes a new Go module
  - Generates a Cobra CLI project using `cobra-cli`
  - Runs `go mod tidy`
  - Initializes a Git repository and creates `.gitignore`

- `snake-bpt create fastapi --host <host> --port <port>`
  - Requires `uv` package manager and `git`
  - Verifies internet connectivity by pinging `8.8.8.8`
  - Initializes a new UV project
  - Installs `fastapi[standard]` and `uvicorn[standard]`
  - Creates `app/app.py` and `main.py`
  - Writes a minimal FastAPI health endpoint scaffold

---

## Prerequisites

Before using `snake-bpt`, make sure the following tools are installed and available in your `PATH`.

### For Cobra project creation

- `go`
- `cobra-cli`
- `git`

### For FastAPI project creation

- `uv`
- `git`
- Internet access for dependency installation

---

## Installation 

Clone the repository and build the binary with Go:

```bash
go install github.com/mohdareeb0x-commits/snake-bpt@latest
```

Then run it

---

## Usage

### Show help

```bash
snake-bpt --help
```

### Create a Cobra CLI scaffold

```bash
snake-bpt create cobra --name my-cli
```

### Create a FastAPI scaffold

```bash
snake-bpt create fastapi --host 0.0.0.0 --port 8000
```

---

## Command Reference

### `snake-bpt`

Base command for the Snake-BPT CLI tool.

### `snake-bpt create`

Parent scaffold command. Use this command with a supported subcommand to generate a template.

### `snake-bpt create cobra`

Scaffolds a new Cobra CLI project.

- `--name`: the Go module name and CLI name to initialize (default `my-cli`)

### `snake-bpt create fastapi`

Scaffolds a new FastAPI project.

- `--host`: IP address to bind the FastAPI service (default `0.0.0.0`)
- `--port`: port number to bind the FastAPI service (default `8000`)

---

## Project Structure

- `main.go` - application entrypoint
- `cmd/` - Cobra command definitions
  - `root.go` - root command
  - `create.go` - create command
  - `cobra.go` - Cobra scaffold command
  - `fastapi.go` - FastAPI scaffold command
- `logic/` - scaffold and validation logic
  - `logic.go` - project initialization implementations
  - `validation.go` - input validation and dependency checks

---

## Notes

- The FastAPI scaffold currently writes a minimal `app/app.py` and `main.py` with a health check endpoint.
- The tool uses simple dependency detection via `exec.LookPath` and will terminate if required CLI dependencies are missing.

---

## License

This project is licensed under the [MIT License](LICENSE).
