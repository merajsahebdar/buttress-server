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

package main

import (
	"os"

	"buttress.io/app/command"
	"buttress.io/app/config"
	"github.com/alecthomas/kong"
	"go.uber.org/zap"
)

// CLI
var CLI struct {
	Serve command.Serve `cmd:"serve" help:"Serve the app APIs."`
}

func main() {
	if parser, err := kong.New(&CLI); err != nil {

	} else {
		if ctx, err := parser.Parse(os.Args[1:]); err != nil {
			config.Log.Fatal("failed to parse cli", zap.Error(err))
		} else {
			if err := ctx.Run(); err != nil {
				config.Log.Fatal("failed to run command", zap.Error(err))
			}
		}
	}
}
