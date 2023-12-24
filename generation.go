package main

import (
	"fmt"
	"strconv"
	"unicode"
	"time"
)

var pieceValues map[string]int
var knightHeuristicMap [64]float32
var generalHeuristicMap [64]float32

func init() {
    pieceValues = map[string]int{
        "p":   -1,
        "P":   1,
        "n":   -3,
        "N":   3,
        "b":   -4,
        "B":   4,
        "r":   -5,
        "R":   5,
        "q":   -10,
        "Q":   10,
        "k":   -999,
        "K":   999,
    }
	knightHeuristicMap = [64]float32{
		0.33, 0.33, 0.5, 0.5, 0.5, 0.5, 0.33, 0.33,
		0.33, 0.5, 0.75, 0.75, 0.75, 0.75, 0.5, 0.33,
		0.5, 1.0, 1.0, 1.0,1.0, 1.0, 1.0, 0.5, 
		0.5, 1.0, 1.0, 1.0,1.0, 1.0, 1.0, 0.5, 
		0.5, 1.0, 1.0, 1.0,1.0, 1.0, 1.0, 0.5, 
		0.5, 1.0, 1.0, 1.0,1.0, 1.0, 1.0, 0.5, 
		0.33, 0.5, 0.75, 0.75, 0.75, 0.75, 0.5, 0.33,
		0.33, 0.33, 0.5, 0.5, 0.5, 0.5, 0.33, 0.33,
	}
	generalHeuristicMap = [64]float32{
		0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90,
		0.90, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95,
		0.95, 1.0, 1.0, 1.0, 1.0 , 1.0, 1.0, 0.95,
		1.0, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.0,
		1.0, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.0,
		0.95, 1.0, 1.0, 1.0, 1.0 , 1.0, 1.0, 0.95,
		0.90, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95,
		0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90,
	}
}
func numberOfPieces(board [64]int) int{
	var nmbr int = 0
	for _, piece := range board{
		if piece != 0 {
			nmbr++
		} 
	}

	return nmbr
}
func isBlackInCheck(board [64]int) bool{
	////fmt.Println("----------IsBlackInCheck()----------")
	var blackKingPosition int 
	for index, piece := range board {
		if piece == -999 {
			blackKingPosition = index
			break
		}
	}
	var new_index int = blackKingPosition
	//diagonala stanga sus
	new_index = blackKingPosition - 9
	////fmt.Println("diagonala stanga sus")
	for new_index % 8 != 7 && new_index >= 0 {
		if board[new_index] != 0 {
			if board[new_index] == 4 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break	
		}
		new_index -= 9
	}


	new_index = blackKingPosition - 7
	//diagonala dreapta sus
	////fmt.Println("diagonala dreapta sus")
	for new_index % 8 != 0 && new_index >= 0 {
		if board[new_index] != 0 {
			if board[new_index] == 4 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index -= 7
	}

	new_index = blackKingPosition + 7 
	//diagonala stanga jos
	////fmt.Println("diagonala stanga jos")
	for new_index % 8 != 7 && new_index <= 63 {
		if board[new_index] != 0 {
			if board[new_index] == 4 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index += 7
	}
	
	new_index = blackKingPosition + 9
	//diagonala dreapta jos
	////fmt.Println("diagonala dreapta jos")
	for new_index % 8 != 0 && new_index <= 63 {
		if board[new_index] != 0 {
			if board[new_index] == 4 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index += 9
	}
	//sus
	new_index = blackKingPosition - 8
	for new_index >= 0 {
		if board[new_index] != 0 {
			if board[new_index] == 5 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index -= 8
	}
	//jos
	////fmt.Println("jos")
	new_index = blackKingPosition + 8
	for new_index <= 63 {
		if board[new_index] != 0 {
			if board[new_index] == 5 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index += 8
	}

	//stanga
	////fmt.Println("stanga")
	new_index = blackKingPosition - 1
	for new_index % 8 != 7 && new_index >= 0{
		if board[new_index] != 0 {
			if board[new_index] == 5 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index -= 1
	}
	//dreapta
	new_index = blackKingPosition + 1
	for new_index % 8 != 0 {
		if board[new_index] != 0 {
			if board[new_index] == 5 || board[new_index] == 10 {
				////fmt.Println(new_index)
				return true
			}
			break
		}
		new_index += 1
	}

	if blackKingPosition % 8 >= 2 {
		if blackKingPosition / 8 >= 1 && board[blackKingPosition-10] == 3 {
			return true
		}
		
		if blackKingPosition / 8 <= 6 && board[blackKingPosition+6] == 3 {
			return true
		}
	}
	if blackKingPosition % 8 >= 1 {
		if blackKingPosition / 8  >= 2 && board[blackKingPosition-17] == 3 {
			return true
		}
		if blackKingPosition / 8 <= 5 && board[blackKingPosition+15] == 3 {
			return true
		}
	}
	if blackKingPosition % 8 <= 5 {
		if blackKingPosition / 8 >= 1 && board[blackKingPosition-6] == 3 {
			return true
		}

		if blackKingPosition / 8 <= 6 && board[blackKingPosition+10] == 3 {
			return true
		}
	}
	
	if blackKingPosition % 8 <= 6 {
		if blackKingPosition / 8 >= 2 && board[blackKingPosition - 15] == 3 {
			return true
		}

		if blackKingPosition / 8 <= 5 && board[blackKingPosition + 17] == 3 {
			return true
		}
	}
	if blackKingPosition + 7 < 63 && (blackKingPosition+7) % 8 != 7 && board[blackKingPosition + 7] == 1{
		return true
	}

	if blackKingPosition + 9 <= 63 && (blackKingPosition+9) % 8 > 0 && board[blackKingPosition + 9] == 1{
		return true
	}
	////fmt.Println("------------------------------")
	return false
}

func isWhiteInCheck(board [64]int) bool{
	var whiteKingPosition int 
	for index, piece := range board {
		if piece == 999 {
			whiteKingPosition = index
			break
		}
	}
	var new_index int = whiteKingPosition
	//diagonala stanga sus
	new_index = whiteKingPosition - 9
	////fmt.Println("diagonala stanga sus")
	for new_index % 8 != 7 && new_index >= 0 {
		if board[new_index] != 0 {
			if board[new_index] == -4 || board[new_index] == -10 {
				return true
			}
			break	
		}
		new_index -= 9
	}


	new_index = whiteKingPosition - 7
	//diagonala dreapta sus
	////fmt.Println("diagonala dreapta sus")
	for new_index % 8 != 0 && new_index >= 0 {
		if board[new_index] != 0 {
			if board[new_index] == -4 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index -= 7
	}

	new_index = whiteKingPosition + 7 
	//diagonala stanga jos
	////fmt.Println("diagonala stanga jos")
	for new_index % 8 != 7 && new_index <= 63 {
		if board[new_index] != 0 {
			if board[new_index] == -4 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index += 7
	}
	
	new_index = whiteKingPosition + 9
	//diagonala dreapta jos
	////fmt.Println("diagonala dreapta jos")
	for new_index % 8 != 0 && new_index <= 63 {
		if board[new_index] != 0 {
			if board[new_index] == -4 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index += 9
	}
	//sus
	new_index = whiteKingPosition - 8
	for new_index >= 0 {
		if board[new_index] != 0 {
			if board[new_index] == -5 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index -= 8
	}
	//jos
	////fmt.Println("jos")
	new_index = whiteKingPosition + 8
	for new_index <= 63 {
		if board[new_index] != 0 {
			if board[new_index] == -5 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index += 8
	}

	//stanga
	////fmt.Println("stanga")
	new_index = whiteKingPosition - 1
	for new_index % 8 != 7 && new_index >= 0{
		if board[new_index] != 0 {
			if board[new_index] == -5 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index -= 1
	}
	//dreapta
	new_index = whiteKingPosition + 1
	for new_index % 8 != 0 {
		if board[new_index] != 0 {
			if board[new_index] == -5 || board[new_index] == -10 {
				return true
			}
			break
		}
		////fmt.Println(new_index)
		new_index += 1
	}
	
	if whiteKingPosition % 8 >= 2 {
		if whiteKingPosition / 8 >= 1 && board[whiteKingPosition-10] == -3 {
			return true
		}
		
		if whiteKingPosition / 8 <= 6 && board[whiteKingPosition+6] == -3 {
			return true
		}
	}
	if whiteKingPosition % 8 >= 1 {
		if whiteKingPosition / 8  >= 2 && board[whiteKingPosition-17] == -3 {
			return true
		}
		if whiteKingPosition / 8 <= 5 && board[whiteKingPosition+15] == -3 {
			return true
		}
	}
	if whiteKingPosition % 8 <= 5 {
		if whiteKingPosition / 8 >= 1 && board[whiteKingPosition-6] == -3 {
			return true
		}

		if whiteKingPosition / 8 <= 6 && board[whiteKingPosition+10] == -3 {
			return true
		}
	}
	
	if whiteKingPosition % 8 <= 6 {
		if whiteKingPosition / 8 >= 2 && board[whiteKingPosition - 15] == -3 {
			return true
		}

		if whiteKingPosition / 8 <= 5 && board[whiteKingPosition + 17] == -3 {
			return true
		}
	}
	if whiteKingPosition - 9 >= 0 && (whiteKingPosition - 9) % 8 != 7 && board[whiteKingPosition - 9] == -1 {
		return true
	}
	if whiteKingPosition - 7 >= 0 && (whiteKingPosition - 7) % 8 > 0 && board[whiteKingPosition - 7] == -1 {
		return true
	}

	return false
}



func fenToBoard(fen_input string) [64]int {
	var output_arr [64]int

	var index int = 0
	var boardIndex int = 0

	for index < len(fen_input) {
		char := rune(fen_input[index])
		if char == ' ' {
			break
		}
		if char == '/' {
			index++
			continue
		}

		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			for i := 0; i < num; i++ {
				output_arr[boardIndex] = 0
				boardIndex++
			}
		} else {
			output_arr[boardIndex] = pieceValues[string(char)]
			boardIndex++
		}
		index++
	}

	return output_arr
}

func makeMove(board [64]int, a, b int) [64]int {
    copy := board
    copy[b] = copy[a]
    copy[a] = 0
    return copy
}

func abs(x int) int{
	if x < 0{
		return x*-1
	}
	return x
}
func boardIndexOf(board [64]int, piece int) int {
	for i, val := range board {
		if val == piece {
			return i
		}
	}
	return -1
}
func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func generateBlackMoves(board [64]int) [][64]int{ //-1 for black && 1 for white
	var boards [][64]int
	////fmt.Println("-------------generateBlackMoves------------")
	for index, piece := range board{
		switch piece{
		case -1: 
			if index + 8 < 64 && board[index+8] == 0{ //una in fata
				new_move := makeMove(board, index, index+8)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}		
			}
			if  index/8 == 1 && board[index+16] == 0 { //doua in fata
				new_move := makeMove(board, index, index+16)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}	
			}
			if index+7 < 64 && board[index+7] > 0 { //atac stanga
				new_move := makeMove(board, index, index+7)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}	
			}
			if index+9 < 64 && board[index+9] > 0 { //atac dreapta
				new_move := makeMove(board, index, index+9)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}	
			}
		case -3:
			if index % 8 >= 2 {
				if index / 8 >= 1 && board[index-10] >= 0 {
					new_move := makeMove(board, index, index-10)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}	
				}
				
				if index / 8 <= 6 && board[index+6] >= 0 {
					new_move := makeMove(board, index, index+6)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}	
				}
			}
			if index % 8 >= 1 {
				if index / 8  >= 2 && board[index-17] >= 0 {
					new_move := makeMove(board, index, index-17)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
				if index / 8 <= 5 && board[index+15] >= 0 {
					new_move := makeMove(board, index, index+15)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
			}
			if index % 8 <= 5 {
				if index / 8 >= 1 && board[index-6] >= 0 {
					new_move := makeMove(board, index, index-6)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}
				}

				if index / 8 <= 6 && board[index+10] >= 0 {
					new_move := makeMove(board, index, index+10)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
			}
			
			if index % 8 <= 6 {
				if index / 8 >= 2 && board[index - 15] >= 0 {
					new_move := makeMove(board, index, index - 15)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}
				}

				if index / 8 <= 5 && board[index + 17] >= 0 {
					new_move := makeMove(board, index, index + 17)
					if !isBlackInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
			}
		case -4:
			var new_index int = index

			//diagonala stanga sus
			new_index = index - 9
			//////fmt.Println("diagonala stanga sus")
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break	
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 9
			}


			new_index = index - 7
			//diagonala dreapta sus
			////fmt.Println("diagonala dreapta sus")
			for new_index % 8 != 0 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 7
			}

			new_index = index + 7 
			//diagonala stanga jos
			////fmt.Println("diagonala stanga jos")
			for new_index % 8 != 7 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index += 7
			}
			
			new_index = index + 9
			//diagonala dreapta jos
			////fmt.Println("diagonala dreapta jos")
			for new_index % 8 != 0 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				////fmt.Println(new_index)
				new_index += 9
			}

		
		case -5:
			var new_index int = index
			//sus
			////fmt.Println("sus")
			new_index = index - 8
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 8
			}
			//jos
			////fmt.Println("jos")
			new_index = index + 8
			for new_index % 8 != 0 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index += 8
			}

			//stanga
			////fmt.Println("stanga")
			new_index = index - 1
			for new_index % 8 != 7 && new_index >= 0{
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 1
			}
			//dreapta
			////fmt.Println("dreapta")
			new_index = index + 1
			//dreapta
			new_index = index + 1
			for new_index % 8 != 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index += 1
			}
		case -10:
			var new_index int = index - 8
			//sus
			////fmt.Println("sus")
			new_index = index - 8
			for new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						////fmt.Println(new_index)
						new_move:= makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards,new_move)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index -= 8
			}
			//jos
			////fmt.Println("jos")
			new_index = index + 8
			for new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 8
			}

			//stanga
			////fmt.Println("stanga")
			new_index = index - 1
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					////fmt.Println(new_index)
					boards = append(boards, new_move)
				}
				new_index -= 1
			}
			//dreapta
			////fmt.Println("dreapta")
			new_index = index + 1
			//dreapta
			new_index = index + 1
			for new_index % 8 != 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							////fmt.Println(new_index)
							boards = append(boards, new_move)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 1
			}

			//diagonala stanga sus
			new_index = index - 9
			////fmt.Println("diagonala stanga sus")
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							////fmt.Println(new_index)
							boards = append(boards, new_move)
						}
					}
					break	
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index -= 9
			}


			new_index = index - 7
			//diagonala dreapta sus
			////fmt.Println("diagonala dreapta sus")
			for new_index % 8 != 0 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index -= 7
			}

			new_index = index + 7 
			//diagonala stanga jos
			////fmt.Println("diagonala stanga jos")
			for new_index % 8 != 7 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 7
			}
			
			new_index = index + 9
			//diagonala dreapta jos
			////fmt.Println("diagonala dreapta jos")
			for new_index % 8 != 0 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] > 0 {
						new_move := makeMove(board, index, new_index)
						if !isBlackInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 9
			}
		case -999:
			var kingMoves []int
			if index % 8 > 0 {
				kingMoves = append(kingMoves, index-1)
				if index/8 > 0 {
					kingMoves = append(kingMoves, index-9)
				}
				if index/8 < 7 {
					kingMoves = append(kingMoves, index + 7)
				}
			}
			if index/8 >= 1{
				kingMoves = append(kingMoves, index-8)
			}
			if index / 8 < 7 {
				kingMoves = append(kingMoves, index + 8)

				if index % 8 < 7 {
					kingMoves = append(kingMoves, index+9)
				}
			}
			
			if index%8 <7{
				kingMoves = append(kingMoves, index+1)
				if index / 8 > 0 {
					kingMoves = append(kingMoves, index - 7)
				}
			}
			for _, new_index := range kingMoves {
				if new_index >= 0 && new_index < 64 {
					whiteKingSquare := boardIndexOf(board, 999)
					if abs(new_index%8-whiteKingSquare%8) > 1 || abs(new_index/8-whiteKingSquare/8) > 1 {
						// Exclude the current square
						if new_index != index {
							if board[new_index] == 0 || board[new_index] > 0 {
								new_move := makeMove(board, index, new_index)
								if !isBlackInCheck(new_move) {
									//fmt.Println(new_index)
									boards = append(boards, new_move)
								}
							}
						}
					}
				}
			}
		}	
	}
	return boards
}

