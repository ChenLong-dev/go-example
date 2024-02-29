//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vugu/vugu/simplehttp"
)

//func main() {
//	l := "127.0.0.1:8844"
//	log.Printf("Starting HTTP Server at %q", l)
//
//	wc := devutil.NewWasmCompiler().SetDir(".")
//	mux := devutil.NewMux()
//	mux.Match(devutil.NoFileExt, devutil.DefaultAutoReloadIndex.Replace(
//		`<!-- styles -->`,
//		`<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">`))
//	mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
//	mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
//	mux.Default(devutil.NewFileServer().SetDir("."))
//
//	log.Fatal(http.ListenAndServe(l, mux))
//}

// go run devserver.go
func main() {
	vuguDir := "./vugufile/welcometovugu"
	//vuguDir := "./vugufile/vgif"
	//vuguDir := "./vugufile/vgattr"

	wd, _ := os.Getwd()
	dir := filepath.Join(wd, vuguDir)
	log.Println("exec dir:", dir)
	l := ":8844"
	log.Printf("Starting HTTP Server at %q", l)
	h := simplehttp.New(dir, true)

	log.Fatal(http.ListenAndServe(l, h))
}
