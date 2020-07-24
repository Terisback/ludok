package main

import (
	"image/color"
	"strings"

	"github.com/Terisback/ludok"
)

const (
	letterWidth = 8
	lineHeight  = 10
)

type Tooltip struct {
	padding       int
	elements      []string
	textColor     color.RGBA
	bgColor       color.RGBA
	textMaxLength int

	Width  int
	Height int
}

func NewTooltip(padding int, textColor color.RGBA, backgroundColor color.RGBA) *Tooltip {
	return &Tooltip{padding: padding, textColor: textColor, bgColor: backgroundColor}
}

func (t *Tooltip) Update(value string) {
	t.elements = strings.Split(value, "\n")
	t.textMaxLength = 0
	for _, v := range t.elements {
		length := len(v)
		if length >= t.textMaxLength {
			t.textMaxLength = length
		}
	}
	t.Width = (t.textMaxLength * letterWidth) + t.padding*2
	t.Height = (len(t.elements) * lineHeight) + t.padding*2 - 4
}

func (t *Tooltip) Draw(x, y int) {
	ludok.Graphics.Box(x, y, x+t.Width, y+t.Height, t.bgColor)
	for i, v := range t.elements {
		ludok.Graphics.Print(x+t.padding, y+t.padding+i*lineHeight, v, t.textColor)
	}
}
