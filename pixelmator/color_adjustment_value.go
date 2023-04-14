package pixelmator

import "fmt"

type ExposureValue struct {
	rangeValue
}

func (ex *ExposureValue) getTerms() []term {
	return []term{Exposure.getTerm()}
}

func (ex *ExposureValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Exposure.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				ex.Value = i
			}
		}
	}

	return nil
}

type TintValue struct {
	rangeValue
}

func (t *TintValue) getTerms() []term {
	return []term{Tint.getTerm()}
}

func (t *TintValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Tint.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				t.Value = i
			}
		}
	}

	return nil
}

type HueValue struct {
	rangeValue
}

func (h *HueValue) getTerms() []term {
	return []term{Hue.getTerm()}
}

func (h *HueValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Hue.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				h.Value = i
			}
		}
	}

	return nil
}

type SaturationValue struct {
	rangeValue
}

func (s *SaturationValue) getTerms() []term {
	return []term{Saturation.getTerm()}
}

func (s *SaturationValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Saturation.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				s.Value = i
			}
		}
	}

	return nil
}

type TemperatureValue struct {
	rangeValue
}

func (t *TemperatureValue) getTerms() []term {
	return []term{Temperature.getTerm()}
}

func (t *TemperatureValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Temperature.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				t.Value = i
			}
		}
	}

	return nil
}

type HighlightsValue struct {
	rangeValue
}

func (h *HighlightsValue) getTerms() []term {
	return []term{Highlights.getTerm()}
}

func (h *HighlightsValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Highlights.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				h.Value = i
			}
		}
	}

	return nil
}

type ShadowsValue struct {
	rangeValue
}

func (s *ShadowsValue) getTerms() []term {
	return []term{Shadows.getTerm()}
}

func (s *ShadowsValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Shadows.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				s.Value = i
			}
		}
	}

	return nil
}

type ContrastValue struct {
	rangeValue
}

func (c *ContrastValue) getTerms() []term {
	return []term{Contrast.getTerm()}
}

func (c *ContrastValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Contrast.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				c.Value = i
			}
		}
	}

	return nil
}

type BlackPointValue struct {
	rangeValue
}

func (b *BlackPointValue) getTerms() []term {
	return []term{BlackPoint.getTerm()}
}

func (b *BlackPointValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case BlackPoint.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				b.Value = i
			}
		}
	}

	return nil
}

type FadeValue struct {
	rangeValue
}

func (f *FadeValue) getTerms() []term {
	return []term{Fade.getTerm()}
}

func (f *FadeValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Fade.getTerm().osascriptVariable:
			if i, ok := v.(int); !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			} else {
				f.Value = i
			}
		}
	}

	return nil
}

type VignetteValue struct {
	rangeGroup
}

func (vig *VignetteValue) getTerms() []term {
	return []term{
		Vignette.getTerm(),
		VignetteExposure.getTerm(),
		VignetteBlackPoint.getTerm(),
		VignetteSoftness.getTerm(),
	}
}

func (vig *VignetteValue) setValues(m map[string]any) error {
	for k, v := range m {
		switch k {
		case Vignette.getTerm().osascriptVariable:
			b, ok := v.(bool)
			if !ok {
				return fmt.Errorf("invalid value: %s must be boolean (got %v)", k, v)
			}

			vig.IsApplied = b

		case VignetteExposure.getTerm().osascriptVariable:
			i, ok := v.(int)
			if !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			}

			vig.Group[VignetteExposure] = i

		case VignetteBlackPoint.getTerm().osascriptVariable:
			i, ok := v.(int)
			if !ok {
				return fmt.Errorf("invalid value: %s must be integer (got %v)", k, v)
			}

			vig.Group[VignetteBlackPoint] = i

		case VignetteSoftness.getTerm().osascriptVariable:
			i, ok := v.(int)
			if !ok {
				return fmt.Errorf("invalid value: must be integer (got %v)", v)
			}

			vig.Group[VignetteSoftness] = i
		}
	}

	return nil
}
