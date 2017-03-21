package goi

import "strings"
import "fmt"

type Color struct {
	R byte
	G byte
	B byte
	A byte
}

func Clr(R byte, G byte, B byte, A byte) Color {
	return Color{
		R: R,
		G: G,
		B: B,
		A: A,
	}
}

func Transparent() Color {
	return Clr(0, 0, 0, 0)
}

func ParseColor(s string) Color {
	if strings.HasPrefix(s, "#") {
		if len(s) == 4 {
			return Clr(getColorBytes(s[1], s[1]), getColorBytes(s[2], s[2]),
				getColorBytes(s[3], s[3]), 255)
		} else if len(s) == 5 {
			return Clr(getColorBytes(s[1], s[1]), getColorBytes(s[2], s[2]),
				getColorBytes(s[3], s[3]), getColorBytes(s[4], s[4]))
		} else if len(s) == 7 {
			return Clr(getColorBytes(s[1], s[2]), getColorBytes(s[3], s[4]),
				getColorBytes(s[5], s[6]), 255)
		} else if len(s) == 9 {
			return Clr(getColorBytes(s[1], s[2]), getColorBytes(s[3], s[4]),
				getColorBytes(s[5], s[6]), getColorBytes(s[7], s[8]))
		}
	}

	return Transparent()
}

func getColorByte(c byte) byte {
	fmt.Printf(":%v\n", c)
	if c >= '0' && c <= '9' {
		return c - '0'
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	}
	return 0
}

func getColorBytes(c1 byte, c2 byte) byte {
	fmt.Printf("%v %v\n", c1, c2)
	return (getColorByte(c1) << 4) + getColorByte(c2)
}
