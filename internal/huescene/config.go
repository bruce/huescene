package huescene

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/amimof/huego"
	"gopkg.in/yaml.v2"
)

var defaultuser = "huescene"

// Config models the configuration information.
type Config struct {
	Username string
	Key      string
	Scenes   []SceneConfig
}

// SceneConfig models the scene-specific configuration information.
type SceneConfig struct {
	Name       string
	Color      string
	Brightness uint8
	Power      YAMLPower
	Lights     []SceneLightConfig
}

// SceneLightConfig models the light-specific configuration information for a scene.
type SceneLightConfig struct {
	Name       string
	Color      string
	Power      YAMLPower
	Brightness uint8
}

// UnmarshalConfig unmarshals raw config and returns a Config.
func UnmarshalConfig(data []byte) (*Config, error) {
	cfg := Config{Username: defaultuser}

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func PrintConfig(cfg Config) {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		fmt.Println("Can not print config.")
		log.Panic(err)
	}

	fmt.Println(string(data))
}

// ReadConfig reads a file and returns a
func ReadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return UnmarshalConfig(data)
}

func ReadConfigFromStdin() (*Config, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		content, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		return UnmarshalConfig(content)

	} else {
		return nil, errors.New("no configuration given")
	}
}

// stateConfig extracts the configuration for a scene by name.
func stateConfig(cfg Config, scene string) *SceneConfig {
	for i := range cfg.Scenes {
		scfg := cfg.Scenes[i]
		if scfg.Name == scene {
			return &scfg
		}
	}
	return nil
}

// lightConfig extracts the configuration for a scene's light.
func lightConfig(l huego.Light, scfg SceneConfig) *SceneLightConfig {
	for i := range scfg.Lights {
		lcfg := scfg.Lights[i]
		if lcfg.Name == l.Name {
			return &lcfg
		}
	}
	return nil
}
