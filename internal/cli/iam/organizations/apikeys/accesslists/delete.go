// Copyright 2020 MongoDB Inc
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

package accesslists

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
	cli.GlobalOpts
	apiKey string
	store  store.OrganizationAPIKeyAccessListWhitelistDeleter
}

func (opts *DeleteOpts) init() error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()))
	return err
}

func (opts *DeleteOpts) Run() error {
	useAccessList, err := shouldUseAccessList(opts.store)
	if err != nil {
		return err
	}

	var f func(string, string, string) error
	if useAccessList {
		f = opts.store.DeleteOrganizationAPIKeyAccessList
	} else {
		f = opts.store.DeleteOrganizationAPIKeyWhitelist
	}

	return opts.Delete(f, opts.ConfigOrgID(), opts.apiKey)
}

// mongocli iam organizations|orgs apiKey(s)|apikey(s) accesslist delete <IP> [--orgId orgId] [--apiKey apiKey] --force.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Access list entry '%s' deleted\n", "Access list entry not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Aliases: []string{"rm"},
		Short:   "Delete an IP access list from your API Key.",
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateOrgID, opts.init); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.apiKey, flag.APIKey, "", usage.APIKey)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	return cmd
}
