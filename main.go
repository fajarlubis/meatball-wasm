package main

import (
	"log"
	"syscall/js"

	"github.com/fajarlubis/meatball-wasm/pdf"
)

func main() {
	p := pdf.NewInvoice(pdf.DefaultA4, pdf.Config{
		Orientation: pdf.Landscape,
	}, map[string]interface{}{})

	invoice, err := p.Generate()
	if err != nil {
		log.Fatal(err)
	}

	jsPDFContent := js.Global().Get("Uint8Array").New(len(invoice.Bytes()))
	js.CopyBytesToJS(jsPDFContent, invoice.Bytes())

	js.Global().Call("displayPDF", jsPDFContent, "output.pdf")
}
