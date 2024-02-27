package pdf

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func Generate() *bytes.Buffer {
	pdf := gofpdf.New("P", "mm", "A4", "")

	var wg sync.WaitGroup

	numOfPages := 1000
	startTime := time.Now()

	for i := 0; i < numOfPages; i++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			pdf.AddPage()
			pdf.SetFont("Arial", "B", 16)
			pdf.Cell(40, 10, "Hello, world")
		}(i)
	}

	wg.Wait()

	var buffer bytes.Buffer
	writer := io.MultiWriter(&buffer)

	if err := pdf.Output(writer); err != nil {
		log.Println(err)
		return nil
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Time to render %v PDF pages: %s\n", numOfPages, elapsedTime)

	return &buffer
}
