package fonts

import (
	"fmt"

	"eliasnaur.com/font/roboto/robotoblack"
	"eliasnaur.com/font/roboto/robotoblackitalic"
	"eliasnaur.com/font/roboto/robotobold"
	"eliasnaur.com/font/roboto/robotobolditalic"
	"eliasnaur.com/font/roboto/robotolight"
	"eliasnaur.com/font/roboto/robotolightitalic"
	"eliasnaur.com/font/roboto/robotomedium"
	"eliasnaur.com/font/roboto/robotomediumitalic"
	"eliasnaur.com/font/roboto/robotoregular"
	"eliasnaur.com/font/roboto/robotothin"
	"eliasnaur.com/font/roboto/robotothinitalic"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func RobotoTheme() *material.Theme {
	v := material.NewTheme()

	var faces []font.FontFace
	if regular, err := robotoFontFace(); err != nil {
		fmt.Println(err)
		return v
	} else {
		faces = append(faces, regular...)
	}
	if italic, err := robotoItalicFontFace(); err != nil {
		fmt.Println(err)
		return v
	} else {
		faces = append(faces, italic...)
	}

	v.Shaper = text.NewShaper(text.WithCollection(faces))

	return v
}

func robotoFontFace() ([]font.FontFace, error) {
	var faces []font.FontFace

	if x, err := opentype.ParseCollection(robotothin.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotolight.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotoregular.TTF); err != nil {
		return faces, err

	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotomedium.TTF); err != nil {
		return faces, err

	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotobold.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotoblack.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}

	return faces, nil
}

func robotoItalicFontFace() ([]font.FontFace, error) {
	var faces []font.FontFace

	if x, err := opentype.ParseCollection(robotothinitalic.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotolightitalic.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotomediumitalic.TTF); err != nil {
		return faces, err

	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotobolditalic.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}
	if x, err := opentype.ParseCollection(robotoblackitalic.TTF); err != nil {
		return faces, err
	} else {
		faces = append(faces, x...)
	}

	return faces, nil
}
