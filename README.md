# gitadd-tui

A Terminal User Interface (TUI) for interactively staging multiple files with `git add`.
![render1770717659111](https://github.com/user-attachments/assets/3c6862e8-0a82-48d4-b17e-3e466ef0856b)



## What is gitadd-tui?

`gitadd-tui` is a simple, interactive terminal tool that makes staging files for Git commits more intuitive and visual. Instead of typing `git add` commands repeatedly or staging all files at once, you can:

- Browse your modified/untracked files in a clean interface
- Select multiple files using keyboard navigation
- Stage files with a simple keypress
- See your git status visually in the terminal

Perfect for developers who want a faster workflow without leaving the command line or switching to a GUI.

## Installation

### Quick Install (curl)

```bash
curl -fsSL https://raw.githubusercontent.com/shrikant23codes/gitadd-tui/main/install.sh | bash
```

### Install from source (requires Go)

```bash
go install github.com/shrikant23codes/gitadd-tui@latest
```

### Build from source

```bash
git clone https://github.com/shrikant23codes/gitadd-tui.git
cd gitadd-tui
go build -o gitadd-tui main.go
sudo mv gitadd-tui /usr/local/bin/
```

## Usage

Navigate to any Git repository and run:

```bash
gitadd-tui
```

Use arrow keys to navigate, Space to select files, and Enter to stage the selected files.

## Requirements

- Go 1.16 or higher (for building from source)
- Git installed on your system
- A Git repository to work with

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is open source. Check the repository for license details.

## Author

Created by [shrikant23codes](https://github.com/shrikant23codes)
