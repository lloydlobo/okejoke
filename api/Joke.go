package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/fatih/color"
	// "github.com/lloydlobo/okejoke/api"
)

// Joke struct
//
// # Example
// $ curl -H "Accept: application/json" https://icanhazdadjoke.com/
//
//	{
//		 "id":      "aFtzPRSnbxc",
//		 "joke":    "Two fish are in a tank, one turns to the other and says,
//		            \"how do  you drive this thing?\"",
//		 "status":  200
//	}
type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// SearchResult struct
//
// Shape of the data to work with API response.
//
// Results:
// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type SearchResult struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

// /////////////////////////////////////////////////////
// API Calls | http package
// /////////////////////////////////////////////////////

// Function getOkejoke makes an API call to fetch joke.
//
// # It uses the `http` package.
// API Calls.
func GetOkejoke() {
	url := "https://icanhazdadjoke.com/" // Store API url in variable url.

	// Channels for passing Go routine pointers.
	chanBytes := make(chan []byte)
	chanJoke := make(chan string)

	// GO routine, do this while the next in line do their job.
	go func() {
		time.Sleep(time.Millisecond * 300)
		chanBytes <- FetchApiData(url)
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		chanJoke <- GetJokeParseData()
	}()

	// Save data into pointer 2nd arg `joke` Joke struct when the response is unmarshal.
	responseBytes := fetchJokeData(url)                          // Pass url into fetchData() method & store returned `response` bytes.
	joke := Joke{}                                               // Variable to store search results in `Joke` struct.
	if err := json.Unmarshal(responseBytes, &joke); err != nil { // Unmarshal returns an InvalidUnmarshalError.
		log.Printf("Could not unmarshal responseBytes. %v", err) // Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer,
	}

	// Save data into pointer 2nd arg `joke` Joke struct when the response is unmarshal.
	color.Set(color.FgRed, color.Bold) // defer color.Unset() // use it in your function
	fmt.Println(string(joke.Joke))     // Convert joke.Joke to a string, and print it to the terminal.
	color.Unset()

	// Returning the pointer. Wait for end of program without any other means of sync...
	resBytesGoRout := <-chanBytes
	resStringGoRout := <-chanJoke

	// Printing response from go routine channels
	fmt.Printf("\nchanDataFetched: %v\n", string(resBytesGoRout))

	color.Set(color.FgRed, color.Bold)
	fmt.Printf("chanresStringGoRout: %v\n", resStringGoRout)
	color.Unset()
}

// /////////////////////////////////////////////////////
// Flags for cobra commands.
// /////////////////////////////////////////////////////

// GetRandomJokeWithTerm searches through jokes with a search term.
//
// Usage:
// $ okejoke random --term=hipster
//
// Example:
// $ curl -H "Accept: application/json" "https://icanhazdadjoke.com/search?term=hipster"
// Main concern in the json response are the following terms:
// "search_term":"hipster","status":200,"total_jokes":3,"total_pages":1
//
// fmt.Printf("Search term: %v", jokeTerm) // for testing function.
func GetRandomJokeWithTerm(jokeTerm string) {
	totalJokes, results := GetJokeDataWithTerm(jokeTerm)
	randomizeJokesList(totalJokes, results, jokeTerm)
}

// randomizeJokesList()
//
// If Seed is not called, the generator behaves as if seeded by Seed\(1\)\. Seed values that have the same remainder when
// divided by 2³¹\-1 generate the same pseudo\-random sequence\.
// Seed, unlike the Rand\.Seed method, is safe for concurrent use\.
//
// Without Unix os current time. return j[rand.Intn(totalJokes)]
func randomizeJokesList(lengthJokes int, j []Joke, jokeTerm string) Joke {
	// Seed uses the provided seed value to INITIALIZE the default Source to a deterministic state\.
	rand.Seed(time.Now().Unix()) // Unix returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC\. The result does not depend on the location associated with t\.
	var output Joke

	min, max := 0, lengthJokes-1 // results arrray is 0 index, so do not exceed total length.
	jokeWithTermNotFound := lengthJokes <= 0

	if jokeWithTermNotFound {
		err := fmt.Errorf("No jokes found with the term: %v", jokeTerm) // Throw an error if joke with terms not found.

		color.Set(color.FgHiYellow) // defer color.Unset() // use it in your function.
		fmt.Println(err.Error())    // Print error to the console.
		color.Unset()               // Unset color printed to console for next print.
	} else {
		randInt := min + rand.Intn(max-min)
		output = j[randInt]
		color.Set(color.FgRed, color.Bold) // defer color.Unset() // use it in your function.
		fmt.Println(j[randInt].Joke)
		color.Unset()
	}
	return output
}

