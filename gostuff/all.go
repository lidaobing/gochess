package gostuff

// import "fmt"
//makes all possible black pawn moves and checks whether or not the target king is still in check
//x and y is source of piece/pawn
func (table *Table) blackPawn(x int8, y int8) bool {

	if x == 1 { //then pawn can move two squares
		if table.ChessBoard[x+2][y] == "-" { //make sure square is not blocked
			capturedPiece := table.makeMove(x, y, x+2, y, "bP")
			if table.isBlackInCheck() == false { //if black is no longer in check then its not mate
				table.undoMove(x, y, x+2, y, "bP", capturedPiece)
				return false
			}
			table.undoMove(x, y, x+2, y, "bP", capturedPiece)
		}
	}
	if x+1 <= 7 && table.ChessBoard[x+1][y] == "-" { //checking pawn movement of one square
		capturedPiece := table.makeMove(x, y, x+1, y, "bP")
		if table.isBlackInCheck() == false { //if white is no longer in check then its not mate
			table.undoMove(x, y, x+1, y, "bP", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y, "bP", capturedPiece) //regardless if in check, move is undone
	}

	if y+1 <= 7 && x+1 <= 7 && (table.ChessBoard[x+1][y+1])[0:1] == "w" { //right capture

		capturedPiece := table.makeMove(x, y, x+1, y+1, "bP")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y+1, "bP", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y+1, "bP", capturedPiece)
	}

	if y-1 >= 0 && x+1 <= 7 && (table.ChessBoard[x+1][y-1])[0:1] == "w" { //left capture

		capturedPiece := table.makeMove(x, y, x+1, y-1, "bP")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y-1, "bP", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y-1, "bP", capturedPiece)
	}
	if table.blackPass[y] { //checking if black can enpassent with this pawn
		if y+1 <= 7 && x+1 <= 7 && table.ChessBoard[x][y+1] == "wP" { //checking if white pawn exist on the right

			capturedPiece := table.makeMove(x, y, x+1, y+1, "bP")
			if table.isBlackInCheck() == false {
				table.undoMove(x, y, x+1, y+1, "bP", capturedPiece)
				return false
			}
			table.undoMove(x, y, x+1, y+1, "bP", capturedPiece)
		}
		if y-1 >= 0 && x+1 <= 7 && table.ChessBoard[x][y-1] == "wP" { //checking left

			capturedPiece := table.makeMove(x, y, x+1, y-1, "bP")
			if table.isBlackInCheck() == false {
				table.undoMove(x, y, x+1, y-1, "bP", capturedPiece)
				return false
			}
			table.undoMove(x, y, x+1, y-1, "bP", capturedPiece)
		}

	}
	//returns true if all possible move of THIS pawn lead to white king still being in check
	return true
}
func (table *Table) whitePawn(x int8, y int8) bool { //x is row y is col

	if x == 6 { //then pawn can move two squares
		if table.ChessBoard[x-2][y] == "-" { //make sure square is not blocked
			capturedPiece := table.makeMove(x, y, x-2, y, "wP")
			if table.isWhiteInCheck() == false { //if white is no longer in check then its not mate
				table.undoMove(x, y, x-2, y, "wP", capturedPiece)
				return false
			}
			table.undoMove(x, y, x-2, y, "wP", capturedPiece)
		}
	}
	if x-1 >= 0 && table.ChessBoard[x-1][y] == "-" { //checking pawn movement of one square
		capturedPiece := table.makeMove(x, y, x-1, y, "wP")
		if table.isWhiteInCheck() == false { //if white is no longer in check then its not mate
			table.undoMove(x, y, x-1, y, "wP", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y, "wP", capturedPiece) //regardless if in check, move is undone
	}

	if x-1 >= 0 && y-1 >= 0 && (table.ChessBoard[x-1][y-1])[0:1] == "b" { //left capture

		capturedPiece := table.makeMove(x, y, x-1, y-1, "wP")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y-1, "wP", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y-1, "wP", capturedPiece)
	}

	if x-1 >= 0 && y+1 <= 7 && (table.ChessBoard[x-1][y+1])[0:1] == "b" { //right capture

		capturedPiece := table.makeMove(x, y, x-1, y+1, "wP")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y+1, "wP", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y+1, "wP", capturedPiece)
	}

	if table.whitePass[y] { //checking if white can enpassent with this pawn

		if y+1 <= 7 && x-1 >= 0 && table.ChessBoard[x][y+1] == "bP" { //checking if black pawn exist on the right

			capturedPiece := table.makeMove(x, y, x-1, y+1, "wP")
			if table.isWhiteInCheck() == false {
				table.undoMove(x, y, x-1, y+1, "wP", capturedPiece)
				return false
			}
			table.undoMove(x, y, x-1, y+1, "wP", capturedPiece)
		}
		if y-1 >= 0 && x-1 >= 0 && table.ChessBoard[x][y-1] == "bP" { //checking left

			capturedPiece := table.makeMove(x, y, x-1, y-1, "wP")
			if table.isWhiteInCheck() == false {

				table.undoMove(x, y, x-1, y-1, "wP", capturedPiece)
				return false
			}
			table.undoMove(x, y, x-1, y-1, "wP", capturedPiece)
		}

	}

	//returns true if all possible move of THIS pawn lead to white king still being in check
	return true
}

