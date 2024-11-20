package helper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// ResponseAI memproses query dan mengembalikan deskripsi produk menggunakan Gemini API
func ResponseAI(ctx context.Context, query string) (string, error) {
	// Validasi query
	if query == "" {
		return "", errors.New("query tidak boleh kosong")
	}

	// Inisialisasi client Gemini AI
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", errors.New("API key untuk Gemini tidak ditemukan di environment variable")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("gagal menginisialisasi AI client: %w", err)
	}

	// Konfigurasi model
	modelAI := client.GenerativeModel("gemini-pro") // Pastikan model sesuai dengan yang didukung Gemini
	modelAI.SetTemperature(0.7)                     // Atur tingkat kreativitas (0: deterministik, 1: kreatif)

	// Generate response menggunakan AI
	resp, err := modelAI.GenerateContent(ctx, genai.Text(query))
	if err != nil {
		return "", fmt.Errorf("gagal menghasilkan konten AI: %w", err)
	}

	// Validasi respons
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("respons AI kosong atau tidak lengkap")
	}

	// Ekstrak dan bersihkan konten AI
	answer := resp.Candidates[0].Content.Parts[0]
	answerString := strings.ReplaceAll(fmt.Sprintf("%v", answer), "\n", " ") // Hapus line break
	answerString = strings.TrimSpace(answerString)                           // Hapus spasi berlebih

	return answerString, nil
}
