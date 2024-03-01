compile:
	GOOS=js GOARCH=wasm go build -o meatball.wasm

copy:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

serve:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'

preview:
	GOOS=js GOARCH=wasm go build -o meatball.wasm && goexec 'http.ListenAndServe("127.0.0.1:8080", http.FileServer(http.Dir(".")))'

cpu-pprof:
	go tool pprof -http=:8080 -seconds=30 pdf/cpu.pprof

mem-pprof:
	go tool pprof -http=:8080 -seconds=30 pdf/mem.pprof

.PHONY: compile copy serve preview cpu-pprof mem-pprof