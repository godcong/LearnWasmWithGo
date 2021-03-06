package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	var numbers string
	calc := func(i []js.Value) {
		//js.Global.Set("output", numbers)

		fmt.Println("calculator:", calc(numbers))
	}

	number := func(i []js.Value) {
		numbers += strconv.Itoa(i[0].Int())
		fmt.Println(numbers)
	}

	symbol := func(i []js.Value) {
		if i[0].String() == "clr" {
			numbers = ""
			return
		}
		numbers += i[0].String()
		fmt.Println(numbers)
	}
	js.Global.Set("symbol", js.NewCallback(symbol))
	js.Global.Set("calc", js.NewCallback(calc))
	js.Global.Set("number", js.NewCallback(number))
	<-c
}

func calc(numbers string) float64 {
	var value []string
	if strings.Index(numbers, "+") >= 0 {
		value = strings.Split(numbers, "+")
		if len(value) >= 2 {
			rlt := 0.0
			for i, v := range value {
				if canSplit(v) {
					if i == 0 {
						rlt = calc(v)
					} else {
						rlt += calc(v)
					}
				} else {
					vv, _ := strconv.Atoi(v)
					if i == 0 {
						rlt = float64(vv)
					} else {
						rlt += float64(vv)
					}
				}
			}
			return rlt
		}

	} else if strings.Index(numbers, "-") >= 0 {
		value = strings.Split(numbers, "-")
		if len(value) >= 2 {
			rlt := 0.0
			for i, v := range value {
				if canSplit(v) {
					if i == 0 {
						rlt = calc(v)
					} else {
						rlt -= calc(v)
					}
				} else {
					vv, _ := strconv.Atoi(v)
					if i == 0 {
						rlt = float64(vv)
					} else {
						rlt -= float64(vv)
					}
				}
			}
			return rlt
		}
	} else if strings.Index(numbers, "*") >= 0 {
		value = strings.Split(numbers, "*")
		if len(value) >= 2 {
			rlt := 0.0
			for i, v := range value {
				if canSplit(v) {
					if i == 0 {
						rlt = calc(v)
					} else {
						rlt *= calc(v)
					}
				} else {
					vv, _ := strconv.Atoi(v)
					if i == 0 {
						rlt = float64(vv)
					} else {
						rlt *= float64(vv)
					}
				}
			}
			return rlt
		}
	} else if strings.Index(numbers, "/") >= 0 {
		value = strings.Split(numbers, "/")
		if len(value) >= 2 {
			rlt := 0.0
			for i, v := range value {
				if canSplit(v) {
					if i == 0 {
						rlt = calc(v)
					} else {
						rlt /= calc(v)
					}
				} else {
					vv, _ := strconv.Atoi(v)
					if i == 0 {
						rlt = float64(vv)
					} else {
						rlt /= float64(vv)
					}
				}
			}
			return rlt
		}

	}
	if len(value) == 1 {
		left, _ := strconv.Atoi(value[0])
		return float64(left)
	}
	return 0

}

func canSplit(v string) bool {
	if strings.Index(v, "+") >= 0 {
		return true
	} else if strings.Index(v, "-") >= 0 {
		return true
	} else if strings.Index(v, "*") >= 0 {
		return true
	} else if strings.Index(v, "/") >= 0 {
		return true
	}
	return false
}
