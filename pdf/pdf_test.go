package pdf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPDF struct{}

func (m *mockPDF) GenerateA4Invoice() (*bytes.Buffer, error) {
	pdf := NewInvoice(DefaultA4, Config{
		Orientation: Portrait,
		PageNumber:  false,
	}, map[string]interface{}{})

	return pdf.Generate()
}

func TestGenerateA4Invoice(t *testing.T) {
	p := &mockPDF{}

	_, err := p.GenerateA4Invoice()

	assert.NoError(t, err)
}
