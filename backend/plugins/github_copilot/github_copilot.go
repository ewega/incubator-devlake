/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/apache/incubator-devlake/core/runner"
	"github.com/apache/incubator-devlake/plugins/github_copilot/impl"
	"github.com/spf13/cobra"
)

// PluginEntry exports a symbol for the plugin
var PluginEntry impl.GithubCopilot //nolint

// standalone mode for debugging
func main() {
	cmd := &cobra.Command{Use: "github_copilot"}
	organizationName := cmd.Flags().StringP("organization-name", "o", "", "github organization name")
	connectionId := cmd.Flags().Uint64P("connection-id", "c", 0, "github copilot connection id")
	timeAfter := cmd.Flags().StringP("time-after", "a", "", "collect data that are created after specified time, ie 2006-01-02")
	_ = cmd.MarkFlagRequired("organization-name")
	_ = cmd.MarkFlagRequired("connection-id")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		runner.DirectRun(cmd, args, PluginEntry, map[string]interface{}{
			"organizationName": *organizationName,
			"connectionId":     *connectionId,
		}, *timeAfter)
	}

	runner.RunCmd(cmd)
}
