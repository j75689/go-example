package plugin

import (
	"fmt"
	"io/ioutil"
	"log"
	"plugin"
	"strings"
	"sync"
)

type PluginFunc func(interface{}, map[string]interface{}) (map[string]interface{}, error)

var (
	pluginfuncs = &sync.Map{}
)

// load all .so plugin file
func Load(path string) {

	// Fix Path
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	// read files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("[Init] ", err)
		return
	}

	// load plugin
	for _, f := range files {
		if !f.IsDir() {
			var runFuncName = f.Name()

			if !strings.HasSuffix(f.Name(), ".so") {
				continue
			}

			if strings.LastIndexAny(runFuncName, ".") > -1 {
				runFuncName = runFuncName[0:strings.LastIndexAny(runFuncName, ".")]
			}

			p, err := plugin.Open(path + f.Name())
			if err != nil {
				log.Println("[Init] ", err)
				continue
			}

			function, err := p.Lookup(runFuncName)
			if err != nil {
				log.Println("[Init] ", err)
				continue
			}

			if f, ok := function.(func(interface{}, map[string]interface{}) (map[string]interface{}, error)); ok {
				ff := PluginFunc(f)
				pluginfuncs.Store(runFuncName, &ff)
			} else {
				log.Println("[Init] load plugin [%s] failed.\n", runFuncName)
			}

		}
	}
}

func Execute(pluginName string, input interface{}, variables map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := pluginfuncs.Load(pluginName); ok {
		plugin := v.(*PluginFunc)
		return (*plugin)(input, variables)
	}
	return variables, fmt.Errorf("plugin [%s] not found.\n", pluginName)
}
