package echo

import (
	"fmt"
	"os"
	"strings"
)

// echo2 неоптимально, т.к. каждый раз наращивает строчку
func echo2() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// echo4 выводит так же имя выполняемой команды
func echo4() {
	fmt.Println(strings.Join(os.Args, " "))
}

// echo5 выводит индекс и значение каждого аргумента по одному аргументу в каждой строке
func echo5() {
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
}

// vs echo3 and echo4
func echo6() {
	var strBuilder strings.Builder
	for _, v := range os.Args {
		strBuilder.WriteString(v)
	}
	fmt.Println(strBuilder.String())
}
