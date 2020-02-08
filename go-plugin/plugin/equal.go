package plugin

import (
	"encoding/json"
	"reflect"
)

type EqualPluginConfig struct {
	Target interface{} `json:"target"`
	Value  interface{} `json:"value"`
}

func Equal(input interface{}, variables map[string]interface{}) (map[string]interface{}, error) {
	var (
		err error
	)

	var config EqualPluginConfig
	param, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(param, &config)
	if err != nil {
		return nil, err
	}

	variables["result"] = reflect.DeepEqual(config.Target, config.Value)

	return variables, err
}
