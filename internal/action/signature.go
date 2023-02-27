package action

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fsrv-xyz/keyboardchamp/internal/helper"
	"github.com/fsrv-xyz/keyboardchamp/internal/keycode"
)

func init() {
	RegistryInstance.Register("signature_command", func() Action {
		return &SignatureCommandAction{GenericAction{Requirements: []int{
			keycode.KeyCodeSignature,
		}}}
	})
}

type SignatureCommandAction struct{ GenericAction }

func (j *SignatureCommandAction) Execute() {
	user := os.Getenv("USER")
	text := fmt.Sprintf("%v %v", time.Now().Format(time.RFC822), user)
	err := helper.RunCommandDisplayZero(exec.Command("/usr/bin/ydotool", "type", text))
	if err != nil {
		log.Println(err)
	}
}
