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
	"io"
	"io/ioutil"
	"os"

	"github.com/creasty/defaults"
	"github.com/markbates/pkger"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
)

func init() {
	// Search for the environment.
	switch os.Getenv("APP_ENV") {
	case "production":
		CurrentEnv = Prod
	case "test":
		CurrentEnv = Test
	default:
		CurrentEnv = Dev
	}

	// Create a logger instance.
	Log = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		os.Stderr,
		zap.NewAtomicLevel(),
	))

	// Read the config content.
	var content []byte
	if f, err := GetEtcPath("/config.yml"); err != nil {
		// Try to load the packed `default config.yml`.
		if pak, err := pkger.Open("/config-default.yml"); err != nil {
			Log.Fatal("failed to load the default config", zap.Error(err))
		} else {
			defer pak.Close()

			var c buffer.Buffer
			io.Copy(&c, pak)
			content = c.Bytes()
		}
	} else {
		if c, err := ioutil.ReadFile(f); err != nil {
			Log.Fatal("failed to load the config", zap.Error(err))
		} else {
			content = c
		}
	}

	// Fill environment variables in the config content.
	content = []byte(os.ExpandEnv(string(content)))

	// Feed `Cog`.
	if err := yaml.Unmarshal(content, &Cog); err != nil {
		Log.Fatal(err.Error())
	}

	// Provide default values for non-exsisting values.
	defaults.Set(&Cog)
}
