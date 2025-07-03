# gcloud-configurations-loader

A CLI tool to manage gcloud configurations from YAML files.

This tool eliminates the repetitive manual setup of gcloud configurations by allowing you to define them declaratively in YAML and apply them consistently across environments.

## Why This Tool?

Managing multiple gcloud configurations manually is tedious and error-prone:

```bash
# The manual way (repetitive and error-prone)
$ gcloud config configurations create dev-environment
$ gcloud config set account dev@example.com --configuration dev-environment
$ gcloud config set project my-dev-project --configuration dev-environment
$ gcloud config set compute/region us-central1 --configuration dev-environment
# ... repeat for each environment
```

With `gcloud-configurations-loader`, you define everything once in YAML and apply consistently:

```bash
$ gcloud-configurations-loader -config environments.yaml
```

## Features

- **YAML-driven configuration**: Define all your gcloud configurations in a single YAML file
- **Safe by default**: Only creates new configurations, skips existing ones unless explicitly told to replace
- **Flexible property support**: Supports any gcloud configuration property (account, project, compute/region, etc.)
- **Dry-run capability**: Preview changes before applying them

## Installation

```bash
$ go install github.com/zinrai/gcloud-configurations-loader@latest
```

## Usage

Apply configurations (create new, skip existing)

```bash
$ gcloud-configurations-loader -config config.yaml
```

Replace existing configurations

```bash
$ gcloud-configurations-loader -config config.yaml -replace
```

Dry run to see what would be done

```bash
$ gcloud-configurations-loader -config config.yaml -dry-run
```

Verbose output

```bash
$ gcloud-configurations-loader -config config.yaml -verbose
```

## Configuration File Format

See the `examples/` directory for sample configuration files:

- `basic-config.yaml`: Simple setup with account and project
- `advanced-config.yaml`: Complex setup with multiple environments and various properties

The `properties` section supports any property that `gcloud config set` accepts. For a full list of available properties, run:

```bash
$ gcloud config set --help
```

## How It Works

1. **Load**: Reads and parses the YAML configuration file
2. **Analyze**: Checks which configurations already exist using `gcloud config configurations describe`
3. **Plan**: Creates an execution plan (create new, skip existing, or replace if `--replace` specified)
4. **Execute**: Runs the appropriate `gcloud` commands:
   - `gcloud config configurations create <name>`
   - `gcloud config set <property> <value> --configuration <name>`

## Safety Features

- **Non-destructive by default**: Won't modify existing configurations unless `--replace` is specified
- **Validation**: Checks for required fields and basic structure before execution
- **Dry-run mode**: Preview changes with `--dry-run`
- **Clear reporting**: Shows exactly what was created, replaced, or skipped

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
