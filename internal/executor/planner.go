package executor

import (
	"github.com/zinrai/gcloud-configurations-loader/internal/config"
	"github.com/zinrai/gcloud-configurations-loader/internal/gcloud"
)

// Represents what actions will be taken
type ExecutionPlan struct {
	ToCreate  []config.Configuration // New configurations to create
	ToSkip    []config.Configuration // Existing configurations to skip
	ToReplace []config.Configuration // Existing configurations to replace
}

type Planner struct {
	gcloudManager *gcloud.Manager
}

func NewPlanner(gcloudManager *gcloud.Manager) *Planner {
	return &Planner{
		gcloudManager: gcloudManager,
	}
}

// Creates an execution plan based on current state
func (p *Planner) AnalyzeConfigurations(configs []config.Configuration, replaceMode bool) ExecutionPlan {
	var plan ExecutionPlan

	for _, cfg := range configs {
		if p.gcloudManager.ConfigurationExists(cfg.Name) {
			if replaceMode {
				plan.ToReplace = append(plan.ToReplace, cfg)
			} else {
				plan.ToSkip = append(plan.ToSkip, cfg)
			}
		} else {
			plan.ToCreate = append(plan.ToCreate, cfg)
		}
	}

	return plan
}
