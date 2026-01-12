# GitHub Copilot Dashboard Preview

This document provides a visual representation of the GitHub Copilot dashboards created.

## Dashboard 1: GitHub Copilot User Data Dashboard

### Layout Overview

```
┌────────────────────────────────────────────────────────────────────────────┐
│                        GITHUB COPILOT USER DATA DASHBOARD                   │
├────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │              📊 Overall Usage Statistics (Number Cards)              │  │
│  │                                                                       │  │
│  │  Active Users: 125    Accepted Lines: 45,230    Acceptance Rate: 68% │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ┌──────────────────┐ ┌──────────────────┐ ┌───────────────────────────┐ │
│  │  📈 AVERAGE      │ │  👥 ENGAGED      │ │   💻 EDITOR DISTRIBUTION  │ │
│  │  ACCEPTANCE RATE │ │  USERS %         │ │                           │ │
│  │                  │ │                  │ │      ┌─────────────┐      │ │
│  │      68%         │ │      73%         │ │      │ VS Code: 45%│      │ │
│  │                  │ │                  │ │      │ IntelliJ: 30│      │ │
│  │   🟢 Good        │ │   🟢 ↑ +5%       │ │      │ Vim: 15%    │      │ │
│  │                  │ │   (vs prev)      │ │      │ Other: 10%  │      │ │
│  └──────────────────┘ └──────────────────┘ └───────────────────────────┘ │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │           📈 Daily AI Code Line Changes (Time Series)                │  │
│  │  Lines                                                                │  │
│  │  5000 ┤                                      ╭─╮                      │  │
│  │  4000 ┤                          ╭─╮      ╭─╯ ╰─╮                    │  │
│  │  3000 ┤              ╭─╮      ╭─╯ ╰─╮  ╭─╯     ╰─╮                  │  │
│  │  2000 ┤      ╭─╮  ╭─╯ ╰─╮  ╭─╯     ╰──╯         ╰─╮                │  │
│  │  1000 ┤  ╭─╯ ╰──╯     ╰──╯                        ╰─╮              │  │
│  │     0 ┴──────────────────────────────────────────────────            │  │
│  │       Jan  Feb  Mar  Apr  May  Jun  Jul  Aug  Sep  Oct              │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │           🔄 Daily AI Interactions (Time Series)                      │  │
│  │  [Shows: Messages, Acceptances, Suggestions over time]               │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ┌──────────────────────────────────┐ ┌──────────────────────────────┐   │
│  │  🔍 Code Review Metrics          │ │  📊 Daily Acceptance Rate    │   │
│  │  (Time Series)                   │ │  (Time Series)               │   │
│  └──────────────────────────────────┘ └──────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │                     📋 User Interactions (Table)                      │  │
│  │ ┌────────┬──────────┬───────────┬───────────────┬─────────────────┐ │  │
│  │ │ User   │ Accepted │ Generated │ Acceptance %  │ Last Activity   │ │  │
│  │ ├────────┼──────────┼───────────┼───────────────┼─────────────────┤ │  │
│  │ │ Alice  │ 12,450   │ 15,200    │ 81.91%        │ 2026-01-12      │ │  │
│  │ │ Bob    │ 8,320    │ 14,100    │ 59.01%        │ 2026-01-11      │ │  │
│  │ │ Carol  │ 10,100   │ 13,800    │ 73.19%        │ 2026-01-12      │ │  │
│  │ └────────┴──────────┴───────────┴───────────────┴─────────────────┘ │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
└────────────────────────────────────────────────────────────────────────────┘
```

## Key Features Implemented

### 1. Average Acceptance Rate Card ✓
- **Location**: Second row, left position
- **Visualization**: Large number display (stat panel)
- **Metric**: `AVG(inline_acceptance_count / NULLIF(inline_suggestions_count, 0))`
- **Display Format**: Percentage (e.g., "68%")
- **Color Coding**:
  - 🔴 Red: < 30% (Low - needs improvement)
  - 🟡 Yellow: 30-60% (Moderate)
  - 🟢 Green: > 60% (Good)
