package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	keyword := "DEVSISTERS"
	relativePath := "./nomad-lectures/chapter2"

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
	// 전체 파일/디렉토리 순회
	// 절대경로에 존재하는 모든 파일, 디렉토리 찾기
	for _, file := range getFiles(absPath) {
		// 디렉토리일 경우
		if file.IsDir() {
			join := filepath.Join(absPath, file.Name())
			recursiveSearch(join, keyword)
		}

		// 파일일 경우
		fullFilePath := filepath.Join(absPath, file.Name())
		data, err := os.Open(fullFilePath)
		checkErr(err)

		// 파일을 스캔
		scanner := bufio.NewScanner(data)
		currentLine := 0
		for scanner.Scan() {
			currentLine++
			line := scanner.Text()
			if searchText(keyword, line) {
				fmt.Println(currentLine, "|", file.Name())
				break
			}
		}

		closeFileExplicitly(data, fullFilePath)
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

func mai2n() {

	// 파일을 읽으려면 꼭 절대경로를 다붙여서 넣어야하나?
	path, err := filepath.Abs("sample.text")
	println(path)

	open, err := os.Open("./sample.text")
	if err != nil {
		fmt.Println("pen", open, err)
	}
	fmt.Println("드디어 ㅅㅂ", open)
	// "dgrep {keyword} {relative path}"
	relativePath := "nomad-lectures/chapter2" //"devsisters" //../defer
	keyword := "keyword"
	fmt.Println("relative Path: ", relativePath)
	fmt.Println("keyword: ", keyword)

	dir := filepath.Dir("../nomad-lectures/chapter2/banking")
	fmt.Println("dir", dir)
	absolutePath := strings.Split(dir, "../")
	if readDir, err := os.ReadDir(absolutePath[1]); err != nil {
		fmt.Println("err: ", err)
	} else {
		for i := 0; i < len(readDir); i++ {
			entry := readDir[i]
			if !entry.IsDir() {
				fmt.Println("files:", entry.Name())
				abs, _ := filepath.Abs(entry.Name())
				fmt.Println("abs", abs)
				if file, err := os.Open(abs + "/" + entry.Name()); err == nil {
					fmt.Println("open file", abs, file)
				} else {
					fmt.Println("can't open file", err)
				}
			}
		}
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
