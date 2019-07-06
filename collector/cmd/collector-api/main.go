package main

import (
	"errors"
	"os"

	"github.com/deislabs/cnab-dashboard/collector/pkg/api"
	"github.com/deislabs/cnab-dashboard/collector/pkg/collector"
	"github.com/spf13/cobra"
)

func main() {
	var opts collector.Options

	cmd := cobra.Command{
		Use:   "collector-api CLAIM_SOURCES",
		Short: "Serve statistics about a set of CNAB claim sources",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("CLAIM_SOURCES is a required argument")
			}
			opts.ClaimSources = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return api.Run(opts)
		},
	}

	cmd.Flags().StringVar(&opts.Recent, "recent", "7d", "Time since a bundle was touched to consider it 'recent'")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