func generateWhiteMoves(board [64]int) [][64]int{ //-1 for black && 1 for white
	var boards [][64]int
	////fmt.Println("-------------generateWhiteMoves------------")
	for index, piece := range board{
		switch piece{
		case 1: 
			if index - 8 >= 0 && board[index-8] == 0{ //una in fata
				new_move := makeMove(board, index, index-8)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}		
			}
			if  index/8 == 6 && board[index-16] == 0 { //doua in fata
				new_move := makeMove(board, index, index-16)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}	
			}
			if (index-7)%8 != 0 && index-7 >= 0 && board[index-7] < 0 { //atac dreapta
				new_move := makeMove(board, index, index-7)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}	
			}
			if (index - 9)%8 != 7 && index-9 >= 0 && board[index-9] < 0 { //atac stanga
				new_move := makeMove(board, index, index-9)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}	
			}
		case 3:
			if index % 8 >= 2 {
				if index / 8 >= 1 && board[index-10] <= 0 {
					new_move := makeMove(board, index, index-10)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}	
				}
				
				if index / 8 <= 6 && board[index+6] <= 0 {
					new_move := makeMove(board, index, index+6)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}	
				}
			}
			if index % 8 >= 1 {
				if index / 8  >= 2 && board[index-17] <= 0 {
					new_move := makeMove(board, index, index-17)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
				if index / 8 <= 5 && board[index+15] <= 0 {
					new_move := makeMove(board, index, index+15)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
			}
			if index % 8 <= 5 {
				if index / 8 >= 1 && board[index-6] <= 0 {
					new_move := makeMove(board, index, index-6)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}
				}

				if index / 8 <= 6 && board[index+10] <= 0 {
					new_move := makeMove(board, index, index+10)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
			}
			
			if index % 8 <= 6 {
				if index / 8 >= 2 && board[index - 15] <= 0 {
					new_move := makeMove(board, index, index - 15)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}
				}

				if index / 8 <= 5 && board[index + 17] <= 0 {
					new_move := makeMove(board, index, index + 17)
					if !isWhiteInCheck(new_move){
						boards = append(boards, new_move)
					}
				}
			}
		case 4:
			var new_index int = index

			//diagonala stanga sus
			new_index = index - 9
			//////fmt.Println("diagonala stanga sus")
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break	
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 9
			}


			new_index = index - 7
			//diagonala dreapta sus
			////fmt.Println("diagonala dreapta sus")
			for new_index % 8 != 0 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 7
			}

			new_index = index + 7 
			//diagonala stanga jos
			////fmt.Println("diagonala stanga jos")
			for new_index % 8 != 7 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index += 7
			}
			
			new_index = index + 9
			//diagonala dreapta jos
			////fmt.Println("diagonala dreapta jos")
			for new_index % 8 != 0 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				////fmt.Println(new_index)
				new_index += 9
			}

		
		case 5:
			var new_index int = index
			//sus
			////fmt.Println("sus")
			new_index = index - 8
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 8
			}
			//jos
			////fmt.Println("jos")
			new_index = index + 8
			for new_index % 8 != 0 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index += 8
			}

			//stanga
			////fmt.Println("stanga")
			new_index = index - 1
			for new_index % 8 != 7 && new_index >= 0{
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isBlackInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index -= 1
			}
			//dreapta
			////fmt.Println("dreapta")
			new_index = index + 1
			//dreapta
			new_index = index + 1
			for new_index % 8 != 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
				}
				new_index += 1
			}
		case 10:
			var new_index int = index - 8
			//sus
			////fmt.Println("sus")
			new_index = index - 8
			for new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						////fmt.Println(new_index)
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)	
						}

					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index -= 8
			}
			//jos
			////fmt.Println("jos")
			new_index = index + 8
			for new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				////fmt.Println(new_index)
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 8
			}

			//stanga
			////fmt.Println("stanga")
			new_index = index - 1
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					////fmt.Println(new_index)
					boards = append(boards, new_move)
				}
				new_index -= 1
			}
			//dreapta
			////fmt.Println("dreapta")
			//dreapta
			new_index = index + 1
			for new_index % 8 != 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							////fmt.Println(new_index)
							boards = append(boards, new_move)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 1
			}

			//diagonala stanga sus
			new_index = index - 9
			////fmt.Println("diagonala stanga sus")
			for new_index % 8 != 7 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							////fmt.Println(new_index)
							boards = append(boards, new_move)
						}
					}
					break	
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index -= 9
			}


			new_index = index - 7
			//diagonala dreapta sus
			////fmt.Println("diagonala dreapta sus")
			for new_index % 8 != 0 && new_index >= 0 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index -= 7
			}

			new_index = index + 7 
			//diagonala stanga jos
			////fmt.Println("diagonala stanga jos")
			for new_index % 8 != 7 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 7
			}
			
			new_index = index + 9
			//diagonala dreapta jos
			////fmt.Println("diagonala dreapta jos")
			for new_index % 8 != 0 && new_index <= 63 {
				if board[new_index] != 0 {
					if board[new_index] < 0 {
						new_move := makeMove(board, index, new_index)
						if !isWhiteInCheck(new_move){
							boards = append(boards, new_move)
							////fmt.Println(new_index)
						}
					}
					break
				}
				new_move := makeMove(board, index, new_index)
				if !isWhiteInCheck(new_move){
					boards = append(boards, new_move)
					////fmt.Println(new_index)
				}
				new_index += 9
			}
		case 999:
			var kingMoves []int
			if index % 8 > 0 {
				kingMoves = append(kingMoves, index-1)
				if index/8 > 0 {
					kingMoves = append(kingMoves, index-9)
				}
				if index/8 < 7 {
					kingMoves = append(kingMoves, index + 7)
				}
			}
			if index/8 >= 1{
				kingMoves = append(kingMoves, index-8)
			}
			if index / 8 < 7 {
				kingMoves = append(kingMoves, index + 8)

				if index % 8 < 7 {
					kingMoves = append(kingMoves, index+9)
				}
			}
			
			if index%8 <7{
				kingMoves = append(kingMoves, index+1)
				if index / 8 > 0 {
					kingMoves = append(kingMoves, index - 7)
				}
			}
			for _, new_index := range kingMoves {
				if new_index >= 0 && new_index < 64 {
					whiteKingSquare := boardIndexOf(board, -999)
					if abs(new_index%8-whiteKingSquare%8) > 1 || abs(new_index/8-whiteKingSquare/8) > 1 {
						// Exclude the current square
						if new_index != index {
							if board[new_index] <= 0 {
								new_move := makeMove(board, index, new_index)
								if !isWhiteInCheck(new_move) {
									boards = append(boards, new_move)
								}
							}
						}
					}
				}
			}
		}	
	}
	return boards
}

