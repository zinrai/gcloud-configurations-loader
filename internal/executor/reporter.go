package executor

import (
	"fmt"

	"github.com/zinrai/gcloud-configurations-loader/internal/config"
)

type Reporter struct{}

func NewReporter() *Reporter {
	return &Reporter{}
}

// Prints what would be done without executing
func (r *Reporter) PrintDryRun(plan ExecutionPlan) {
	fmt.Println("Dry run - no changes will be made:")
	fmt.Println()

	if len(plan.ToCreate) > 0 {
		fmt.Println("The following configurations would be created:")
		for _, cfg := range plan.ToCreate {
			fmt.Printf("  * %s\n", cfg.Name)
		}
		fmt.Println()
	}

	if len(plan.ToReplace) > 0 {
		fmt.Println("The following configurations would be replaced:")
		for _, cfg := range plan.ToReplace {
			fmt.Printf("  - %s\n", cfg.Name)
		}
		fmt.Println()
	}

	if len(plan.ToSkip) > 0 {
		fmt.Println("The following configurations would be skipped:")
		for _, cfg := range plan.ToSkip {
			fmt.Printf("  | %s (use -replace to update)\n", cfg.Name)
		}
		fmt.Println()
	}

	fmt.Printf("Summary: %d to create, %d to replace, %d to skip\n",
		len(plan.ToCreate), len(plan.ToReplace), len(plan.ToSkip))
}

// Prints the execution plan
func (r *Reporter) PrintPlan(plan ExecutionPlan) {
	if len(plan.ToCreate) > 0 {
		fmt.Println("Creating configurations:")
		for _, cfg := range plan.ToCreate {
			fmt.Printf("  %s\n", cfg.Name)
		}
	}

	if len(plan.ToReplace) > 0 {
		fmt.Println("Replacing configurations:")
		for _, cfg := range plan.ToReplace {
			fmt.Printf("  %s\n", cfg.Name)
		}
	}
	fmt.Println()
}

// Prints skipped configurations
func (r *Reporter) PrintSkipped(skipped []config.Configuration) {
	fmt.Println()
	for _, cfg := range skipped {
		fmt.Printf("| Skipped existing configuration: %s (use -replace to update)\n", cfg.Name)
	}
}

// Prints the execution summary
func (r *Reporter) PrintSummary(created, replaced, skipped int) {
	fmt.Println()
	if created > 0 {
		fmt.Printf("* Created %d configuration(s)\n", created)
	}
	if replaced > 0 {
		fmt.Printf("- Replaced %d configuration(s)\n", replaced)
	}
	if skipped > 0 {
		fmt.Printf("| Skipped %d existing configuration(s)\n", skipped)
	}

	fmt.Printf("\nSummary: %d created, %d replaced, %d skipped\n", created, replaced, skipped)
}
