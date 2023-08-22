/*
Copyright 2020 The Kubernetes Authors.

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

package provider

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRequestBackoff(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	az := GetTestCloud(ctrl)
	az.CloudProviderBackoff = true
	az.ResourceRequestBackoff = wait.Backoff{Steps: 3}

	backoff := az.RequestBackoff()
	assert.Equal(t, wait.Backoff{Steps: 3}, backoff)

}
