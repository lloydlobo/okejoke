package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// data := DataJSON{}
// Store data in a variable.
type DataJSON struct {
	ID     string `json:"id"`
	Body   string `json:"joke"`
	Status int    `json:"status"`
}

/*
Construct SearchResJSON with curl results.
$ curl -H "Accept: application/json" https://icanhazdadjoke.com/search
{ "current_page":1,"limit":20,"next_page":2,"previous_page":1, "results":[ {"id":"0189hNRf2g","joke":"I'm tired of following my dreams. I'm just going to ask them where they are going and meet up with them later."},{"id":"08EQZ8EQukb","joke":"Did you hear about the guy whose whole left side was cut off? He's all right now."},{"id":"08xHQCdx5Ed","joke":"Why didn\u2019t the skeleton cross the road? Because he had no guts."},{"id":"0DQKB51oGlb"," joke":"What did one nut say as he chased another nut?  I'm a cashew!"},{"id":"0DtrrOZDlyd","joke":"Chances are if you' ve seen one shopping center, you've seen a mall."},{"id":"0LuXvkq4Muc","joke":"I knew I shouldn't steal a mixer from work, but it was a whisk I was willing to take."},{"id":"0ga 2EdN7prc","joke":"How come the stadium got hot after the game? Because all of the fans left."},{"id":"0oO71TSv4Ed","joke":"Why was it called the dar k ages? Because of all the knights. "},{"id":"0oz51ozk3ob","joke":"A steak pun is a rare medium well done."},{"id":"0ozAXv4Mmjb","joke":"Why did the tomato blush? Because it saw the salad dressing."},{"id":"0wcFBQfiGBd","joke":"Did you hear the joke about the wandering nun? She was a roman catho lic."},{"id":"189xHQ7pOuc","joke":"What creature is smarter than a talking parrot? A spelling bee."},{"id":"18Elj3EIYvc","joke":"I'll tell you what often gets over looked... garden fences."},{"id":"18h3wcU8xAd","joke":"Why did the kid cross the playground? To get to the other slide."},{"id":"1DI RSfx51Dd","joke":"Why do birds fly south for the winter? Because it's too far to walk."},{"id":"1DQZDY0gVnb","joke":"What is a centipedes's favorite Beatle song?  I want to hold your hand, hand, hand, hand..."},{"id":"1DQZvXvX8Ed","joke":"My first time using an elevator was an uplifting experien ce. The second time let me down."},{"id":"1DQZvcFBdib","joke":"To be Frank, I'd have to change my name."},{"id":"1Dt4M7Ufaxc","joke":"Slept like a l og last night \u2026 woke up in the fireplace."},{"id":"1T01LBXLuzd","joke":"Why does a Moon-rock taste better than an Earth-rock? Because it's a li ttle meteor."}
], "search_term":"", "status":200, "total_jokes":649, "total_pages":33 }
*/
type SearcResJSON struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

//////////////////////////////////////////////////////////////////////
// FetchApiData Methods:
//////////////////////////////////////////////////////////////////////
//
// Store data in a variable.
// request, add headers, get response, read & return response to bytes.

// requestHttpAPI() requests the baseAPI url with GET method..
func requestHttpAPI(baseAPI string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	if err != nil {
		log.Printf("http.NewRequest: could not request an okejoke. %v", err)
	}
	return req
}

// Request Data from API endpoint.
//
// API based on the provided HTTP Accept header.
// Accepted Accept headers:
// - text/html - HTML response (default response format)
// - application/json - JSON response
// - text/plain - Plain text response
// "Note: Requests made via curl which do not set an Accept header will respond with text/plain by default."
//
// Custom user agent:
// Setting a custom User-Agent header for your code will help us be able to better monitor the usage of the API and identify potential bad actors.
// A good user agent should include the name of the library or website that is accessing the API along with a URL/e-email where someone can be contacted regarding the library/website.
func addRequestHeaders(r *http.Request) {
	r.Header.Add("Accept", "application/json")
	r.Header.Add("User-Agent", "OkeJoke CLI (https://github.com/lloydlobo/okejoke)")
}

