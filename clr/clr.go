package clr

import "image/color"

func NRGBA(clr color.Color) color.NRGBA {
	r, g, b, a := clr.RGBA()
	return color.NRGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func RGBA(r, g, b uint8, a float32) color.NRGBA {
	if a > 1 {
		a = 1
	}
	if a < 0 {
		a = 0
	}
	return color.NRGBA{
		R: r,
		G: g,
		B: b,
		A: uint8(255 * a),
	}
}

func RGB(r, g, b uint8) color.NRGBA {
	return RGBA(r, g, b, 1)
}

func GRAYA(v uint8, a float32) color.NRGBA {
	return RGBA(v, v, v, a)
}

func GRAY(v uint8) color.NRGBA {
	return GRAYA(v, 1)
}
