package main

import "fmt"

type Graph struct {
	numberOfNodes int
	adjacentList  map[int][]int
}

func (g *Graph) addVertex(value int) {
	if g.adjacentList[value] != nil {
		return
	} else {
		if g.adjacentList == nil {
			g.adjacentList = make(map[int][]int)
		}
		g.adjacentList[value] = []int{}
		g.numberOfNodes++
	}
}

func checkIfExists(array []int, value int) bool {
	for _, num := range array {
		if num == value {
			return true
		}
	}
	return false
}

func (g *Graph) addEdge(fromNode int, toNode int) {
	if g.adjacentList[fromNode] == nil || g.adjacentList[toNode] == nil {
		return
	}

	if checkIfExists(g.adjacentList[fromNode], toNode) {
		return
	}

	g.adjacentList[fromNode] = append(g.adjacentList[fromNode], toNode)
	g.adjacentList[toNode] = append(g.adjacentList[toNode], fromNode)
}

func (g *Graph) printConnections() {
	for vertex, edges := range g.adjacentList {
		for _, edge := range edges {
			fmt.Printf("Vertex %d is connected to Vertex %d\n", vertex, edge)
		}
	}
}

func (g *Graph) printAdjacencyList() {
	for vertex, edges := range g.adjacentList {
		fmt.Printf("%d: %v\n", vertex, edges)
	}
}
