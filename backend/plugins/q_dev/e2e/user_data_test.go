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

package e2e

import (
	"testing"

	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/apache/incubator-devlake/helpers/e2ehelper"
	"github.com/apache/incubator-devlake/plugins/q_dev/impl"
	"github.com/apache/incubator-devlake/plugins/q_dev/models"
)

// TestQDevS3FileMeta tests the S3 file metadata model
func TestQDevS3FileMeta(t *testing.T) {
	var plugin impl.QDev
	dataflowTester := e2ehelper.NewDataFlowTester(t, "q_dev", plugin)

	// Flush tables before testing
	dataflowTester.FlushTabler(&models.QDevS3FileMeta{})

	// Import CSV data into the tool table
	dataflowTester.ImportCsvIntoTabler("./snapshot_tables/_tool_q_dev_s3_file_meta.csv", &models.QDevS3FileMeta{})

	// Verify the file meta data
	dataflowTester.VerifyTableWithOptions(
		models.QDevS3FileMeta{},
		e2ehelper.TableOptions{
			CSVRelPath:  "./snapshot_tables/_tool_q_dev_s3_file_meta.csv",
			IgnoreTypes: []interface{}{common.NoPKModel{}},
			IgnoreFields: []string{
				"processed_time",
			},
		},
	)
}

// TestQDevUserData tests the user data model
func TestQDevUserData(t *testing.T) {
	var plugin impl.QDev
	dataflowTester := e2ehelper.NewDataFlowTester(t, "q_dev", plugin)

	// Flush tables before testing
	dataflowTester.FlushTabler(&models.QDevUserData{})

	// Import CSV data into the tool table
	dataflowTester.ImportCsvIntoTabler("./snapshot_tables/_tool_q_dev_user_data.csv", &models.QDevUserData{})

	// Verify the user data
	dataflowTester.VerifyTableWithOptions(
		models.QDevUserData{},
		e2ehelper.TableOptions{
			CSVRelPath:  "./snapshot_tables/_tool_q_dev_user_data.csv",
			IgnoreTypes: []interface{}{common.Model{}},
			IgnoreFields: []string{
				"created_at",
				"updated_at",
				// Ignore all other metrics fields that are not in our test CSV
				"code_review_succeeded_event_count",
				"inline_chat_dismissal_event_count",
				"inline_chat_dismissed_line_additions",
				"inline_chat_dismissed_line_deletions",
				"inline_chat_rejected_line_additions",
				"inline_chat_rejected_line_deletions",
				"inline_chat_rejection_event_count",
				"inline_chat_accepted_line_deletions",
				"code_review_failed_event_count",
				"dev_acceptance_event_count",
				"dev_accepted_lines",
				"dev_generated_lines",
				"dev_generation_event_count",
				"doc_generation_accepted_file_updates",
				"doc_generation_accepted_files_creations",
				"doc_generation_accepted_line_additions",
				"doc_generation_accepted_line_updates",
				"doc_generation_event_count",
				"doc_generation_rejected_file_creations",
				"doc_generation_rejected_file_updates",
				"doc_generation_rejected_line_additions",
				"doc_generation_rejected_line_updates",
				"test_generation_accepted_lines",
				"test_generation_accepted_tests",
				"test_generation_event_count",
				"test_generation_generated_lines",
				"test_generation_generated_tests",
				"transformation_event_count",
				"transformation_lines_generated",
				"transformation_lines_ingested",
			},
		},
	)
}
