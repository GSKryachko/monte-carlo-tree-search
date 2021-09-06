package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Game interface {
	Init() Game
	GetMoves() []Game
	CountLegalMoves() int
	GetWinner() int
	PickRandom() Game
}

type MonteCarloTreeNode struct {
	parent             *MonteCarloTreeNode
	exploredChildren   map[Game]MonteCarloTreeNode
	unexploredChildren []Game
	gamesWon           int
	gamesPlayed        int
	state              Game
}

func (m *MonteCarloTreeNode) PropagateStats(won bool) {
	if won {
		m.gamesWon += 1
	}
	if m.parent != nil {
		m.parent.PropagateStats(won)
	}
}

type NodeStats struct {
	gamesWon    int
	gamesPlayed int
}

func (m *MonteCarloTreeNode) Iterate() {
	node := m
	// 1. Selection
	for len(node.unexploredChildren) == 0 {
		node = node.PickBest()
	}
	// 2. Expansion
	game = game.PickRandomUnexplored()
	// 3. Simulation
	next := Game.PickRandom()
	for next != nil {
		next = Game.PickRandom()
	}
	// 4.Propagation
	game.UpdateStats()
}

func (m *MonteCarloTreeNode) PickBest() Game {
	var bestScore float64
	var bestMove Game
	for child, move := range m.exploredChildren {
		//explored children always have at least one played game
		score := child.gamesWon / child.gamesPlayed
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	if bestScore == 0 {
		return m.state.PickRandom()
	}
	return bestMove
}
