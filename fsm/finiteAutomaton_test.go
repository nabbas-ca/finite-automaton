package fsm

import (
	"fmt"
	"testing"
)

func TestNewFiniteAutomaton(t *testing.T) {

	threeModFA, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2"),
		NewSet("0", "1"), "S0", NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})

	tests := []struct {
		name        string
		Q           Set[string]
		Sigma       Set[string]
		q0          string
		F           Set[string]
		Delta       map[string]map[string]string
		want        *FiniteAutomaton
		expectedErr error
	}{
		{
			name:  "Q empty",
			Q:     NewSet[string](),
			Sigma: NewSet("0", "1"),
			q0:    "S0",
			F:     NewSet("S0", "S1", "S2"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("Q(list of acceptable states) is empty"),
		},
		{
			name:  "Sigma empty",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet[string](),
			q0:    "S0",
			F:     NewSet("S0", "S1", "S2"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("Σ(Sigma)=(list of acceptable input) is empty"),
		},
		{
			name:  "q0 invalid",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet("0", "1"),
			q0:    "S3",
			F:     NewSet("S0", "S1", "S2"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("q0(initial state) is not one of the acceptable states"),
		},
		{
			name:  "F is empty",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet("0", "1"),
			q0:    "S0",
			F:     NewSet[string](),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("F(list of acceptable final states) is empty"),
		},
		{
			name:  "F is not a subset of Q",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet("0", "1"),
			q0:    "S0",
			F:     NewSet("S0", "S1", "S3"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("F(list of acceptable final states) is not a subset of Q"),
		},
		{
			name:  "Delta doesn't contain all the states",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet("0", "1"),
			q0:    "S0",
			F:     NewSet("S0", "S1", "S2"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("delta doesn't contain all the possible state possibilities"),
		},
		{
			name:  "Delta doesn't contain all the inputs",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet("0", "1"),
			q0:    "S0",
			F:     NewSet("S0", "S1", "S2"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        nil,
			expectedErr: fmt.Errorf("delta doesn't contain all the possible input possibilities"),
		},
		{
			name:  "threeModFA green test creation",
			Q:     NewSet("S0", "S1", "S2"),
			Sigma: NewSet("0", "1"),
			q0:    "S0",
			F:     NewSet("S0", "S1", "S2"),
			Delta: map[string]map[string]string{
				"S0": {"0": "S0", "1": "S1"},
				"S1": {"0": "S2", "1": "S0"},
				"S2": {"0": "S1", "1": "S2"},
			},
			want:        threeModFA,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFiniteAutomaton(tt.Q, tt.Sigma, tt.q0, tt.F, tt.Delta)
			if (err == nil) != (tt.expectedErr == nil) {
				t.Errorf("NewFSM() error mismatch. err = %v, expectedErr %v", err, tt.expectedErr)
			}
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("NewFSM() error mismatch: err:%v, expected error:%v", err, tt.expectedErr)
			}

			if got != nil && !got.Equals(tt.want) {
				t.Errorf("NewFSM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiniteAutomaton_String(t *testing.T) {

	threeModFA, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2"),
		NewSet("0", "1"), "S0", NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})
	tests := []struct {
		name string
		fa   *FiniteAutomaton
		want string
	}{
		{
			name: "threeModFA example",
			fa:   threeModFA,
			want: `FA:
	Q=(S0, S1, S2)
	Σ=(0, 1)
	q0=S0
	F=(S0, S1, S2)
	δ=map[S0:map[0:S0 1:S1] S1:map[0:S2 1:S0] S2:map[0:S1 1:S2]]
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.fa.String(); got != tt.want {
				t.Errorf("FSM.String() = \n%v\n want \n%v", got, tt.want)
			}
		})
	}
}

func TestFiniteAutomaton_NewFiniteStateMachine(t *testing.T) {
	threeModFA, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2"),
		NewSet("0", "1"), "S0", NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})

	tests := []struct {
		name        string
		fa          *FiniteStateMachine
		expectedFSM *FiniteStateMachine
	}{
		{
			name:        "init FSM green test",
			fa:          threeModFA.NewFiniteStateMachine(),
			expectedFSM: &FiniteStateMachine{FA: threeModFA, currentState: threeModFA.q0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output := tt.fa.FA.NewFiniteStateMachine()

			if output == nil && tt.expectedFSM == nil {
				// good, both match
			} else if output == nil || tt.expectedFSM == nil {
				t.Errorf("FiniteAutomaton.NewFiniteStateMachine() mismatch. output= %v, expected %v", output, tt.expectedFSM)
			} else if output.currentState != tt.expectedFSM.currentState { // compare current states which should be initial state
				t.Errorf("FiniteAutomaton.NewFiniteStateMachine().currentState=%v, expected %v", output.currentState, tt.expectedFSM.currentState)
			}
		})
	}
}
