package huescene

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
)

func SetScene(cfg Config, bridge huego.Bridge, scene string) error {

	scfg := stateConfig(cfg, scene)
	if scfg == nil {
		return fmt.Errorf("no scene configuration found for scene %s", scene)
	}

	ls, err := bridge.GetLights()
	if err != nil {
		log.Println("no lights found")
		return nil
	}

	errs := make(chan error, len(ls))

	for _, l := range ls {
		go func(l huego.Light) {
			errs <- setLightScene(l, *scfg)
		}(l)
	}
	for range ls {
		err = <-errs
		if err != nil {
			return err
		}
	}

	return nil
}

func setLightScene(l huego.Light, scfg SceneConfig) error {
	lcfg := lightConfig(l, scfg)
	if lcfg == nil {
		log.Printf("no config for scene=%s, light=%s", scfg.Name, l.Name)
		return nil
	}

	power := lightPower(scfg, *lcfg)
	if power != nil {
		if *power {
			l.On()
		} else {
			// It's supposed to be off. Just turn it off.
			l.Off()
			return nil
		}
	}

	err := setLightColor(l, scfg, *lcfg)
	if err != nil {
		return err
	}

	setLightBrightness(l, scfg, *lcfg)

	return nil
}

func setLightColor(l huego.Light, scfg SceneConfig, lcfg SceneLightConfig) error {
	color := lightColor(scfg, lcfg)
	if color == "" {
		return nil
	}

	ri, gi, bi, err := rgbStringToRgbInt(color)
	if err != nil {
		return err
	}

	r, g, b, err := rgbIntToRgbFloat(ri, gi, bi)
	if err != nil {
		return err
	}

	x, y, err := rgbFloatToXy(r, g, b)
	if err != nil {
		return err
	}

	err = l.Xy([]float32{x, y})
	if err != nil {
		// We just log these
		log.Printf("Error setting color for light=%s: %v\n", l.Name, err)
	}

	return nil
}

func setLightBrightness(l huego.Light, scfg SceneConfig, lcfg SceneLightConfig) {
	bri := lightBrightness(scfg, lcfg)
	err := l.Bri(bri)
	if err != nil {
		// We just log these
		log.Printf("Error setting brightness for light=%s: %v\n", l.Name, err)
	}
}

func lightColor(scfg SceneConfig, lcfg SceneLightConfig) string {
	if lcfg.Color != "" {
		return lcfg.Color
	}
	return scfg.Color
}

func lightBrightness(scfg SceneConfig, lcfg SceneLightConfig) uint8 {
	if lcfg.Brightness != 0 {
		return lcfg.Brightness
	} else if scfg.Brightness != 0 {
		return scfg.Brightness
	}
	return 255
}

func lightPower(scfg SceneConfig, lcfg SceneLightConfig) *bool {
	var result bool
	if lcfg.Power.IsSet {
		result = lcfg.Power.Value
	} else if lcfg.Color != "" || lcfg.Brightness > 0 {
		result = true
	} else if scfg.Power.IsSet {
		result = scfg.Power.Value
	} else if scfg.Color != "" || scfg.Brightness > 0 {
		result = true
	}
	return &result
}
