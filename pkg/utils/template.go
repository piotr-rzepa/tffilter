package utils

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"text/template"

	tfj "github.com/hashicorp/terraform-json"
)

const TfTargetTemplate = "-target={{.Type}}.{{.Name}}{{ if not (eq .Index nil) }}[{{.Index}}]{{end}} "
const TfTargetTemplateName = "tftarget"

type TemplateObject struct {
	Type   string
	Name   string
	Index  interface{}
	Change tfj.Actions
}

func findIntersection(filters []string, actions tfj.Actions) []string {
	intersection := make([]string, 0)

	// Traverse each element in arr1
	for _, num1 := range filters {
		// Check if num1 exists in arr2
		for _, num2 := range actions {
			fmt.Println(num1, "=", num2)
			if num1 == string(num2) {
				intersection = append(intersection, num1)
				break
			}
		}
	}

	return intersection
}

func parseTemplate() (*template.Template, error) {
	return template.New(TfTargetTemplateName).Parse(TfTargetTemplate)
}

func ProcessPlanChanges(changes []*tfj.ResourceChange, filters Filters) (*bytes.Buffer, error) {
	var templatedObjects bytes.Buffer
	tmpl, err := parseTemplate()
	if err != nil {
		log.Fatal(err)
	}
	resources := []TemplateObject{}
	c := 0
	for i, a := range changes {
		commonActions := findIntersection(filters.ActionFilters, a.Change.Actions)
		fmt.Println("common actions", commonActions)
		if len(commonActions) < 1 {
			fmt.Println("No common actions found between filters and plan changes")
			c = c + 1
			continue
		}
		// fmt.Println(a.Type, a.Name, a.Index, a.Change.Actions)
		if a.Change.Actions.NoOp() {
			//TODO: Handle moved operation
			fmt.Printf("WARNING: No-op operation detected for resource %s %s - it might be no-change but also moved operation!", a.Type, a.Name)
			c = c + 1
			continue
		}
		resources = append(resources, TemplateObject{Type: a.Type, Name: a.Name, Index: a.Index, Change: a.Change.Actions})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(TemplateObject{Type: a.Type, Name: a.Name, Index: a.Index, Change: a.Change.Actions})
		err = tmpl.Execute(&templatedObjects, resources[i-c])
		if err != nil {
			log.Fatal(err)
		}
	}
	// If we didn't add any objects to an array, it means that no actions should be performed
	if len(resources) < 1 {
		return nil, errors.New("found 0 resources after filtering - no action will be performed")
	}
	return &templatedObjects, nil
}