// GetJokeDataWithTerm
//
// API Call.
func GetJokeDataWithTerm(jokeTerm string) (totalJokes int, jokeList []Joke) {
	// Sprintf formats according to a format specifier and returns the resulting string.
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm) // Construct API url with search term user passes in.
	responseBytes := fetchJokeData(url)                                       // Pass url into fetchData() method & store returned `response` bytes.
	jokeListRaw := SearchResult{}                                             // Variable to STORE search results in SearchResult struct.

	// Part 1: Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Printf("Could not unmarshal responseBytes - %v", err) // Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer,
	} // NOTE: If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

	// Part 2: Unmarshall the result data and store in slice of `Joke`.
	jokes := []Joke{} // STORE
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("Could not unmarshal jokeListRaw results - %v", err) // Unmarshall returns error so handle it.
	}
	return jokeListRaw.TotalJokes, jokes
}

// /////////////////////////////////////////////////////
// API Calls with http package.
// /////////////////////////////////////////////////////

// fetchJokeData requests server with its custom headers with net/http package.
//
// To make a request with custom headers. use `NewRequest` and DefaultClient.Do“.
// `byte` is an alias for uint8 and is equivalent to uint8 in all ways.
//
// Common HTTP methods.
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
// [`http.MethodGet` on pkg.go.dev](https://pkg.go.dev/net/http?utm_source=gopls#MethodGet).
//
//	$  curl -H "Accept: application/json" https://icanhazdadjoke.com/
//
//	{
//	       "id":"aFtzPRSnbxc",
//	       "joke":"Two fish are in a tank, one turns to the other and says, \"how do
//	 you drive this thing?\"",
//	       "status":200
//	   }
func fetchJokeData(baseAPI string) []byte {
	// NewRequest wraps NewRequestWithContext using context.Background.
	// Create a new request. args: (HTTP method, url, request body).
	request, err := http.NewRequest(
		http.MethodGet, // method.
		baseAPI,        // url.
		nil,            // request body.
	)
	// Handle error if returned from http.NewRequest().
	if err != nil {
		log.Printf("Could not request an okejoke. %v", err)
	}
	// Add Header to request the API for response data as JSON.
	request.Header.Add("Accept", "application/json") // Add adds the key, value pair to the header.
	// Add a custom User-Agent to tell the API mantainers, how the API is being used.
	request.Header.Add("User-Agent", "Okejoke CLI (https://github.com/lloydlobo/okejoke)") // The key is case insensitive; it is canonicalized by CanonicalHeaderKey.
	// Pass the request to http.DefaultClient.Do(), method to get a response.
	// DefaultClient is the default Client, and is used by Get, Head, and Post.
	response, err := http.DefaultClient.Do(request) // do sends an http request and returns an http response, following policy (such as redirects, cookies, auth) as configured on the client.
	if err != nil {
		log.Printf("could not make a request. %v", err)
	} // Handle error returned from http.DefaultClient.Do() method.
	// Pass response.Body to ioutil.ReadAll()
	// to read into bytes.
	// Package ioutil implements some I/O utility functions.
	responseBytes, err := ioutil.ReadAll(response.Body) // ReadAll reads from r until an error or EOF and returns the data it read.
	if err != nil {
		log.Printf("Could not read the response body. %v", err)
	} // Handle error returned from ioutil.ReadAll() method.
	// Return response as bytes: responseBytes []byte.
	return responseBytes
}

func searchResultExample() string {
	jsonStr := `{ 
    "current_page":1,
    "limit":20, 
    "next_page":1, 
    "previous_page":1,
    "results":[
      {"id":"xc21Lmbxcib","joke":"How did the hipster burn the roof of his mouth? He ate the pizza before it was cool."},
      {"id":"GlGBIY0wAAd","joke":"How much does a hipster weigh? An instagram."},
      {"id":"NRuHJYgaUDd","joke":"How many h ipsters does it take to change a lightbulb? Oh, it's a really obscure number. You've probably never heard of it."} 
    ], 
    "search_term":"hipster",
    "status":200,
    "total_jokes":3,
    "total_pages":1}`
	return jsonStr
}

// Header contains the request header fields either received by the server or to be sent by the client.
//
// If a server received a request with header lines,
//     Host: example.com
//     accept-encoding: gzip, deflate
//     Accept-Language: en-us
//     fOO: Bar
//     foo: two
// then
//     Header = map[string][]string{
//     	"Accept-Encoding": {"gzip, deflate"},
//     	"Accept-Language": {"en-us"},
//     	"Foo": {"Bar", "two"},
//     }
//
// For incoming requests, the Host header is promoted to the Request.Host field and removed from the Header map.
// HTTP defines that header names are case-insensitive. The request parser implements this by using CanonicalHeaderKey, making the first character and any characters following a hyphen uppercase and the rest lowercase.
// For client requests, certain headers such as Content-Length and Connection are automatically written when needed and values in Header may be ignored. See the documentation for the Request.Write method.

// /////////////////////////////////////////////////////
// Logging Errors | Info
// /////////////////////////////////////////////////////

// 2022/09/18 18:50:48 fork/exec echo 'curl -H "Accept: application/json" https://icanhazdadjoke .com/': no such file or directory
// exit status 1
