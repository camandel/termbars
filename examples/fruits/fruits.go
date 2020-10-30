package main

import (
	"fmt"

	"github.com/camandel/termbars"
)

func main() {

	config := `{
					"title":     "Fruits example",
					"percwidth":  50,
					"showvalues": false
			   }`

	data := ` [
					{ 
						"label": "cherries",
						"value": 1,
						"color": "1"
						  
					},
					{ 
						"label": "apples",
						"value": 3,
						"color": "2"
					},
					{ 
						"label": "bananas",
						"value": 2,
						"color": "3"
					},
					{ 
						"label": "blueberries",
						"value": 5,
						"color": "4"
					}
   			 ]`

	b := termbars.NewConfig(config, data)
	b.Draw()
}