//brute forces all possible knight moves for white
func (table *Table) whiteKnight(x int8, y int8) bool {

	//starting at top left going clockwise
	if x-2 >= 0 && y-1 >= 0 && (table.ChessBoard[x-2][y-1])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x-2, y-1, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-2, y-1, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-2, y-1, "wN", capturedPiece)
	} //white can't capture own piece
	if x-2 >= 0 && y+1 <= 7 && (table.ChessBoard[x-2][y+1])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x-2, y+1, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-2, y+1, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-2, y+1, "wN", capturedPiece)
	}
	if x-1 >= 0 && y+2 <= 7 && (table.ChessBoard[x-1][y+2])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x-1, y+2, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y+2, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y+2, "wN", capturedPiece)
	}
	if x+1 <= 7 && y+2 <= 7 && (table.ChessBoard[x+1][y+2])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x+1, y+2, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+1, y+2, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y+2, "wN", capturedPiece)
	}
	if x+2 <= 7 && y+1 <= 7 && (table.ChessBoard[x+2][y+1])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x+2, y+1, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+2, y+1, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+2, y+1, "wN", capturedPiece)
	}
	if x+2 <= 7 && y-1 >= 0 && (table.ChessBoard[x+2][y-1])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x+2, y-1, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+2, y-1, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+2, y-1, "wN", capturedPiece)
	}
	if x+1 <= 7 && y-2 >= 0 && (table.ChessBoard[x+1][y-2])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x+1, y-2, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+1, y-2, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y-2, "wN", capturedPiece)
	}
	if x-1 >= 0 && y-2 >= 0 && (table.ChessBoard[x-1][y-2])[0:1] != "w" {

		capturedPiece := table.makeMove(x, y, x-1, y-2, "wN")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y-2, "wN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y-2, "wN", capturedPiece)
	}
	return true
}

func (table *Table) blackKnight(x int8, y int8) bool {
	if x-2 >= 0 && y-1 >= 0 && (table.ChessBoard[x-2][y-1])[0:1] != "b" { //starting at top left going clockwise

		capturedPiece := table.makeMove(x, y, x-2, y-1, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-2, y-1, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-2, y-1, "bN", capturedPiece)
	}
	if x-2 >= 0 && y+1 <= 7 && (table.ChessBoard[x-2][y+1])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x-2, y+1, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-2, y+1, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-2, y+1, "bN", capturedPiece)
	}
	if x-1 >= 0 && y+2 <= 7 && (table.ChessBoard[x-1][y+2])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x-1, y+2, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-1, y+2, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y+2, "bN", capturedPiece)
	}
	if x+1 <= 7 && y+2 <= 7 && (table.ChessBoard[x+1][y+2])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x+1, y+2, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y+2, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y+2, "bN", capturedPiece)
	}
	if x+2 <= 7 && y+1 <= 7 && (table.ChessBoard[x+2][y+1])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x+2, y+1, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+2, y+1, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+2, y+1, "bN", capturedPiece)
	}
	if x+2 <= 7 && y-1 >= 0 && (table.ChessBoard[x+2][y-1])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x+2, y-1, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+2, y-1, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+2, y-1, "bN", capturedPiece)
	}
	if x+1 <= 7 && y-2 >= 0 && (table.ChessBoard[x+1][y-2])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x+1, y-2, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y-2, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y-2, "bN", capturedPiece)
	}
	if x-1 >= 0 && y-2 >= 0 && (table.ChessBoard[x-1][y-2])[0:1] != "b" {

		capturedPiece := table.makeMove(x, y, x-1, y-2, "bN")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-1, y-2, "bN", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y-2, "bN", capturedPiece)
	}
	return true
}

