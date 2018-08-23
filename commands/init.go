package commands

import (
	"fmt"
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
		log.Errorf("init failed with %v", err.Error())
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

	err = generateFileFromTemplate(templates.WebpackConfig, filepath.Join(appPath, "webpack.config.js"), struct {
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
	err = os.MkdirAll(filepath.Join(appPath, "src/containers"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/containers folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/components"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/components folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/reducers"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/reducers folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/interfaces"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/interfaces folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/sagas"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/sagas folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/actions"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/actions folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/models"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/models folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(appPath, "src/html"), 0755)
	if err != nil {
		log.Errorf("init failed. Unable to create src/html folder. Error: %v", err.Error())
		return 1
	}

	templateData := struct{ AppName string }{name}

	if err = generateFileFromTemplate(templates.SrcIndexTsxTemplate, filepath.Join(appPath, "src", "index.tsx"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcAppContainerTemplate, filepath.Join(appPath, "src", "containers", fmt.Sprintf("%v.ts", name)), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcAppComponentTemplate, filepath.Join(appPath, "src", "components", fmt.Sprintf("%v.tsx", name)), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcRootReducerTemplate, filepath.Join(appPath, "src", "reducers", "index.ts"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcInterfaceWindow, filepath.Join(appPath, "src", "interfaces", "EnhancedWindow.ts"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcInterfaceAction, filepath.Join(appPath, "src", "interfaces", "Action.ts"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcSagaIndex, filepath.Join(appPath, "src", "sagas", "index.ts"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcActionIndex, filepath.Join(appPath, "src", "actions", "index.ts"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcModelServerConfig, filepath.Join(appPath, "src", "models", "ServerConfig.ts"), templateData); err != nil {
		return 1
	}

	if err = generateFileFromTemplate(templates.SrcHtmlIndex, filepath.Join(appPath, "src", "html", "index.html"), templateData); err != nil {
		return 1
	}

	return 0
}

func generateFileFromTemplate(templateString string, filepath string, values interface{}) error {
	configTemplate, err := template.New("").Parse(templateString)
	if err != nil {
		log.Errorf("init failed. Unable to generate %v. Error: %v", filepath, err.Error())
		return err
	}

	config := strings.Builder{}
	configTemplate.Execute(&config, values)
	err = ioutil.WriteFile(filepath, []byte(config.String()), 0644)

	if err != nil {
		log.Errorf("init failed. Unable to generate %v. Error: %v", filepath, err.Error())
		return err
	}

	return nil
}