type KeyType [64]int
var myMap = make(map[KeyType]float32)
func addValue(key KeyType, value float32) {
    myMap[key] = value
}

func findValue(key KeyType) (float32, bool) {
    value, found := myMap[key]
    return value, found
}

func eval(board [64]int) float32{
	value, found := findValue(board)
	if found{
		return value
	}
	if len(generateBlackMoves(board)) == 0 && isBlackInCheck(board){
		addValue(board, 9999.99)
		return 9999.99
	}
	
	var whiteBishops, blackBishops int 
	var score float32
	for index, piece := range board{
		switch abs(piece) {
		case 1:
			score += generalHeuristicMap[index] * float32(piece)
		case 3:
			score+= knightHeuristicMap[index] * 3.25 * generalHeuristicMap[index] * float32(piece/3)
		case 4:
			score += 3.25 * float32(piece/4) *  generalHeuristicMap[index]
			if piece < 0 {
				blackBishops++
			} else {
				whiteBishops++
			}
		case 5:
			score += 5.0 * float32(piece/5) * generalHeuristicMap[index]
		case 10:
			score += 9.0 * float32(piece/10) * generalHeuristicMap[index]
		}
	}

	if blackBishops >= 2 {
		score -= 0.5
	}	
	if whiteBishops >= 2 {
		score += 0.5
	}

	var stackedWhitePawns, stackedBlackPanws int = -1,-1
	for col_index := 0; col_index <= 7; col_index++{
		var new_col_index int
		for new_col_index < 64 {
			if board[new_col_index] == -1{
				stackedBlackPanws++;
			} else if board[new_col_index] == 1{
				stackedWhitePawns++;
			}
			new_col_index += 8
		}
	}
	score -= float32(stackedWhitePawns /2)
	score += float32(stackedBlackPanws /2)
	addValue(board, score)
	return score
}


	
func Minimax(board [64]int, depth int, alpha, beta float32, color int) float32 {
    if depth == 0 {
        return eval(board)
    }

    var moves []([64]int)
    if color == 1 {
        moves = generateWhiteMoves(board)
    } else {
        moves = generateBlackMoves(board)
    }

    if len(moves) == 0 {
        // No legal moves, evaluate current board
        return eval(board)
    }

    if color == 1 {
        maxScore := float32(-9999.99)
        for _, move := range moves {
            score := Minimax(move, depth-1, alpha, beta, -color)
            if score > maxScore {
                maxScore = score
            }
            alpha = max(alpha, score)
            if beta <= alpha {
                break // Beta cut-off
            }
        }
        return maxScore
    } else {
        minScore := float32(9999.0)
        for _, move := range moves {
            score := Minimax(move, depth-1, alpha, beta, -color)
            if score < minScore {
                minScore = score
            }
            beta = min(beta, score)
            if beta <= alpha {
                break // Alpha cut-off
            }
        }
        return minScore
    }
}

