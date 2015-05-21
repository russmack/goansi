package goansi

import (
	"strconv"
	"strings"
)

type Anstring string

const (
	ESC             = "\x1b"    // Escape character ASCII decimal 27
	CSI             = ESC + "[" // Control Sequence Introducer
	SGRSuffix       = "m"       // Select Graphic Rendition
	SGRTmpl         = CSI + "%s" + SGRSuffix
	TextTmpl        = SGRTmpl + "%s" + AllOff
	AllOff          = "\x1b[0m" // Reset all attributes
	BoldIntensity   = 1         //
	Underline       = 4         //
	BlinkSlow       = 5         // less than 150 per minute
	ImageNegative   = 7         // inverse or reverse; swap foreground and background
	NormalIntensity = 22        // Neither bold nor faint
	// Set text color (foreground) = 30–37 ; 30 + x
	FgBlack   = 30
	FgRed     = 31
	FgGreen   = 32
	FgYellow  = 33
	FgBlue    = 34
	FgMagenta = 35
	FgCyan    = 36
	FgGray    = 37
	// Set background color	 = 	40–47 ; 40 + x
	BgBlack   = 40
	BgRed     = 41
	BgGreen   = 42
	BgYellow  = 43
	BgBlue    = 44
	BgMagenta = 45
	BgCyan    = 46
	BgGray    = 47
)

func NewAnstring(s string) Anstring {
	return Anstring(s)
}
func (s Anstring) Raw() string {
	return strings.Replace(string(s), "\x1b", "", -1)
}

func meld(code int, s Anstring) Anstring {
	return CSI + Anstring(strconv.Itoa(code)) + SGRSuffix + s + AllOff
}

func (s Anstring) Bold() Anstring {
	return meld(BoldIntensity, s)
}
func (s Anstring) Underline() Anstring {
	return meld(Underline, s)
}
func (s Anstring) Normal() Anstring {
	return meld(NormalIntensity, s)
}

func (s Anstring) Black() Anstring {
	return meld(FgBlack, s)
}
func (s Anstring) Red() Anstring {
	return meld(FgRed, s)
}
func (s Anstring) Green() Anstring {
	return meld(FgGreen, s)
}
func (s Anstring) Yellow() Anstring {
	return meld(FgYellow, s)
}
func (s Anstring) Blue() Anstring {
	return meld(FgBlue, s)
}
func (s Anstring) Magenta() Anstring {
	return meld(FgMagenta, s)
}
func (s Anstring) Cyan() Anstring {
	return meld(FgCyan, s)
}
func (s Anstring) Gray() Anstring {
	return meld(FgGray, s)
}

func (s Anstring) BgBlack() Anstring {
	return meld(BgBlack, s)
}
func (s Anstring) BgRed() Anstring {
	return meld(BgRed, s)
}
func (s Anstring) BgGreen() Anstring {
	return meld(BgGreen, s)
}
func (s Anstring) BgYellow() Anstring {
	return meld(BgYellow, s)
}
func (s Anstring) BgBlue() Anstring {
	return meld(BgBlue, s)
}
func (s Anstring) BgMagenta() Anstring {
	return meld(BgMagenta, s)
}
func (s Anstring) BgCyan() Anstring {
	return meld(BgCyan, s)
}
func (s Anstring) BgGray() Anstring {
	return meld(BgGray, s)
}

//func (s Anstring) () Anstring {
//	return meld(Fg, s)
//}

/* Omitted as poorly supported.

Reset / Normal	 = 	0	  // 	all attributes off
Faint (decreased intensity)	 = 	2	  // 	Not widely supported.
Italic: on	 = 	3	  // 	Not widely supported. Sometimes treated as inverse.
Blink: Rapid	 = 	6	  // 	MS-DOS ANSI.SYS; 150+ per minute; not widely supported
Conceal	 = 	8	  // 	Not widely supported.
Crossed-out	 = 	9	  // 	Characters legible, but marked for deletion. Not widely supported.
Primary(default) font	 = 	10	  //
n-th alternate font	 = 	11–19	  // 	Select the n-th alternate font (14 being the fourth alternate font, up to 19 being the 9th alternate font).
Fraktur	 = 	20	  // 	hardly ever supported
Bold: off or Underline: Double	 = 	21	  // 	Bold off not widely supported; double underline hardly ever supported.
Not italic, not Fraktur	 = 	23	  //
Underline: None	 = 	24	  // 	Not singly or doubly underlined
Blink: off	 = 	25	  //
Reserved	 = 	26	  //
Image: Positive	 = 	27	  //
Reveal	 = 	28	  // 	conceal off
Not crossed out	 = 	29	  //
Reserved for extended set foreground color	 = 	38	  // 	typical supported next arguments are 5;x where x is color index (0..255) or 2;r;g;b where r,g,b are red, green and blue color channels (out of 255)
Default text color (foreground)	 = 	39	  // 	implementation defined (according to standard)
Reserved for extended set background color	 = 	48	  // 	typical supported next arguments are 5;x where x is color index (0..255) or 2;r;g;b where r,g,b are red, green and blue color channels (out of 255)
Default background color	 = 	49	  // 	implementation defined (according to standard)
Reserved	 = 	50	  //
Framed	 = 	51	  //
Encircled	 = 	52	  //
Overlined	 = 	53	  //
Not framed or encircled	 = 	54	  //
Not overlined	 = 	55	  //
Reserved	 = 	56–59	  //
ideogram underline or right side line	 = 	60	  // 	hardly ever supported
ideogram double underline or double line on the right side	 = 	61	  // 	hardly ever supported
ideogram overline or left side line	 = 	62	  // 	hardly ever supported
ideogram double overline or double line on the left side	 = 	63	  // 	hardly ever supported
ideogram stress marking	 = 	64	  // 	hardly ever supported
ideogram attributes off	 = 	65	  // 	hardly ever supported, reset the effects of all of 60–64
Set foreground text color, high intensity	 = 	90–97	  // 	aixterm (not in standard)
Set background color, high intensity	 = 	100–107	  // 	aixterm (not in standard)

*/
