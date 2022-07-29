package action

import (
	"log"
	"os/exec"

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
	err := exec.Command("xdotool", "key", "--", "Shift+F6").Run()
	if err != nil {
		log.Println(err)
	}
}
