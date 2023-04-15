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

func (ca ColorAdjustment) getChildren() []ColorAdjustment {
	switch ca {
	case Vignette:
		return []ColorAdjustment{VignetteExposure, VignetteBlackPoint, VignetteSoftness}
	default:
		return nil
	}
}

func (ca ColorAdjustment) getRange() (int, int) {
	switch ca {
	case Exposure:
		return -400, 400

	case Tint:
		return -200, 200

	case Hue:
		return -100, 100

	case Saturation:
		return -100, 100

	case Temperature:
		return -200, 200

	case Highlights:
		return -100, 100

	case Shadows:
		return -100, 100

	case Contrast:
		return -200, 200

	case BlackPoint:
		return -200, 200

	case Fade:
		return 0, 200

	case VignetteExposure:
		return -200, 200

	case VignetteBlackPoint:
		return -200, 200

	case VignetteSoftness:
		return 0, 100

	default:
		return 0, 0
	}
}

func (ca ColorAdjustment) newValue() ColorAdjustmentValue {
	switch ca {
	case Vignette:
		v, err := newRangeGroup(ca)
		if err != nil {
			return nil
		}

		return v

	default:
		return newRangeValue(ca)
	}
}

type term struct {
	osascriptVariable string
	osascriptTerm     string
}
