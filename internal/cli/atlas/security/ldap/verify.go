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

package ldap

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type VerifyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	hostname           string
	port               int
	bindUsername       string
	bindPassword       string
	caCertificate      string
	authzQueryTemplate string
	store              store.LDAPConfigurationVerifier
}

func (opts *VerifyOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()))
	return err
}

var verifyTemplate = `REQUEST ID	PROJECT ID	STATUS
{{.RequestID}}	{{.GroupID}}	{{.Status}}
`

func (opts *VerifyOpts) Run() error {
	r, err := opts.store.VerifyLDAPConfiguration(opts.ConfigProjectID(), opts.newLDAP())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *VerifyOpts) newLDAP() *atlas.LDAP {
	return &atlas.LDAP{
		Hostname:           opts.hostname,
		Port:               opts.port,
		BindUsername:       opts.bindUsername,
		BindPassword:       opts.bindPassword,
		CaCertificate:      opts.caCertificate,
		AuthzQueryTemplate: opts.authzQueryTemplate,
	}
}

// mongocli security atlas ldap verify --hostname hostname --port port --bindUsername bindUsername --bindPassword bindPassword --caCertificate caCertificate --authzQueryTemplate authzQueryTemplate [--projectId projectId].
func VerifyBuilder() *cobra.Command {
	opts := &VerifyOpts{}
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Request verification of an LDAP configuration.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), verifyTemplate))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.hostname, flag.Hostname, "", usage.LDAPHostname)
	cmd.Flags().IntVar(&opts.port, flag.Port, defaultLDAPPort, usage.LDAPPort)
	cmd.Flags().StringVar(&opts.bindUsername, flag.BindUsername, "", usage.BindUsername)
	cmd.Flags().StringVar(&opts.bindPassword, flag.BindPassword, "", usage.BindPassword)
	cmd.Flags().StringVar(&opts.caCertificate, flag.CaCertificate, "", usage.CaCertificate)
	cmd.Flags().StringVar(&opts.authzQueryTemplate, flag.AuthzQueryTemplate, "", usage.AuthzQueryTemplate)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Hostname)
	_ = cmd.MarkFlagRequired(flag.BindUsername)
	_ = cmd.MarkFlagRequired(flag.BindPassword)

	cmd.AddCommand(
		StatusBuilder(),
	)

	return cmd
}
