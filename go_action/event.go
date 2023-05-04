package main

// https://github.com/actions/toolkit/blob/main/packages/github/src/context.ts
// https://docs.github.com/en/actions/learn-github-actions/environment-variables

import (
	"encoding/json"
	"github.com/google/go-github/v50/github"
	"os"
	"strconv"
)

type getEnvFunc func(key string) string

type actionEvent struct {
	getEnv            getEnvFunc
	Action            string
	ActionPath        string
	ActionRepository  string
	Actions           bool
	Actor             string
	ActorId           string
	ApiUrl            string
	BaseRef           string
	GithubEnvPath     string
	EventName         string
	EventPath         string
	RawEventPayload   []byte
	EventPayload      interface{}
	GraphQlUrl        string
	HeadRef           string
	Job               string
	Path              string
	Ref               string
	RefName           string
	RefProtected      bool
	RefType           string
	Repository        string
	RepositoryId      string
	RepositoryOwner   string
	RepositoryOwnerId string
	RetentionDays     int64
	RunAttempt        int64
	RunId             int64
	RunNumber         int64
	RunnerArch        string
	RunnerDebug       bool
	RunnerName        string
	RunnerOs          string
	RunnerTemp        string
	RunnerToolCache   string
	ServerUrl         string
	Sha               string
	StepSummary       string
	Workflow          string
	WorkflowRef       string
	WorkflowSha       string
	Workspace         string
}

