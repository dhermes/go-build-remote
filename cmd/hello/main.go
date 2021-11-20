// Copyright 2021 Danny Hermes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	Anything string
}

func run() error {
	c := Config{}
	cmd := &cobra.Command{
		Use:           "hello",
		Short:         "No-op cobra CLI",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Printf("c = %#v\n", c)
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(
		&c.Anything,
		"anything",
		c.Anything,
		"Anything goes...here",
	)

	required := []string{"anything"}
	for _, name := range required {
		err := cobra.MarkFlagRequired(cmd.PersistentFlags(), name)
		if err != nil {
			return err
		}
	}
	return cmd.Execute()
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
