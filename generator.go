//go:build ignore
// +build ignore

/*
Copyright 2022. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"
	"path/filepath"
	"text/template"
)

const (
	crdTemplate = `// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package crd

var {{ .ExportedVar }}File = {{ printf "%q" .File }}
var {{ .ExportedVar }}CRD = []byte({{- printf "%s" .CRD -}})
`
)

func generate(filename, outputFilename, crdName string) {
	// Get CustomResourceDefinition file
	fileAbs, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(fileAbs)
	if err != nil {
		panic(err)
	}
	contentStr := "`" + string(content) + "`"

	// Find the output.
	crd, err := os.Create(outputFilename + ".go")
	if err != nil {
		panic(err)
	}
	defer crd.Close()

	// Store file contents.
	type CRDInfo struct {
		CRD         string
		File        string
		ExportedVar string
	}
	mi := CRDInfo{
		CRD:         contentStr,
		File:        filename,
		ExportedVar: crdName,
	}

	// Generate template.
	manifest := template.Must(template.New("crd-generate").Parse(crdTemplate))
	if err := manifest.Execute(crd, mi); err != nil {
		panic(err)
	}
}

func main() {
	classifierFile := "../../config/crd/bases/lib.projectsveltos.io_classifiers.yaml"
	generate(classifierFile, "classifiers", "Classifier")

	classifierReportFile := "../../config/crd/bases/lib.projectsveltos.io_classifierreports.yaml"
	generate(classifierReportFile, "classifierreports", "ClassifierReport")

	debuggingConfigurationFile := "../../config/crd/bases/lib.projectsveltos.io_debuggingconfigurations.yaml"
	generate(debuggingConfigurationFile, "debuggingconfigurations", "DebuggingConfiguration")
}
