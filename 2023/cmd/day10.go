package cmd

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type tile struct {
	coord
	pipe pipe
}

func (t *tile) isStartTile() bool {
	return t.pipe.isStarting
}

func (t tile) String() string {
	return fmt.Sprintf("(%v, %v): %v", t.x, t.y, t.pipe.value)
}

type coord struct {
	x int
	y int
}

type pipe struct {
	value      string
	N          bool
	E          bool
	S          bool
	W          bool
	isStarting bool
}

func (t *tile) connectedTo(other tile) bool {
	if other.pipe.value == "." {
		return false
	}

	switch t.pipe.value {
	case "|":
		if t.x == other.x {
			if (t.y - 1) == other.y { // north of
				return other.pipe.S
			} else if (t.y + 1) == other.y { // south of
				return other.pipe.N
			}
		}
	case "-":
		if t.y == other.y {
			if (t.x - 1) == other.x { // west of of
				return other.pipe.E
			} else if (t.x + 1) == other.x { // east of
				return other.pipe.W
			}
		}
	case "L":
		if t.y-1 == other.y && t.x == other.x { // north of
			return other.pipe.S
		} else if t.x+1 == other.x && t.y == other.y { // east of
			return other.pipe.W
		}
	case "J":
		if t.y == other.y && t.x-1 == other.x { // west of
			return other.pipe.E
		} else if t.y-1 == other.y && t.x == other.x { // north of
			return other.pipe.S
		}
	case "7":
		if t.y == other.y && t.x-1 == other.x { // west of
			return other.pipe.E
		} else if t.x == other.x && t.y+1 == other.y { // south of
			return other.pipe.N
		}
	case "F":
		if t.x == other.x && t.y+1 == other.y { // south of
			return other.pipe.N
		} else if t.x+1 == other.x && t.y == other.y { // east of
			return other.pipe.W
		}
	case ".":
		return false
	case "S":
		if t.x == other.x {
			if (t.y - 1) == other.y { // north of
				return other.pipe.S
			} else if (t.y + 1) == other.y { // south of
				return other.pipe.N
			}
		} else if t.y == other.y {
			if (t.x - 1) == other.x { // west of of
				return other.pipe.E
			} else if (t.x + 1) == other.x { // east of
				return other.pipe.W
			}
		}
	default:
		return false
	}

	return false
}

func tilesBordering(x, y int, field [][]tile) []tile {
	tiles := make([]tile, 0)

	if y-1 >= 0 {
		tiles = append(tiles, field[y-1][x])
	}
	if x+1 < len(field[y]) {
		tiles = append(tiles, field[y][x+1])
	}
	if y+1 < len(field) {
		tiles = append(tiles, field[y+1][x])
	}
	if x-1 >= 0 {
		tiles = append(tiles, field[y][x-1])
	}

	return tiles
}

func connectedTilebordering(x, y int, field [][]tile) []tile {
	curTile := field[y][x]
	connectedTiles := make([]tile, 0, 2)

	for _, o := range tilesBordering(x, y, field) {
		if curTile.connectedTo(o) {
			connectedTiles = append(connectedTiles, o)
		}
	}
	return connectedTiles
}

type step struct {
	curTile  tile
	prevTile tile
	distance int
}

func (s step) String() string {
	return fmt.Sprintf("Cur %v, prev %v, distance: %v\n", s.curTile, s.prevTile, s.distance)
}

func day10(inputData []string) int {
	var S tile

	field := make([][]tile, len(inputData))
	pipes := make(map[string]pipe)
	pipes["|"] = pipe{value: "|", N: true, S: true}
	pipes["-"] = pipe{value: "-", E: true, W: true}
	pipes["L"] = pipe{value: "L", N: true, E: true}
	pipes["J"] = pipe{value: "J", N: true, W: true}
	pipes["7"] = pipe{value: "7", S: true, W: true}
	pipes["F"] = pipe{value: "F", S: true, E: true}
	pipes["."] = pipe{value: "."}
	pipes["S"] = pipe{value: "S", N: true, E: true, S: true, W: true, isStarting: true}

	for y, row := range inputData {
		tilesRaw := strings.Split(row, "")
		tiles := make([]tile, len(tilesRaw))

		for x, t := range tilesRaw {
			tiles[x] = tile{coord: coord{x: x, y: y}, pipe: pipes[t]}

			if tiles[x].isStartTile() {
				S = tiles[x]
			}
		}

		field[y] = tiles
	}

	connectedTiles := connectedTilebordering(S.x, S.y, field)
	dir1 := step{curTile: connectedTiles[0], distance: 1, prevTile: S}
	dir2 := step{curTile: connectedTiles[1], distance: 1, prevTile: S}

	for dir1.curTile != dir2.curTile {
		dir1Connected := connectedTilebordering(dir1.curTile.x, dir1.curTile.y, field)
		dir1Connected = slices.DeleteFunc(dir1Connected, func(t tile) bool { return t == dir1.prevTile })
		dir2Connected := connectedTilebordering(dir2.curTile.x, dir2.curTile.y, field)
		dir2Connected = slices.DeleteFunc(dir2Connected, func(t tile) bool { return t == dir2.prevTile })

		dir1.prevTile = dir1.curTile
		dir1.distance++
		dir1.curTile = dir1Connected[0]
		dir2.prevTile = dir2.curTile
		dir2.curTile = dir2Connected[0]
	}

	return dir1.distance
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day10",
		Run: func(_ *cobra.Command, _ []string) {
			start := time.Now()

			res := day10(inputData)

			fmt.Printf("Result: %v ran in %v\n", res, time.Since(start))
		},
	},
	)
}
