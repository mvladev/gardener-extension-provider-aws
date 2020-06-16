// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1_test

import (
	"testing"

	"github.com/gardener/gardener-extension-provider-aws/pkg/apis/config/v1alpha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/utils/pointer"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config V1alpha1 Suite")
}

var _ = Describe("Defaults", func() {

	Describe("SetDefaults_ControllerConfiguration", func() {

		It("correct defaults are set", func() {
			given := &v1alpha1.ControllerConfiguration{}
			expected := &v1alpha1.ControllerConfiguration{
				ShootStorageClassConfig: &v1alpha1.StorageClass{
					Encrypted: pointer.BoolPtr(false),
				},
			}

			v1alpha1.SetDefaults_ControllerConfiguration(given)

			Expect(given).To(BeEquivalentTo(expected))
		})

	})

})
