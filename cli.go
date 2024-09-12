package main

import (
	"context"
	"log"
	"os"

	tree "github.com/charmbracelet/lipgloss/tree"
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic"
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic/client/native"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type model struct {
	ctx           *ctx
	diff          string
	commitMessage string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	style := getRootStyle(width, height, m.ctx.theme)

	t := tree.New().
		EnumeratorStyle(enumeratorStyle).
		Root("# Table of Contents").
		Child(
			tree.New().
				Root("## diff").
				Child(m.diff),
		).
		Child(
			tree.New().
				Root("## Chapter 2").
				Child("Chapter 2.1").
				Child("Chapter 2.2"),
		)

	return style.Render(t.String())
}

type getCommitMessageMsg string

func getCommitMessage(diff string) getCommitMessageMsg {
	ctx := context.Background()
	client, err := native.MakeClient(native.Config{
		APIKey: os.Getenv("ANTHROPIC_API_KEY"),
	})
	if err != nil {
		panic(err)
	}

	request := anthropic.NewMessageRequest(
		[]anthropic.MessagePartRequest{
			{Role: "system", Content: []anthropic.ContentBlock{anthropic.NewTextContentBlock("You are a helpful assistant that generates commit messages for a git diff.")}},
			{Role: "user", Content: []anthropic.ContentBlock{anthropic.NewTextContentBlock(diff)}},
		},
		anthropic.WithModel[anthropic.MessageRequest](anthropic.Claude35Sonnet),
		anthropic.WithMaxTokens[anthropic.MessageRequest](20),
	)

	response, err := client.Message(ctx, request)
	if err != nil {
		panic(err)
	}

	return getCommitMessageMsg(response.Content[0].Text)
}