- **Description**: "Average acceptance rate of AI suggestions over the selected time period"

### 2. Engaged Users % Card with Delta ✓
- **Location**: Second row, center position
- **Visualization**: Large number display with trend indicator (stat panel)
- **Metric**: `COUNT(DISTINCT CASE WHEN inline_acceptance_count > 0 THEN user_id END) / COUNT(DISTINCT user_id)`
- **Display Format**: Percentage with delta (e.g., "73% ↑ +5%")
- **Delta Feature**: Shows change from previous period (green up arrow for increase, red down arrow for decrease)
- **Color Coding**:
  - 🔴 Red: < 40% (Low engagement)
  - 🟡 Yellow: 40-70% (Moderate engagement)
  - 🟢 Green: > 70% (High engagement)
- **Description**: "Percentage of engaged users (users with at least one accepted suggestion) with change from previous period"

### 3. Editor Distribution Panel ✓
- **Location**: Second row, right position (spans wider)
- **Visualization**: Pie chart
- **Metric**: Distribution of users by editor/IDE
- **Display Format**: 
  - **Labels on pie slices**: Show percentages (e.g., "45%")
  - **Legend**: Shows both counts and percentages (e.g., "VS Code: 125 users (45%)")
  - **Tooltips**: When hovering over a slice, shows percentage
- **Configuration**:
  - `displayLabels: ["percent"]` - Percentages on pie slices
  - `legend.values: ["value", "percent"]` - Both metrics in legend
  - `tooltip.mode: "single"` - Shows details when highlighting
- **Description**: "Distribution of GitHub Copilot usage across different editors/IDEs. Values shown as percentages."

## Dashboard 2: GitHub Copilot + DORA Correlation

### Purpose
Correlates GitHub Copilot usage metrics with DORA (DevOps Research and Assessment) performance indicators to understand the impact of AI-assisted development on engineering efficiency.

### Layout Overview

```
┌────────────────────────────────────────────────────────────────────────────┐
│                  GITHUB COPILOT + DORA CORRELATION DASHBOARD                │
├────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │                        📝 Dashboard Introduction                      │  │
│  │  This dashboard correlates GitHub Copilot usage metrics with DORA    │  │
│  │  performance indicators to help understand the impact of AI-assisted │  │
│  │  development on engineering efficiency.                              │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ┌──────────────────────────────── Overview Statistics ────────────────┐  │
│  │                                                                       │  │
│  │  ┌────────────┐ ┌────────────┐ ┌────────────┐ ┌────────────┐       │  │
│  │  │👥 Active   │ │📝 AI Code  │ │✅ AI       │ │🚀 Deploy   │       │  │
│  │  │   Users    │ │   Lines    │ │ Accept %   │ │   Count    │       │  │
│  │  │    125     │ │   45,230   │ │    68%     │ │     142    │       │  │
│  │  └────────────┘ └────────────┘ └────────────┘ └────────────┘       │  │
│  │                                                                       │  │
│  │  ┌────────────┐ ┌────────────┐                                       │  │
│  │  │⏱️  Lead    │ │❌ Change   │                                       │  │
│  │  │   Time     │ │  Failure   │                                       │  │
│  │  │  4.2 hrs   │ │    8.5%    │                                       │  │
│  │  └────────────┘ └────────────┘                                       │  │
│  └───────────────────────────────────────────────────────────────────────┘│
│                                                                             │
│  ┌─────────────────── AI Usage vs DORA Metrics Correlation ────────────┐  │
│  │                                                                       │  │
│  │  📈 AI Code Generation vs Lead Time Trend                            │  │
│  │  📈 AI Acceptance Rate vs Deployment Frequency                       │  │
│  │  📈 AI Test Generation vs Change Failure Rate                        │  │
│  │  📈 GitHub Copilot Users vs Code Review Findings                     │  │
│  └───────────────────────────────────────────────────────────────────────┘│
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │              📋 Monthly GitHub Copilot vs DORA Comparison            │  │
│  │ ┌──────┬──────┬────────┬─────────┬───────┬──────────┬────────────┐ │  │
│  │ │Month │Users │AI Lines│Accept % │Deploy │Lead Time │Change Fail │ │  │
│  │ ├──────┼──────┼────────┼─────────┼───────┼──────────┼────────────┤ │  │
│  │ │Dec   │ 125  │ 45,230 │  68%    │  142  │  4.2 hrs │    8.5%    │ │  │
│  │ │Nov   │ 118  │ 42,100 │  65%    │  138  │  4.8 hrs │    9.2%    │ │  │
│  │ │Oct   │ 112  │ 38,900 │  63%    │  135  │  5.1 hrs │    9.8%    │ │  │
│  │ └──────┴──────┴────────┴─────────┴───────┴──────────┴────────────┘ │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
└────────────────────────────────────────────────────────────────────────────┘
```

