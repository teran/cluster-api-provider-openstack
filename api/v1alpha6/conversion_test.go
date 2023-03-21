/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha6

import (
	"testing"

	utilconversion "sigs.k8s.io/cluster-api/util/conversion"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha7"
)

func TestFuzzyConversion(t *testing.T) {
	t.Run("for OpenStackCluster", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:   &infrav1.OpenStackCluster{},
		Spoke: &OpenStackCluster{},
	}))

	t.Run("for OpenStackClusterTemplate", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:   &infrav1.OpenStackClusterTemplate{},
		Spoke: &OpenStackClusterTemplate{},
	}))

	t.Run("for OpenStackMachine", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:   &infrav1.OpenStackMachine{},
		Spoke: &OpenStackMachine{},
	}))

	t.Run("for OpenStackMachineTemplate", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:   &infrav1.OpenStackMachineTemplate{},
		Spoke: &OpenStackMachineTemplate{},
	}))
}