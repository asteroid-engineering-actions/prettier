package actioncontext

import "context"

type ActionContext struct {
	Action           string
	ActionPath       string
	ActionRepository string
	Actor            string
	ActorId          string
	BaseRef          string
	Ci               bool
	EventName        string
	EventPath        string
	HeadRef          string
	Job              string
	GithubPath       string
	Ref              GithubRef
	Repository       GithubRepository
	//RetentionDays     int64
	//RunAttempt        int64
	//RunId             int64
	//RunNumber         int64
	//RunnerArch        string
	//RunnerDebug       bool
	//RunnerName        string
	//RunnerOs          string
	//RunnerTemp        string
	//RunnerToolCache   string
	//Sha               string
	//StepSummary       string
	//Workflow          string
	//WorkflowRef       string
	//WorkflowSha       string
	//Workspace         string
}

type GithubRef struct {
	FullRef      string
	RefName      string
	RefProtected bool
	RefType      string
}

type GithubRepository struct {
	Name    string
	Id      string
	Owner   string
	OwnerId string
}

var (
	Actions         bool
	ApiUrl          string
	GithubEnvPath   string
	GraphQlUrl      string
	GitHubServerUrl string
)

type key struct{}

var contextKey = &key{}

func init() {}

func NewContext(parent context.Context) context.Context {
	return nil
}
