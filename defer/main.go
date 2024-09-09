package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//func main() {
//	log.Println("#1")
//	argumentEvaluation1()
//
//	log.Println("#2")
//	argumentEvaluation2()
//
//	log.Println("#3")
//	argumentEvaluation3()
//}
//
//func argumentEvaluation1() {
//	for i := 0; i < 3; i++ {
//		defer fmt.Printf("%d \n", i) // deferred: [0, 1, 2]
//	}
//}
//
//func argumentEvaluation2() {
//	for i := 0; i < 3; i++ {
//		defer func() {
//			fmt.Printf("%d \n", i) // deferred: [3, 3, 3]
//		}()
//	}
//}
//
//func argumentEvaluation3() {
//	for i := 0; i < 3; i++ {
//		defer func(i int) {
//			fmt.Printf("%d \n", i) // deferred: [0, 1, 2]
//		}(i)
//	}
//}

func main() {
	log.Println("\n#1 - 기본 defer")
	basicDefer()

	log.Println("\n#2 - 클로저를 이용한 defer")
	deferWithClosure()

	log.Println("\n#3 - 인자를 전달하는 defer")
	deferWithArgument()
}

func basicDefer() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%d \n", i)
	}
	fmt.Println("기본 defer 실행 완료")
}

func deferWithClosure() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Printf("%d \n", i)
		}()
	}
	fmt.Println("클로저를 이용한 defer 실행 완료")
}

func deferWithArgument() {
	for i := 0; i < 3; i++ {
		defer func(n int) {
			fmt.Printf("%d \n", n)
		}(i)
	}
	fmt.Println("인자를 전달하는 defer 실행 완료")
}
func handleInnerPanic() {
	defer fmt.Println("(4) reachable")
	fmt.Println("(1) reachable")
	defer func() {
		v := recover()
		fmt.Println("(3) recovered:", v)
	}()
	defer fmt.Println("(2) reachable")
	panic("panic here")
	defer fmt.Println("(5) unreachable")
	fmt.Println("unreachable")
}

func Contents(filename string) (string, error) {
	f, err := os.Open(filename) // 자원 할당
	if err != nil {
		return "", err
	}
	defer f.Close() // 함수 종료 시점에 자원 해제
	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f.Close() 호출
		}
	}
	return string(result), nil // f.Close() 호출
}
