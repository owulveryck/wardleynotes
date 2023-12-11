package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type color = string
type element = string

func main() {
	sets := make([]map[element]color, 0)
	dec := json.NewDecoder(os.Stdin)
	var err error
	for err == nil {
		var content map[element]color
		err = dec.Decode(&content)
		if err != nil {
			continue
		}
		sets = append(sets, content)
	}
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	// color[0]: nocolor, color[1]: red, color[2]: green, color[3] c
	const (
		nocolor = 0
		red     = 1
		orange  = 2
		green   = 3
	)
	colorCount := make(map[element][]int, 0)
	for _, s := range sets {
		for k, v := range s {
			if _, ok := colorCount[k]; !ok {
				colorCount[k] = make([]int, 4)
			}
			switch v {
			case "green":
				colorCount[k][green]++
			case "red":
				colorCount[k][red]++
			case "orange":
				colorCount[k][orange]++
			case "nocolor":
				colorCount[k][nocolor]++
			default:
				log.Println("unknown color", v)
			}
		}
	}
	// compute average colors
	average := make(map[element]color, len(colorCount))
	for k, v := range colorCount {
		numberOfDeclaredColors := v[red] + v[orange] + v[green]
		if numberOfDeclaredColors != 0 {
			redComponent := (v[orange]*255 + v[red]*255) / numberOfDeclaredColors
			greenComponent := (v[green]*255 + v[orange]*165) / numberOfDeclaredColors
			blueComponent := 0
			average[k] = fmt.Sprintf("rgb(%v,%v,%v)", redComponent, greenComponent, blueComponent)
		} else {
			average[k] = "grey"
		}
	}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(average)
}
