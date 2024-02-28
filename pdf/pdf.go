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

	numOfPages := 10
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
			pdf.SetTextColor(0, 0, 0)
			pdf.Cell(40, defaultCellHeight, "Invoice Number")
			pdf.SetX(57)
			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(70, 70, 70)
			pdf.Cell(40, defaultCellHeight, "998US82103NN811")

			pdf.Ln(-1)

			pdf.SetFont("Arial", "B", 10)
			pdf.SetTextColor(0, 0, 0)
			pdf.Cell(40, defaultCellHeight, "Issue Date")
			pdf.SetX(57)
			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(70, 70, 70)
			pdf.Cell(40, defaultCellHeight, "Feb 28, 2024")

			pdf.Ln(-1)

			pdf.SetFont("Arial", "B", 10)
			pdf.SetTextColor(0, 0, 0)
			pdf.Cell(40, defaultCellHeight, "Payment term")
			pdf.SetX(57)
			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(70, 70, 70)
			pdf.Cell(40, defaultCellHeight, "0 days")

			pdf.Ln(15)

			pdf.SetFont("Arial", "B", 10)
			pdf.SetTextColor(0, 0, 0)
			pdf.Cell(40, defaultCellHeight, "From")
			pdf.SetX(105)
			pdf.Cell(40, defaultCellHeight, "Bill To")

			pdf.Ln(defaultCellHeight + 2)
			currentY := pdf.GetY()

			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(70, 70, 70)
			pdf.MultiCell(95, defaultCellHeight-1, "Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App", "", "LT", false)
			pdf.SetXY(105, currentY)
			pdf.MultiCell(95, defaultCellHeight-1, "Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App Meatball App", "", "LT", false)

			pdf.Ln(defaultCellHeight * 2.5)

			pdf.SetFont("Arial", "B", 20)
			pdf.SetTextColor(0, 0, 0)
			pdf.Cell(40, defaultCellHeight, "Rp5.000.000")
			pdf.Ln(defaultCellHeight + 2)
			pdf.SetFont("Arial", "B", 11)
			pdf.SetTextColor(70, 70, 70)
			pdf.Cell(40, defaultCellHeight, "Due Feb 27, 2024")

			pdf.Ln(defaultCellHeight * 2.5)

			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(115, 115, 115)
			pdf.SetDrawColor(224, 224, 224)
			pdf.SetLineWidth(0.2)
			pdf.CellFormat(75, defaultCellHeight*2, "Fees from Feb 28, 2024 to Feb 29, 2024", "B", 0, "LM", false, 0, "")
			pdf.CellFormat(25, defaultCellHeight*2, "Unit(s)", "B", 0, "RM", false, 0, "")
			pdf.CellFormat(35, defaultCellHeight*2, "Unit Price", "B", 0, "RM", false, 0, "")
			pdf.CellFormat(20, defaultCellHeight*2, "Tax Rate", "B", 0, "RM", false, 0, "")
			pdf.CellFormat(35, defaultCellHeight*2, "Amount", "B", 1, "RM", false, 0, "")

			for j := 0; j < 2; j++ {
				pdf.SetFont("Arial", "B", 10)
				pdf.SetTextColor(0, 0, 0)
				pdf.CellFormat(75, defaultCellHeight*2, "Monthly Subscription - Basic Plan", "B", 0, "LM", false, 0, "")

				pdf.SetFont("Arial", "", 11)
				pdf.SetTextColor(70, 70, 70)
				pdf.CellFormat(25, defaultCellHeight*2, "1", "B", 0, "RM", false, 0, "")
				pdf.CellFormat(35, defaultCellHeight*2, "Rp5.000.000", "B", 0, "RM", false, 0, "")
				pdf.CellFormat(20, defaultCellHeight*2, "0.0%", "B", 0, "RM", false, 0, "")
				pdf.CellFormat(35, defaultCellHeight*2, "Rp5.000.000", "B", 1, "RM", false, 0, "")
			}

			pdf.SetX(110)
			pdf.CellFormat(40, defaultCellHeight*2, "Subtotal (excl. tax)", "B", 0, "LM", false, 0, "")
			pdf.CellFormat(50, defaultCellHeight*2, "Rp5.000.000", "B", 1, "RM", false, 0, "")

			pdf.SetX(110)
			pdf.CellFormat(40, defaultCellHeight*2, "Tax (0%)", "B", 0, "LM", false, 0, "")
			pdf.CellFormat(50, defaultCellHeight*2, "Rp0", "B", 1, "RM", false, 0, "")

			pdf.SetX(110)
			pdf.CellFormat(40, defaultCellHeight*2, "Subtotal (Incl. tax)", "B", 0, "LM", false, 0, "")
			pdf.CellFormat(50, defaultCellHeight*2, "Rp5.000.000", "B", 1, "RM", false, 0, "")

			pdf.SetX(110)
			pdf.SetFont("Arial", "B", 11)
			pdf.CellFormat(40, defaultCellHeight*2, "Total due", "", 0, "LM", false, 0, "")
			pdf.CellFormat(50, defaultCellHeight*2, "Rp5.000.000", "", 1, "RM", false, 0, "")

			pdf.SetFooterFunc(func() {
				pdf.SetY(-15)
				pdf.SetFont("Arial", "", 8)
				pdf.SetTextColor(70, 70, 70)
				pdf.CellFormat(0, 10, fmt.Sprintf("Page %d of %d", pdf.PageNo(), numOfPages),
					"", 0, "L", false, 0, "")
				pdf.CellFormat(0, 10, "Powered by Meatball Realtime Subscription Engine",
					"", 0, "R", false, 0, "")
			})
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
