package action

import (
	"log"
	"os/exec"

	"github.com/fsrv-xyz/keyboardchamp/internal/helper"
	"github.com/fsrv-xyz/keyboardchamp/internal/keycode"
)

func init() {
	RegistryInstance.Register("update_general_command", func() Action {
		return &JetBrainsRenameAction{GenericAction{Requirements: []int{
			keycode.KeyCodeUpdate,
			keycode.KeyCode1,
		}}}
	})
}

type UpdateGeneralCommandAction struct{ GenericAction }

func (j *UpdateGeneralCommandAction) Execute() {
	err := helper.RunCommandDisplayZero(exec.Command("/usr/bin/xdotool", "type", "apt update -y && apt upgrade -y && apt autoremove -y; pending_packages; puppet agent -t; sophia\n"))
	if err != nil {
		log.Println(err)
	}
}
