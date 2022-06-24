package apps

import (
	"fmt"
	"github.com/bestbug456/epaper"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"rpi-house-energy-display/domain/model"
)

type wavesahre213Display struct {
	width   int
	height  int
	black   color.RGBA
	white   color.RGBA
	Display *image.RGBA
	Context *draw2dimg.GraphicContext
	Epd     *epaper.Epd
}

func (display *wavesahre213Display) Newwavesahre213Display() wavesahre213Display {
	registerFonts()
	waveSahreScreen := wavesahre213Display{
		width:   104,
		height:  212,
		Display: nil,
		black:   color.RGBA{A: 0xff},
		white:   color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		Context: nil,
	}
	waveSahreScreen.Display = image.NewRGBA(image.Rect(0, 0, waveSahreScreen.height, waveSahreScreen.width))
	waveSahreScreen.Context = draw2dimg.NewGraphicContext(waveSahreScreen.Display)

	waveSahreScreen.Context.SetFillColor(waveSahreScreen.white)
	waveSahreScreen.Context.Fill()
	epd := epaper.CreateEpd()
	waveSahreScreen.Epd = &epd
	waveSahreScreen.Epd.Init()

	return waveSahreScreen

}

func (display *wavesahre213Display) Draw(lecture model.EnergyLecture) {
	graphicContext := display.Context
	graphicContext.SetFillColor(image.Black)

	graphicContext.SetDPI(72) // 16 m3x6
	// gc.SetDPI(96) // 12 m3x6
	graphicContext.SetFontSize(16)

	graphicContext.SetFontData(draw2d.FontData{
		Name: "m3x6",
	})

	row := 6.0 + 4.0

	graphicContext.FillStringAt(lecture.LectureValue, 1, row-3)
	display.show(lecture)

}

func (display *wavesahre213Display) show(lecture model.EnergyLecture) {
	dataToshow := display.Epd.GetBuffer(display.Display)
	display.Epd.Display(dataToshow)
}

func drawRect(gc *draw2dimg.GraphicContext, x, y, w, h float64) {
	gc.BeginPath()
	gc.MoveTo(x, y)
	gc.LineTo(x+w, y)
	gc.LineTo(x+w, y+h)
	gc.LineTo(x, y+h)
	gc.Close()
}

func registerFonts() {
	m3x6Font := parseFont("m3x6")
	draw2d.RegisterFont(draw2d.FontData{
		Name: "m3x6",
	}, m3x6Font)

	fmt.Println("fonts registered")
}

func parseFont(name string) (f *truetype.Font) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if dir[0:4] == "/tmp" {
		dir = "."
	}
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/assets/fonts/%s.ttf", dir, name))
	if err != nil {
		return nil
	}
	f, err = truetype.Parse(b)
	if err != nil {
		return nil
	}

	return f
}
