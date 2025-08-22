package fsm

import (
	"fmt"
	"strconv"
	"strings"
)

// FiniteStateMachine represents a Finite State machine.
//
//	It is possible to parametrize FiniteAutomaton with T(for states) and U(for input elements), but I think this could be an overkill
type FiniteStateMachine struct {
	FA              *FiniteAutomaton // FA represents the FiniteAutomaton that configures this FSM. It is equivalent to FSM config object
	currentState    string
	OutputConverter func(string) (int, error) // OutputConverter, converts a state to an output. This is a flexible way to convert states to outputs
}

// NewFiniteAutomaton creates a new FSM with the tuple (Q,Σ,q0,F,δ). Does initial error checking as well.
func NewFiniteStateMachine(inputFA FiniteAutomaton) (*FiniteStateMachine, error) {
	f := FiniteStateMachine{FA: &inputFA}

	// TODO: add optional output converter func as an argument to this method
	f.OutputConverter = DefaultOutputCoverter

	// return
	return &f, nil
}

// DefaultOutputCoverter is the default output converter from state to final output
func DefaultOutputCoverter(state string) (int, error) {
	return strconv.Atoi(strings.TrimPrefix(state, "S"))

}

// String returns a string representing the FiniteStateMachine as a string
func (f *FiniteStateMachine) String() string {
	return fmt.Sprintf("FSM: \n\tFA=%s\n", f.FA.String())
}

// GetNextState gets the next state based on next input rune and currentState. This func corresponds to a func equivalent of delta instead of a double map.
//
//	Potentially we can make this extensible by allowing user to provide this func
func (f *FiniteStateMachine) ProcessInputRune(inputRune string) error {
	// check if the rune is acceptable
	if !f.FA.Sigma.Contains(string(inputRune)) {
		return fmt.Errorf("rune %v is not an acceptable input", inputRune)
	}
	f.currentState = f.FA.Delta[f.currentState][inputRune]
	return nil
}

// GetFSMOutput gets the output of the FSM based on the given inputs. returns an error if it encounters an error in processing
func (f *FiniteStateMachine) GetFSMOutput(input string) (int, error) {

	for _, r := range input { // go over the runes of the input string
		err := f.ProcessInputRune(string(r)) // Let FSM process rune and go to next state
		if err != nil {
			return 0, err // if we encounter an error, pass it back and return 0
		}
	}

	// finalState is here when done with processing input
	finalState := f.currentState

	// check if final state is one of the accepted states
	if !f.FA.F.Contains(finalState) {
		return 0, fmt.Errorf("state %s is not one of the accepted final states. F=%v", finalState, f.FA.F)
	}
	// convert the finalState to an output string using the output converter func
	return f.OutputConverter(finalState)
}
