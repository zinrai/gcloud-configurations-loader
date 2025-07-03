package cmd

import (
	"flag"
	"fmt"

	"github.com/zinrai/gcloud-configurations-loader/internal/config"
	"github.com/zinrai/gcloud-configurations-loader/internal/executor"
	"github.com/zinrai/gcloud-configurations-loader/internal/gcloud"
)

var (
	configFile = flag.String("config", "config.yaml", "Configuration file path")
	replace    = flag.Bool("replace", false, "Replace existing configurations")
	dryRun     = flag.Bool("dry-run", false, "Show what would be done without executing")
	verbose    = flag.Bool("verbose", false, "Verbose output")
	help       = flag.Bool("help", false, "Show help message")
)

// Entry point for the CLI
func Execute() error {
	flag.Parse()

	if *help {
		printHelp()
		return nil
	}

	// Load and validate configuration file
	configData, err := config.LoadConfig(*configFile)
	if err != nil {
		return fmt.Errorf("failed to load config file %s: %w", *configFile, err)
	}

	if err := config.ValidateConfigFile(configData); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	// Initialize components
	gcloudManager := gcloud.NewManager()
	exec := executor.NewExecutor(gcloudManager)

	// Execute configurations
	return exec.Execute(configData.Configurations, *replace, *dryRun, *verbose)
}

func printHelp() {
	fmt.Println("gcloud-configurations-loader - Manage gcloud configurations from YAML")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  gcloud-configurations-loader [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Apply configurations (create new, skip existing)")
	fmt.Println("  gcloud-configurations-loader -config config.yaml")
	fmt.Println()
	fmt.Println("  # Replace existing configurations")
	fmt.Println("  gcloud-configurations-loader -config config.yaml -replace")
	fmt.Println()
	fmt.Println("  # Dry run to see what would be done")
	fmt.Println("  gcloud-configurations-loader -config config.yaml -dry-run")
	fmt.Println()
	fmt.Println("  # Verbose output")
	fmt.Println("  gcloud-configurations-loader -config config.yaml -verbose")
}
