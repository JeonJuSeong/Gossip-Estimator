package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	allNodes := flag.Int("nodes", 8044124, "Number of nodes in the entire network")
	fanOut := flag.Int("fanout", 20, "Number of times to propagate to other nodes")
	loopCount := flag.Int("loop", 2, "Number of loop")
	flag.Parse()

	file, err := os.Create("measurement.txt")
	if err != nil {
		fmt.Println("having trouble opening or creating a file.:", err)
		return
	}
	defer file.Close()
	for x := 0; x < *loopCount; x++ {
		startTest(file, allNodes, fanOut)
	}
}

func startTest(file *os.File, allNodes *int, fanOut *int) {
	redunTable, list20 := initTest(allNodes, fanOut)
	redundant := 0
	nodes := 1
	gossip := 0

	for {
		gossip++
		for j := 0; j < nodes; j++ {
			for i := 0; i < *fanOut; i++ {
				randomInt := 0
				for z := 0; z < i; z++ {
					randomInt = rand.Intn(*allNodes)
					if list20[z] == randomInt {
						z--
					}
				}
				list20[i] = randomInt
				redunTable[randomInt]++
			}
			for z := 0; z < *fanOut; z++ {
				list20[z] = 0
			}
		}
		nodes = 0
		for k := 0; k < *allNodes; k++ {
			if redunTable[k] != 0 {
				redundant += (redunTable[k] - 1)
				nodes++
			}
		}

		//if all nodes were propagated
		if nodes == *allNodes {
			break
		}
	}
	for i := 0; i < *allNodes; i++ {
		if redunTable[i] != 1 {
			redundant += (redunTable[i] - 1)
		}
	}
	file.WriteString("redundant: " + strconv.Itoa(redundant) + "\t| gossip: " + strconv.Itoa(gossip) + "\n")
	fmt.Printf("redundant: %d, gossip: %d\n", redundant, gossip)
}

func initTest(allNodes *int, fanOut *int) ([]int, []int) {
	list := make([]int, *allNodes)
	list2 := make([]int, *fanOut)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return list, list2
}
