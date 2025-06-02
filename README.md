

# Rangefetch

**Status:** ⚠️ This project is actively under development and not yet stable. But it’s improving every day!

## What is Rangefetch?

Rangefetch is a lightweight, elegant, and cross-platform system information tool designed to quickly show you essential details about your computer. Whether you’re on Linux, macOS, or Windows, Rangefetch lets you easily check your system’s hardware and software specs.

Curious about your OS version, CPU details, memory usage, or kernel version? Rangefetch is here to give you a clean, simple overview.

## Why Rangefetch?

There are many system info tools out there, but Rangefetch stands out with its minimal and user-friendly design. No clutter, no confusing details—just the important info you want, fast.

Plus, it’s open source, so you can customize it however you like and even contribute to make it better.

## Features

* Operating system and version details.
* CPU model, core count, and frequency.
* Total and available RAM.
* Kernel version.
* Screen resolution (on supported platforms).
* Disk usage .
* Minimal resource usage and fast startup.
* Cross-platform support (Linux, macOS, Windows).
* Simple configuration via a JSON file.

## Getting Started

### Requirements

* Go programming language version 1.20 or newer installed on your machine.
* Git (optional, for cloning the repository).

### Installation

There is no official stable release yet. For now, you can build the tool from source using Go. Official packages and easy installers will be available in future releases.

### Building from Source Example

```bash
git clone https://github.com/your-username/rangefetch.git
cd rangefetch
go build -o rangefetch
./rangefetch
```

### Usage

Currently, Rangefetch displays basic system info with no command-line options or GUI. More advanced features are planned for upcoming versions.

## Configuration

Rangefetch uses a `config.json` file to store user preferences. The exact location and structure of this file will be finalized in later versions. For now, it runs with default settings.

## How Can You Contribute?

Rangefetch grows stronger with your help! If you find this project useful, feel free to:

* Report issues or bugs,
* Suggest new features,
* Contribute code and improvements.

We appreciate every kind of contribution and your ideas. Let’s build this together!

## License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0). See the [LICENSE](./LICENSE) file for details.

