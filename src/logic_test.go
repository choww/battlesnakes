package main

import (
  "testing"
  "reflect"
)


type Scenario struct {
  Desc string
  Head Coord
  Avoid Coord
  AvoidMany []Coord
  Expected interface{}
}

func TestAvoidance(t *testing.T) {
  t.Run("avoid moving out of bounds", func (t *testing.T) {
	  board := Board{
      Width: 8,
      Height: 8,
	  }
    scenarios := []Scenario {
      {
        Desc: "avoid up",
        Head: Coord{3,7},
        Expected: []string{"up"},
      },
      {
        Desc: "avoid down",
        Head: Coord{3,0},
        Expected: []string{"down"},
      },
      {
        Desc: "avoid left",
        Head: Coord{0,3},
        Expected: []string{"left"},
      },
      {
        Desc: "avoid right",
        Head: Coord{7,3},
        Expected: []string{"right"},
      },
      {
        Desc: "avoid up & left",
        Head: Coord{0,7},
        Expected: []string{"up", "left"},
      },
      {
        Desc: "avoid up & right",
        Head: Coord{7,7},
        Expected: []string{"up", "right"},
      },
      {
        Desc: "avoid down & right",
        Head: Coord{7,0},
        Expected: []string{"down", "right"},
      },
      {
        Desc: "avoid down & left",
        Head: Coord{0,0},
        Expected: []string{"down", "left"},
      },
    }

    for _, scenario := range scenarios {
      got := avoidWalls(board, scenario.Head)
      if (!reflect.DeepEqual(got, scenario.Expected)) {
        t.Errorf("Scenario %s: expected to avoid %#v, got %#v", scenario.Desc, scenario.Expected, got)
      }
    }
  })

  t.Run("avoid anything right next to my head", func(t *testing.T) {
    scenarios := []Scenario {
      {
        Desc: "avoid up",
        Head: Coord{4,8},
        Avoid: Coord{4,9},
        Expected: "up",
      },
      {
        Desc: "avoid down",
        Head: Coord{4,8},
        Avoid: Coord{4,7},
        Expected: "down",
      },
      {
        Desc: "avoid left",
        Head: Coord{4,8},
        Avoid: Coord{3,8},
        Expected: "left",
      },
      {
        Desc: "avoid right",
        Head: Coord{4,8},
        Avoid: Coord{5,8},
        Expected: "right",
      },
      {
        Desc: "2 cells up",
        Head: Coord{4,7},
        Avoid: Coord{4,9},
        Expected: "",
      },
    }

    for _, scenario := range scenarios {
      got := avoidNeighbour(scenario.Avoid, scenario.Head)
      if (got != scenario.Expected) {
        t.Errorf("Scenario %s: expected to avoid %s, got %s", scenario.Desc, scenario.Expected, got)
      }
    }
  })

  t.Run("find closest neighbour", func(t *testing.T) {
	  board := Board{
      Width: 7,
      Height: 7,
	  }
    scenarios := []Scenario{
      {
        Desc: "scenario 1",
        Head: Coord{X: 4, Y:3},
        AvoidMany: []Coord{
          {X: 2, Y: 0},
          {X: 1, Y: 4},
          {X: 3, Y: 3},
        },
        Expected: Coord{X: 3, Y: 3},
      },
      {
        Desc: "scenario 2",
        Head: Coord{X: 0, Y: 1},
        AvoidMany: []Coord{
          {X: 1, Y: 1},
          {X: 1, Y: 0},
          {X: 2, Y: 0},
        },
        Expected: Coord{X: 1, Y: 1},
      },
    }

    for _, scenario := range scenarios {
      neighboursInPath := findNeighboursInPath(scenario.Head, scenario.AvoidMany)
      got := findClosestNeighbour(board, scenario.Head, neighboursInPath)
      if (!reflect.DeepEqual(got, scenario.Expected)) {
        t.Errorf("Scenario %s: expected to find %#v, got %#v", scenario.Desc, scenario.Expected, got)
      }
    }
  })

  // t.Run("avoid being stuck", func(t *testing.T) {
	//   board := Board{
  //     Width: 7,
  //     Height: 7,
	//   }

  //   scenarios := []Scenario{
  //     {
  //       Desc: "scenario 1",
  //       Head: Coord{X: 6, Y: 1},
  //       AvoidMany: []Coord{
  //         {X: 5, Y: 1},
  //         {X: 5, Y: 0},
  //         {X: 4, Y: 0},
  //       },
  //       Expected: "down",
  //     },
  //     {
  //       Desc: "scenario 2",
  //       Head: Coord{X: 0, Y: 1},
  //       AvoidMany: []Coord{
  //         {X: 1, Y: 1},
  //         {X: 1, Y: 0},
  //         {X: 2, Y: 0},
  //       },
  //       Expected: "down",
  //     },
  //   }

  //   for _, scenario := range scenarios {
  //     neighboursInPath := findNeighboursInPath(scenario.Head, scenario.AvoidMany)
  //     got := findClosestNeighbour(board, scenario.Head, neighboursInPath)
  //     if (!reflect.DeepEqual(got, scenario.Expected)) {
  //       t.Errorf("Scenario %s: expected to find %#v, got %#v", scenario.Desc, scenario.Expected, got)
  //     }
  //   }
  // })
}

// TODO e2e test
// func TestNeckAvoidance(t *testing.T) {
// 	// Arrange
// 	me := Battlesnake{
// 		// Length 3, facing right
// 		Head: Coord{X: 2, Y: 0},
// 		Body: []Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
// 	}
// 	state := GameState{
// 		Board: Board{
// 			Snakes: []Battlesnake{me},
// 		},
// 		You: me,
// 	}
//
// 	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
// 	for i := 0; i < 100; i++ {
// 		nextMove := move(state)
// 		// Assert never move left
// 		if nextMove.Move == "left" {
// 			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
// 		}
// 	}
// }

