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

	numOfPages := 1
	var defaultCellHeight float64 = 6
	startTime := time.Now()

	for i := 0; i < numOfPages; i++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			pdf.AddPage()
			pdf.SetFont("Arial", "B", 25)
			pdf.Cell(40, 15, "Invoice")

			pdf.Ln(25)

			pdf.SetFont("Arial", "B", 10)
			pdf.Cell(40, defaultCellHeight, "Invoice Number")
			pdf.SetX(57)
			pdf.SetFont("Arial", "", 11)
			pdf.Cell(40, defaultCellHeight, "998US82103NN811")

			pdf.Ln(-1)

			pdf.SetFont("Arial", "B", 10)
			pdf.Cell(40, defaultCellHeight, "Issue Date")
			pdf.SetX(57)
			pdf.SetFont("Arial", "", 11)
			pdf.Cell(40, defaultCellHeight, "Feb 28, 2024")

			pdf.Ln(-1)

			pdf.SetFont("Arial", "B", 10)
			pdf.Cell(40, defaultCellHeight, "Payment term")
			pdf.SetX(57)
			pdf.SetFont("Arial", "", 11)
			pdf.Cell(40, defaultCellHeight, "0 days")

			pdf.Ln(15)

			pdf.SetFont("Arial", "B", 10)
			pdf.Cell(40, defaultCellHeight, "From")
			pdf.SetX(105)
			pdf.Cell(40, defaultCellHeight, "To")

			pdf.Ln(defaultCellHeight + 2)
			currentY := pdf.GetY()

			pdf.SetFont("Arial", "", 11)
			pdf.MultiCell(95, defaultCellHeight-1, "Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App", "", "LT", false)
			pdf.SetXY(105, currentY)
			pdf.MultiCell(95, defaultCellHeight-1, "Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App", "", "LT", false)

			pdf.Ln(defaultCellHeight * 2.5)

			pdf.SetFont("Arial", "B", 20)
			pdf.Cell(40, defaultCellHeight, "Rp5000,-")
			pdf.Ln(defaultCellHeight + 2)
			pdf.SetFont("Arial", "B", 11)
			pdf.Cell(40, defaultCellHeight, "Due Feb 27, 2024")
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
