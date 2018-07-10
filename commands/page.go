package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/template"
	"github.com/sparkymat/tsagen/templates"
)

func GeneratePage(name string, folder string, appName string) int {
	log.Infof("Generating page %v", name)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Errorf("page generation failed with %v", err.Error())
		return 1
	}

	if appName == "" {
		appName = filepath.Base(currentPath)
	}

	pagePath := filepath.Join(currentPath, folder, name)
	if _, err = os.Stat(pagePath); err == nil {
		log.Error("page generation failed. Folder already exists")
		return 1
	}

	err = os.MkdirAll(pagePath, 0755)
	if err != nil {
		log.Errorf("page generation failed. Unable to create page folder. Error: %v", err.Error())
		return 1
	}
	err = os.MkdirAll(filepath.Join(pagePath, "components"), 0755)
	if err != nil {
		log.Errorf("page generation failed. Unable to create page folder. Error: %v", err.Error())
		return 1
	}

	templateData := struct {
		PageName string
		AppName  string
	}{name, appName}

	// Create container
	if err = generateFileFromTemplate(templates.SrcPageContainerTemplate, filepath.Join(pagePath, "index.ts"), templateData); err != nil {
		return 1
	}

	// Create root component
	if err = generateFileFromTemplate(templates.SrcPageComponentTemplate, filepath.Join(pagePath, "components", "index.tsx"), templateData); err != nil {
		return 1
	}

	// Create reducer
	if err = generateFileFromTemplate(templates.SrcPageReducerTemplate, filepath.Join(pagePath, "reducer.ts"), templateData); err != nil {
		return 1
	}

	// Create actions
	if err = generateFileFromTemplate(templates.BlankFile, filepath.Join(pagePath, "actions.ts"), templateData); err != nil {
		return 1
	}

	// Create sagas
	if err = generateFileFromTemplate(templates.SrcPageSagaTemplate, filepath.Join(pagePath, "sagas.ts"), templateData); err != nil {
		return 1
	}

	// Helpful messages
	log.Infof("Successfully generated page %v", name)

	helpTemplateString := `

Remember to:

1. Add your reducer to the app's root reducer and state

	import {{.PageName}}Reducer, { {{.PageName}}State } from "../{{.PageName}}/reducer";

	..

	export interface {{.PageName}}State {
		..
		{{.PageName}}: {{.PageName}}State;
	}

	..

	const rootReducer = combineReducers({
		..
		{{.PageName}}: {{.PageName}}Reducer,
	})

2. Include your sagas in the root saga

	import {{.PageName}}Sagas from "../{{.PageName}}/sagas";

	..

	sagas = sagas.concat({{.PageName}}Sagas);

	..
`
	helpTemplate, err := template.New("").Parse(helpTemplateString)
	if err != nil {
		return 0
	}
	helpText := strings.Builder{}
	helpTemplate.Execute(&helpText, struct {
		PageName string
		AppName  string
	}{name, appName})
	fmt.Println(helpText.String())

	return 0
}
