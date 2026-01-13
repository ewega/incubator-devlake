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

package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

const RAW_USAGE_TABLE = "github_copilot_usage"

var _ plugin.SubTaskEntryPoint = CollectUsage

var CollectUsageMeta = plugin.SubTaskMeta{
	Name:             "CollectUsage",
	EntryPoint:       CollectUsage,
	EnabledByDefault: true,
	Description:      "Collect GitHub Copilot usage metrics from API",
	DomainTypes:      []string{},
}

func CollectUsage(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*GithubCopilotTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("Collecting GitHub Copilot usage data for organization: %s", data.Options.OrganizationName)

	// Build API URL template
	var urlTemplate string
	if data.Options.EnterpriseName != "" {
		urlTemplate = fmt.Sprintf("/enterprises/%s/copilot/usage", data.Options.EnterpriseName)
	} else {
		urlTemplate = fmt.Sprintf("/orgs/%s/copilot/usage", data.Options.OrganizationName)
	}

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: GithubCopilotApiParams{
				ConnectionId:     data.Options.ConnectionId,
				OrganizationName: data.Options.OrganizationName,
			},
			Table: RAW_USAGE_TABLE,
		},
		ApiClient:   data.ApiClient,
		UrlTemplate: urlTemplate,
		PageSize:    100,
		Query: func(reqData *helper.RequestData) (url.Values, errors.Error) {
			query := url.Values{}
			if data.Options.Since != nil {
				query.Set("since", data.Options.Since.Format("2006-01-02"))
			}
			query.Set("per_page", "100")
			if reqData.Pager != nil {
				query.Set("page", fmt.Sprintf("%v", reqData.Pager.Page))
			}
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var items []json.RawMessage
			err := helper.UnmarshalResponse(res, &items)
			if err != nil {
				return nil, err
			}
			return items, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}

// GithubCopilotApiParams holds the API params for raw data
type GithubCopilotApiParams struct {
	ConnectionId     uint64
	OrganizationName string
}
