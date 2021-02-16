// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package internal

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/storskegg/ent/entc/gen"

	"github.com/stretchr/testify/require"
)

func TestSnapshot_Restore(t *testing.T) {
	t.Log("Running snapshot-restore integration test")
	const testPackage = "../integration/privacy/ent"
	err := addConflicts(testPackage)
	require.NoError(t, err)
	storage, err := gen.NewStorage("sql")
	require.NoError(t, err)
	snap := &Snapshot{
		Path: filepath.Join(testPackage, "internal/schema.go"),
		Config: &gen.Config{
			Storage: storage,
			Target:  testPackage,
			Schema:  filepath.Join(testPackage, "schema"),
			Header: `
			// Copyright 2019-present Facebook Inc. All rights reserved.
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated by entc, DO NOT EDIT.
			`,
		}}
	require.NoError(t, snap.Restore())
	err = exec.Command("go", "generate", testPackage).Run()
	require.NoError(t, err)
}

// addConflicts adds VCS conflicts to the files that match the given patterns.
func addConflicts(dir string) error {
	rand.Seed(time.Now().UnixNano())
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, info := range infos {
		if info.IsDir() || info.Name() == "generate.go" {
			continue
		}
		path := filepath.Join(dir, info.Name())
		fi, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		if _, err := fi.WriteString("\n" + conflictMarker); err != nil {
			return err
		}
		if err := fi.Close(); err != nil {
			return err
		}
	}
	return nil
}
