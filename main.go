package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	files    []string
	cursor   int
	selected map[int]bool
}

func initialModel() model {
	files := getUnstagedFiles()
	return model{
		files:    files,
		selected: make(map[int]bool),
	}
}

func getUnstagedFiles() []string {
	cmd := exec.Command("git", "diff", "--name-only")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to get unstaged files: %v", err)
		return []string{}
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var files []string
	for _, line := range lines {
		if line != "" {
			files = append(files, line)
		}
	}
	return files
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.files)-1 {
				m.cursor++
			}
		case " ":
			if _, ok := m.selected[m.cursor]; ok {
				m.selected[m.cursor] = false
			} else {
				m.selected[m.cursor] = true
			}
		case "enter":
			var selectedFiles []string
			for i := range m.selected {
				if m.selected[i] {
					selectedFiles = append(selectedFiles, m.files[i])
				}
			}
			if len(selectedFiles) > 0 {
				cmd := "git add " + strings.Join(selectedFiles, " ")
				err := clipboard.WriteAll(cmd)
				if err != nil {
					log.Fatalf("Failed to copy to clipboard: %v", err)
				}
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if len(m.files) == 0 {
		return "No unstaged files.\n"
	}
	s := "Select files to add (Space to select, Enter to copy command, q to quit):\n\n"
	for i, file := range m.files {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.selected[i] {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, file)
	}

	s += "\nSelected: "
	if len(m.selected) > 0 {
		s += fmt.Sprintf("%d file(s)", len(m.selected))
	} else {
		s += "None.."
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error while running program: %v", err)
		os.Exit(1)
	}

}
