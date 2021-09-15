/*
 * Copyright 2021 Meraj Sahebdar
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrAssetNotFound = fmt.Errorf("not found")

	vars = []string{
		"/var/buttress",
	}

	etcs = []string{
		"/etc/buttress",
	}
)

// GetPath Returns the valid path to the requested asset from the provided set of dirs.
//
// Errors:
//   - config/ErrAssetNotFound in case of non-existing asset with the given path.
func GetPath(dir string, dirs []string, path ...string) (string, error) {
	fin := strings.Join(path, "/")

	for _, dir := range dirs {
		if _, err := os.Stat(dir + fin); err == nil {
			return dir + fin, nil
		}
	}

	return "", errors.Wrap(ErrAssetNotFound, fmt.Sprintf("asset %s: %s not found", dir, fin))
}

// GetVarPath Returns the valid path to the requested asset from `vars`.
//
// ErrorsRefs:
//   - config/GetPath
func GetVarPath(path ...string) (string, error) {
	if found, err := GetPath("var", vars, path...); err != nil {
		return "", err
	} else {
		return found, nil
	}
}

// GetEtcPath Returns the valid path to the requested asset from `vars`.
//
// ErrorsRefs:
//   - config/GetPath
func GetEtcPath(path ...string) (string, error) {
	if found, err := GetPath("etc", etcs, path...); err != nil {
		return "", err
	} else {
		return found, nil
	}
}
