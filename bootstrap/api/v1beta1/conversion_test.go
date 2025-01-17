/*
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
package v1beta1

import (
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	"k8s.io/apimachinery/pkg/runtime"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"

	cabp3v1 "github.com/k3s-io/cluster-api-k3s/bootstrap/api/v1beta2"
)

func TestFuzzyConversion(t *testing.T) {
	g := NewWithT(t)
	scheme := runtime.NewScheme()
	g.Expect(AddToScheme(scheme)).To(Succeed())
	g.Expect(cabp3v1.AddToScheme(scheme)).To(Succeed())

	t.Run("for KThreesConfig", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme:      scheme,
		Hub:         &cabp3v1.KThreesConfig{},
		Spoke:       &KThreesConfig{},
		FuzzerFuncs: []fuzzer.FuzzerFuncs{},
	}))
	t.Run("for KThreesConfigTemplate", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme:      scheme,
		Hub:         &cabp3v1.KThreesConfigTemplate{},
		Spoke:       &KThreesConfigTemplate{},
		FuzzerFuncs: []fuzzer.FuzzerFuncs{},
	}))
}
