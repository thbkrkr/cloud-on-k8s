// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:          "deployer",
		Short:        "Tool for interfacing with managed k8s cluster providers.",
		SilenceUsage: true,
	}

	rootCmd.AddCommand(ExecuteCommand())
	rootCmd.AddCommand(CreateCommand())
	rootCmd.AddCommand(GetCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerFileFlags(cmd *cobra.Command) (*string, *string) {
	var plansFile, configFile string
	cmd.PersistentFlags().StringVar(&plansFile, "base-config-dir", "./config", "Directory containing base configuration.")
	cmd.PersistentFlags().StringVar(&configFile, "config-file", "config/deployer-config.yml", "File containing run config.")
	return &plansFile, &configFile
}
