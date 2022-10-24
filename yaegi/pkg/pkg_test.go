package pkg

import (
	"bytes"
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"os"
	"path/filepath"
	"testing"
)

func fatalStderrf(t *testing.T, format string, args ...interface{}) {
	t.Helper()

	fmt.Fprintf(os.Stderr, format+"\n", args...)
	t.FailNow()
}

func TestPackages_pkg(t *testing.T) {
	desc := "vendor"
	goPath := "./_pkg/"
	expected := "root Fromage"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg0(t *testing.T) {
	desc := "sub-subpackage"
	goPath := "./_pkg0/"
	expected := "root Fromage Cheese"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg1(t *testing.T) {
	desc := "subpackage"
	goPath := "./_pkg1/"
	expected := "root Fromage!"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg2(t *testing.T) {
	desc := "multiple vendor folders"
	goPath := "./_pkg2/"
	expected := "root Fromage Cheese!"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg3(t *testing.T) {
	desc := "multiple vendor folders and subpackage in vendor"
	goPath := "./_pkg3/"
	expected := "root Fromage Couteau Cheese!"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg4(t *testing.T) {
	desc := "multiple vendor folders and multiple subpackages in vendor"
	goPath := "./_pkg4/"
	expected := "root Fromage Cheese Vin! Couteau"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg5(t *testing.T) {
	desc := "vendor flat"
	goPath := "./_pkg5/"
	expected := "root Fromage Cheese Vin! Couteau"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg6(t *testing.T) {
	desc := "fallback to GOPATH"
	goPath := "./_pkg6/"
	expected := "root Fromage Cheese Vin! Couteau"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg7(t *testing.T) {
	desc := "recursive vendor"
	goPath := "./_pkg7/"
	expected := "root vin cheese fromage"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg8(t *testing.T) {
	desc := "named subpackage"
	goPath := "./_pkg8/"
	expected := "root Fromage!"
	topImport := "github.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg9(t *testing.T) {
	desc := "xx different packages in the same directory xx"
	goPath := "./_pkg9/"
	expected := `1:21: import "github.com/foo/pkg" error: found packages pkg and pkgfalse in ` + filepath.FromSlash("_pkg9/src/github.com/foo/pkg")
	topImport := "github.com/foo/pkg"

	i := interp.New(interp.Options{GoPath: goPath})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	_, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport))
	if err == nil {
		t.Fatalf("got no error, want %q", expected)
	}

	//_, err := i.Eval(`pkg.NewSample()`)
	//if err == nil {
	//	t.Fatal(err)
	//}

	//fn := value.Interface().(func() string)
	//
	//msg := fn()
	t.Logf("[gopath:%s][desc:%s] err:%s", goPath, desc, err.Error())
	if err.Error() != expected {
		t.Errorf("got %q, want %q", err.Error(), expected)
	}
}

func TestPackages_pkg10(t *testing.T) {
	desc := "at the project root"
	goPath := "./_pkg10/"
	expected := "root Fromage"
	topImport := "github.com/foo"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg11(t *testing.T) {
	desc := "eval main that has vendored dep"
	goPath := "./_pkg11/"
	expected := "Fromage"
	//topImport := "github.com/foo/pkg"
	evalFile := "./_pkg11/src/foo/foo.go"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	//if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
	//	t.Fatal(err)
	//}
	//
	//value, err := i.Eval(`pkg.NewSample()`)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fn := value.Interface().(func() string)
	//
	//msg := fn()

	if _, err := i.EvalPath(filepath.FromSlash(evalFile)); err != nil {
		fatalStderrf(t, "%v", err)
	}
	msg := stdout.String()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg12(t *testing.T) {
	desc := "vendor dir is a sibling or an uncle"
	goPath := "./_pkg12/"
	expected := "Yo hello"
	topImport := "guthib.com/foo/pkg"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`pkg.NewSample()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)

	msg := fn()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg121(t *testing.T) {
	desc := "eval main with vendor as a sibling"
	goPath := "./_pkg12/"
	expected := "Yo hello"
	//topImport := "github.com/foo/pkg"
	evalFile := "./_pkg12/src/guthib.com/foo/main.go"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.EvalPath(filepath.FromSlash(evalFile)); err != nil {
		fatalStderrf(t, "%v", err)
	}
	msg := stdout.String()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg13(t *testing.T) {
	desc := "eval main with vendor"
	goPath := "./_pkg13/"
	expected := "foobar"
	//topImport := "github.com/foo/pkg"
	evalFile := "./_pkg13/src/guthib.com/foo/bar/main.go"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	path := filepath.FromSlash(evalFile)
	if _, err := i.EvalPath(path); err != nil {
		fatalStderrf(t, "%v, path:%s", err, path)
	}
	msg := stdout.String()
	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}
}

func TestPackages_pkg14(t *testing.T) {
	desc := "==== chen long ===="
	goPath := "./_pkg14/"
	expected := "new a file plugin --> new a file"
	topImport := "plugins"

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{GoPath: goPath, Stdout: &stdout, Stderr: &stderr})
	// Use binary standard library
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(fmt.Sprintf(`import "%s"`, topImport)); err != nil {
		t.Fatal(err)
	}

	value, err := i.Eval(`plugins.NewPlugin()`)
	if err != nil {
		t.Fatal(err)
	}

	fn := value.Interface().(func() string)
	msg := fn()

	t.Logf("[gopath:%s][desc:%s] msg:%s", goPath, desc, msg)
	if msg != expected {
		fatalStderrf(t, "Got %q, want %q", msg, expected)
	}

}
