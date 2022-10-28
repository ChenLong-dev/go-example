package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

func fatalStderrf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

type T2 struct {
	C string
	D int
}

func process() {
	desc := "==== chen long ===="
	goPath := "../../service/plugins/"
	//expected := "new a file plugin --> new a file"
	topImport := "fileplugin"

	logx.Info(desc)
	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		fmt.Println(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		fmt.Println(err)
	}

	value, err := i.Eval(`fileplugin.NewPlugin()`)
	if err != nil {
		fmt.Println(err)
	}

	fn := value.Interface().(func() interface{})
	msg := fn()

	resByre, resByteErr := json.Marshal(msg)
	if resByteErr != nil {
		fmt.Printf("%v", resByteErr)
		return
	}
	var newData T2
	jsonRes := json.Unmarshal(resByre, &newData)
	if jsonRes != nil {
		fmt.Printf("%v", jsonRes)
		return
	}

	fmt.Printf("[gopath:%s][desc:%s] newData.C:%s, newData.D:%d\n", goPath, desc, newData.C, newData.D)
	//if msg != expected {
	//	fatalStderrf("Got %q, want %q", msg, expected)
	//}

}

func main() {
	for {
		process()
		time.Sleep(time.Second * 15)
		fmt.Println("--------------------------------------------------")
	}
}
