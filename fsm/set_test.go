package fsm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {

	tests := []struct {
		name     string
		elements []string
		want     Set[string]
	}{
		{
			name:     "green test",
			elements: []string{"S0", "S1", "S2"},
			want:     Set[string]{"S0": struct{}{}, "S1": struct{}{}, "S2": struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(tt.elements...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {

	tests := []struct {
		name           string
		inputString    string
		expectedOutput Set[string]
		expectedError  error
		convertFunc    func(string) (string, error)
	}{
		{
			name:           "green test",
			inputString:    "(S0,S1,S2)",
			expectedOutput: NewSet("S0", "S1", "S2"),
			expectedError:  nil,
			convertFunc:    func(s string) (string, error) { return s, nil },
		},
		{
			name:           "no paranthese test",
			inputString:    "S0,S1,S2",
			expectedOutput: NewSet[string](),
			expectedError:  fmt.Errorf("input string doesn't have surrounding parentheses"),
			convertFunc:    func(s string) (string, error) { return s, nil },
		},
		{
			name:           "error convertFunc",
			inputString:    "(S0,S1,S2)",
			expectedOutput: NewSet[string](),
			expectedError:  fmt.Errorf("convertFunc error"),
			convertFunc:    func(s string) (string, error) { return "", fmt.Errorf("convertFunc error") },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputSet := NewSet[string]() // create an empty string set
			err := outputSet.Parse(tt.inputString, tt.convertFunc)
			if err != nil {
				if tt.expectedError == nil {
					t.Errorf("Parse() got error %v, expected no error", err)
				} else {
					// we are good so far, let's compare errors
					if err.Error() != tt.expectedError.Error() {
						t.Errorf("Parse() got error %v, expected error %v", err, tt.expectedError)
					}
				}
			} else {
				// no error
				if tt.expectedError == nil {
					// we are good so far, let's compare output
					if !outputSet.IsSubset(tt.expectedOutput) || !tt.expectedOutput.IsSubset(outputSet) { //to be equivalent, both sets need to be subsets of each other
						t.Errorf("Parse() got %v, expected %v", outputSet, tt.expectedOutput)
					}
				} else {
					// we are not good, let's output expected error
					t.Errorf("Parse() got no error, expected error %v", tt.expectedError)
				}
			}
		})
	}
}

func TestSet_AddTest(t *testing.T) {

	tests := []struct {
		name        string
		initialSet  Set[string]
		newVal      string
		expectedSet Set[string]
	}{
		{
			name:        "add new value",
			initialSet:  NewSet([]string{"S0", "S1", "S2"}...),
			newVal:      "S3",
			expectedSet: NewSet([]string{"S0", "S1", "S2", "S3"}...),
		},
		{
			name:        "add existing value",
			initialSet:  NewSet([]string{"S0", "S1", "S2"}...),
			newVal:      "S2",
			expectedSet: NewSet([]string{"S0", "S1", "S2"}...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initialSet.Add(tt.newVal) // add new val to initialSet

			if !tt.initialSet.IsSubset(tt.expectedSet) || !tt.expectedSet.IsSubset(tt.initialSet) { // if initialset and expectedSet are not equivalent
				t.Errorf("Set.Add(%v) = %v, expected %v", tt.newVal, tt.initialSet, tt.expectedSet)
			}
		})
	}
}

func TestSet_RemoveTest(t *testing.T) {

	tests := []struct {
		name        string
		initialSet  Set[string]
		removeVal   string
		expectedSet Set[string]
	}{
		{
			name:        "delete existing value",
			initialSet:  NewSet([]string{"S0", "S1", "S2"}...),
			removeVal:   "S2",
			expectedSet: NewSet([]string{"S0", "S1"}...),
		},
		{
			name:        "delete non-existing value, should do nothing ", //TODO: maybe make it throw an error
			initialSet:  NewSet([]string{"S0", "S1", "S2"}...),
			removeVal:   "S3",
			expectedSet: NewSet([]string{"S0", "S1", "S2"}...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initialSet.Remove(tt.removeVal) // add new val to initialSet

			if !tt.initialSet.IsSubset(tt.expectedSet) || !tt.expectedSet.IsSubset(tt.initialSet) { // if initialset and expectedSet are not equivalent
				t.Errorf("Set.Remove(%v) = %v, expected %v", tt.removeVal, tt.initialSet, tt.expectedSet)
			}
		})
	}
}
