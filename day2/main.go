package main

import (
	_ "embed"
)

//go:embed ex.txt
var ex string

/*
A  | rock     1p | Y  paper  2p      win 6p
B  | paper    2p | X  rock  1p       loss 0p
C  | scissors 3p | Z  scissors 3p    draw 3p

Y draw
X loose
Z win
*/

func main() {

}
