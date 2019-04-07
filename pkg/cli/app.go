/*
Copyright 2019 Sensible

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
package cli

import (
	ucli "github.com/urfave/cli"
)

func App(version string) *ucli.App {
	app := ucli.NewApp()
	app.Name = "Sensible SDT"
	app.Version = version
	app.Usage = "Schema Definition Translation"

	return app
}