func (table *Table) whiteBishop(x int8, y int8) bool { //moves all possible moves with this one wB

	var i int8
	var j int8
	for i, j = x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 { //left upper

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by same piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "wB")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wB", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for i, j = x-1, y+1; i >= 0 && j <= 7; i, j = i-1, j+1 { //right upper

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wB")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wB", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}
	for i, j = x+1, y+1; i <= 7 && j <= 7; i, j = i+1, j+1 { //right lower

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wB")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wB", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}
	for i, j = x+1, y-1; i <= 7 && j >= 0; i, j = i+1, j-1 { //left lower

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wB")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wB", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}

	return true
}

func (table *Table) blackBishop(x int8, y int8) bool {
	var i int8
	var j int8
	for i, j = x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 { //left upper

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "bB")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bB", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
	}
	for i, j = x-1, y+1; i >= 0 && j <= 7; i, j = i-1, j+1 { //right upper

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "bB")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bB", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
	}
	for i, j = x+1, y+1; i <= 7 && j <= 7; i, j = i+1, j+1 { //right lower

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "bB")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bB", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
	}
	for i, j = x+1, y-1; i <= 7 && j >= 0; i, j = i+1, j-1 { //left lower

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "bB")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bB", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bB", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
	}

	return true
}

func (table *Table) whiteRook(x int8, y int8) bool {

	var i int8
	var j int8
	for i = x - 1; i >= 0; i-- { //up

		if (table.ChessBoard[i][y])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, y, "wR")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, y, "wR", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "wR", capturedPiece)
		if (table.ChessBoard[i][y])[0:1] == "b" {
			break
		}
	}
	for i = x + 1; i <= 7; i++ { //down

		if (table.ChessBoard[i][y])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, y, "wR")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, y, "wR", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "wR", capturedPiece)
		if (table.ChessBoard[i][y])[0:1] == "b" {
			break
		}
	}
	for j = y - 1; j >= 0; j-- { //left

		if (table.ChessBoard[x][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, x, j, "wR")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x, j, "wR", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "wR", capturedPiece)
		if (table.ChessBoard[x][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
	}
	for j = y + 1; j <= 7; j++ { //right

		if (table.ChessBoard[x][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, x, j, "wR")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x, j, "wR", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "wR", capturedPiece)
		if (table.ChessBoard[x][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
	}
	return true
}

func (table *Table) blackRook(x int8, y int8) bool {

	var i int8
	var j int8
	for i = x - 1; i >= 0; i-- { //up

		if (table.ChessBoard[i][y])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, y, "bR")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, y, "bR", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "bR", capturedPiece)
		if (table.ChessBoard[i][y])[0:1] == "w" {
			break
		}
	}
	for i = x + 1; i <= 7; i++ { //down

		if (table.ChessBoard[i][y])[0:1] == "b" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, y, "bR")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, y, "bR", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "bR", capturedPiece)
		if (table.ChessBoard[i][y])[0:1] == "w" {
			break
		}
	}
	for j = y - 1; j >= 0; j-- { //left

		if (table.ChessBoard[x][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, x, j, "bR")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x, j, "bR", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "bR", capturedPiece)
		if (table.ChessBoard[x][j])[0:1] == "w" {
			break
		}
	}
	for j = y + 1; j <= 7; j++ { //right

		if (table.ChessBoard[x][j])[0:1] == "b" {
			break
		}
		capturedPiece := table.makeMove(x, y, x, j, "bR")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x, j, "bR", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "bR", capturedPiece)
		if (table.ChessBoard[x][j])[0:1] == "w" {
			break
		}
	}
	return true
}

