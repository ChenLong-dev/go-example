package main

import (
	"flag"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
)

var (
	k float64
	b float64
)

func init() {
	flag.Float64Var(&k, "k", 0, "int flag value")
	flag.Float64Var(&b, "b", 0, "int flag value")
}

func main() {
	flag.Parse()

	fmt.Println("k:", k)
	fmt.Println("b:", b)

	p := plot.New()

	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	kb := plotter.NewFunction(func(x float64) float64 { return k*x + b })
	kb.Color = plotutil.Color(0)

	//square := plotter.NewFunction(func(x float64) float64 { return x * x })
	//square.Color = plotutil.Color(0)
	//
	//sqrt := plotter.NewFunction(func(x float64) float64 { return 10 * math.Sqrt(x) })
	//sqrt.Dashes = []vg.Length{vg.Points(1), vg.Points(2)}
	//sqrt.Width = vg.Points(1)
	//sqrt.Color = plotutil.Color(1)
	//
	//exp := plotter.NewFunction(func(x float64) float64 { return math.Pow(2, x) })
	//exp.Dashes = []vg.Length{vg.Points(2), vg.Points(3)}
	//exp.Width = vg.Points(2)
	//exp.Color = plotutil.Color(2)
	//
	//sin := plotter.NewFunction(func(x float64) float64 { return 10*math.Sin(x) + 50 })
	//sin.Dashes = []vg.Length{vg.Points(3), vg.Points(4)}
	//sin.Width = vg.Points(3)
	//sin.Color = plotutil.Color(3)

	p.Add(kb)
	p.Legend.Add("kx+b", kb)
	//p.Add(square, sqrt, exp, sin)
	//p.Legend.Add("x^2", square)
	//p.Legend.Add("10*sqrt(x)", sqrt)
	//p.Legend.Add("2^x", exp)
	//p.Legend.Add("10*sin(x)+50", sin)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	p.X.Min = 0
	p.X.Max = 10
	p.Y.Min = 0
	p.Y.Max = 100

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "functions.png"); err != nil {
		log.Fatal(err)
	}
}
