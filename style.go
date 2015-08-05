package goansi

import ()

// Style is a list of style codes to be applied to a string.
type Style []uint8

// NewStyle creates a new Style list.
func NewStyle() Style {
	return Style{}
}

// Render applies a list of styles to the provided string.
func (s Style) Render(raw string) string {
	t := NewAnstring(raw)
	for i := len(s) - 1; i >= 0; i-- {
		t.meld(s[i])
	}
	return string(t)
}

// Text Decoration.

func (s Style) Bold() Style {
	return append(s, BoldIntensity)
}
func (s Style) Underline() Style {
	return append(s, Underline)
}
func (s Style) Normal() Style {
	return append(s, NormalIntensity)
}
func (s Style) BlinkSlow() Style {
	return append(s, BlinkSlow)
}
func (s Style) Negative() Style {
	return append(s, ImageNegative)
}

// Foreground Colours.

// Black adds the black foreground style to the style list.
func (s Style) Black() Style {
	return append(s, FgBlack)
}
func (s Style) Red() Style {
	return append(s, FgRed)
}
func (s Style) Green() Style {
	return append(s, FgGreen)
}
func (s Style) Yellow() Style {
	return append(s, FgYellow)
}
func (s Style) Blue() Style {
	return append(s, FgBlue)
}
func (s Style) Magenta() Style {
	return append(s, FgMagenta)
}
func (s Style) Cyan() Style {
	return append(s, FgCyan)
}
func (s Style) Gray() Style {
	return append(s, FgGray)
}

// Background Colours.

func (s Style) BgBlack() Style {
	return append(s, BgBlack)
}
func (s Style) BgRed() Style {
	return append(s, BgRed)
}
func (s Style) BgGreen() Style {
	return append(s, BgGreen)
}
func (s Style) BgYellow() Style {
	return append(s, BgYellow)
}
func (s Style) BgBlue() Style {
	return append(s, BgBlue)
}
func (s Style) BgMagenta() Style {
	return append(s, BgMagenta)
}
func (s Style) BgCyan() Style {
	return append(s, BgCyan)
}
func (s Style) BgGray() Style {
	return append(s, BgGray)
}
