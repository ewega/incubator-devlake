# GitHub Copilot Metrics Plugin - Implementation Tasks

## Overview
This document outlines the phased implementation plan for the GitHub Copilot Metrics Plugin for Apache DevLake. The plugin will collect, transform, and visualize GitHub Copilot usage metrics.

## Phase 1: Core Plugin Infrastructure

### Task 1.1: Create Plugin Skeleton
- [ ] Create `backend/plugins/github_copilot/` directory structure
- [ ] Implement basic plugin interfaces (PluginMeta, PluginTask, PluginModel, PluginMigration)
- [ ] Set up plugin registration in main plugin registry
- [ ] Add plugin to build configuration

### Task 1.2: Define Data Models
- [ ] Create tool layer models for Copilot metrics data
- [ ] Define migration scripts for database schema
- [ ] Implement GetTablesInfo() method listing all models
- [ ] Register migration scripts in register.go

### Task 1.3: API Connection Setup
- [ ] Implement connection API endpoints
- [ ] Add authentication/token management for GitHub Copilot API
- [ ] Create connection test functionality
- [ ] Add scope configuration support

## Phase 2: Data Collection

### Task 2.1: Implement Copilot Metrics Collector
- [ ] Create collector for Copilot usage metrics API endpoint
- [ ] Implement incremental collection with time-based bookmarking
- [ ] Add rate limiting and error handling
- [ ] Create unit tests for collector

### Task 2.2: Implement Data Extraction
- [ ] Create extractor to parse raw API responses
- [ ] Transform API data to tool layer models
- [ ] Handle pagination and batch processing
- [ ] Add extractor unit tests

### Task 2.3: Implement Domain Layer Conversion
- [ ] Create converter to map tool models to domain models
- [ ] Implement any custom domain models needed for Copilot metrics
- [ ] Add converter unit tests
- [ ] Integrate with existing domain layer

## Phase 3: Metrics Calculations

### Task 3.1: Calculate Acceptance Rate Metrics
- [ ] Implement average acceptance rate calculation for Copilot suggestions
- [ ] Add daily/weekly/monthly aggregations
- [ ] Create metrics card for acceptance rate visualization
- [ ] Add SQL queries for Grafana dashboard

### Task 3.2: Calculate Engaged Users Metrics
- [ ] Implement engaged users count (users with active Copilot usage)
- [ ] Add time-based filtering for engaged user metrics
- [ ] Create metrics card for engaged users visualization
- [ ] Add SQL queries for Grafana dashboard

### Task 3.3: Additional Metrics
- [ ] Lines of code suggested by Copilot
- [ ] Lines of code accepted from Copilot
- [ ] Suggestion acceptance by language
- [ ] Suggestion acceptance by editor/IDE

## Phase 4: Grafana Dashboard

### Task 4.1: Create Base Dashboard
- [ ] Create new Grafana dashboard JSON file for Copilot metrics
- [ ] Set up dashboard layout and panels
- [ ] Configure data source connections
- [ ] Add time range selectors and filters

### Task 4.2: Add Metrics Visualizations
- [ ] Create panel for average acceptance rate metric
- [ ] Create panel for engaged users count
- [ ] Add line charts for trend analysis
- [ ] Add bar charts for breakdown by language/editor

### Task 4.3: Modify Editor Distribution Dashboard
- [ ] Update existing editor distribution visualizations to show percentages
- [ ] Add percentage calculations to SQL queries
- [ ] Format display to show both counts and percentages
- [ ] Test dashboard with sample data

### Task 4.4: Dashboard Polish
- [ ] Add dashboard descriptions and documentation
- [ ] Configure panel tooltips and legends
- [ ] Set appropriate color schemes and thresholds
- [ ] Add drill-down capabilities where applicable

## Phase 5: UI Integration

### Task 5.1: Add Copilot Connection UI
- [ ] Create connection configuration form in config-ui
- [ ] Add validation for Copilot API tokens
- [ ] Implement connection testing in UI
- [ ] Add help text and documentation links

### Task 5.2: Add Copilot Scope Selection
- [ ] Create scope selection interface for organizations/repositories
- [ ] Add scope preview/validation
- [ ] Implement scope configuration UI
- [ ] Add bulk selection capabilities

### Task 5.3: Update Navigation
- [ ] Add Copilot plugin to data sources menu
- [ ] Update connection list to include Copilot connections
- [ ] Add appropriate icons and labels
- [ ] Update documentation

## Phase 6: Testing & Documentation

### Task 6.1: E2E Tests
- [ ] Create E2E test fixtures with sample Copilot data
- [ ] Implement collection E2E tests
- [ ] Test full data pipeline (collect → extract → convert)
- [ ] Verify dashboard queries return correct results

### Task 6.2: Documentation
- [ ] Write plugin README with setup instructions
- [ ] Document API requirements and permissions
- [ ] Create user guide for dashboard usage
- [ ] Add troubleshooting section

### Task 6.3: Integration Testing
- [ ] Test with real GitHub Copilot API (staging)
- [ ] Verify data accuracy and completeness
- [ ] Test error scenarios and edge cases
- [ ] Performance testing with large datasets

## Phase 7: Naming Consistency

### Task 7.1: Review and Rename Instances
- [ ] Search codebase for all instances of "Copilot" 
- [ ] Ensure consistent naming convention (e.g., "GitHub Copilot" vs "Copilot")
- [ ] Update variable names, comments, and strings
- [ ] Update UI labels and display names

### Task 7.2: Update Documentation
- [ ] Review all documentation for naming consistency
- [ ] Update dashboard titles and descriptions
- [ ] Ensure API documentation uses consistent terminology
- [ ] Update error messages and logging

## Additional Tasks (New Requirements)

### Task A1: Average Acceptance Rate Metric Card
- [ ] Design metric card layout for acceptance rate
- [ ] Implement SQL query for average acceptance rate across all users
- [ ] Add comparison to previous period (week/month)
- [ ] Display as percentage with trend indicator
- [ ] Add tooltip with detailed breakdown

### Task A2: Engaged Users Metric Card
- [ ] Design metric card layout for engaged users count
- [ ] Implement SQL query for distinct active users
- [ ] Define "engaged" criteria (e.g., minimum suggestions accepted)
- [ ] Add comparison to previous period
- [ ] Display with growth/decline indicator

### Task A3: Editor Distribution Percentage Display
- [ ] Modify existing editor distribution SQL queries to include percentage
- [ ] Update dashboard panel to display both count and percentage
- [ ] Format percentage display (e.g., "123 (45.6%)")
- [ ] Add total count display
- [ ] Update panel title and description

### Task A4: Copilot Naming Standardization
- [ ] Create style guide for "GitHub Copilot" terminology
- [ ] Search and replace inconsistent usage across:
  - [ ] Dashboard titles and panel names
  - [ ] SQL query comments and aliases
  - [ ] UI component labels
  - [ ] Code comments and documentation
  - [ ] Error messages and logs
- [ ] Review PR for any remaining inconsistencies

## Notes
- Each task should be implemented as a separate commit where practical
- All code must include appropriate tests
- Documentation should be updated alongside code changes
- Follow DevLake coding conventions and Apache license requirements
- Ensure backward compatibility where applicable
