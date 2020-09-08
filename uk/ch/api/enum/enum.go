/*
Copyright 2018 Foundation for Open Software Development (www.fosdev.org)

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

package enum

import (
	"gopkg.in/yaml.v2"
)

type ENUM map[string]map[string]string

func (y ENUM) Get(s, k string) string {
	v, ok := y[string(s)][string(k)]
	if !ok {
		return ""
	}
	return v
}

var Constants, FilingHistoryDescriptions, MortgageDescriptions, DisqualifiedOfficerDescriptions ENUM

func init() {
	if err := yaml.Unmarshal([]byte(filingHistoryDescriptionsYAML), &FilingHistoryDescriptions); err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal([]byte(constantsYAML), &Constants); err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal([]byte(mortgageDescriptionsYAML), &MortgageDescriptions); err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal([]byte(disqualifiedOfficerDescriptionsYAML), &DisqualifiedOfficerDescriptions); err != nil {
		panic(err)
	}
}
