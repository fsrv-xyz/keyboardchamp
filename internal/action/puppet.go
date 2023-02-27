package action

import (
	"log"
	"os/exec"

	"github.com/fsrv-xyz/keyboardchamp/internal/helper"
	"github.com/fsrv-xyz/keyboardchamp/internal/keycode"
)

func init() {
	RegistryInstance.Register("puppet_disable", func() Action {
		return &PuppetDisableAction{GenericAction{Requirements: []int{
			keycode.KeyCodePuppet,
			keycode.KeyCode1,
		}}}
	})
	RegistryInstance.Register("puppet_enable", func() Action {
		return &PuppetEnableAction{GenericAction{Requirements: []int{
			keycode.KeyCodePuppet,
			keycode.KeyCode2,
		}}}
	})
	RegistryInstance.Register("puppet_agent_test", func() Action {
		return &PuppetAgentTestAction{GenericAction{Requirements: []int{
			keycode.KeyCodePuppet,
			keycode.KeyCode3,
		}}}
	})
	RegistryInstance.Register("puppet_agent_test_http-debug", func() Action {
		return &PuppetAgentTestHTTPDebugAction{GenericAction{Requirements: []int{
			keycode.KeyCodePuppet,
			keycode.KeyCode4,
		}}}
	})
	RegistryInstance.Register("puppet_agent_test_bob_dev", func() Action {
		return &PuppetAgentTestBobDevAction{GenericAction{Requirements: []int{
			keycode.KeyCodePuppet,
			keycode.KeyCode5,
		}}}
	})
	RegistryInstance.Register("puppet_agent_test_bob_dev_noop", func() Action {
		return &PuppetAgentTestBobDevNoopAction{GenericAction{Requirements: []int{
			keycode.KeyCodePuppet,
			keycode.KeyCode6,
		}}}
	})
}

type PuppetEnableAction struct{ GenericAction }

func (p *PuppetEnableAction) Execute() {
	err := helper.RunCommandDisplayZero(
		exec.Command("/usr/bin/ydotool", "type", "puppet agent --enable\n"),
	)
	if err != nil {
		log.Println(err)
	}
}

type PuppetDisableAction struct{ GenericAction }

func (p *PuppetDisableAction) Execute() {
	err := helper.RunCommandDisplayZero(
		exec.Command("/usr/bin/ydotool", "type", "puppet agent --disable\n"),
	)
	if err != nil {
		log.Println(err)
	}
}

type PuppetAgentTestAction struct{ GenericAction }

func (p *PuppetAgentTestAction) Execute() {
	err := helper.RunCommandDisplayZero(
		exec.Command("/usr/bin/ydotool", "type", "puppet agent -t\n"),
	)
	if err != nil {
		log.Println(err)
	}
}

type PuppetAgentTestHTTPDebugAction struct{ GenericAction }

func (p *PuppetAgentTestHTTPDebugAction) Execute() {
	err := helper.RunCommandDisplayZero(
		exec.Command("/usr/bin/ydotool", "type", "puppet agent -t --http_debug\n"),
	)
	if err != nil {
		log.Println(err)
	}
}

type PuppetAgentTestBobDevAction struct{ GenericAction }

func (p *PuppetAgentTestBobDevAction) Execute() {
	err := helper.RunCommandDisplayZero(
		exec.Command("/usr/bin/ydotool", "type", "puppet agent -t --environment bob_dev\n"),
	)
	if err != nil {
		log.Println(err)
	}
}

type PuppetAgentTestBobDevNoopAction struct{ GenericAction }

func (p *PuppetAgentTestBobDevNoopAction) Execute() {
	err := helper.RunCommandDisplayZero(
		exec.Command("/usr/bin/ydotool", "type", "puppet agent -t --environment bob_dev --noop\n"),
	)
	if err != nil {
		log.Println(err)
	}
}
