package gcloud

import "os/exec"

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

// Builds the command to create a configuration
func (m *Manager) buildCreateCommand(name string) *exec.Cmd {
	return exec.Command("gcloud", "config", "configurations", "create", name)
}

// Builds the command to describe a configuration
func (m *Manager) buildDescribeCommand(name string) *exec.Cmd {
	return exec.Command("gcloud", "config", "configurations", "describe", name)
}

// Builds the command to delete a configuration
func (m *Manager) buildDeleteCommand(name string) *exec.Cmd {
	return exec.Command("gcloud", "config", "configurations", "delete", name, "--quiet")
}

// Builds the command to set a configuration property
func (m *Manager) buildSetCommand(configName, key, value string) *exec.Cmd {
	return exec.Command("gcloud", "config", "set", key, value, "--configuration", configName)
}

// Checks if a configuration exists
func (m *Manager) ConfigurationExists(name string) bool {
	return m.buildDescribeCommand(name).Run() == nil
}

// Creates a new configuration
func (m *Manager) CreateConfiguration(name string) error {
	return m.buildCreateCommand(name).Run()
}

// Deletes a configuration
func (m *Manager) DeleteConfiguration(name string) error {
	return m.buildDeleteCommand(name).Run()
}

// Sets a property for a configuration
func (m *Manager) SetConfigProperty(configName, key, value string) error {
	return m.buildSetCommand(configName, key, value).Run()
}
