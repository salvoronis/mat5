package main

import(
	"fmt"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
	"time"
	"strconv"
)

func createWindow(){
	a := app.App()
	scene := core.NewNode()
	gui.Manager().Set(scene)
	chart := gui.NewChart(0, 0)
	chart.SetMargins(10, 10, 10, 10)
	chart.SetBorders(2, 2, 2, 2)
	chart.SetBordersColor(math32.NewColor("green"))
	chart.SetColor(math32.NewColor("white"))
	chart.SetTitle("Adams", 16)
	chart.SetPosition(0, 0)
	width, height := a.GetSize()
	chart.SetSize(float32(width), float32(height)-100)
	chart.SetScaleY(5, &math32.Color{0.8, 0.8, 0.8})
	chart.SetFontSizeY(13)
	chart.SetRangeY(-5.0,5.0)
	chart.SetScaleX(5, &math32.Color{0.8, 0.8, 0.8})
	chart.SetFontSizeX(13)
	chart.SetRangeX(0.0, 14.0, 70.0)
	scene.Add(chart)

	dd1 := gui.NewDropDown(100, gui.NewImageLabel("func"))
	dd1.SetPosition(10, float32(height)-100)
	dd1.Add(gui.NewImageLabel("x^2-2y"))
	dd1.Add(gui.NewImageLabel("sin(x)"))
	dd1.Add(gui.NewImageLabel("Log(x+1)"))
	dd1.Add(gui.NewImageLabel("x+y"))
	scene.Add(dd1)

	ed1 := gui.NewEdit(100, "error")
	ed1.SetPosition(10, float32(height)-60)
	scene.Add(ed1)

	ed2 := gui.NewEdit(100, "right border")
	ed2.SetPosition(10, float32(height)-80)
	scene.Add(ed2)

	edx := gui.NewEdit(50, "x0")
	edx.SetPosition(10, float32(height)-40)
	scene.Add(edx)

	edy := gui.NewEdit(50, "y0")
	edy.SetPosition(60, float32(height)-40)
	scene.Add(edy)

	var g3 *gui.Graph
	var gt *gui.Graph
	var dots []*gui.Image
	var setted bool = false

	b1 := gui.NewButton("find y")
	b1.SetPosition(120, float32(height)-100)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		if setted {
			chart.RemoveGraph(g3)
			g3 = nil
			chart.RemoveGraph(gt)
			gt = nil
			for i := 0; i < len(dots); i++ {
				scene.Remove(dots[i])
			}
			setted = false
		}
		var f func(float32, float32) float32
		var ft func(float32) float32
		switch dd1.Selected().Text(){
		case "x^2-2y":
			f = one
			ft = one_true
		case "sin(x)":
			f = two
			ft = two_true
		case "Log(x+1)":
			f = three
			ft = three_true
		case "x+y":
			f = four
			ft = four_true
		}
		error, err := strconv.ParseFloat(ed1.Text(),32)
		if err != nil {
			return
		}
		x0, err := strconv.ParseFloat(edx.Text(),32)
		if err != nil {
			return
		}
		y0, err := strconv.ParseFloat(edy.Text(),32)
		if err != nil {
			return
		}
		xn, err := strconv.ParseFloat(ed2.Text(),32)
		if err != nil {
			return
		}
		step := float32(xn/5)
		chart.SetRangeX(0.0, step, float32(xn))
		startPoint := point{x:float32(x0), y:float32(y0)}
		data := eiler(f, startPoint, float32(error*3), float32(error))
		data = Adams(f, data, float32(xn), float32(error))
		dlag := make([]float32, 0)
		truegraph := make([]float32,0)
		for i := 0.0; i < xn; i+=0.2 {
			lagrange := InterpolateLagrangePolynomial(float32(i), len(data), data)
			dlag = append(dlag, lagrange)
			truegraph = append(truegraph, ft(float32(i)))
		}

		g3 = chart.AddLineGraph(&math32.Color{0, 0, 1}, dlag)
		gt = chart.AddLineGraph(&math32.Color{0, 1, 0}, truegraph)
		setted = true

		dots = make([]*gui.Image,0)
		if len(data) != 0 {
			for i := 0; i < len(data); i++ {
				dot, err := gui.NewImage("./dot.png")
				if err != nil {
					fmt.Println("ooops")
				}
				mi,ma := chart.RangeY()
				celly := 105/((math32.Abs(mi)+math32.Abs(ma))/4)
				cellx := float32(146/xn)*5
				dot.SetPosition(49+data[i].x*cellx,251-data[i].y*celly)
				if 49 < (49+data[i].x*cellx) && (49+data[i].x*cellx) < 749 && 41 < (251-data[i].y*celly) && (251-data[i].y*celly) < 461 {
					dots = append(dots, dot)
				}
			}
			for i := 0; i < len(dots); i++ {
				scene.Add(dots[i])
			}
		}
	})
	scene.Add(b1)

	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)

	onResize := func(evname string, ev interface{}) {
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	a.Gls().ClearColor(0.2, 0.7, 0.9, 10.0)

	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}