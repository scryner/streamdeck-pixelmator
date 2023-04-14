package pixelmator

type adjustable interface {
	adjust(delta int) error
}

func (px *Pixelmator) AdjustExposure(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustTemperature(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustHighlights(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustShadows(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustContrast(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustBlackpoint(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustFade(delta int) error {
	return nil
}

func (px *Pixelmator) AdjustVignetteExposure(delta int) error {
	return nil
}
