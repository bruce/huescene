package huescene

type YAMLPower struct {
	Value   bool
	IsValid bool
	IsSet   bool
}

func (i *YAMLPower) UnmarshalYAML(unmarshal func(interface{}) error) error {
	i.IsSet = true

	err := unmarshal(&i.Value)
	if err != nil {
		return err
	}

	return nil
}
