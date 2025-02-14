// Copyright 2021 MongoDB Inc
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

package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_regions.go -package=mocks github.com/mongodb/mongocli/internal/store CloudProviderRegionsLister

type CloudProviderRegionsLister interface {
	CloudProviderRegions(string, string, []*string) (*atlas.CloudProviders, error)
}

// CloudProviderRegions encapsulates the logic to manage different cloud providers.
func (s *Store) CloudProviderRegions(projectID, tier string, providerName []*string) (*atlas.CloudProviders, error) {
	options := &atlas.CloudProviderRegionsOptions{
		Providers: providerName,
		Tier:      tier,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).Clusters.ListCloudProviderRegions(context.Background(), projectID, options)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
