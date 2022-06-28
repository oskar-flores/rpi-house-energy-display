package apps

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	epaper "github.com/oskar-flores/edp_2.13_V3"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"rpi-house-energy-display/domain/model"
	"time"
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

func Newwavesahre213Display() wavesahre213Display {
	registerFonts()
	waveSahreScreen := wavesahre213Display{
		width:   122,
		height:  250,
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

func (display *wavesahre213Display) Draw(lecture *model.EnergyLecture) {
	graphicContext := display.Context

	graphicContext.SetFillColor(image.Black)

	graphicContext.SetDPI(72) // 16 m3x6
	// gc.SetDPI(96) // 12 m3x6
	graphicContext.SetFontSize(16)

	graphicContext.SetFontData(draw2d.FontData{
		Name: "m3x6",
	})

	row := 6.0 + 4.0
	offset := -3.0

	graphicContext.SetFontSize(16)
	graphicContext.FillStringAt("active:", 1, 4*row+offset)
	graphicContext.SetFontSize(32)
	graphicContext.FillStringAt(lecture.LectureValue, 60, 4*row+offset)
	graphicContext.FillStringAt(lecture.LectureValue, 140, 4*row+offset)

	graphicContext.SetFontSize(16)
	graphicContext.FillStringAt("total:", 1, 5*row+offset+6)
	//graphicContext.FillStringAt(strconv.Itoa(stats.cases), 60, 5*row+offset+6)
	//graphicContext.FillStringAt(strconv.Itoa(stats.czCases), 140, 5*row+offset+6)
	//graphicContext.FillStringAt("recovered:", 1, 6*row+offset+6)
	//graphicContext.FillStringAt(strconv.Itoa(stats.recovered), 60, 6*row+offset+6)
	//graphicContext.FillStringAt(strconv.Itoa(stats.czRecovered), 140, 6*row+offset+6)
	//graphicContext.FillStringAt("deaths:", 1, 7*row+offset+6)
	//graphicContext.FillStringAt(strconv.Itoa(stats.deaths), 60, 7*row+offset+6)
	//graphicContext.FillStringAt(strconv.Itoa(stats.czDeaths), 140, 7*row+offset+6)
	//graphicContext.FillStringAt("(+"+strconv.Itoa(stats.czNew)+")", 140, 8*row+offset+6)
	//// if prev.cases > 0 {
	// 	gc.FillStringAt("(+"+strconv.Itoa(stats.cases-prev.cases)+")", 60, 8*row-2)
	// }

	graphicContext.FillStringAt("Last refreshed: "+lecture.LectureDate.Format(time.RFC3339), 1, 103)
	display.show()

}

func (display *wavesahre213Display) show() {
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

func (display *wavesahre213Display) Close() {
	display.Epd.Clear()
	display.Epd.TurnDisplayOff()
	display.Epd.Close()
}