func (a *actionEvent) initializeActionContext() *actionEvent {
	if a.getEnv == nil {
		a.getEnv = os.Getenv
	}

	a.Action = a.getEnv("GITHUB_ACTION")
	a.ActionPath = a.getEnv("GITHUB_ACTION_PATH")
	a.ActionRepository = a.getEnv("GITHUB_ACTION_REPOSITORY")
	a.Actions, _ = strconv.ParseBool(a.getEnv("GITHUB_ACTIONS"))
	a.Actor = a.getEnv("GITHUB_ACTOR")
	a.ActorId = a.getEnv("GITHUB_ACTOR_ID")
	a.ApiUrl = a.getEnv("GITHUB_API_URL")
	a.BaseRef = a.getEnv("GITHUB_BASE_REF")
	a.GithubEnvPath = a.getEnv("GITHUB_ENV")
	a.EventName = a.getEnv("GITHUB_EVENT_NAME")
	a.EventPath = a.getEnv("GITHUB_EVENT_PATH")
	a.GraphQlUrl = a.getEnv("GITHUB_GRAPHQL_URL")
	a.HeadRef = a.getEnv("GITHUB_HEAD_REF")
	a.Job = a.getEnv("GITHUB_JOB")
	a.Path = a.getEnv("GITHUB_PATH")
	a.Ref = a.getEnv("GITHUB_REF")
	a.RefName = a.getEnv("GITHUB_REF_NAME")
	a.RefProtected, _ = strconv.ParseBool(a.getEnv("GITHUB_REF_PROTECTED"))
	a.RefType = a.getEnv("GITHUB_REF_TYPE")
	a.Repository = a.getEnv("GITHUB_REPOSITORY")
	a.RepositoryId = a.getEnv("GITHUB_REPOSITORY_ID")
	a.RepositoryOwner = a.getEnv("GITHUB_REPOSITORY_OWNER")
	a.RepositoryOwnerId = a.getEnv("GITHUB_REPOSITORY_OWNER_ID")
	a.RetentionDays, _ = strconv.ParseInt(a.getEnv("GITHUB_RETENTION_DATS"), 10, 64)
	a.RunAttempt, _ = strconv.ParseInt(a.getEnv("GITHUB_RUN_ATTEMPT"), 10, 64)
	a.RunId, _ = strconv.ParseInt(a.getEnv("GITHUB_RUN_ID"), 10, 64)
	a.RunNumber, _ = strconv.ParseInt(a.getEnv("GITHUB_RUN_NUMBER"), 10, 64)
	a.RunnerArch = a.getEnv("RUNNER_ARCH")
	a.RunnerDebug, _ = strconv.ParseBool(a.getEnv("RUNNER_DEBUG"))
	a.RunnerName = a.getEnv("RUNNER_NAME")
	a.RunnerOs = a.getEnv("RUNNER_OS")
	a.RunnerTemp = a.getEnv("RUNNER_TEMP")
	a.RunnerToolCache = a.getEnv("RUNNER_TOOL_CACHE")
	a.ServerUrl = a.getEnv("GITHUB_SERVER_URL")
	a.Sha = a.getEnv("GITHUB_SHA")
	a.StepSummary = a.getEnv("GITHUB_STEP_SUMMARY")
	a.Workflow = a.getEnv("GITHUB_WORKFLOW")
	a.WorkflowRef = a.getEnv("GITHUB_WORKFLOW_REF")
	a.WorkflowSha = a.getEnv("GITHUB_WORKFLOW_SHA")
	a.Workspace = a.getEnv("GITHUB_WORKSPACE")

	a.RawEventPayload = loadRawPayload(a.EventPath)
	// https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows
	switch a.EventName {
	case "branch_protection_rule":
		a.EventPayload = &github.BranchProtectionRuleEvent{}
	case "check_run":
		a.EventPayload = &github.CheckRunEvent{}
	case "check_suite":
		a.EventPayload = &github.CheckSuiteEvent{}
	case "create":
		a.EventPayload = &github.CreateEvent{}
	case "delete":
		a.EventPayload = &github.DeleteEvent{}
	case "deployment":
		a.EventPayload = &github.DeploymentEvent{}
	case "deployment_status":
		a.EventPayload = &github.DeploymentStatusEvent{}
	case "discussion":
		a.EventPayload = &github.DiscussionEvent{}
	case "discussion_comment":
		a.EventPayload = &github.DiscussionCommentEvent{}
	case "fork":
		a.EventPayload = &github.ForkEvent{}
	case "gollum":
		a.EventPayload = &github.GollumEvent{}
	case "issue_comment":
		a.EventPayload = &github.IssueCommentEvent{}
	case "issues":
		a.EventPayload = &github.IssuesEvent{}
	case "label":
		a.EventPayload = &github.LabelEvent{}
	case "merge_group":
		a.EventPayload = &github.MergeGroupEvent{}
	case "milestone":
		a.EventPayload = &github.MilestoneEvent{}
	case "page_build":
		a.EventPayload = &github.PageBuildEvent{}
	case "project":
		a.EventPayload = &github.ProjectEvent{}
	case "project_card":
		a.EventPayload = &github.ProjectCard{}
	case "project_column":
		a.EventPayload = &github.ProjectColumnEvent{}
	case "public":
		a.EventPayload = &github.PublicEvent{}
	case "pull_request":
		a.EventPayload = &github.PullRequestEvent{}
	case "pull_request_comment":
		// https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#pull_request_comment-use-issue_comment
		a.EventPayload = &github.IssueCommentEvent{}
	case "pull_request_review":
		a.EventPayload = &github.PullRequestReviewEvent{}
	case "pull_request_review_comment":
		a.EventPayload = &github.PullRequestReviewCommentEvent{}
	case "pull_request_target":
		a.EventPayload = &github.PullRequestTargetEvent{}
	case "push":
		a.EventPayload = &github.PushEvent{}
	case "registry_package":
		a.EventPayload = &github.PackageEvent{}
	case "release":
		a.EventPayload = &github.ReleaseEvent{}
	case "repository_dispatch":
		a.EventPayload = &github.RepositoryDispatchEvent{}
	//case "schedule":
	case "status":
		a.EventPayload = &github.StatusEvent{}
	case "watch":
		a.EventPayload = &github.WatchEvent{}
	//case "workflow_call":
	// The workflow being called is not aware it is being called from a separate workflow -- this could be inferred via a variable (like GITHUB_WORKFLOW) which identifies a workflow
	case "workflow_dispatch":
		a.EventPayload = &github.WorkflowDispatchEvent{}
	case "workflow_run":
		a.EventPayload = &github.WorkflowRunEvent{}
	}

	if a.EventPayload != nil {
		err := json.Unmarshal(a.RawEventPayload, &a.EventPayload)

		if err != nil {
			panic(err)
		}
	}

	return a
}

func loadRawPayload(eventPath string) []byte {
	if eventPath == "" {
		return nil
	}

	rawEventPayload, err := os.ReadFile(eventPath)

	if err != nil {
		return nil
	}

	return rawEventPayload
}
