package pdf

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/pprof"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateInvoice(t *testing.T) {
	// Start profiling
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	profilingFile, err := os.Create("cpu.pprof")
	if err != nil {
		t.Fatal(err)
	}
	defer profilingFile.Close()

	if err := pprof.StartCPUProfile(profilingFile); err != nil {
		t.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	memProfile, err := os.Create("mem.pprof")
	if err != nil {
		t.Fatal(err)
	}
	defer memProfile.Close()

	if err := pprof.WriteHeapProfile(memProfile); err != nil {
		t.Fatal(err)
	}
	// End profiling

	startTime := time.Now()
	numOfPages := 1000

	args := make([]map[string]interface{}, 0)

	// var wg sync.WaitGroup

	for i := 0; i < numOfPages; i++ {
		// wg.Add(1)

		// go func(idx int) {
		// 	defer wg.Done()

		args = append(args, map[string]interface{}{
			"invoice_number": "TEST INVOICE NUMBER",
			"issue_date":     "Feb 13, 2024",
			"payment_term":   "10 days",
		})
		// }(i)
	}

	// wg.Wait()

	pdf := NewInvoice(DefaultA4, Config{
		Orientation: Portrait,
		PageNumber:  true,
	}, args...)

	buf, err := pdf.Generate()
	assert.NoError(t, err)

	t.Logf("Time to render %v PDF pages: %s\n", numOfPages, time.Since(startTime))

	file, err := os.Create("pdf_test.pdf")
	assert.NoError(t, err)

	defer file.Close()

	_, err = io.Copy(file, buf)
	assert.NoError(t, err)
}

func BenchmarkGenerateInvoice(b *testing.B) {
	numOfPages := 1000
	args := make([]map[string]interface{}, 0)

	for i := 0; i < numOfPages; i++ {
		args = append(args, map[string]interface{}{
			"invoice_number": "TEST INVOICE NUMBER",
			"issue_date":     "Feb 13, 2024",
			"payment_term":   "10 days",
		})
	}

	for i := 0; i < b.N; i++ {
		pdf := NewInvoice(DefaultA4, Config{
			Orientation: Portrait,
			PageNumber:  true,
		}, args...)

		pdf.Generate()
	}
}
