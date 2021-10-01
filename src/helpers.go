package main

import (
  "math"
  "strings"
)

func eliminateUnselectedMoves(selectedMove string, possibleMoves map[string]bool) {
  for direction, _ := range possibleMoves {
    if (direction == selectedMove) {
      continue
    } else {
      possibleMoves[direction] = false
    }
  }
}

// avoid move if target is right next to head
// axis refers to whether the distance value is along the X or Y axis
func moveToAvoid(distance int, axis string) string {
  directions := map[string]map[string]string{
    "X": {
      "positive": "right",
      "negative": "left",
    },
    "Y": {
      "positive": "up",
      "negative": "down",
    },
  }
  axis = strings.ToUpper(axis)
  isNextToHead := math.Abs(float64(distance)) == 1
  if isNextToHead {
    if distance > 0 {
      return directions[axis]["positive"]
    }

    return directions[axis]["negative"]
  }
  return ""
}

func avoidNeighbour(coord Coord, myHead Coord) (direction string) {
  if (coord.Y != myHead.Y && coord.X != myHead.X) {
    return
  } else if coord.Y == myHead.Y {
    distanceFromHead := coord.X - myHead.X
    direction = moveToAvoid(distanceFromHead, "X")
  } else if coord.X == myHead.X {
    distanceFromHead := coord.Y - myHead.Y
    direction = moveToAvoid(distanceFromHead, "Y")
  }
  return direction
}

func avoidWalls(board Board, myHead Coord) (movesToAvoid []string) {
  boardWidth := board.Width - 1
  boardHeight := board.Height - 1

  isAtTopWall := myHead.Y == boardHeight
  isAtRightWall := myHead.X == boardWidth
  isAtBottomWall := myHead.Y == 0
  isAtLeftWall := myHead.X == 0

  if isAtTopWall {
    movesToAvoid = append(movesToAvoid, "up")
  } else if isAtBottomWall {
    movesToAvoid = append(movesToAvoid, "down")
  }

  if isAtRightWall {
    movesToAvoid = append(movesToAvoid, "right")
  } else if isAtLeftWall {
    movesToAvoid = append(movesToAvoid, "left")
  }

  return movesToAvoid
}

// find coordinates on the same X or Y axis
func findNeighbours(myHead Coord, coordinates []Coord) (neighbours []Coord) {
  for _, coord := range coordinates {
    if (coord.X == myHead.X || coord.Y == myHead.Y) {
      neighbours = append(neighbours, coord)
    }
  }
  return neighbours
}

func findClosestNeighbour(board Board, myHead Coord, neighbours []Coord) (neighbour Coord) {
  // this is the current closest neighbour--start off at max value
  neighbour = Coord{
    X: board.Width,
    Y: board.Height,
  }
  for _, coord := range neighbours {
    targetLocation := float64(coord.X + coord.Y)
    headLocation := float64(myHead.X + myHead.Y)

    currMinValue := float64(neighbour.X + neighbour.Y)
    targetDistance := math.Abs(headLocation - targetLocation)
    comparator := math.Min(currMinValue, targetDistance)

    // set the new cloest coordinate
    if (comparator < currMinValue) {
      neighbour = coord
    }
  }

  return neighbour
}


