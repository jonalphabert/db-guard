package logviewer

import (
	"github.com/fatih/color"
)

var levelColors = map[string]*color.Color{
	"DEBUG": color.New(color.FgHiBlack),
	"INFO":  color.New(color.FgHiBlue),
	"WARN":  color.New(color.FgHiYellow),
	"ERROR": color.New(color.FgHiRed),
	"SUCCESS": color.New(color.FgHiGreen),
}