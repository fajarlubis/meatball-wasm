package main

import (
	"fmt"
	"log"
	"syscall/js"
	"time"

	"github.com/fajarlubis/meatball-wasm/pdf"
)

func main() {
	startTime := time.Now()
	numOfPages := 1000

	data := make([]map[string]interface{}, 0)
	for i := 0; i < numOfPages; i++ {
		data = append(data, map[string]interface{}{
			"invoice_number": "TEST INVOICE NUMBER",
			"issue_date":     "Feb 13, 2024",
			"payment_term":   "10 days",
		})
	}

	invoice := pdf.NewInvoice(pdf.DefaultA4, pdf.Config{
		Orientation: pdf.Portrait,
		PageNumber:  true,
		ForceOrder:  true,
	}, data...)

	pdfBuf, err := invoice.Generate()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Time to render %v PDF pages: %s\n", numOfPages, time.Since(startTime))

	jsPDFContent := js.Global().Get("Uint8Array").New(len(pdfBuf.Bytes()))
	js.CopyBytesToJS(jsPDFContent, pdfBuf.Bytes())

	js.Global().Call("displayPDF", jsPDFContent, "output.pdf")

	div := js.Global().Get("document").Call("getElementById", "msg")
	div.Set("textContent", fmt.Sprintf("Time to render %v PDF pages: %s\n", numOfPages, time.Since(startTime)))
}
