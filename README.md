# Vanilla OS - Apx Package Manager

Apx is the default package manager for Vanilla OS. It is a wrapper around multiple package managers to install packages and run commands inside a managed container.

## Overview

**Apx** (/à·peks/) provides a unified interface for managing packages across different Linux distributions. It uses [distrobox](https://github.com/89luca89/distrobox) to create containerized environments where you can install packages without affecting your host system.

## Features

- **Multi-distribution support**: Install packages from different Linux distributions
- **Container-based**: Packages are installed in isolated containers
- **Package manager abstraction**: Unified interface for different package managers (apt, dnf, pacman, etc.)
- **Desktop integration**: Automatic desktop entry export for installed applications
- **Subsystem management**: Create and manage multiple subsystems with different stacks

## Stack

- **Language**: Go 1.25.0
- **Framework**: Vanilla OS SDK + CLI builder
- **UI**: Charmbracelet (bubbletea, lipgloss) for terminal UI
- **Container**: distrobox integration

## Project Structure

```
cmd/
  main.go              # Application entry point
  main_check.go        # String validation variant
  locales/             # Multi-language support (33+ languages)

internal/cli/
  structs.go           # CLI command definitions
  runtime.go           # Subsystem execution handlers
  subsystems.go        # Subsystem lifecycle management
  stacks.go            # Stack operations
  pkgmanagers.go       # Package manager operations

core/                  # Core business logic (subsystems, stacks, pkg management)
config/                # Configuration files
distrobox/            # distrobox integration scripts
```

## Building

### Prerequisites

- Go 1.25.0 or higher
- make
- podman or docker
- git

### Build Steps

```bash
# Clone and enter directory
git clone --recursive https://github.com/yourusername/Vanilla-.git
cd Vanilla-

# Build
make build

# Install system-wide
sudo make install
sudo make install-manpages

# Install to custom location
make install PREFIX=$HOME/.local
make install-manpages PREFIX=$HOME/.local
```

## Usage

### Basic Commands

```bash
# List available subsystems
apx subsystems list

# Create a new subsystem
apx subsystems new --name=ubuntu-latest

# Enter a subsystem
apx ubuntu-latest enter

# Install packages in subsystem
apx ubuntu-latest install package-name

# Run a command in subsystem
apx ubuntu-latest run command

# List available stacks
apx stacks list

# Manage package managers
apx pkgmanagers list
```

## Dependencies

### Direct Dependencies
- `github.com/google/uuid` - UUID generation
- `github.com/vanilla-os/sdk` - Application framework
- `gopkg.in/yaml.v2` - YAML configuration

### UI Dependencies
- `github.com/charmbracelet/bubbletea` - Terminal UI framework
- `github.com/charmbracelet/lipgloss` - Terminal styling

## Documentation

For detailed documentation, visit: https://docs.vanillaos.org/docs/en/apx

## Translations

Contribute translations via [Weblate](https://hosted.weblate.org/projects/vanilla-os/apx)

## License

GNU General Public License v3.0

## Authors

- Mirko Brombin <brombin94@gmail.com>
- Pietro di Caprio <pietro@fabricators.ltd>
- Vanilla OS Contributors

## Contributing

Contributions are welcome! Please ensure:

1. Dependencies are updated with `go get`, `go mod tidy`, and `go mod vendor`
2. Code follows Go conventions
3. Translations are updated in Weblate
