package pixelmator

type ColorAdjustment int

const (
	Noop ColorAdjustment = iota
	Exposure
	Tint
	Hue
	Saturation
	Temperature
	Highlights
	Shadows
	Contrast
	BlackPoint
	Fade
	Vignette
	VignetteExposure
	VignetteBlackPoint
	VignetteSoftness
)

func (ca ColorAdjustment) getTerm() term {
	switch ca {
	case Exposure:
		return term{"exposureValue", "exposure"}

	case Tint:
		return term{"tintValue", "tint"}

	case Hue:
		return term{"hueValue", "hue"}

	case Saturation:
		return term{"saturationValue", "saturation"}

	case Temperature:
		return term{"temperatureValue", "temperature"}

	case Highlights:
		return term{"highlightsValue", "highlights"}

	case Shadows:
		return term{"shadowsValue", "shadows"}

	case Contrast:
		return term{"contrastValue", "contrast"}

	case BlackPoint:
		return term{"blackpointValue", "black point"}

	case Fade:
		return term{"fadeValue", "fade"}

	case Vignette:
		return term{"vignetteApplied", "vignette"}

	case VignetteExposure:
		return term{"vignetteExposureValue", "vignette exposure"}

	case VignetteBlackPoint:
		return term{"vignetteBlackpointValue", "vignette black point"}

	case VignetteSoftness:
		return term{"vignetteSoftnessValue", "vignette softness"}

	default:
		return term{"", ""}
	}
}

type term struct {
	osascriptTerm     string
	osascriptVariable string
}

type ColorAdjustmentValue interface {
	getTerms() []term
	setValues(map[string]any) error
}

type rangeInterface interface{}

type rangeValue struct {
	Value      int
	MinOfRange int
	MaxOfRange int
	rangeInterface
}

type rangeGroup struct {
	IsApplied bool
	Group     map[ColorAdjustment]rangeInterface
}
