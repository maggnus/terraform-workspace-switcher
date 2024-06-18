# Terraform Workspace Switcher (tfs)

`tfs` is a simple CLI tool that allows users to interactively switch between Terraform workspaces.

## Features

- Lists available Terraform workspaces
- Provides an interactive menu to select a workspace
- Switches to the selected workspace

## Installation

To install `tfs`, make sure you have Go installed and then run:

```sh
go install github.com/maggnus/terraform-workspace-switcher/cmd/tfs@latest
```
## Usage
Navigate to your Terraform project directory and run:
```
tfs
```

## Example
This will display an interactive menu where you can use the arrow keys to navigate and press Enter to select a Terraform workspace. The selected workspace will be set as the current workspace.
```shell
$ tfs
Select Terraform Workspace
> default
  dev
  uat
  prod
Successfully switched to workspace: dev
```

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue to discuss any changes or improvements.

## License
This project is licensed under the MIT License.

## Acknowledgments
- [terraform-exec](https://github.com/hashicorp/terraform-exec) by HashiCorp
- [gocliselect](https://github.com/Nexidian/gocliselect) by Nexidian