var pieceToFEN = map[int]string{
	-1:   "p",
	1:    "P",
	-3:   "n",
	3:    "N",
	-4:   "b",
	4:    "B",
	-5:   "r",
	5:    "R",
	-10:  "q",
	10:   "Q",
	-999: "k",
	999:  "K",
}

func boardToFEN(board [64]int) string {
	fen := ""
	emptyCount := 0

	for i, piece := range board {
		if (i)%8 == 0 && i != 0 {
			if emptyCount > 0 {
				fen += fmt.Sprintf("%d", emptyCount)
				emptyCount = 0
			}

			fen += "/"
		}

		if piece == 0 {
			emptyCount++
		} else {
			if emptyCount > 0 {
				fen += fmt.Sprintf("%d", emptyCount)
				emptyCount = 0
			}

			fen += pieceToFEN[piece]
		}
	}

	if emptyCount > 0 {
		fen += fmt.Sprintf("%d", emptyCount)
	}
	fen += " w - - 0 1"
	return fen 
}

func PrettyPrintBoard(board [64]int) {
	fmt.Println("  a b c d e f g h")
	fmt.Println(" +----------------")
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d|", rank+1)
		for file := 0; file < 8; file++ {
			index := rank*8 + file
			piece := pieceToFEN[board[index]]
			fmt.Printf(" %s", piece)
		}
		fmt.Println()
	}
	fmt.Println(" +----------------")
}
//-5-3-4-10-999-4-3-5-1-1-1-1-1-1-1-1000000000000000000000000000000001111111153410999435 
func main(){
	//"8/8/1n6/8/8/8/8/8"

	//var fennot string = "r1bqkbnr/pp1pp1pp/2n5/5p2/2p1P3/8/PPPP1PPP/R1BQKBNR"
    //board := fenToBoard(fennot)

    //x := generateBlackMoves(board)
	/*
	for _, y := range x{
		k := generateBlackMoves(y)
		for _, u := range k {
			generateBlackMoves(u)
		}
	}
	*/
	////fmt.Println(len(x))
	////fmt.Println(isBlackInCheck(board))

	for true {
		var fennot string
		fmt.Scan(&fennot)
		board := fenToBoard(fennot)
	
		//8/p7/5k2/p4p1p/P1N2b2/1PP5/4R1PP/3r3K
		//fmt.Println(boardToFEN(board))
		depth := 6
		var nb [64]int
		maxScore := -9999.99
		startTime := time.Now()
		for _, move := range generateWhiteMoves(board) {
			score := Minimax(move, depth-1, -9999.9, 9999.9, -1)
			if score > float32(maxScore){
				maxScore = float64(score)
				nb = move
			}
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
	
		fmt.Println(maxScore)
		fmt.Println(boardToFEN(nb))

		fmt.Printf("Time taken: %s\n", elapsed)
	
		
	}
}