func (table *Table) whiteQueen(x int8, y int8) bool {

	var i int8
	var j int8

	for i, j = x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 { //left upper

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wQ", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}
	for i, j = x-1, y+1; i >= 0 && j <= 7; i, j = i-1, j+1 { //right upper

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wQ", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}
	for i, j = x+1, y+1; i <= 7 && j <= 7; i, j = i+1, j+1 { //right lower

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wQ", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}
	for i, j = x+1, y-1; i <= 7 && j >= 0; i, j = i+1, j-1 { //left lower

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, j, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, j, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "wQ", capturedPiece)
		if (table.ChessBoard[i][j])[0:1] == "b" {
			break
		}
	}

	//rook moves are now checked
	for i = x - 1; i >= 0; i-- { //up

		if (table.ChessBoard[i][y])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, y, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, y, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "wQ", capturedPiece)
		if (table.ChessBoard[i][y])[0:1] == "b" {
			break
		}
	}
	for i = x + 1; i <= 7; i++ { //down

		if (table.ChessBoard[i][y])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, i, y, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, i, y, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "wQ", capturedPiece)
		if (table.ChessBoard[i][y])[0:1] == "b" {
			break
		}
	}
	for j = y - 1; j >= 0; j-- { //left

		if (table.ChessBoard[x][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, x, j, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x, j, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "wQ", capturedPiece)
		if (table.ChessBoard[x][j])[0:1] == "b" {
			break
		}
	}
	for j = y + 1; j <= 7; j++ { //right

		if (table.ChessBoard[x][j])[0:1] == "w" {
			break
		}
		capturedPiece := table.makeMove(x, y, x, j, "wQ")
		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x, j, "wQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "wQ", capturedPiece)
		if (table.ChessBoard[x][j])[0:1] == "b" {
			break
		}
	}

	return true
}

func (table *Table) blackQueen(x int8, y int8) bool {
	var i int8
	var j int8

	for i, j = x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 { //left upper

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by same color
		}
		capturedPiece := table.makeMove(x, y, i, j, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bQ", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for i, j = x-1, y+1; i >= 0 && j <= 7; i, j = i-1, j+1 { //right upper

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bQ", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for i, j = x+1, y+1; i <= 7 && j <= 7; i, j = i+1, j+1 { //right lower

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bQ", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for i, j = x+1, y-1; i <= 7 && j >= 0; i, j = i+1, j-1 { //left lower

		if (table.ChessBoard[i][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, j, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, j, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, j, "bQ", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}

	//rook moves are now checked
	for i = x - 1; i >= 0; i-- { //up

		if (table.ChessBoard[i][y])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, i, y, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, y, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "bQ", capturedPiece)

		if (table.ChessBoard[i][y])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for i = x + 1; i <= 7; i++ { //down

		if (table.ChessBoard[i][y])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by same piece
		}
		capturedPiece := table.makeMove(x, y, i, y, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, i, y, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, i, y, "bQ", capturedPiece)

		if (table.ChessBoard[i][y])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for j = y - 1; j >= 0; j-- { //left

		if (table.ChessBoard[x][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, x, j, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x, j, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "bQ", capturedPiece)

		if (table.ChessBoard[i][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by enemy piece
		}
	}
	for j = y + 1; j <= 7; j++ { //right

		if (table.ChessBoard[x][j])[0:1] == "b" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
		capturedPiece := table.makeMove(x, y, x, j, "bQ")
		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x, j, "bQ", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, j, "bQ", capturedPiece)

		if (table.ChessBoard[x][j])[0:1] == "w" {
			break //no need to check rest of squares if the square is already occupied by a piece
		}
	}

	return true
}

func (table *Table) whiteKing(x int8, y int8) bool {
	//starting from top left corner
	if x-1 >= 0 && y-1 >= 0 && (table.ChessBoard[x-1][y-1])[0:1] != "w" { //top left

		capturedPiece := table.makeMove(x, y, x-1, y-1, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x - 1
		table.whiteKingY = y - 1
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y-1, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y-1, "wK", capturedPiece)

	}

	if x-1 >= 0 && (table.ChessBoard[x-1][y])[0:1] != "w" { //top middle

		capturedPiece := table.makeMove(x, y, x-1, y, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x - 1
		table.whiteKingY = y
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y, "wK", capturedPiece)
	}

	if x-1 >= 0 && y+1 <= 7 && (table.ChessBoard[x-1][y+1])[0:1] != "w" { //top right

		capturedPiece := table.makeMove(x, y, x-1, y+1, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x - 1
		table.whiteKingY = y + 1
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x-1, y+1, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x-1, y+1, "wK", capturedPiece)
	}

	if y-1 >= 0 && (table.ChessBoard[x][y-1])[0:1] != "w" { //middle left

		capturedPiece := table.makeMove(x, y, x, y-1, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x
		table.whiteKingY = y - 1
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x, y-1, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, y-1, "wK", capturedPiece)
	}

	if y+1 <= 7 && (table.ChessBoard[x][y+1])[0:1] != "w" { //middle right

		capturedPiece := table.makeMove(x, y, x, y+1, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x
		table.whiteKingY = y + 1
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x, y+1, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x, y+1, "wK", capturedPiece)
	}

	if x+1 <= 7 && y-1 >= 0 && (table.ChessBoard[x+1][y-1])[0:1] != "w" { //bottom left

		capturedPiece := table.makeMove(x, y, x+1, y-1, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x + 1
		table.whiteKingY = y - 1
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+1, y-1, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y-1, "wK", capturedPiece)
	}

	if x+1 <= 7 && (table.ChessBoard[x+1][y])[0:1] != "w" { //bottom middle

		capturedPiece := table.makeMove(x, y, x+1, y, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x + 1
		table.whiteKingY = y
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+1, y, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y, "wK", capturedPiece)
	}

	if x+1 <= 7 && y+1 <= 7 && (table.ChessBoard[x+1][y+1])[0:1] != "w" { //bottom right

		capturedPiece := table.makeMove(x, y, x+1, y+1, "wK")

		table.whiteOldX = x //storing king information
		table.whiteOldY = y

		table.whiteKingX = x + 1
		table.whiteKingY = y + 1
		table.kingUpdate = true

		if table.isWhiteInCheck() == false {
			table.undoMove(x, y, x+1, y+1, "wK", capturedPiece)
			return false
		}
		table.undoMove(x, y, x+1, y+1, "wK", capturedPiece)
	}

	return true
}

func (table *Table) blackKing(x int8, y int8) bool {
	//starting from top left corner
	if x-1 >= 0 && y-1 >= 0 && (table.ChessBoard[x-1][y-1])[0:1] != "b" { //top left

		capturedPiece := table.makeMove(x, y, x-1, y-1, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x - 1
		table.blackKingY = y - 1
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-1, y-1, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x-1, y-1, "bK", capturedPiece)
	}

	if x-1 >= 0 && (table.ChessBoard[x-1][y])[0:1] != "b" { //top middle

		capturedPiece := table.makeMove(x, y, x-1, y, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x - 1
		table.blackKingY = y
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-1, y, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x-1, y, "bK", capturedPiece)
	}

	if x-1 >= 0 && y+1 <= 7 && (table.ChessBoard[x-1][y+1])[0:1] != "b" { //top right

		capturedPiece := table.makeMove(x, y, x-1, y+1, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x - 1
		table.blackKingY = y + 1
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x-1, y+1, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x-1, y+1, "bK", capturedPiece)
	}

	if y-1 >= 0 && (table.ChessBoard[x][y-1])[0:1] != "b" { //middle left

		capturedPiece := table.makeMove(x, y, x, y-1, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x
		table.blackKingY = y - 1
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x, y-1, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x, y-1, "bK", capturedPiece)
	}

	if y+1 <= 7 && (table.ChessBoard[x][y+1])[0:1] != "b" { //middle right

		capturedPiece := table.makeMove(x, y, x, y+1, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x
		table.blackKingY = y + 1
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x, y+1, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x, y+1, "bK", capturedPiece)
	}

	if x+1 <= 7 && y-1 >= 0 && (table.ChessBoard[x+1][y-1])[0:1] != "b" { //bottom left

		capturedPiece := table.makeMove(x, y, x+1, y-1, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x + 1
		table.blackKingY = y - 1
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y-1, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x+1, y-1, "bK", capturedPiece)
	}

	if x+1 <= 7 && (table.ChessBoard[x+1][y])[0:1] != "b" { //bottom middle

		capturedPiece := table.makeMove(x, y, x+1, y, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x + 1
		table.blackKingY = y
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x+1, y, "bK", capturedPiece)
	}

	if x+1 <= 7 && y+1 <= 7 && (table.ChessBoard[x+1][y+1])[0:1] != "b" { //bottom right

		capturedPiece := table.makeMove(x, y, x+1, y+1, "bK")

		table.blackOldX = x //storing king information
		table.blackOldY = y

		table.blackKingX = x + 1
		table.blackKingY = y + 1
		table.kingUpdate = true

		if table.isBlackInCheck() == false {
			table.undoMove(x, y, x+1, y+1, "bK", capturedPiece)

			return false
		}
		table.undoMove(x, y, x+1, y+1, "bK", capturedPiece)
	}

	return true
}
