# finite-automaton
A repository to create FAs(Finite Automaton) and FSMs(Finite State Machines)

## To compile

make build

## To run unit tests

make test

fsm.coverage file will be at top folder of the repository

## To use API within golang

- First, create an FA using fsm.NewFiniteAutomaton(Q,Sigma,q0,F,Delta as a map[string(state)]map[string(input)][string(state)]) func like this:
```
threeModFA, err := fsm.NewFiniteAutomaton(
		fsm.NewSet("S0", "S1", "S2"),
		fsm.NewSet("0", "1"), "S0", fsm.NewSet("S0", "S1", "S2"),
		map[string]map[string]string{
			"S0": {"0": "S0", "1": "S1"},
			"S1": {"0": "S2", "1": "S0"},
			"S2": {"0": "S1", "1": "S2"},
		})
```
- Then, initialize an FSM from the FA created above like this:
```
fsm := threeModFA.NewFiniteStateMachine()
```
- At last, process the input through the FSM like this:
```
output, err := fsm.GetFSMOutput(input)
```
- You could process the FSM at each input character using fsm.ProcessInputRune(inputRune string) if desired. Usually not needed.
- Function main.modThree in main.go provides a good example on how to use the API