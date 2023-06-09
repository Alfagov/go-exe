package stateoftictactoe

import "errors"

func checkHorizontal(board []string) (bool, error) {

	oX := 0
	oO := 0
	tmp := ""

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp += string(board[i][j])
		}

		if tmp == lineX {
			oX++
		}

		if tmp == lineO {
			oO++
		}

		tmp = ""
	}

	if oX+oO > 1 {
		return false, errors.New("2 win lines, bad")
	}

	if oX+oO == 1 {
		return true, nil
	}

	return false, nil
}

func checkVertical(board []string) (bool, error) {

	oX := 0
	oO := 0
	tmp := ""

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp += string(board[j][i])
		}

		if tmp == lineX {
			oX++
		}

		if tmp == lineO {
			oO++
		}

		tmp = ""
	}

	if oX+oO > 1 {
		return false, errors.New("2 win lines, bad")
	}

	if oX+oO == 1 {
		return true, nil
	}

	return false, nil
}

func checkDiagonal(board []string) (bool, error) {
	oX := 0
	oO := 0
	tmp1 := string(board[0][0]) + string(board[1][1]) + string(board[2][2])
	tmp2 := string(board[0][2]) + string(board[1][1]) + string(board[2][0])
	if tmp1 == lineX {
		oX++
	}
	if tmp2 == lineX {
		oX++
	}
	if tmp1 == lineO {
		oO++
	}
	if tmp2 == lineO {
		oO++
	}
	if oX+oO > 1 {
		return false, errors.New("2 win lines, bad")
	}
	if oX+oO == 1 {
		return true, nil
	}
	return false, nil
}

func onGoing(board []string) (finish bool, err error) {

	cX := 0
	cO := 0

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				cX++
			}
			if board[i][j] == 'O' {
				cO++
			}
		}
	}

	if cX-cO <= 1 && 0 <= cX-cO {
		if cO+cX == 9 {
			return true, nil
		}
		return false, nil
	}

	return false, errors.New("bad game")
}
