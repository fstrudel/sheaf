/*
 * Copyright 2020 Sheaf Authors
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package sheaf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/bryanl/sheaf/pkg/images"
)

const (
	// BundleConfigFilename is the filename for a fs config.
	BundleConfigFilename = "bundle.json"

	bundleConfigDefaultVersion = "0.1.0"
)

// UserDefinedImageType is the type of user defined image.
type UserDefinedImageType string

const (
	// MultiResult states the JSON path will return multiple values.
	MultiResult UserDefinedImageType = "multiple"
	// SingleResult states the JSON path will return a single value.
	SingleResult UserDefinedImageType = "single"
)

// UserDefinedImage is a user defined image. These allow sheaf to find more
// images.
type UserDefinedImage struct {
	APIVersion string               `json:"apiVersion"`
	Kind       string               `json:"kind"`
	JSONPath   string               `json:"jsonPath"`
	Type       UserDefinedImageType `json:"type"`
}

// BundleConfig is a fs configuration.
type BundleConfig struct {
	// Name is the name of the fs.
	Name string `json:"name"`
	// Version is the version of the fs.
	Version string `json:"version"`
	// SchemaVersion is the version of the schema this fs uses.
	SchemaVersion string `json:"schemaVersion"`
	// Images is a set of images required by the fs.
	Images images.Set `json:"images"`
	// UserDefinedImages is a list of user defined image locations.
	UserDefinedImages []UserDefinedImage `json:"userDefinedImages,omitempty"`
}

// NewBundleConfig creates a BundleConfig.
func NewBundleConfig(name, version string) BundleConfig {
	if version == "" {
		version = bundleConfigDefaultVersion
	}

	return BundleConfig{
		Name:          name,
		Version:       version,
		SchemaVersion: "v1alpha1",
	}
}

// LoadBundleConfig loads a BundleConfig from a file.
func LoadBundleConfig(filename string) (BundleConfig, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return BundleConfig{}, fmt.Errorf("read %q: %w", filename, err)
	}

	var bc BundleConfig
	if err := json.Unmarshal(data, &bc); err != nil {
		return BundleConfig{}, err
	}

	return bc, nil
}

// StoreBundleConfig saves a BundleConfig to a file, destructively.
func StoreBundleConfig(bc BundleConfig, filename string) error {
	jbc, err := json.Marshal(bc)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, jbc, 0644)
}

// Filename returns the fs archive file name for this BundleConfig.
func (bc *BundleConfig) Filename(dir string) string {
	return filepath.Join(dir, fmt.Sprintf("%s-%s.tgz", bc.Name, bc.Version))
}
