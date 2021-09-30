package main

import (
	"testing"
)

type Scenario struct {
  Me Battlesnake
  State GameState
}

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 2, Y: 0},
		Body: []Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 100; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
		}
	}
}

// TODO tests
func TestAvoidWalls(t *testing.T) {
  testCases := []Scenario {
    // SCENARIO 1 - snake facing up
    {
      Me: Battlesnake{
        Head: Coord{X: 7, Y: 4},
        Body: []Coord{{X: 7, Y: 4}, {X: 7, Y: 3}, {X: 7, Y: 2}},
      },
      State: GameState{
        Board: Board{
          Height: 8,
          Width: 8,
        },
        Snakes: []Battlesnake{}
      },
    },
    // SCENARIO 2 - snake facing down
    {
      Me: Battlesnake{
        Head: Coord{X: 6, 4},
        Body: []Coord{{X: 6, Y: 4}, {X:6, Y: 5}, {X: 6, Y: 6}},
      },
      State: GameState{
        Board: {
          Height: 8,
          Width: 8,
        },
        Snakes: []Battlesnake{},
      },
    },
  }

  for _, scenario := range testCases {
    move(scenario.State)
  }
}

