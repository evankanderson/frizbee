//
// Copyright 2023 Stacklok, Inc.
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

package ghactions

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/stacklok/frizbee/pkg/ghactions"
)

// CmdList represents the one sub-command
func CmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists the used github actions",
		Long: `This utility lists all the github actions used in the workflows

Example: 
	frizbee ghactions list
`,
		RunE:         list,
		SilenceUsage: true,
	}

	cmd.Flags().StringP("dir", "d", ".github/workflows", "workflows directory")

	return cmd
}

func list(cmd *cobra.Command, _ []string) error {
	dir := cmd.Flag("dir").Value.String()

	actions, err := ghactions.ListActionsInDirectory(dir)
	if err != nil {
		return fmt.Errorf("failed to list actions: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(actions, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal actions: %w", err)
	}
	jsonString := string(jsonBytes)

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", jsonString)
	return nil
}
