package main

import (
	"log"
	"syscall/js"

	"github.com/fajarlubis/meatball-wasm/pdf"
)

func main() {
	data := make([]map[string]interface{}, 0)
	for i := 0; i < 1000; i++ {
		data = append(data, map[string]interface{}{
			"invoice_number": "TEST INVOICE NUMBER",
			"issue_date":     "Feb 13, 2024",
			"payment_term":   "10 days",
		})
	}

	invoice := pdf.NewInvoice(pdf.DefaultA4, pdf.Config{
		Orientation: pdf.Landscape,
		ForceOrder:  true,
	}, data...)

	pdfBuf, err := invoice.Generate()
	if err != nil {
		log.Fatal(err)
	}

	jsPDFContent := js.Global().Get("Uint8Array").New(len(pdfBuf.Bytes()))
	js.CopyBytesToJS(jsPDFContent, pdfBuf.Bytes())

	js.Global().Call("displayPDF", jsPDFContent, "output.pdf")
}
