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

package azclient_test

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/recording"
)

var _ = Describe("Auth", Ordered, func() {
	var tenantID string = "tenantid"
	var clientID string = "clientid"

	var cred azcore.TokenCredential
	var authProvider *azclient.AuthProvider
	var httpRecorder *recording.Recorder
	var err error
	BeforeAll(func() {
		httpRecorder, err = recording.NewRecorder("testdata/auth")
		Expect(err).NotTo(HaveOccurred())
		tenantID = httpRecorder.TenantID()
		clientID = httpRecorder.ClientID()
	})

	When("AADClientSecret is set", func() {
		It("should return a valid token", func() {
			authProvider, err = azclient.NewAuthProvider(azclient.AzureAuthConfig{
				TenantID:        tenantID,
				AADClientID:     clientID,
				AADClientSecret: httpRecorder.ClientSecret(),
			}, &arm.ClientOptions{
				ClientOptions: azpolicy.ClientOptions{
					Transport: httpRecorder.HTTPClient(),
				},
			})
			Expect(err).NotTo(HaveOccurred())
			cred, err = authProvider.GetAzIdentity()
			Expect(err).NotTo(HaveOccurred())
			Expect(cred).NotTo(BeNil())
			token, err := cred.GetToken(context.Background(), azpolicy.TokenRequestOptions{
				Scopes:   []string{"https://management.azure.com/.default"},
				TenantID: tenantID,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(token.Token).NotTo(BeEmpty())
		})
	})
	AfterAll(func() {
		err := httpRecorder.Stop()
		Expect(err).NotTo(HaveOccurred())
	})
})
