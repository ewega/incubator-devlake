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

package api

import (
	"github.com/apache/incubator-devlake/core/errors"
	coreModels "github.com/apache/incubator-devlake/core/models"
	"github.com/apache/incubator-devlake/core/plugin"
)

// MakePipelinePlanV200 creates a pipeline plan for the github_copilot plugin
func MakePipelinePlanV200(
	subtaskMetas []plugin.SubTaskMeta,
	connectionId uint64,
	scopes []*coreModels.BlueprintScope,
) (coreModels.PipelinePlan, []plugin.Scope, errors.Error) {
	// For now, return a simple plan with just the subtasks
	// This can be expanded later to support more complex pipeline generation
	var plan coreModels.PipelinePlan
	var pluginScopes []plugin.Scope
	
	// Create a single stage with all subtasks
	if len(scopes) > 0 {
		stage := make(coreModels.PipelineStage, 0)
		for _, scope := range scopes {
			// Create task options from scope
			options := make(map[string]interface{})
			options["connectionId"] = connectionId
			options["organizationName"] = scope.ScopeId
			
			// Convert subtask metas to subtask names
			subtaskNames := make([]string, len(subtaskMetas))
			for i, meta := range subtaskMetas {
				subtaskNames[i] = meta.Name
			}
			
			stage = append(stage, &coreModels.PipelineTask{
				Plugin:   "github_copilot",
				Subtasks: subtaskNames,
				Options:  options,
			})
		}
		plan = append(plan, stage)
	}

	return plan, pluginScopes, nil
}
