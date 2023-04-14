package pixelmator

func syncColorAdjustments(m map[ColorAdjustment]ColorAdjustmentValue) error {
	return nil
}

const syncScript = `if application "Pixelmator Pro" is running then
	tell application "Pixelmator Pro"
		tell the color adjustments of the current layer of the front document
			set exposureValue to its exposure
			set tintValue to its tint
			set hueValue to its hue
			set saturationValue to its saturation
			set temperatureValue to its temperature
			set highlightsValue to its highlights
			set shadowsValue to its shadows
			set contrastValue to its contrast
			set blackpointValue to its black point
			set fadeValue to its fade
			set vignetteEnable to its vignette
			set vignetteExposureValue to its vignette exposure
		end tell
	end tell
end if

if application "Pixelmator Pro" is running then
	{exposureValue, tintValue, hueValue, saturationValue, temperatureValue, highlightsValue, shadowsValue, contrastValue, blackpointValue, fadeValue, vignetteEnable, vignetteExposureValue}
else
	{}
end if`
