package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"log"
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
	// myNeck := state.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")

	// if myNeck.X < myHead.X { // neck is on the left
	// 	possibleMoves["left"] = false
	// } else if myNeck.X > myHead.X { // neck is on the right
	// 	possibleMoves["right"] = false
	// } else if myNeck.Y < myHead.Y { // neck is below head
	// 	possibleMoves["down"] = false
	// } else if myNeck.Y > myHead.Y { // neck is above head
	// 	possibleMoves["up"] = false
	// }

  // Step 1 - Don't hit walls.
  // Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
  wallMoves := avoidWalls(state.Board, myHead)
  for _, move := range wallMoves {
    possibleMoves[move] = false
  }

  // : Step 2 - Don't hit yourself.
  // Use information in GameState to prevent your Battlesnake from colliding with itself.
  // TODO avoid getting stuck on the board OR eliminate all all moves when in certain situations
  myBody := state.You.Body[1:] // the first coordinate is always the head, no need to include it
  for _, coord := range myBody {
    var direction string = avoidNeighbour(coord, myHead)

    if len(direction) > 0 {
      possibleMoves[direction] = false
    }
  }

  //TODO can i combine the coordinates for steps 2 & 3?

  // Step 3 - Don't collide with others.
  // Use information in GameState to prevent your Battlesnake from colliding with others.
  snakes := state.Board.Snakes
  myName := state.You.Name

  var otherSnakes []Battlesnake
  for _, snake := range snakes {
    if (snake.Name == myName) {
      continue
    }
    otherSnakes = append(otherSnakes, snake)
  }

  for _, snake := range otherSnakes {
    for _, coord := range snake.Body {
      var direction string = avoidNeighbour(coord, myHead)

      if len(direction) > 0 {
        possibleMoves[direction] = false
      }
    }
  }

  //  Step 4 - Find food.
  // Use information in GameState to seek out and find food.
  food := state.Board.Food
  // find food we can get to right now
  var foodInPath []Coord = findNeighbours(myHead, food)
  var closestFood Coord = findClosestNeighbour(state.Board, myHead, foodInPath)

   // eliminate all moves that aren't the desired direction
  if (closestFood.X == myHead.X) {
    yDistance := closestFood.Y - myHead.Y
    if yDistance > 0 {
      eliminateUnselectedMoves("up", possibleMoves)
    } else {
      eliminateUnselectedMoves("down", possibleMoves)
    }
  } else if (closestFood.Y == myHead.Y) {
    xDistance := closestFood.X - myHead.X
    if xDistance > 0 {
      eliminateUnselectedMoves("right", possibleMoves)
    } else {
      eliminateUnselectedMoves("left", possibleMoves)
    }
  }

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
    Shout: "YAY",
	}
}
