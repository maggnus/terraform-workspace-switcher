package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/nexidian/gocliselect"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	// Initialize Terraform in the current directory
	tf, err := tfexec.NewTerraform(".", "terraform")
	if err != nil {
		log.Fatalf("failed to create Terraform instance: %v", err)
	}

	currentWorkspace, err := tf.WorkspaceShow(ctx)
	if err != nil {
		log.Fatalf("failed to get current workspace: %v", err)
	}

	// List available workspaces
	workspaces, _, err := tf.WorkspaceList(ctx)

	// Create the selection menu
	menu := gocliselect.NewMenu("Select Terraform Workspace")
	for n, workspace := range workspaces {
		menu.AddItem(workspace, workspace)
		if workspace == currentWorkspace {
			menu.CursorPos = n
		}
	}

	selected := menu.Display()

	if selected == "" {
		os.Exit(0)
	}

	// Select the chosen workspace
	if err := tf.WorkspaceSelect(context.Background(), selected); err != nil {
		log.Fatalf("failed to select workspace: %v", err)
	}

	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("Successfully switched to workspace: %s\n", red(selected))
}
