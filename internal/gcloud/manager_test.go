package gcloud

import (
	"reflect"
	"testing"
)

func TestBuildCommands(t *testing.T) {
	manager := NewManager()

	tests := []struct {
		name     string
		function func() []string
		expected []string
	}{
		{
			name:     "create configuration",
			function: func() []string { return manager.buildCreateCommand("test").Args },
			expected: []string{"gcloud", "config", "configurations", "create", "test"},
		},
		{
			name:     "set property",
			function: func() []string { return manager.buildSetCommand("cfg", "project", "my-proj").Args },
			expected: []string{"gcloud", "config", "set", "project", "my-proj", "--configuration", "cfg"},
		},
		{
			name:     "describe configuration",
			function: func() []string { return manager.buildDescribeCommand("test").Args },
			expected: []string{"gcloud", "config", "configurations", "describe", "test"},
		},
		{
			name:     "delete configuration",
			function: func() []string { return manager.buildDeleteCommand("test").Args },
			expected: []string{"gcloud", "config", "configurations", "delete", "test", "--quiet"},
		},
		{
			name:     "set property with special characters",
			function: func() []string { return manager.buildSetCommand("test-cfg", "compute/zone", "us-central1-a").Args },
			expected: []string{"gcloud", "config", "set", "compute/zone", "us-central1-a", "--configuration", "test-cfg"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.function()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
