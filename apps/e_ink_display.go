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
	outputModlels "rpi-house-energy-display/apps/model"
	"strconv"
	"time"
)

type Waveshare213Display struct {
	width   int
	height  int
	black   color.RGBA
	white   color.RGBA
	Display *image.RGBA
	Context *draw2dimg.GraphicContext
	Epd     *epaper.Epd
}

func Newwavesahre213Display() Waveshare213Display {
	registerFonts()
	screen := Waveshare213Display{
		width:   122, //y
		height:  250, //x
		Display: nil,
		black:   color.RGBA{A: 0xff},
		white:   color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		Context: nil,
	}
	screen.Display = image.NewRGBA(image.Rect(0, 0, screen.height, screen.width))
	screen.Context = draw2dimg.NewGraphicContext(screen.Display)

	screen.Context.SetFillColor(screen.white)
	screen.Context.Fill()
	epd := epaper.CreateEpd()
	screen.Epd = &epd
	screen.Epd.Init()

	return screen

}

func (display *Waveshare213Display) Draw(data outputModlels.DisplayModel) {
	graphicContext := display.Context

	//V3 driver buffer generation is priting  "mirror mode"
	//so we flip the image arong x axis, and then is correctly displayed
	//int the device.
	flipImage(graphicContext, display.width, 2)

	graphicContext.SetFillColor(image.Black)
	graphicContext.SetDPI(72) // 16 m3x6
	graphicContext.SetFontSize(16)

	graphicContext.SetFontData(draw2d.FontData{
		Name: "inconsolata",
	})

	row := 6.0 + 4.0
	offset := -3.0

	graphicContext.SetFontSize(16)
	graphicContext.FillStringAt("Current Measure:", 1, 4*row+offset)
	graphicContext.FillStringAt("Current Cost:", 1, 6*row+offset)
	graphicContext.FillStringAt(strconv.FormatFloat(data.CurrentLecture, 'f', -1, 32)+" watts", 125, 4*row+offset)
	graphicContext.FillStringAt(strconv.FormatFloat(data.CurrentPrice, 'f', -1, 32)+" E/w", 125, 6*row+offset)

	graphicContext.SetFontSize(16)
	graphicContext.FillStringAt("total:", 1, 8*row+offset+6)
	graphicContext.FillStringAt(strconv.FormatFloat(data.CurrentCostInPVC, 'f', -1, 32)+" Eur", 60, 8*row+offset+6)
	graphicContext.FillStringAt("Last refreshed: "+time.Now().Format(time.RFC3339), 1, 103)

	display.show()

}

func (display *Waveshare213Display) show() {
	draw2dimg.SaveToPngFile("display.png", display.Display)
	dataAsBuffer := display.Epd.GetBuffer(display.Display)
	display.Epd.Display(dataAsBuffer)
}

func drawRect(gc *draw2dimg.GraphicContext, x, y, width, height float64) {
	gc.BeginPath()
	gc.MoveTo(x, y)
	gc.LineTo(x+width, y)
	gc.LineTo(x+width, y+height)
	gc.LineTo(x, y+height)
	gc.Close()
}

func registerFonts() {
	inconsolata := parseFont("inconsolata")
	draw2d.RegisterFont(draw2d.FontData{
		Name: "inconsolata",
	}, inconsolata)

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

func (display *Waveshare213Display) Close() {
	display.Epd.Clear()
	display.Epd.TurnDisplayOff()
	display.Epd.Close()
}

// Flips the image around the X axis, and adjusts offset
func flipImage(gc draw2d.GraphicContext, width, offset int) {
	gc.Translate(float64(width*offset), 0)
	gc.Scale(-1, 1)
}
