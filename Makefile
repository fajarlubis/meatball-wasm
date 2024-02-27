compile:
	GOOS=js GOARCH=wasm go build -o meatball.wasm

copy:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

serve:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'

.PHONY: compile copy serve