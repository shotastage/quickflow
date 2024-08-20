# Supported Operating Systems for QuickFlow

QuickFlow is designed to run efficiently on specific operating systems. This document outlines the supported environments and provides additional information for optimal performance.

## Officially Supported Operating Systems

### macOS
- Versions: 10.15 (Catalina) and later
- Architecture: Intel and Apple Silicon (M1/M2)

#### Notes for macOS:
- Ensure you have the latest version of Xcode Command Line Tools installed.
- We recommend using Homebrew to install dependencies like PostgreSQL.

### Linux Distributions

#### Ubuntu
- Versions: 20.04 LTS and later
- Architecture: x86_64 and ARM64

#### Fedora
- Versions: Fedora 33 and later
- Architecture: x86_64 and ARM64

#### Red Hat Enterprise Linux (RHEL)
- Versions: RHEL 8 and later
- Architecture: x86_64 and ARM64

#### Notes for Linux:
- Ensure your system is up to date with the latest packages.
- You may need to add the official PostgreSQL repository to get the latest version, especially for Ubuntu and Fedora.
- For RHEL, ensure you have the appropriate subscriptions and repositories enabled.

## Unsupported Operating Systems

### Windows and Windows Server
QuickFlow does not currently officially support Windows or Windows Server environments. However, there are potential workarounds:

#### Windows Subsystem for Linux (WSL)
It is theoretically possible to run QuickFlow on Windows using WSL with a supported Linux distribution. This approach allows you to run a Linux environment directly on Windows, potentially enabling QuickFlow operation.

**Important Note:**
- Running QuickFlow on WSL is not officially supported.
- This configuration has not been thoroughly tested by our team.
- Users choosing to run QuickFlow via WSL do so at their own risk and may encounter unexpected issues.
- We cannot guarantee full functionality or provide support for WSL-based setups.

If you need to run QuickFlow on a Windows machine, we recommend the following options in order of preference:
1. Use a virtual machine with a fully supported Linux distribution.
2. If using a virtual machine is not feasible, consider using WSL2 with a supported Linux distribution, understanding the limitations and lack of official support.

## General Requirements

Regardless of the operating system, ensure your environment meets these general requirements:

1. Go version 1.16 or later
2. PostgreSQL 12 or later
3. Sufficient disk space for your expected database size
4. Minimum 4GB RAM (8GB or more recommended for production use)

## Testing and Reporting Issues

If you encounter any issues with QuickFlow on the supported operating systems, please report them to our support team. Include the following information:

- Operating System name, version, and architecture
- Go version
- PostgreSQL version
- Error messages or logs

## Future Support

We are constantly evaluating our support for different operating systems. While we currently do not officially support Windows, this may change in future releases. We are also considering expanding our Linux distribution support. Stay tuned to our release notes and documentation for updates on supported environments.