## Technical Implementation Details

### SQL Query Examples

#### Average Acceptance Rate
```sql
SELECT
  AVG(inline_acceptance_count / NULLIF(inline_suggestions_count, 0)) as 'Average Acceptance Rate'
FROM lake._tool_github_copilot_user_data
WHERE $__timeFilter(date)
```

#### Engaged Users % (with time series for delta calculation)
```sql
SELECT
  DATE_FORMAT(date, '%Y-%m-%d') as time,
  COUNT(DISTINCT CASE WHEN inline_acceptance_count > 0 THEN user_id END) / 
    NULLIF(COUNT(DISTINCT user_id), 0) as 'Engaged Users %'
FROM lake._tool_github_copilot_user_data
WHERE $__timeFilter(date)
GROUP BY DATE_FORMAT(date, '%Y-%m-%d')
ORDER BY time
```

#### Editor Distribution
```sql
SELECT
  COALESCE(editor_name, 'Unknown') as editor,
  COUNT(DISTINCT user_id) as requests
FROM lake._tool_github_copilot_user_data
WHERE $__timeFilter(date)
GROUP BY editor_name
ORDER BY requests DESC
```

## Naming Conventions ✓

All instances use "GitHub Copilot" (not just "Copilot"):
- ✓ Dashboard titles
- ✓ Panel descriptions  
- ✓ Table references (`_tool_github_copilot_user_data`)
- ✓ File names (`GithubCopilot.json`, `GithubCopilotDORA.json`)
- ✓ UID identifiers (`github_copilot_user_data`, `github_copilot_dora`)

## Color Scheme & Thresholds

### Acceptance Rate Thresholds
Based on industry best practices for AI code completion tools:
- 🔴 **< 30%**: Low - Indicates low trust or poor suggestions
- 🟡 **30-60%**: Moderate - Room for improvement
- 🟢 **> 60%**: Good - Healthy adoption

### Engagement Thresholds
Based on typical enterprise software adoption rates:
- 🔴 **< 40%**: Low - Most users not engaging
- 🟡 **40-70%**: Moderate - Fair adoption
- 🟢 **> 70%**: High - Strong user engagement

## Installation & Usage

1. Import dashboards into Grafana
2. Configure MySQL datasource connection
3. Ensure `_tool_github_copilot_user_data` table is populated
4. Use time range picker to analyze different periods
5. Hover over Editor Distribution pie slices to see percentage breakdowns

## Data Requirements

The dashboards expect the following fields in `_tool_github_copilot_user_data`:
- `user_id`, `display_name`, `date`, `editor_name`
- `inline_acceptance_count`, `inline_suggestions_count`
- `inline_ai_code_lines`, `chat_ai_code_lines`
- Code review and fix metrics
- Other GitHub Copilot usage metrics
