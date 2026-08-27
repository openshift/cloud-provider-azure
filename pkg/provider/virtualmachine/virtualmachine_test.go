/*
Copyright 2024 The Kubernetes Authors.

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

package virtualmachine

import (
	"testing"

	"sigs.k8s.io/cloud-provider-azure/pkg/consts"
)

func TestNilVirtualMachine(t *testing.T) {
	var vm *VirtualMachine

	if vm.IsVirtualMachine() {
		t.Error("nil VirtualMachine should return false for IsVirtualMachine()")
	}
	if vm.IsVirtualMachineScaleSetVM() {
		t.Error("nil VirtualMachine should return false for IsVirtualMachineScaleSetVM()")
	}
	if vm.ManagedByVMSS() {
		t.Error("nil VirtualMachine should return false for ManagedByVMSS()")
	}
	if vm.GetInstanceViewStatus() != nil {
		t.Error("nil VirtualMachine should return nil for GetInstanceViewStatus()")
	}
	if vm.GetProvisioningState() != consts.ProvisioningStateUnknown {
		t.Errorf("nil VirtualMachine should return %q for GetProvisioningState(), got %q", consts.ProvisioningStateUnknown, vm.GetProvisioningState())
	}
}
