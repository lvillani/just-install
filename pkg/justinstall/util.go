// just-install - The simple package installer for Windows
// Copyright (C) 2020 just-install authors.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package justinstall

import (
	"bytes"
	"os"
	"strings"
	"text/template"
)

// expandString expands any environment variable in the given string, with additional variables
// coming from the given context.
func expandString(s string, context map[string]string) string {
	data := environMap()

	// Merge the given context
	for k, v := range context {
		data[k] = v
	}

	var buf bytes.Buffer

	template.Must(template.New("expand").Parse(s)).Execute(&buf, data)

	return buf.String()
}

// environMap returns the current environment variables as a map.
func environMap() map[string]string {
	ret := make(map[string]string)
	env := os.Environ()

	for _, v := range env {
		split := strings.SplitN(v, "=", 2)

		if split[0] == "" && split[1] == "" {
			continue
		}

		split[0] = strings.ToUpper(split[0]) // Normalize variable names to upper case
		split[0] = strings.Replace(split[0], "(X86)", "_X86", -1)

		ret[split[0]] = split[1]
	}

	return ret
}
