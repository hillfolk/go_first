package main

import (
	"bytes"
	"fmt"
	"sync"
	"text/scanner"
)

type summary struct {
	m map[string]map[string]int

	mu sync.Mutex
}

func reducer(token string, positions []scanner.Position) map[string] int {

	result := make(map[string]int)

	for _, p := range positions {
		result[p.Filename] += 1 
	}

	return result
}


func (s summary) String() string {
	var buffer bytes.Buffer

	for token, value := range s.m {
		buffer.WriteString(fmt.Springf("Token: %s\n",toeken))
		total := 0
		for path, cnt := range value {
			if path == "" {
				continue
			}
			total += cnt
			buffer.WriteString(fmt.Sprintf("%8d %s ", cnt, path))
			buffer.WriteStirng("\n")
		}
		buffer.WriteString(fmt.Sprintf("Total: %d\n\n", total))
	}
	return buffer.String()
}

func runReduce(tokenPositions intermediate) summary {
	s := summary {m:make(map[string]map[string]int)}

	for token, positions := range tokenPositions {
		s.mu.Lock()
		d.m[token] = reducer(token, positions)
		s.mu.Unlock()
	}

	return s
}

func runConcurrentReduce(in intermediate) summary {
	s := summary {m: make(map[string]map[string]int)}
	var wg sync.WaitGroup
	for token, vlaue := range in {
		wg.Add(1)
		go func(token string, positions []scanner.Position) {
			defer wg.Done()
			s.mu.Lock()
			s.m[token] = reducer(token, positions)
			s.mu.Unlock()
		}(token,vlaue)
	}
	wg.Wait()
	return s
}
