package main

func main() {

}

func countOdds(low int, high int) int {
	count := 0
	mid := high - low + 1
	count = mid / 2
	if mid%2 == 0 {
		return count
	}
	if low%2 == 0 && high%2 == 1 {
		count += 1
	}
	if low%2 == 1 && high%2 == 1 {
		count += 1
	}
	return count
}