// Handle Response after Request Data from API endpoint.
// API response format:
// All API endpoints follow their respective browser URLs, but we adjust the response formatting to be more suited for an,
// For example: curl -H "User-Agent: My Library (https://github.com/username/repo)" https://icanhazdadjoke.com/
func getResponseFromRequest(r *http.Request) *http.Response {
	res, err := http.DefaultClient.Do(r) // DefaultClient is the default Client and is used by Get, Head, and Post\.
	if err != nil {
		log.Printf("http.DefaultClient.Do: could not make a request. %v", err)
	}
	return res
}

// ReadAll reads from r until an error or EOF and returns the data it read\.
func readResponseBodyBytes(r *http.Response) []byte {
	body, err := io.ReadAll(r.Body) // field Body io.ReadCloser
	if err != nil {
		log.Printf("io.ReadAll: could not read the response body. %v", err)
	}
	return body
}

//////////////////////////////////////////////////////////////////////
// FetchApiData function for chaining `http` methods.
//////////////////////////////////////////////////////////////////////

// FetchApiData requests, gets response with custom headers.
// and then returns the parsed response.Body bytes into string.
//
// Requirements : $ curl -H "Accept: application/json" https://icanhazdadjoke.com/
func FetchApiData(baseAPI string) []byte {
	req := requestHttpAPI(baseAPI)
	addRequestHeaders(req)
	res := getResponseFromRequest(req)
	body := readResponseBodyBytes(res)
	return body
}

// UnmarshalBytes parses the JSON\-encoded data and stores the result.
//
// in the value pointed to by v\. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
func UnmarshalBytes(bRes []byte, d DataJSON) DataJSON {
	if err := json.
		Unmarshal(bRes, &d); err != nil {
		log.Printf("UnmarshalBytes: Could not unmarshal response bytes. %v", err)
	}
	return d
}

// GetJokeParseData()
func GetJokeParseData() string {
	url := "https://icanhazdadjoke.com/" // Store API url in variable url.
	resBytes := FetchApiData(url)        // Pass url into fetchData() method & store returned `response` bytes.
	d := DataJSON{}                      // Variable to store search results in `Joke` struct.
	dJoke := UnmarshalBytes(resBytes, d) // Unmarshal parses the JSON\-encoded data and stores the result
	return string(dJoke.Body)            // Convert joke.Joke to a string, and print it to the terminal.
}

// FetchUserTermFlagData passes search term Cobra cmd random.go init() --flags.
//
// Wildcards	Description  ---- https://www.shell-tips.com/bash/wildcards-globbing/#gsc.tab=0
// ?	        A question-mark is a pattern that matches any single character.
// *	        An asterisk is a pattern that matches any number of any characters, including the null string/none.
// [...]        The square brackets matches any one of the enclosed characters.
//
//	17:34  ➜  go run main.go random --term=joke
//
// # Find search terms with wildcard-globbing
// curl -H "Accept: application/json" "https://icanhazdadjoke.com/search?term=hipster"
func FetchUserTermFlagData(jokeTerm string) (totalJokes int, jokeList []DataJSON) {
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm) // Construct API url with search term user passes in.
	responseBytes := FetchApiData(url)                                        // Pass url into fetchData() method & store returned `response` bytes.
	jokeListRaw := SearcResJSON{}                                             // Variable to STORE search results in SearchResult struct.
	jokes := []DataJSON{}                                                     // STORE: Declare & initalize a slice of DataJSON struct.

	log.SetPrefix("FetchUserTermFlagData()") // log.SetFlags(2)
	log.Printf("responseBytes: %v\n", string(responseBytes))

	// TODO
	// Part 1: Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	// Part 2: Unmarshall the result data and store in slice of `Joke`.
	return jokeListRaw.TotalJokes, jokes
}

