package leetcode815

import (
	fifo "github.com/foize/go.fifo"
)

func GetTrainsCount(r [][]int, s int, d int) int {

	jun := make(map[int][]int)

	//add all the junctions
	//i represet the train number
	//j represents the station of each trains
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			jun[r[i][j]] = append(jun[r[i][j]], i)
		}
	}

	type trainCount struct {
		station int
		count   int
	}

	//map for visited train
	visitedTrains := make(map[int]bool)

	//routes = [[1, 2, 7], [5, 3, 6, 7]]
	bfsQueue := fifo.NewQueue()
	bfsQueue.Add(trainCount{count: 0, station: s})
	for bfsQueue.Len() != 0 {
		st := bfsQueue.Next().(trainCount)
		// get all the trains coming to station st.station
		ts := jun[st.station]
		//get the all routes of the train t
		for t := range ts {
			// if we already explore the train t then ignore it
			if !visitedTrains[t] {
				for i := 0; i < len(r[t]); i++ {
					// if ith station is destination
					if r[t][i] == d {
						return st.count + 1
					}
					//if not add into bfs queue
					bfsQueue.Add(trainCount{count: st.count + 1, station: r[t][i]})
				}
				visitedTrains[t] = true
			}
		}
	}
	return -1
}
