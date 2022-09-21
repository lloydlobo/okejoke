package api

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"testing"
)

func TestFetchApiDataStaus200(t *testing.T) {
	url := "https://icanhazdadjoke.com/" // Store API url in variable url.
	res, want := FetchApiData(url), 200
	log.SetPrefix("TestFetchApiDataStaus200 ➜ ")
	log.SetFlags(5)

	got := DataJSON{} // Variable to store search results in `Joke` struct.
	if err := json.Unmarshal(res, &got); err != nil {
		log.Println(err)
	}
	log.Printf("➜ status: %v\n", got.Status)
	if got.Status != 200 {
		t.Errorf("TestFetchApiDataStaus200: got %v; want %v", got, want)
	}

}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// SIMPLE BENCHMARKING SUITE
// https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func primeNumbers(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// b.N specifies the number of iterations; the value is not fixed, but dynamically allocated, ensuring that the benchmark runs for at least one second by default.
func BenchmarkPrimeNumbers(b *testing.B) {
	num := 121
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}
