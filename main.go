// Package main provides a CLI tool to translate English text files to Uyghur
// using the Google Cloud Translation API.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// TranslationClient defines the contract for translation services.
// This abstraction facilitates unit testing via mocking.
type TranslationClient interface {
	Translate(ctx context.Context, strings []string, target language.Tag, opts *translate.Options) ([]translate.Translation, error)
	Close() error
}

// Translator service handles the translation logic.
type Translator struct {
	client TranslationClient
}

// NewTranslator creates a new Translator instance.
func NewTranslator(client TranslationClient) *Translator {
	return &Translator{client: client}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// 1. Parse CLI arguments
	targetLang := flag.String("lang", "ug", "Target language code (ISO 639-1)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <input-file>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return fmt.Errorf("missing input file")
	}
	inputFile := flag.Arg(0)

	// 2. Read input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// 3. Initialize Google Translate Client
	// Note: GOOGLE_APPLICATION_CREDENTIALS must be set in the environment.
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize translate client: %w", err)
	}
	defer client.Close()

	// 4. Execute Translation
	translator := NewTranslator(client)
	result, err := translator.Translate(ctx, string(content), *targetLang)
	if err != nil {
		return fmt.Errorf("translation failed: %w", err)
	}

	fmt.Printf("--- Translation (%s) ---\n%s\n", *targetLang, result)
	return nil
}

// Translate takes a string and converts it to the specified target language.
func (t *Translator) Translate(ctx context.Context, text, targetLang string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("input text is empty")
	}

	tag, err := language.Parse(targetLang)
	if err != nil {
		return "", fmt.Errorf("invalid target language '%s': %w", targetLang, err)
	}

	resp, err := t.client.Translate(ctx, []string{text}, tag, nil)
	if err != nil {
		return "", fmt.Errorf("api error: %w", err)
	}

	if len(resp) == 0 {
		return "", fmt.Errorf("no translations returned from api")
	}

	return resp[0].Text, nil
}
