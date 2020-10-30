package termbars

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/muesli/termenv"
	"golang.org/x/sys/unix"
)

const (
	Black  string = "0"
	Red    string = "1"
	Green  string = "2"
	Yellow string = "3"
	Blue   string = "4"
	Magent string = "5"
	Cyan   string = "6"
	White  string = "7"
)

type BarChartConfig struct {
	Title      string `json:"title"`
	PercWidth  int    `json:"percwidth"`
	ShowValues bool   `json:"showvalues"`
}

type barData struct {
	Label string `json:"label"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

type BarChart struct {
	config *BarChartConfig
	data   []barData
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Create new chart with default values
func New() *BarChart {
	b := new(BarChart)
	c := new(BarChartConfig)

	// Set defaults
	c.Title = ""
	c.PercWidth = 80
	c.ShowValues = true

	b.config = c
	b.data = []barData{}

	return b
}

// Create new chart based on json config
func NewConfig(config, data string) *BarChart {
	b := New()

	c := new(BarChartConfig)
	if err := json.Unmarshal([]byte(config), &c); err != nil {
		fmt.Println("Error unmarshal config")
	}

	if err := json.Unmarshal([]byte(data), &b.data); err != nil {
		fmt.Println("Error unmarshal data", err)
	}

	// Set config from json
	if c.Title != "" {
		b.config.Title = c.Title
	}
	if c.PercWidth != 0 {
		b.config.PercWidth = c.PercWidth
	}
	if c.ShowValues {
		b.config.ShowValues = true
	} else {
		b.config.ShowValues = false
	}

	return b
}

// Return the number of elements into the chart
func (b BarChart) Len() int {
	return len(b.data)
}

// Add new element to dataset (label and value)
func (b *BarChart) Add(label string, value int) {
	e := barData{label, value, getRndColor()}
	b.data = append(b.data, e)
}

// Add new element to dataset (label, value and color)
func (b *BarChart) AddColor(label string, value int, color string) {
	e := barData{label, value, color}
	b.data = append(b.data, e)
}

// Set chart title
func (b *BarChart) SetTitle(title string) {
	b.config.Title = title
}

// Set percentage of witdh screen to use
func (b *BarChart) SetPercWidth(perc int) {
	b.config.PercWidth = perc
}

// Show/hide values after label
func (b *BarChart) SetShowValues(showValues bool) {
	b.config.ShowValues = showValues
}

// Return max value
func (b BarChart) MaxValue() int {
	m := 0
	for _, v := range b.data {
		if v.Value > m {
			m = v.Value
		}
	}
	return m
}

// Return string lenght of max value
func (b BarChart) MaxValueLenght() int {
	return len(strconv.Itoa(b.MaxValue()))
}

// Return string lenght of largest label
func (b BarChart) MaxLabelLenght() int {
	m := 0
	for _, v := range b.data {
		if l := len(v.Label); l > m {
			m = l
		}
	}
	return m
}

// Return max label lenght
func (b BarChart) Draw() {

	p := termenv.ColorProfile()

	maxCols, err := getWidth()
	if err != nil {
		fmt.Println("error")
	}

	// Resize maxCols based on PercWith (default 75%)
	if b.config.PercWidth != 100 {
		maxCols = maxCols * b.config.PercWidth / 100
	}

	maxLabelLenght := b.MaxLabelLenght()
	maxValueLenght := b.MaxValueLenght()
	maxValue := b.MaxValue()

	if !b.config.ShowValues {
		maxValueLenght = 0
	}
	maxWidth := maxCols - maxLabelLenght - maxValueLenght - 3
	increment := float64(maxValue) / float64(maxWidth)

	if b.config.Title != "" {
		fmt.Printf("%s\n\n", b.config.Title)
	}

	for _, r := range b.data {
		bar_chunks := int((float64(r.Value) * 8 / increment) / 8)
		reminder := int(float64(r.Value)*8/increment) % 8

		var lastChunk rune

		if reminder != 0 {
			lastChunk = '█' + rune(8-reminder)
		}

		chunks := strings.Repeat("█", bar_chunks) + string(lastChunk)
		bar := termenv.String(chunks).Foreground(p.Color(r.Color))

		if b.config.ShowValues {
			padding := strings.Repeat(" ", maxLabelLenght-len(r.Label)+maxValueLenght-len(strconv.Itoa(r.Value)))
			fmt.Printf("%s: %s%d %s\n", r.Label, padding, r.Value, bar.String())
		} else {
			padding := strings.Repeat(" ", maxLabelLenght-len(r.Label))
			fmt.Printf("%s %s %s\n", r.Label, padding, bar.String())
		}
	}
}

// Return terminal width (number of columns)
func getWidth() (int, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, err
	}

	return int(ws.Col), nil
}

// Get a random color
func getRndColor() string {

	var min, max int

	switch p := termenv.ColorProfile(); p {
	case termenv.Ascii:
		min = 1
		max = 1
	case termenv.ANSI:
		min = 1
		max = 15
	default:
		min = 17
		max = 231
	}

	return fmt.Sprintf("%d", rand.Intn(max-min)+min)
}
