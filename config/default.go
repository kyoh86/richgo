package config

import "github.com/wacul/ptr"

var defaultConfig Config

func init() {
	var long = LabelTypeLong
	defaultConfig = Config{
		LabelType: &long,
		BuildStyle: &Style{
			Bold: ptr.Bool(true),
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Yellow,
			},
		},
		StartStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: LightBlack,
			},
		},
		PassStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Green,
			},
		},
		FailStyle: &Style{
			Bold: ptr.Bool(true),
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Red,
			},
		},
		SkipStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: LightBlack,
			},
		},
		CoverThreshold: ptr.Int(50),
		CoveredStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Green,
			},
		},
		UncoveredStyle: &Style{
			Bold: ptr.Bool(true),
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Yellow,
			},
		},
		FileStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Cyan,
			},
		},
		LineStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Magenta,
			},
		},
		PassPackageStyle: &Style{
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Green,
			},
			Hide: ptr.True(),
		},
		FailPackageStyle: &Style{
			Hide: ptr.True(),
			Bold: ptr.Bool(true),
			Foreground: &Color{
				Type: ColorTypeName,
				Name: Red,
			},
		},
	}
}
