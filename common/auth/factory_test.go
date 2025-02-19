/**
 * Copyright 2020 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package auth ...
package auth

import (
	"testing"

	"github.com/IBM/ibmcloud-volume-interface/config"
	vpcconfig "github.com/IBM/ibmcloud-volume-vpc/block/vpcconfig"
	"github.com/stretchr/testify/assert"
)

func TestNewContextCredentialsFactory(t *testing.T) {
	conf := &vpcconfig.VPCBlockConfig{
		VPCConfig: &config.VPCProviderConfig{
			Enabled:         true,
			EndpointURL:     "test-iam-url",
			VPCTimeout:      "30s",
			IamClientID:     "test-iam_client_id",
			IamClientSecret: "test-iam_client_secret",
		},
	}

	_, err := NewVPCContextCredentialsFactory(conf)
	assert.NotNil(t, err)
}
