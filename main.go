package main

import (
	"syscall/js"

	"github.com/fajarlubis/meatball-wasm/pdf"
)

func main() {
	w := pdf.Generate()

	jsPDFContent := js.Global().Get("Uint8Array").New(len(w.Bytes()))
	js.CopyBytesToJS(jsPDFContent, w.Bytes())

	js.Global().Call("displayPDF", jsPDFContent, "output.pdf")
}
