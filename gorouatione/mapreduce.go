package main

import "text/scanner"
import "sync"
import "fmt"
import "os"

type partial struct {
	token string
	scanner.Position
}


func mapper(path string, out chan<- partial) {
	file, err := os.Open(path)
	if err != nil {
		return 
	}

	defer file.Close()

	//정상 파일 아닌 경우 바로 반환

	if info, err := file.Stat(); err != nil || info.Model().IsDir() {
		return
	}

	var s scanner.Scanner
	s.Filename = path
	s.Init(file)

	// 파일의 모든 토큰을 스캔하여 out 채널로 전송
	tok := s.Scan()
	for tok != scanner.EOF {
		fmt.Println(s.Pos())
		out <- partial{s.TokenText(),s.Pos()}
		tok  = s.Scan()
	}
	
}

func numMap(paths <-chan string) <-chan partial {
	out := make(chan partial, BUF_SIZE)

	go func() {
		for path := range pahts {
			mapper(path, out)
		}
		close(out)
	}()
	return out
}

type intermediate map[string][] scanner.Position

func (m intermediate) addPartial(p partial) {
	postions, ok := m[p.token]
	if !ok {
		postions = make([]scanner.Position,1)
		
	}

	positions = append(positions, p.Position)
	m[p.token] = positions
}


func reducer(token string, positions []scanner.Position) map[string]int {
	result := make(map[string]int)

	for _,p := range positions {
		result[p.Filename] +1
		
	}
	return result
}

type summary struct {
	//키: token
	//값: map[string]int
	//          키:file path
	//          값:token count

	m map[string]map[string]int

	mu sync.Mutex
}

func (s summary) String() string {
	var buffer bytes.Buffer

	for token, value := range s.m {
		buffer.WirteString(fmt.Sprintf("Token: %s\n", token))
		total := 0

		for path,cnt := range value {
			if path == "" {
				continue
			}

			total += cnt
			buffer.WriteString(fmt.Sprintf("%8d %s ", cnt, path))
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("Total: %d\n\n", totla))
	}
	return buffer.String()
}

func runReduce(toeknPositions intermediate) summary {
	s := summary {m: make(map[string]map[string]int)}
	for token, positions := range tokenPositions {
		s.mu.Lock()
		s.m[token] = reducer(tokenm,positions)
		s.mu.Unlock()
	}
	return s
}

func main() {
	paths := find(parseargs())
	fmt.Println(runReduce(collect(runMap(paths))))
}


func runConcurrentMap(paths <-chan string) <-chan partial {
	out := make(chan partial, BUF_SIZE)

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				mapper(path,out)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func runConcurrentReduce(in intermediate) summary {
	result := make(summary)

	var wg sync.WaitGroup
	for token, value := range in {
		wg.Add(1)
		go func(token string, positions []scanner.Position) {
			defer wg.Done()
			result[token] = reducer(token, positions)
		}(token,value)		
	}
	wg.Wait()
	return result
}

func main() {
	paths := find(parseArgs())
	fmt.Println(runConcurrentReduce(collect(runConcurrentMap(paths))))
}