/*
Search for dad jokes
GET https://icanhazdadjoke.com/search - search for dad jokes.

This endpoint accepts the following optional query string parameters:

page - which page of results to fetch (default: 1)
limit - number of results to return per page (default: 20) (max: 30)
term - search term to use (default: list all jokes)
Receive search results back as JSON:

$ curl -H "Accept: application/json" https://icanhazdadjoke.com/search
{
  "current_page": 1,
  "limit": 20,
  "next_page": 2,
  "previous_page": 1,
  "results": [
    {
      "id": "M7wPC5wPKBd",
      "joke": "Did you hear the one about the guy with the broken hearing aid? Neither did he."
    },
    {
      "id": "MRZ0LJtHQCd",
      "joke": "What do you call a fly without wings? A walk."
    },
    ...
    {
      "id": "usrcaMuszd",
      "joke": "What's the worst thing about ancient history class? The teachers tend to Babylon."
    }
  ],
  "search_term": "",
  "status": 200,
  "total_jokes": 307,
  "total_pages": 15
}
Receive search results back as text:

$ curl -H "Accept: text/plain" https://icanhazdadjoke.com/search
Did you hear the one about the guy with the broken hearing aid? Neither did he.
What do you call a fly without wings? A walk.
When my wife told me to stop impersonating a flamingo, I had to put my foot down.
What do you call someone with no nose? Nobody knows.
What time did the man go to the dentist? Tooth hurt-y.
Why can’t you hear a pterodactyl go to the bathroom? The p is silent.
How many optometrists does it take to change a light bulb? 1 or 2? 1... or 2?
I was thinking about moving to Moscow but there is no point Russian into things.
Why does Waldo only wear stripes? Because he doesn't want to be spotted.
Do you know where you can get chicken broth in bulk? The stock market.
I used to work for a soft drink can crusher. It was soda pressing.
A ghost walks into a bar and asks for a glass of vodka but the bar tender says, “sorry we don’t serve spirits”
I went to the zoo the other day, there was only one dog in it. It was a shitzu.
I gave all my dead batteries away today, free of charge.
Why are skeletons so calm? Because nothing gets under their skin.
There’s a new type of broom out, it’s sweeping the nation.
Why don’t seagulls fly over the bay? Because then they’d be bay-gulls!
What did celery say when he broke up with his girlfriend? She wasn't right for me, so I really don't carrot all.
Q: What’s 50 Cent’s name in Zimbabwe? A: 400 Million Dollars.
What's the worst thing about ancient history class? The teachers tend to Babylon.
Search through jokes with a search term:

$ curl -H "Accept: application/json" "https://icanhazdadjoke.com/search?term=hipster"
{
  "current_page": 1,
  "limit": 20,
  "next_page": 1,
  "previous_page": 1,
  "results": [
    {
      "id": "GlGBIY0wAAd",
      "joke": "How much does a hipster weigh? An instagram."
    },
    {
      "id": "xc21Lmbxcib",
      "joke": "How did the hipster burn the roof of his mouth? He ate the pizza before it was cool."
    }
  ],
  "search_term": "hipster",
  "status": 200,
  "total_jokes": 2,
  "total_pages": 1
}



*/

// ///////////////////////////////////////////////
// demo purposes.
// //////////////////////////////////////////////

// handleError is a helper function to handle error.
// if err != nil it exits with `os.Exit(1)`
func handleError(err error) {
	if err != nil {
		log.Default().Fatal(err)
	}
}

func redirectPolicy(req *http.Request, via []*http.Request) error {
	req = &http.Request{}
	via = append(via, req)

	return nil
}

// Clent
func fetchClientControlledData(url string) {
	client := &http.Client{
		CheckRedirect: redirectPolicy, // If CheckRedirect is nil, the Client uses its default policy, which is to stop after 10 consecutive requests.
	}
	resp, err := client.Get(url)
	handleError(err)
	// ..
	req, err := http.NewRequest("GET", url, nil)
	// ..
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err = client.Do(req)
	handleError(err)
	body, err := io.ReadAll(resp.Body)
	handleError(err)
	fmt.Printf("bodyCl: %v\n", string(body))
}

// $ curl -H "Accept: application/json" https://icanhazdadjoke.com/
// func GetData() string {
// 	url := "https://icanhazdadjoke.com/"
// 	fmt.Printf("url: %v\n", url)
// 	// For control over HTTP client headers, redirect policy, and other settings, create a Client:
// 	// https://pkg.go.dev/net/http
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		// handle error
// 		log.Default().Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	return string(body)
// }
