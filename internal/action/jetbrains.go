package action

import (
	"keyboardchamp/internal/keycode"
	"log"
	"os/exec"
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
