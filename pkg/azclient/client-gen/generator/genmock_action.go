/*
Copyright 2023 The Kubernetes Authors.

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

// Package generator
package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
)

func generateMock(ctx *genall.GenerationContext, pkg *loader.Package, typeName string, _ ClientGenConfig, headerText string) error {
	var mockCache bytes.Buffer
	//nolint:gosec // G204 ignore this!
	cmd := exec.Command("mockgen", "-package", "mock_"+pkg.Name, pkg.PkgPath, typeName)
	cmd.Stdout = &mockCache
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(pkg.Name+"/mock_"+pkg.Name, 0755); err != nil {
		return err
	}
	mockFile, err := ctx.Open(pkg, "mock_"+pkg.Name+"/interface.go")
	if err != nil {
		return err
	}
	defer mockFile.Close()
	_, err = mockFile.Write([]byte(headerText + "\n"))
	if err != nil {
		return err
	}
	_, err = mockFile.Write(mockCache.Bytes())
	fmt.Println("Generated mock in " + pkg.Name)
	return err
}
