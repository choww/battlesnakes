package main

import (
  "math"
)

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

func avoid(coordinate Coord, myHead Coord) (direction string) {
  if (coordinate.Y != myHead.Y && coordinate.X != myHead.X) {
    return
  } else if coordinate.Y == myHead.Y {
    distanceFromHead := myHead.X - coordinate.X
    direction = moveToEliminate(distanceFromHead, "X")
  } else if coordinate.X == myHead.X {
    distanceFromHead := myHead.Y - coordinate.Y
    direction = moveToEliminate(distanceFromHead, "Y")
  }
  return direction
}
