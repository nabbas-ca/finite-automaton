package fsm

import (
	"strconv"
	"testing"
)

// TestFiniteStateMachine_RangeFSMTestTestFSMRange tests ThreeMod FA and FourMod FA from 0-1000 generating a test for each iteration
func TestFiniteStateMachine_RangeFSMTestTestFSMRange(t *testing.T) {

	threeModFA, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2"),
		NewSet("0", "1"), "S0", NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})
	fourModFA, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2", "S3"),
		NewSet("0", "1"), "S0", NewSet("S0", "S1", "S2", "S3"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S3"},
			"S2": {"0": "S0", "1": "S1"},
			"S3": {"0": "S2", "1": "S3"},
		})

	tests := []struct {
		name   string
		fa     *FiniteAutomaton
		maxInt int
		mod    int // this the mod this FSM is solving
	}{
		{
			name:   "Mod-Three FA range test",
			fa:     threeModFA,
			maxInt: 1000,
			mod:    3,
		},
		{
			name:   "Mod-Four FA range test",
			fa:     fourModFA,
			maxInt: 1000,
			mod:    4,
		},
	}

	for _, tt := range tests {
		// generate maxInt testcases per test
		for i := 0; i <= tt.maxInt; i++ {

			t.Run(tt.name, func(t *testing.T) {
				// generate expected output using regular mod operator
				expectedOutput := i % tt.mod
				// generate binary string from i
				binaryStr := strconv.FormatInt(int64(i), 2)

				fsm := tt.fa.NewFiniteStateMachine()    // initialize a new FSM
				got, err := fsm.GetFSMOutput(binaryStr) // process binary input through FSM
				if err != nil {
					t.Errorf("FiniteStateMachine.GetFSMOutput() error = %v", err)
					return
				}
				if got != expectedOutput {
					t.Errorf("FiniteStateMachine.GetFSMOutput() = %v, want %v", got, expectedOutput)
				}
			})
		}

	}
}
func TestFiniteStateMachine_GetFSMOutput(t *testing.T) {

	threeModFA, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2"),
		NewSet("0", "1"), "S0", NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})

	threeModFA_MissingFinalState, _ := NewFiniteAutomaton(
		NewSet("S0", "S1", "S2"),
		NewSet("0", "1"), "S0", NewSet("S0", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})

	tests := []struct {
		name    string
		fa      *FiniteAutomaton
		input   string
		want    int
		wantErr bool
	}{
		{
			name:    "Mod-Three FA example 1",
			fa:      threeModFA,
			input:   "110",
			want:    0,
			wantErr: false,
		},
		{
			name:    "Mod-Three FA example 2",
			fa:      threeModFA,
			input:   "1101",
			want:    1,
			wantErr: false,
		},
		{
			name:    "Mod-Three FA bad input",
			fa:      threeModFA,
			input:   "1102",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Mod-Three FA bad final state",
			fa:      threeModFA_MissingFinalState,
			input:   "1101",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fsm := tt.fa.NewFiniteStateMachine()
			got, err := fsm.GetFSMOutput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FSM.GetFSMOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FSM.GetFSMOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultOutputCoverter(t *testing.T) {
	tests := []struct {
		name    string
		state   string
		want    int
		wantErr bool
	}{
		{
			name:    "S0",
			state:   "S0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "S1",
			state:   "S1",
			want:    1,
			wantErr: false,
		},
		{
			name:    "F0",
			state:   "F0",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefaultOutputCoverter(tt.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultOutputCoverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultOutputCoverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
