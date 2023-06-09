package stateoftictactoe

type State string

const (
	Win     State  = "win"
	Ongoing State  = "ongoing"
	Draw    State  = "draw"
	lineX   string = "XXX"
	lineO   string = "OOO"
)

func StateOfTicTacToe(board []string) (State, error) {

	var state State

	finish, err := onGoing(board)
	if err != nil {
		return state, err
	}

	checkH, err := checkHorizontal(board)
	if err != nil {
		return state, err
	}
	if checkH {
		return Win, nil
	}

	checkV, err := checkVertical(board)
	if err != nil {
		return state, err
	}
	if checkV {
		return Win, nil
	}

	checkD, err := checkDiagonal(board)
	if err != nil {
		return Win, nil
	}
	if checkD {
		return Win, nil
	}
	if finish {
		return Draw, nil
	}

	return Ongoing, nil
}
