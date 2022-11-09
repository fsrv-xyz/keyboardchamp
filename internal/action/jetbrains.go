package action

import (
	"log"
	"os/exec"

	"github.com/fsrv-xyz/keyboardchamp/internal/helper"
	"github.com/fsrv-xyz/keyboardchamp/internal/keycode"
)

func init() {
	RegistryInstance.Register("jetbrains_rename", func() Action {
		return &JetBrainsRenameAction{GenericAction{Requirements: []int{
			keycode.KeyCodeJetBrains,
			keycode.KeyCodeEscape,
		}}}
	})
}

type JetBrainsRenameAction struct{ GenericAction }

func (j *JetBrainsRenameAction) Execute() {
	err := helper.RunCommandDisplayZero(exec.Command("/usr/bin/xdotool", "key", "--", "Shift+F6"))
	if err != nil {
		log.Println(err)
	}
}
