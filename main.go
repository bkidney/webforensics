// Copyright © 2016 Brian Kidney <bkidney@briankidney.ca>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"
	"net/http"
	"os"
)

type Status int

const (
	new = iota
	good
	bad
)

func main() {

	if len(os.Args) < 2 {
		// print usage
		os.Exit(1)
	}

	root := os.Args[1]

	// Get URL from Args
	//	Arg validation
	pages := make(map[string]Status)
	pages[root] = new

	toProcess := make([]string, 0)
	toProcess = append(toProcess, root)

	// For each URL
	for len(toProcess) > 0 {
		url := toProcess[len(toProcess)-1]
		toProcess = toProcess[:len(toProcess)-1]

		//	Fetch page
		resp, err := http.Get(url)

		//	If fetch error
		if err != nil {
			//		Catalog Error
			pages[url] = bad
			fmt.Println(err)
			//		Return
			continue
		}
		defer resp.Body.Close()
		pages[url] = good

		//	Extract New URLs

		//	Add URLS to stucture graph
		//	Add URLs to list of URLs to be processed
	}

	var nGood, nBad, nNew int
	for _, status := range pages {
		switch status {
		case good:
			nGood++
		case bad:
			nBad++
		case new:
			nNew++
		}
	}

	fmt.Printf("Done! Visited %d good and %d bad links. %d links were not visted \n", nGood, nBad, nNew)
}
