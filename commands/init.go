package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/template"
	logger "github.com/apsdehal/go-logger"
	"github.com/markbates/inflect"
	"github.com/sparkymat/tsagen/templates"
)

var log, _ = logger.New("tsagen", 1, os.Stdout)

func InitialiseApp(name string, forceCreation bool) int {
	log.Infof("Initialising %v", name)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Errorf("init failed with %V", err.Error())
		return 1
	}

	appPath := filepath.Join(currentPath, name)
	if _, err = os.Stat(appPath); err == nil && !forceCreation {
		log.Error("init failed. Folder already exists")
		return 1
	}

	err = os.MkdirAll(appPath, 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create app folder. Error: %v", err.Error())
		return 1
	}

	// Create config files
	jsFileName := strings.Replace(inflect.Underscore(name), "_", "-", -1)
	appTitle := inflect.Titleize(strings.Replace(jsFileName, "-", " ", -1))
	err = generateFileFromTemplate(templates.PackageJSONTemplate, filepath.Join(appPath, "package.json"), struct{ AppName string }{name})
	if err != nil {
		log.Errorf("init failed. Unable to generate package.json. Error: %v", err.Error())
		return 1
	}

	err = generateFileFromTemplate(templates.TsConfigTemplate, filepath.Join(appPath, "tsconfig.json"), nil)
	if err != nil {
		log.Errorf("init failed. Unable to generate tsconfig.json. Error: %v", err.Error())
		return 1
	}

	err = generateFileFromTemplate(templates.WebpackConfigTemplate, filepath.Join(appPath, "webpack.config.js"), struct {
		JsFileName string
		AppTitle   string
	}{jsFileName, appTitle})
	if err != nil {
		log.Errorf("init failed. Unable to generate webpack.config.js. Error: %v", err.Error())
		return 1
	}

	// Generate source files
	err = os.MkdirAll(filepath.Join(appPath, "src"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src folder. Error: %v", err.Error())
		return 1
	}

	err = generateFileFromTemplate(templates.SrcIndexTsxTemplate, filepath.Join(appPath, "src", "index.tsx"), struct {
		AppName string
	}{name})
	if err != nil {
		log.Errorf("init failed. Unable to generate src/index.tsx. Error: %v", err.Error())
		return 1
	}

	return 0
}

func generateFileFromTemplate(templateString string, filepath string, values interface{}) error {
	configTemplate, err := template.New("").Parse(templateString)
	if err != nil {
		return err
	}

	config := strings.Builder{}
	configTemplate.Execute(&config, values)
	err = ioutil.WriteFile(filepath, []byte(config.String()), 0644)
	return err
}
