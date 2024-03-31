package clr

import "image/color"

// check

func IsTransparent(clr color.NRGBA) bool {
	return clr.A <= 0
}

func IsTranslucent(clr color.NRGBA) bool {
	return clr.A < 255
}

func IsEmpty(clr color.NRGBA) bool {
	return clr.R <= 0 && clr.G <= 0 && clr.B <= 0 && clr.A <= 0
}

func IsBlack(clr color.NRGBA) bool {
	return clr.R <= 0 && clr.G <= 0 && clr.B <= 0 && clr.A > 0
}

func IsWhite(clr color.NRGBA) bool {
	return clr.R >= 255 && clr.G >= 255 && clr.B >= 255 && clr.A > 0
}

// conditional

func Or(target, dest color.NRGBA) color.NRGBA {
	if IsEmpty(target) {
		return dest
	}
	return target
}

// combining

func SetOpacity(target color.NRGBA, opacity float32) color.NRGBA {
	target.A = uint8(255 * opacity)
	return target
}

func AddOpacity(target color.NRGBA, opacity float32) color.NRGBA {
	target.A += uint8(255 * opacity)
	return target
}
