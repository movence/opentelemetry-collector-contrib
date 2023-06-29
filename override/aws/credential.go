// Copyright The OpenTelemetry Authors
// Portions of this file Copyright 2018-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws // import "github.com/open-telemetry/opentelemetry-collector-contrib/override/aws"
import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type credentialsProvider []credentials.Provider
type CredentialsChainOverride struct {
	credentialsProvider credentialsProvider
}

var credentialsChainOverride *CredentialsChainOverride

func GetCredentialsChainOverride() *CredentialsChainOverride {
	if credentialsChainOverride == nil {
		credentialsChainOverride = &CredentialsChainOverride{credentialsProvider: make([]credentials.Provider, 0)}
	}
	return credentialsChainOverride
}

func (c *CredentialsChainOverride) AppendCredentialsChain(credentialsProvider credentials.Provider) {
	c.credentialsProvider = append(c.credentialsProvider, credentialsProvider)
}

func (c *CredentialsChainOverride) GetCredentialsChain() []credentials.Provider {
	return c.credentialsProvider
}