package main

import (
	"fmt"
	"os"

	"github.com/nabbas-ca/finite-automaton/fsm"
)

// modThree is the function that implements the FSM/FA api. Need to initialize FA and then create an FSM from it, which will process the input via GetFSMOutput func
func modThree(input string) int {
	threeModFA, err := fsm.NewFiniteAutomaton(
		fsm.NewSet("S0", "S1", "S2"),
		fsm.NewSet("0", "1"), "S0", fsm.NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating three mod FSM. Error:= %v\n", err)
		os.Exit(1) // shouldn't happen
	}

	fsm := threeModFA.NewFiniteStateMachine()
	output, err := fsm.GetFSMOutput(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting FSM output. Error: %v\n", err)
		os.Exit(1) // shouldn't happen
	}
	return output

}
func main() {
	// TODO: use Cobra (possibly viper) to parse command line and create config file for the following

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: fsm <input>\n")
	}
	input := os.Args[1] // input string is first arg

	output := modThree(input)
	fmt.Printf("%d\n", output)

}
