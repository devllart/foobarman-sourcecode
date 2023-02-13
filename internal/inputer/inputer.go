package inputer

import (
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/fmtc"

	"github.com/manifoldco/promptui"
)

var fisrt = true

func GetCommand(label string) (command string) {
	if fisrt {
		fisrt = false
		return ""
	}

	fmtc.Printf("\n")
	prompt := promptui.Prompt{
		Label:    fmtc.Sprintf(label),
		Validate: CheckInput,
	}

	command, err := prompt.Run()

	if err != nil {
		panic(err.Error())
	}

	return command
}

func CheckInput(input string) error {
	state.SetCommand(input)
	return commands.FakeExec()
}
