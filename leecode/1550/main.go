package main

import "golang_practice/log_for_app/log"

func main() {

}

func threeConsecutiveOdds(arr []int) bool {
	oddConsecutiveOdds := make([]int, 0)
	log.Println("count start!!")
	for _, number := range arr {
		if number%2 != 0 {
			oddConsecutiveOdds = append(oddConsecutiveOdds, number)
		} else {
			oddConsecutiveOdds = []int{}
		}

		if len(oddConsecutiveOdds) >= 3 {
			log.Print("right")
			log.Println("count over!!")
			return true
		}
	}
	log.Println("count over!!")
	return false
}
