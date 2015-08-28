/*
 * g2z - Zabbix module adapter for Go
 * Copyright (C) 2015 - Ryan Armstrong <ryan@cavaliercoder.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */
package g2z

import (
	"bytes"
	"regexp"
	"strings"
)

type DiscoveryData []DiscoveryItem

type DiscoveryItem map[string]string

var macroIllegalPattern = regexp.MustCompile(`[^A-Z0-9_]+`)

// JSON returns a JSON encoded discovery data string, compatible with Zabbix
// Low-Level discovery rules for v2.2.0 and above.
func (c DiscoveryData) JSON() string {
	b := bytes.Buffer{}

	b.WriteString("{\n\t\"data\":[")

	for i, item := range c {
		if i > 0 {
			b.WriteString(",")
		}

		b.WriteString("\n\t{")

		firstMacro := true
		for macro, val := range item {
			if firstMacro {
				firstMacro = false
			} else {
				b.WriteString(",")
			}

			b.WriteString("\n\t\t\"{#")
			b.WriteString(macroName(macro))
			b.WriteString("}\":\"")
			b.WriteString(jsonEscape(val))
			b.WriteString("\"")
		}

		b.WriteString("}")
	}

	b.WriteString("]}")

	return b.String()
}

func jsonEscape(a string) string {
	return strings.Replace(a, "\"", "\\\"", -1)
}

func macroName(name string) string {
	name = strings.ToUpper(name)
	name = strings.Replace(name, " ", "_", -1)
	name = macroIllegalPattern.ReplaceAllString(name, "")
	return name
}
