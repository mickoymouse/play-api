package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Settings : structure for settings
type Settings struct {
	AWSRegion    string
	AWSAccessKey string
	AWSSecretKey string
}

var settings = Settings{}

//Init : initialization of settings
func Init() {
	var env = os.Getenv("GO_ENV")
	LoadSettingsByEnv(env)
}

//LoadSettingsByEnv : loading settings base on your GO_ENV variables
func LoadSettingsByEnv(env string) {
	configFile := GetExecDirectory() + "/settings.json"
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

//Get : get current settings variables
func Get() Settings {
	Init()
	return settings
}

//GetExecDirectory : getting current executable directory
func GetExecDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
