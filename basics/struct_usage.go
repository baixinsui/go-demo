package basics

import "fmt"

//StructUsage
/*
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
type struct_variable_type struct {
   member1 definition
   member2 definition
   ...
   memberN definition
}
*/

type Student struct {
	name          string
	studentNumber string
	age           int8
	sex           Sex
	className     string
	books         []Books
}

type Sex struct {
	SexStr string
}

const (
	Male    = "Male"
	Female  = "Female"
	Unknown = "Unknown"
)

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

// ToString
/*定义Student类型ToString方法*/
func (student Student) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d, StudentNumber: %s, Sex: %s, ClassName: %s, Books: %v", student.name, student.age, student.studentNumber, student.sex.SexStr, student.studentNumber, student.books)
}

func StructUsage() {

	var book1 Books /* 声明 book1 为 Books 类型 */
	/* book 1 描述 */
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	/* 打印 book1 信息 */
	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.bookId)

	/* 声明 book2 为 Books 类型 并初始化, 按照结构体内参数顺序传入全部值 */
	var book2 = Books{"Python 教程", "www.runoob.com", "Python 语言教程", 6495700}
	/* 打印 book2 信息 */
	fmt.Printf("Book 2 title : %s\n", book2.title)
	fmt.Printf("Book 2 author : %s\n", book2.author)
	fmt.Printf("Book 2 subject : %s\n", book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", book2.bookId)

	var books1 = []Books{book1, book2}

	var student1 = Student{"John", "050101", 10, Sex{Male}, "5年级1班", books1}
	fmt.Println(student1)
	var books2 = []Books{book2, book1}

	/* 声明 book2 为 Books 类型 并初始化, key：value传入部分值 */
	var student2 = Student{name: "Lee", age: 10, studentNumber: "050201", sex: Sex{Female}, books: books2}
	fmt.Println(student2.ToString())

	var student3 Student
	student3.name = "Yarn"
	student3.sex = Sex{Unknown}
	student3.age = 11
	student3.className = "6年级1班"
	student3.studentNumber = "060112"
	fmt.Println(student3.ToString())
}
