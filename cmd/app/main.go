package main

import (
	"devllart/foobarman/internal/alert"
	"devllart/foobarman/internal/commands"
	"devllart/foobarman/internal/inputer"
	"devllart/foobarman/internal/sale"
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/internal/state"
	"devllart/foobarman/src/funcs"
)

/**
 * This game is simulator barman.
 * By playing it, you yourself may become a barman.
 * So be careful... Good luck ;)
 */

func main() {
	state.Load()
	funcs.CliClear() // Clear console
	scenes.Hello()   // Run scene "Hello", In the scene ask name's barman

	// Run game cycle
	for state.Run == true {
		sale.Sale()                              // If coctail is ready then sale to client the coctail
		scenes.Show(state.Scene)                 // In scenes.Hello global scene's context was changed to "Store"
		alert.ClearInfo()                        // Clear hints and warning message
		command := inputer.GetCommand("%Y \\> ") // Get command from input of user
		state.SetCommand(command)                // Set command
		commands.Exec()                          // Try execute user command's
		state.Save()                             // Save state game
	}
}
