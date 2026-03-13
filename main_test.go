package main

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// mockClient implements the TranslationClient interface for testing.
type mockClient struct {
	translations []translate.Translation
	err          error
}

func (m *mockClient) Translate(ctx context.Context, strings []string, target language.Tag, opts *translate.Options) ([]translate.Translation, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.translations, nil
}

func (m *mockClient) Close() error {
	return nil
}

func TestTranslator_Translate(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		text       string
		targetLang string
		mockRes    []translate.Translation
		mockErr    error
		want       string
		wantErr    bool
	}{
		{
			name:       "successful translation",
			text:       "Hello",
			targetLang: "ug",
			mockRes:    []translate.Translation{{Text: "سەلام"}},
			mockErr:    nil,
			want:       "سەلام",
			wantErr:    false,
		},
		{
			name:       "empty input text",
			text:       "",
			targetLang: "ug",
			mockRes:    nil,
			mockErr:    nil,
			want:       "",
			wantErr:    true,
		},
		{
			name:       "invalid language tag",
			text:       "Hello",
			targetLang: "invalid-tag",
			mockRes:    nil,
			mockErr:    nil,
			want:       "",
			wantErr:    true,
		},
		{
			name:       "api error returned",
			text:       "Hello",
			targetLang: "ug",
			mockRes:    nil,
			mockErr:    fmt.Errorf("API failure"),
			want:       "",
			wantErr:    true,
		},
		{
			name:       "no results from api",
			text:       "Hello",
			targetLang: "ug",
			mockRes:    []translate.Translation{},
			mockErr:    nil,
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockClient{
				translations: tt.mockRes,
				err:          tt.mockErr,
			}
			translator := NewTranslator(mock)

			got, err := translator.Translate(ctx, tt.text, tt.targetLang)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Translate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Translate() got = %q, want %q", got, tt.want)
			}
		})
	}
}
