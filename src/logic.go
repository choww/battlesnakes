package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"log"
  "math"
	"math/rand"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "choww",
		Color:      "#f5b900",
		Head:       "scarf",
		Tail:       "mouse",
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// get move to eliminate based on distance from head
// axis refers to whether the distance value is along the X or Y axis
func moveToEliminate(distance int, axis string) string {
  direction := map[string]map[string]string{
    "X": {
      "positive": "left",
      "negative": "right",
    },
    "Y": {
      "positive": "down",
      "negative": "up",
    },
  }
  isNextToHead := math.Abs(float64(distance)) == 1
  if isNextToHead {
    if distance > 0 {
      return direction[axis]["positive"]
    }

    return direction[axis]["negative"]
  }
  return ""
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")
	if myNeck.X < myHead.X {
		possibleMoves["left"] = false
	} else if myNeck.X > myHead.X {
		possibleMoves["right"] = false
	} else if myNeck.Y < myHead.Y {
		possibleMoves["down"] = false
	} else if myNeck.Y > myHead.Y {
		possibleMoves["up"] = false
	}

	// Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := state.Board.Width - 1
	boardHeight := state.Board.Height - 1

  // distance of head from the walls
  distanceFromTop := boardHeight - myHead.Y
  distanceFromRight := boardWidth - myHead.X

  if distanceFromTop == 0 {
    possibleMoves["up"] = false
  } else if myHead.Y == 0 {  // at the bottom wall
    possibleMoves["down"] = false
  }

  if distanceFromRight == 0 {
    possibleMoves["right"] = false
  } else if myHead.X == 0 {  // at the left wall
    possibleMoves["left"] = false
  }


	// : Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
  myBody := state.You.Body[1:] // the first coordinate is always the head, no need to include it
  for _, coord := range myBody {
    var direction string
    // we only need to worry about coordinates that have the same Y coordinate and is X +/- 1 away
    if coord.Y == myHead.Y {
      distanceFromHead := myHead.X - coord.X
      direction = moveToEliminate(distanceFromHead, "X")
    } else if coord.X == myHead.X {
      distanceFromHead := myHead.Y - coord.Y
      direction = moveToEliminate(distanceFromHead, "Y")
    }

    if len(direction) > 0 {
      possibleMoves[direction] = false
    }
  }

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}


	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}

  log.Printf("🐍 body: %v", state.You.Body)
  log.Printf("save moves: %v\n", safeMoves)

	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
