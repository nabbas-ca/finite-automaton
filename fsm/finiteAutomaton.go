package fsm

import (
	"fmt"
	"reflect"
)

// FiniteAutomaton represents a Finite Automaton.
//
//	It is possible to parametrize FiniteAutomaton with T(for states) and U(for input elements), but I think this could be an overkill
type FiniteAutomaton struct {
	Q     Set[string]                  // Q is the set of acceptable FSM states.
	Sigma Set[string]                  // Sigma is the acceptable set of inputs for the FSM.
	q0    string                       // Q0 is initial state. TODO: maybe make it handle runes instead of just strings
	F     Set[string]                  // F is the Set of final states. Using map[string]struct{} as a way to store a set
	Delta map[string]map[string]string // Delta is a map between initial state(T) , acceptable input (U) and resulting in next state (T)
	//   It is possible to make this a func(string,string) string, but a map seemed sufficient. We can go either way
}

// NewFiniteAutomaton creates a new FSM with the tuple (Q,Σ,q0,F,δ). Does initial error checking as well.
func NewFiniteAutomaton(Q Set[string], Sigma Set[string], q0 string, F Set[string], Delta map[string]map[string]string) (*FiniteAutomaton, error) {
	fa := FiniteAutomaton{}

	// initial error checking for Q
	if len(Q) == 0 {
		return nil, fmt.Errorf("Q(list of acceptable states) is empty")
		// TODO: maybe handle all errors together in a method by itself to help users instead of one at a time. Possible future enhancement
	}
	fa.Q = Q.DeepCopy()

	// initial error checking for Sigma
	if len(Sigma) == 0 {
		return nil, fmt.Errorf("Σ(Sigma)=(list of acceptable input) is empty")

	}
	fa.Sigma = Sigma.DeepCopy()

	// check if q0 is in one of the elements in Q
	if !fa.Q.Contains(q0) {
		return nil, fmt.Errorf("q0(initial state) is not one of the acceptable states")
	}
	fa.q0 = q0

	// initial error checking for F
	if len(F) == 0 {
		return nil, fmt.Errorf("F(list of acceptable final states) is empty")
	}
	// check if F is subset of Q
	if !F.IsSubset(fa.Q) {
		return nil, fmt.Errorf("F(list of acceptable final states) is not a subset of Q")
	}
	fa.F = F.DeepCopy()

	// check if Delta covers all states in Q
	// check if Delta covers all inputs

	// deep copy Delta
	fa.Delta = make(map[string]map[string]string, len(Delta)) // initialize outer map
	outerKeySet := NewSet[string]()                           // outerKeySet will hold all the outer map keys, to check that it has to match Q
	for k, v := range Delta {                                 //iterate over outer map
		fa.Delta[k] = make(map[string]string, len(v)) // initialize inner map
		outerKeySet.Add(k)                            //  populate innerkeyset for error checking later
		innerKeySet := NewSet[string]()               // innerKeySet will hold all the inner map keys, to check that it has to match Sigma
		for k1, v1 := range v {                       // iterate over inner map
			fa.Delta[k][k1] = v1
			innerKeySet.Add(k1) // populate innerkeyset for error checking later
		}
		// here we have an inner set of keys, which should be a equivalent to Sigma
		if !innerKeySet.IsSubset(fa.Sigma) || !fa.Sigma.IsSubset(innerKeySet) { // innerKeySet should be a subset of Sigma and vice versa to be equivalent
			return nil, fmt.Errorf("delta doesn't contain all the possible input possibilities")
		}

	}
	// here we have an outer set of keys, which should be a equivalent to Q
	if !outerKeySet.IsSubset(fa.Q) || !fa.Q.IsSubset(outerKeySet) { // innerKeySet should be a subset of Sigma and vice versa to be equivalent
		return nil, fmt.Errorf("delta doesn't contain all the possible state possibilities")
	}

	// ------Finished deepcopy of Delta --------

	// return
	return &fa, nil
}

// String returns a string representing the FiniteAutomaton as a string
func (f *FiniteAutomaton) String() string {
	return fmt.Sprintf("FA:\n\tQ=%s\n\tΣ=%s\n\tq0=%v\n\tF=%s\n\tδ=%v\n", f.Q.String(), f.Sigma.String(), f.q0, f.F.String(), f.Delta)
}

// NewFiniteStateMachine returns a new FiniteStateMachine with initialized state
func (f *FiniteAutomaton) NewFiniteStateMachine() *FiniteStateMachine {
	return &FiniteStateMachine{FA: f, currentState: f.q0, OutputConverter: DefaultOutputCoverter}
}

// Equals returns whether 2 FSMs are equivalent
func (f *FiniteAutomaton) Equals(otherFSM *FiniteAutomaton) bool {

	if !reflect.DeepEqual(f.Q, otherFSM.Q) {
		return false
	}
	if !reflect.DeepEqual(f.Sigma, otherFSM.Sigma) {
		return false
	}
	if !reflect.DeepEqual(f.q0, otherFSM.q0) {
		return false
	}
	if !reflect.DeepEqual(f.F, otherFSM.F) {
		return false
	}
	if !reflect.DeepEqual(f.Delta, otherFSM.Delta) {
		return false
	}
	return true
}
