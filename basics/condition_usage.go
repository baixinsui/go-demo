package basics

import "fmt"

func SwitchTest() {
	/* 定义局部变量 */
	var grade string
	var marks int
	fmt.Println("Enter your mark:")
	_, err := fmt.Scanln(&marks)
	if err != nil {
		println(err)
		fmt.Printf("Input value of mark:%v is invalid!", &age)
		return
	}

	switch marks / 10 {
	case 9, 10:
		grade = "A"
	case 8:
		grade = "B"
	case 6, 7:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

func IfElseTest() {
	var youngAge = 18
	var olderAge = 60
	var yourAge int
	fmt.Println("Enter your age:")
	_, err := fmt.Scanln(&yourAge)
	if err != nil {
		println(err)
		fmt.Printf("Input value of age:%v is invalid!", &age)
		return
	}
	if yourAge < youngAge {
		fmt.Println("You are minor")
	} else if yourAge >= youngAge {
		fmt.Println("You are adult")
		if yourAge >= olderAge {
			fmt.Println("You are old people")
		}
	}

}
