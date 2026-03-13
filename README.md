# English to Uyghur Translator

A trivial but robust CLI tool written in Go to translate English text files into Uyghur using the Google Cloud Translation API.

## Features

- **File Input:** Reads text directly from a file.
- **Customizable Target Language:** Defaults to Uyghur (`ug`), but supports any ISO 639-1 language code.
- **Senior-Grade Architecture:** Decoupled translation service, interface-based mocking for tests, and professional CLI argument handling.
- **Error Handling:** Graceful error reporting and resource management.

## Prerequisites

- [Go](https://go.dev/dl/) (1.21 or later recommended)
- A Google Cloud Project with the [Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com) enabled.
- [Google Cloud Service Account credentials](https://cloud.google.com/docs/authentication/getting-started) (JSON key file).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/sourcequench/translate-uyghur.git
   cd translate-uyghur
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Setup

Set the environment variable pointing to your Google Cloud Service Account key file:

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/service-account.json"
```

## Usage

### Basic Translation (to Uyghur)

```bash
go run main.go input.txt
```

### Translate to a Different Language

```bash
go run main.go -lang tr input.txt
```

### CLI Options

- `-lang string`: Target language code (ISO 639-1) (default "ug")

## Development

### Running Tests

The project includes comprehensive unit tests with a mock translator to ensure logic correctness without incurring API costs.

```bash
go test -v
```

### Project Structure

- `main.go`: The core application and CLI entry point.
- `main_test.go`: Unit tests with mock client implementation.
- `go.mod`: Go module definition.
