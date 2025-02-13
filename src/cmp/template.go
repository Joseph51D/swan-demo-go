/* ****************************************************************************
 * Copyright 2020 51 Degrees Mobile Experts Limited (51degrees.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 * ***************************************************************************/

package cmp

import (
	"strconv"
)

type EmailTemplate struct {
	Salt           []byte
	PreferencesUrl string
}

func (t EmailTemplate) Show(i string) bool {
	b, err := strconv.ParseInt(i, 10, 8)
	if err != nil {
		return false
	}

	for _, v := range t.Salt {
		if v == byte(b-1) {
			return true
		}
	}
	return false
}

func (t EmailTemplate) Number(i string) string {
	b, err := strconv.ParseInt(i, 10, 8)
	if err != nil {
		return ""
	}

	for n, v := range t.Salt {
		if v == byte(b-1) {
			return strconv.Itoa(n + 1)
		}
	}
	return ""
}
