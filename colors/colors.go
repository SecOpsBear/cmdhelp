package colors

import (
	"runtime"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[1;31m"
	ColorGreen  = "\033[1;32m"
	ColorYellow = "\033[1;33m"
	ColorBlue   = "\033[1;34m"
	ColorPurple = "\033[1;35m"
	ColorCyan   = "\033[1;36m"
	ColorWhite  = "\033[1;37m"
)

type VarColors struct {
	ColorReset  string
	ColorRed    string
	ColorGreen  string
	ColorYellow string
	ColorBlue   string
	ColorPurple string
	ColorCyan   string
	ColorWhite  string
}

type Color struct {
	os  string
	col VarColors
}

func GetOSColors() *Color {
	var c Color
	osType := runtime.GOOS
	c.os = osType
	switch osType {
	case "windows", "darwin":
		c.col = VarColors{
			ColorReset:  "",
			ColorRed:    "",
			ColorGreen:  "",
			ColorYellow: "",
			ColorBlue:   "",
			ColorPurple: "",
			ColorCyan:   "",
			ColorWhite:  "",
		}
		return &c
	case "linux":
		ck := VarColors{
			ColorReset:  "\033[0m",
			ColorRed:    "\033[1;31m",
			ColorGreen:  "\033[1;32m",
			ColorYellow: "\033[1;33m",
			ColorBlue:   "\033[1;34m",
			ColorPurple: "\033[1;35m",
			ColorCyan:   "\033[1;36m",
			ColorWhite:  "\033[1;37m",
		}
		c.col = ck
		return &c
	default:
		c.col = VarColors{
			ColorReset:  "",
			ColorRed:    "",
			ColorGreen:  "",
			ColorYellow: "",
			ColorBlue:   "",
			ColorPurple: "",
			ColorCyan:   "",
			ColorWhite:  "",
		}
		return &c
	}

}

// ResetColor prefix's and sufix's the string with reset parameter (\033[0m)
func (c *Color) ResetColor(str string) string {
	newStr := string(c.col.ColorReset) + str + string(c.col.ColorReset)
	return newStr
}

// RedColor prefix's with \033[1;31m and sufix's the string with reset parameter (\033[0m)
func (c *Color) RedColor(str string) string {
	newStr := string(c.col.ColorRed) + str + string(c.col.ColorReset)
	return newStr
}

// GreenColor prefix's \033[1;32m and sufix's the string with reset parameter (\033[0m)
func (c *Color) GreenColor(str string) string {
	newStr := string(c.col.ColorGreen) + str + string(c.col.ColorReset)
	return newStr
}

// YellowColor prefix's \033[1;33m and sufix's the string with reset parameter (\033[0m)
func (c *Color) YellowColor(str string) string {
	newStr := string(c.col.ColorYellow) + str + string(c.col.ColorReset)
	return newStr
}

// BlueColor prefix's \033[1;34m and sufix's the string with reset parameter (\033[0m)
func (c *Color) BlueColor(str string) string {
	newStr := string(c.col.ColorBlue) + str + string(c.col.ColorReset)
	return newStr
}

// PurpleColor prefix's \033[1;35m and sufix's the string with reset parameter (\033[0m)
func (c *Color) PurpleColor(str string) string {
	newStr := string(c.col.ColorPurple) + str + string(c.col.ColorReset)
	return newStr
}

// CyanColor prefix's \033[1;36m and sufix's the string with reset parameter (\033[0m)
func (c *Color) CyanColor(str string) string {
	newStr := string(c.col.ColorCyan) + str + string(c.col.ColorReset)
	return newStr
}

// WhiteColor prefix's \033[1;37m and sufix's the string with reset parameter (\033[0m)
func (c *Color) WhiteColor(str string) string {
	newStr := string(c.col.ColorWhite) + str + string(c.col.ColorReset)
	return newStr
}

// func ResetColor(str string) string {
// 	newStr := string(ColorReset) + str + string(ColorReset)
// 	return newStr
// }

// func RedColor(str string) string {
// 	newStr := string(ColorRed) + str + string(ColorReset)
// 	return newStr
// }

// func GreenColor(str string) string {
// 	newStr := string(ColorGreen) + str + string(ColorReset)
// 	return newStr
// }

// func YellowColor(str string) string {
// 	newStr := string(ColorYellow) + str + string(ColorReset)
// 	return newStr
// }

// func BlueColor(str string) string {
// 	newStr := string(ColorBlue) + str + string(ColorReset)
// 	return newStr
// }

// func PurpleColor(str string) string {
// 	newStr := string(ColorPurple) + str + string(ColorReset)
// 	return newStr
// }

// func CyanColor(str string) string {
// 	newStr := string(ColorCyan) + str + string(ColorReset)
// 	return newStr
// }

// func WhiteColor(str string) string {
// 	newStr := string(ColorWhite) + str + string(ColorReset)
// 	return newStr
// }
