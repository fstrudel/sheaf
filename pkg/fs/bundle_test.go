// +build !integration

/*
 * Copyright 2020 Sheaf Authors
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package fs

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bryanl/sheaf/internal/testutil"
	"github.com/bryanl/sheaf/pkg/sheaf"
)

func TestBundle_Path(t *testing.T) {
	testutil.WithBundleDir(t, func(dir string) {
		testutil.StageFile(t, sheaf.BundleConfigFilename, filepath.Join(dir, sheaf.BundleConfigFilename))

		bundle, err := NewBundle(dir)
		require.NoError(t, err)

		actual := bundle.Path()
		require.Equal(t, dir, actual)
	})
}

func TestBundle_Config(t *testing.T) {
	testutil.WithBundleDir(t, func(dir string) {
		configRaw := testutil.StageFile(t, sheaf.BundleConfigFilename, filepath.Join(dir, sheaf.BundleConfigFilename))
		bcc := BundleConfigCodec{}
		r := bytes.NewReader(configRaw)

		wanted, err := bcc.Decode(r)
		require.NoError(t, err)

		bundle, err := NewBundle(dir)
		require.NoError(t, err)

		actual := bundle.Config()
		require.Equal(t, wanted, actual)
	})
}

func TestLoadBundleConfig(t *testing.T) {
	testutil.WithBundleDir(t, func(dir string) {
		testutil.StageFile(t, sheaf.BundleConfigFilename, filepath.Join(dir, sheaf.BundleConfigFilename))

		c, err := LoadBundleConfig(dir)
		require.NoError(t, err)

		require.Equal(t, "knative-serving-0.12", c.GetName())
	})

}
