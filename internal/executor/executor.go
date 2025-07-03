package executor

import (
	"fmt"

	"github.com/zinrai/gcloud-configurations-loader/internal/config"
	"github.com/zinrai/gcloud-configurations-loader/internal/gcloud"
)

type Executor struct {
	gcloudManager *gcloud.Manager
	planner       *Planner
	reporter      *Reporter
}

func NewExecutor(gcloudManager *gcloud.Manager) *Executor {
	return &Executor{
		gcloudManager: gcloudManager,
		planner:       NewPlanner(gcloudManager),
		reporter:      NewReporter(),
	}
}

// Applies the configurations according to the execution plan
func (e *Executor) Execute(configs []config.Configuration, replaceMode, dryRun, verbose bool) error {
	plan := e.planner.AnalyzeConfigurations(configs, replaceMode)

	if dryRun {
		e.reporter.PrintDryRun(plan)
		return nil
	}

	if verbose {
		e.reporter.PrintPlan(plan)
	}

	// Execute new configurations
	for _, cfg := range plan.ToCreate {
		if err := e.createConfiguration(cfg, verbose); err != nil {
			return fmt.Errorf("failed to create configuration %s: %w", cfg.Name, err)
		}
	}

	// Execute replacements
	for _, cfg := range plan.ToReplace {
		if err := e.replaceConfiguration(cfg, verbose); err != nil {
			return fmt.Errorf("failed to replace configuration %s: %w", cfg.Name, err)
		}
	}

	// Report skipped configurations
	if len(plan.ToSkip) > 0 {
		e.reporter.PrintSkipped(plan.ToSkip)
	}

	e.reporter.PrintSummary(len(plan.ToCreate), len(plan.ToReplace), len(plan.ToSkip))
	return nil
}

// Creates a new configuration and sets its properties
func (e *Executor) createConfiguration(cfg config.Configuration, verbose bool) error {
	if verbose {
		fmt.Printf("Creating configuration: %s\n", cfg.Name)
	}

	if err := e.gcloudManager.CreateConfiguration(cfg.Name); err != nil {
		return err
	}

	return e.setConfigurationProperties(cfg, verbose)
}

// Replaces an existing configuration
func (e *Executor) replaceConfiguration(cfg config.Configuration, verbose bool) error {
	if verbose {
		fmt.Printf("Replacing configuration: %s\n", cfg.Name)
	}

	if err := e.gcloudManager.DeleteConfiguration(cfg.Name); err != nil {
		return err
	}

	if err := e.gcloudManager.CreateConfiguration(cfg.Name); err != nil {
		return err
	}

	return e.setConfigurationProperties(cfg, verbose)
}

// Sets all properties for a configuration
func (e *Executor) setConfigurationProperties(cfg config.Configuration, verbose bool) error {
	for key, value := range cfg.Properties {
		if verbose {
			fmt.Printf("  Setting %s = %s\n", key, value)
		}
		if err := e.gcloudManager.SetConfigProperty(cfg.Name, key, value); err != nil {
			return fmt.Errorf("failed to set property %s=%s: %w", key, value, err)
		}
	}
	return nil
}
