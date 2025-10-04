/*
 * Copyright (C) 2020-2022, IrineSistiana
 *
 * This file is part of mosdns.
 *
 * mosdns is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * mosdns is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package utils

import (
	"github.com/go-viper/mapstructure/v2"
	"golang.org/x/exp/constraints"
	"strconv"
)

func SetDefaultNum[K constraints.Integer | constraints.Float](p *K, d K) {
	if *p == 0 {
		*p = d
	}
}

func SetDefaultUnsignNum[K constraints.Integer | constraints.Float](p *K, d K) {
	if *p <= 0 {
		*p = d
	}
}

func SetDefaultString(p *string, d string) {
	if len(*p) == 0 {
		*p = d
	}
}

func CheckNumRange[K constraints.Integer | constraints.Float](v, min, max K) bool {
	if v < min || v > max {
		return false
	}
	return true
}

// WeakDecode decodes args from config to output.
func WeakDecode(in any, output any) error {
	config := &mapstructure.DecoderConfig{
		ErrorUnused:      true,
		Result:           output,
		WeaklyTypedInput: true,
		TagName:          "yaml",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(in)
}

func ParseNameOrNum[T constraints.Integer](s string, m map[string]T) (T, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		v, ok := m[s]
		return v, ok
	}
	return T(i), true
}
