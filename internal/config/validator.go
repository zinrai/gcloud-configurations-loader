package config

import "errors"

// Performs basic structural validation
func ValidateConfiguration(config Configuration) error {
	if config.Name == "" {
		return errors.New("configuration name is required")
	}
	if len(config.Properties) == 0 {
		return errors.New("at least one property must be specified")
	}
	return nil
}

// Validates all configurations in the file
func ValidateConfigFile(configFile *ConfigFile) error {
	if len(configFile.Configurations) == 0 {
		return errors.New("at least one configuration must be specified")
	}

	for _, config := range configFile.Configurations {
		if err := ValidateConfiguration(config); err != nil {
			return err
		}
	}
	return nil
}
