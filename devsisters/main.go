package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s {keyword} {relative path}", os.Args[0])
	}

	// 명령줄 인자에서 키워드와 상대경로를 받아옴
	keyword := os.Args[1]
	relativePath := os.Args[2]

	// 절대경로 찾기
	absPath := findAbsPath(relativePath)
	recursiveSearch(absPath, keyword)
}

func searchText(keyword string, text string) bool {
	return strings.Contains(text, keyword)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error : %v", err)
	}
}

func recursiveSearch(absPath string, keyword string) {
	var wg sync.WaitGroup
	// 전체 파일/디렉토리 순회
	// 절대경로에 존재하는 모든 파일, 디렉토리 찾기
	for _, file := range getFiles(absPath) {
		fullFilePath := filepath.Join(absPath, file.Name())

		// 디렉토리일 경우
		if file.IsDir() {
			join := filepath.Join(absPath, file.Name())
			wg.Add(1)
			go func(path string) {
				defer wg.Done()
				recursiveSearch(path, keyword)
			}(join)
		} else {
			wg.Add(1)
			go func(filePath string) {
				defer wg.Done()
				processFile(filePath, keyword)
			}(fullFilePath)
		}
	}
	wg.Wait()
}

func processFile(filePath string, keyword string) {
	data, err := os.Open(filePath)
	checkErr(err)
	defer closeFileExplicitly(data, filePath)

	// 파일을 스캔
	scanner := bufio.NewScanner(data)
	currentLine := 0
	for scanner.Scan() {
		currentLine++
		line := scanner.Text()
		if searchText(keyword, line) {
			fmt.Println(currentLine, "|", filepath.Base(filePath))
			break
		}
	}

}

func findAbsPath(relativePath string) string {
	absPath, err := filepath.Abs(relativePath)
	checkErr(err)
	return absPath
}

func getFiles(absPath string) []os.DirEntry {
	files, err := os.ReadDir(absPath)
	checkErr(err)
	return files
}

func closeFileExplicitly(file *os.File, fullFilePath string) {
	if err := file.Close(); err != nil {
		log.Printf("Error closing file %s: %v", fullFilePath, err)
	}
}

/*
1. 상대경로 입력
2. 해당 경로 하위의 요소 파악
	- directory
	- file
3. directory 라면 또 재귀식으로 돌려야함.
4. 파일이라면 문서 탐색 => 함수 만들자.
	- 문서 탐색 시 찾으려는 키워드가 몇번째 라인에 있는건 어떻게 알지?
Q. fmt 와 log 패키지의 차이?
*/
