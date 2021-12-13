package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func doKnightPath(ctx *fiber.Ctx) error {

	fmt.Printf("<Welcome Player>")
	initialPosition := ctx.Params("coord")

	fmt.Printf("Coord: [%s]", initialPosition)
	if !isValid(initialPosition) {
		fmt.Printf("posicao INVALIDA!")
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.New("coord invalid"))
	}

	x, y := mapInitialPos(initialPosition)
	if x == 0 || y == 0 {
		fmt.Printf("Coordenadas invalidas")
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.New("coord invalid"))
	}

	result := addAllMoves(x, y)

	RemoveDuplicates := removeDuplicateValues(result)

	fmt.Printf("resultado: %+v", result)

	return ctx.Status(fiber.StatusOK).JSON(RemoveDuplicates)
}

func addMove(x int64, y int64) (valid bool, Xatual int64, Yatual int64) {
	if (x >= 1) && (x <= 8) && (y >= 1) && (y <= 8) {
		return true, x, y
	}
	return false, 0, 0
}

var moves = []struct{ dr, dc int64 }{
	{1, 2}, {2, 1}, {2, -1},
	{1, -2}, {-1, -2}, {-2, -1},
	{-2, +1}, {-1, +2}}

var possibleMoves []string

func addAllMoves(x int64, y int64) []string {
	for _, level1 := range moves {
		valid, nextX, nextY := addMove(x+level1.dr, y+level1.dc)
		if valid == false {
			continue
		}
		for _, level2 := range moves {
			valid, stopedX, stopedY := addMove(nextX+level2.dr, nextY+level2.dc)
			if valid == false {
				continue
			}
			movement := convertToAlg(stopedX, stopedY)
			if len(movement) < 2 {
				continue
			}
			possibleMoves = append(possibleMoves, movement)
		}
	}
	return possibleMoves
}

func mapInitialPos(s string) (line int64, column int64) {
	line = getNumberFromLetter(s[0:1])

	column = getCoordFromString(s[1:2])
	fmt.Printf("LINHA [%d], COLUNA [%d]", line, column)
	return line, column
}

func isValid(s string) bool {
	fields := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	if len(s) < 2 || len(s) > 3 {
		fmt.Printf("len invalid: %d", len(s))
		return false
	}
	coord, err := strconv.Atoi(s[1:2])
	if err != nil {
		fmt.Printf("coord: %d - error: %+v", coord, err)
		return false
	}
	if coord > 8 || coord < 0 {
		fmt.Printf("coord > 8 ou < 0 ")
		return false
	}
	for _, field := range fields {
		if s[0:1] == field {
			return true
		}
		fmt.Printf("field [%s] != [%s] ", s[0:1], field)
	}
	return false
}

func getNumberFromLetter(s string) int64 {
	switch s {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	default:
		return 0
	}
}

func setLetterFromNumber(i int64) string {
	switch i {
	case 1:
		return "a"
	case 2:
		return "b"
	case 3:
		return "c"
	case 4:
		return "d"
	case 5:
		return "e"
	case 6:
		return "f"
	case 7:
		return "g"
	case 8:
		return "h"
	default:
		return ""
	}
}

func getCoordFromString(s string) int64 {
	coord, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	if coord > 8 || coord < 0 {
		return 0
	}
	return coord
}

func convertToAlg(x int64, y int64) string {
	newy := strconv.FormatInt(y, 10)
	s := setLetterFromNumber(x)
	s = s + newy
	return s
}

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
