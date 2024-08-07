package visualization

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"station/internal/model"
	"unicode"
)

// pixelFont defines a simple bitmap font for drawing text
var pixelFont = map[rune][][]bool{
	'0': {{true, true, true}, {true, false, true}, {true, false, true}, {true, false, true}, {true, true, true}},
	'1': {{false, true, false}, {true, true, false}, {false, true, false}, {false, true, false}, {true, true, true}},
	'2': {{true, true, true}, {false, false, true}, {true, true, true}, {true, false, false}, {true, true, true}},
	'3': {{true, true, true}, {false, false, true}, {false, true, true}, {false, false, true}, {true, true, true}},
	'4': {{true, false, true}, {true, false, true}, {true, true, true}, {false, false, true}, {false, false, true}},
	'5': {{true, true, true}, {true, false, false}, {true, true, true}, {false, false, true}, {true, true, true}},
	'6': {{true, true, true}, {true, false, false}, {true, true, true}, {true, false, true}, {true, true, true}},
	'7': {{true, true, true}, {false, false, true}, {false, true, false}, {true, false, false}, {true, false, false}},
	'8': {{true, true, true}, {true, false, true}, {true, true, true}, {true, false, true}, {true, true, true}},
	'9': {{true, true, true}, {true, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
	'A': {{false, true, true, false}, {true, false, false, true}, {true, true, true, true}, {true, false, false, true}, {true, false, false, true}},
	'B': {{true, true, true, false}, {true, false, false, true}, {true, true, true, false}, {true, false, false, true}, {true, true, true, false}},
	'C': {{false, true, true, true}, {true, false, false, false}, {true, false, false, false}, {true, false, false, false}, {false, true, true, true}},
	'D': {{true, true, true, false}, {true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {true, true, true, false}},
	'E': {{true, true, true, true}, {true, false, false, false}, {true, true, true, false}, {true, false, false, false}, {true, true, true, true}},
	'F': {{true, true, true, true}, {true, false, false, false}, {true, true, true, false}, {true, false, false, false}, {true, false, false, false}},
	'G': {{false, true, true, true}, {true, false, false, false}, {true, false, true, true}, {true, false, false, true}, {false, true, true, true}},
	'H': {{true, false, false, true}, {true, false, false, true}, {true, true, true, true}, {true, false, false, true}, {true, false, false, true}},
	'I': {{true, true, true}, {false, true, false}, {false, true, false}, {false, true, false}, {true, true, true}},
	'J': {{false, false, true, true}, {false, false, false, true}, {false, false, false, true}, {true, false, false, true}, {false, true, true, false}},
	'K': {{true, false, false, true}, {true, false, true, false}, {true, true, false, false}, {true, false, true, false}, {true, false, false, true}},
	'L': {{true, false, false, false}, {true, false, false, false}, {true, false, false, false}, {true, false, false, false}, {true, true, true, true}},
	'M': {{true, false, false, true}, {true, true, true, true}, {true, false, false, true}, {true, false, false, true}, {true, false, false, true}},
	'N': {{true, false, false, true}, {true, true, false, true}, {true, false, true, true}, {true, false, false, true}, {true, false, false, true}},
	'O': {{false, true, true, false}, {true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {false, true, true, false}},
	'P': {{true, true, true, false}, {true, false, false, true}, {true, true, true, false}, {true, false, false, false}, {true, false, false, false}},
	'Q': {{false, true, true, false}, {true, false, false, true}, {true, false, false, true}, {true, false, true, false}, {false, true, false, true}},
	'R': {{true, true, true, false}, {true, false, false, true}, {true, true, true, false}, {true, false, true, false}, {true, false, false, true}},
	'S': {{false, true, true, true}, {true, false, false, false}, {false, true, true, false}, {false, false, false, true}, {true, true, true, false}},
	'T': {{true, true, true, true, true}, {false, false, true, false, false}, {false, false, true, false, false}, {false, false, true, false, false}, {false, false, true, false, false}},
	'U': {{true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {false, true, true, false}},
	'V': {{true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {false, true, true, false}, {false, false, true, false}},
	'W': {{true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {true, false, false, true}, {false, true, true, false}},
	'X': {{true, false, false, true}, {false, true, true, false}, {false, false, true, false}, {false, true, true, false}, {true, false, false, true}},
	'Y': {{true, false, false, true}, {false, true, true, false}, {false, false, true, false}, {false, false, true, false}, {false, false, true, false}},
	'Z': {{true, true, true, true}, {false, false, true, false}, {false, true, false, false}, {true, false, false, false}, {true, true, true, true}},
	' ': {{false, false, false}, {false, false, false}, {false, false, false}, {false, false, false}, {false, false, false}},
	'_': {{false, false, false, false}, {false, false, false, false}, {false, false, false, false}, {false, false, false, false}, {true, true, true, true}},
}

// CreateVisualization generates a PNG image of the network and train paths
func CreateVisualization(stations map[string]*model.Station, paths [][]string) error {
	// Define initial canvas size and margins
	width, height := 1000, 800
	margin := 50

	// Calculate the bounding box of the network
	maxX, maxY := 0, 0
	for _, station := range stations {
		if station.X > maxX {
			maxX = station.X
		}
		if station.Y > maxY {
			maxY = station.Y
		}
	}

	// Calculate scaling factor and grid step size based on the bounding box
	scale := int(math.Min(float64(width-margin*2)/float64(maxX), float64(height-margin*2)/float64(maxY))) - 1
	gridStep := scale / 2

	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw grid and axes
	drawGrid(img, margin, width-margin, margin, height-margin, gridStep)
	drawAxes(img, margin, width-margin, margin, height-margin)

	// Draw stations
	for name, station := range stations {
		x, y := margin+station.X*scale, height-margin-station.Y*scale
		drawCircle(img, x, y, 5, color.RGBA{0, 0, 255, 255}) // Blue circle for stations

		// Draw station name
		nameColor := color.RGBA{255, 0, 0, 255} // Red color for station names
		drawLargeText(img, name, x+15, y-10, nameColor, 4)
	}

	// Draw connections between stations
	for _, station := range stations {
		for _, conn := range station.Connections {
			drawLine(img,
				margin+station.X*scale, height-margin-station.Y*scale,
				margin+conn.X*scale, height-margin-conn.Y*scale,
				color.RGBA{100, 100, 100, 255}) // Gray lines for connections
		}
	}

	// Draw paths with different colors
	colors := []color.RGBA{
		{255, 0, 0, 255},   // Red
		{0, 255, 0, 255},   // Green
		{255, 165, 0, 255}, // Orange
		{255, 0, 255, 255}, // Magenta
	}
	for i, path := range paths {
		pathColor := colors[i%len(colors)]
		for j := 1; j < len(path); j++ {
			start := stations[path[j-1]]
			end := stations[path[j]]
			drawLine(img,
				margin+start.X*scale, height-margin-start.Y*scale,
				margin+end.X*scale, height-margin-end.Y*scale,
				pathColor)
		}
	}

	// Save the image
	f, err := os.Create("network_visualization.png")
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, img)
	fmt.Println("Visualization saved as network_visualization.png")
	return nil
}

// drawGrid draws a grid on the image
func drawGrid(img *image.RGBA, left, right, top, bottom, step int) {
	lightGray := color.RGBA{200, 200, 200, 255}
	for x := left; x <= right; x += step {
		drawLine(img, x, top, x, bottom, lightGray)
	}
	for y := top; y <= bottom; y += step {
		drawLine(img, left, y, right, y, lightGray)
	}
}

// drawAxes draws the x and y axes on the image
func drawAxes(img *image.RGBA, left, right, top, bottom int) {
	black := color.RGBA{0, 0, 0, 255}
	drawLine(img, left, bottom, right, bottom, black) // X-axis
	drawLine(img, left, top, left, bottom, black)     // Y-axis
}

// drawCircle draws a filled circle on the image
func drawCircle(img *image.RGBA, x, y, r int, c color.RGBA) {
	for dx := -r; dx <= r; dx++ {
		for dy := -r; dy <= r; dy++ {
			if dx*dx+dy*dy <= r*r {
				img.Set(x+dx, y+dy, c)
			}
		}
	}
}

// drawLine draws a line on the image using Bresenham's line algorithm
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.RGBA) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x1, y1, c)
		if x1 == x2 && y1 == y2 {
			return
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// drawLargeText draws text on the image using the pixelFont
func drawLargeText(img *image.RGBA, text string, x, y int, c color.Color, size int) {
	bgColor := color.RGBA{0, 0, 255, 255}     // Blue background
	fgColor := color.RGBA{255, 255, 255, 255} // White text

	// Calculate text dimensions
	textWidth := len(text) * (4 + 1) * size // 4 pixels wide + 1 pixel space, multiplied by size
	textHeight := 5 * size                  // Each character is 5 pixels tall

	// Draw background
	for dy := 0; dy < textHeight+4; dy++ {
		for dx := 0; dx < textWidth+4; dx++ {
			img.Set(x+dx-2, y+dy-2, bgColor)
		}
	}

	// Draw text
	currentX := x
	for _, char := range text {
		upperChar := unicode.ToUpper(char) // Convert to uppercase
		if pixelMap, ok := pixelFont[upperChar]; ok {
			for i, row := range pixelMap {
				for j, pixel := range row {
					if pixel {
						for dy := 0; dy < size; dy++ {
							for dx := 0; dx < size; dx++ {
								img.Set(currentX+j*size+dx, y+i*size+dy, fgColor)
							}
						}
					}
				}
			}
		} else {
			// Draw a filled rectangle for unknown characters
			for i := 0; i < 5; i++ {
				for j := 0; j < 3; j++ {
					for dy := 0; dy < size; dy++ {
						for dx := 0; dx < size; dx++ {
							img.Set(currentX+j*size+dx, y+i*size+dy, fgColor)
						}
					}
				}
			}
		}
		currentX += (4 + 1) * size // Move to the next character position (4 pixels + 1 space)
	}
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
