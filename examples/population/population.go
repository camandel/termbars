package main

import (
	"github.com/camandel/termbars"
)

func main() {

    b := termbars.New()

	b.SetTitle("Population example")
	b.SetPercWidth(90)
	b.SetShowValues(true)

	b.Add("Tokyo", 37977000)
	b.Add("Jakarta", 34540000)
	b.Add("Delhi", 29617000)
	b.Add("Mumbai", 23355000)
	b.Add("Manila", 23088000)
	b.Add("Shanghai", 22120000)
	b.Add("Sao Paulo", 22046000)
	b.Add("Seoul", 21794000)
	b.Add("Mexico City", 20996000)

	b.Draw()
}
