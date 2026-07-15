package main

import (
	"fmt"
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"go-fiber-test/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	if err := database.DBConn.AutoMigrate(&m.Dogs{}, &m.Company{}); err != nil {
		panic("AutoMigrate failed: " + err.Error())
	}
	fmt.Println("Migration completed!")
}

func No0() {
	i := 2
	fmt.Println("Example-: If Else Condition ")
	if i == 0 {
		fmt.Println("zero")
	} else if i == 1 {
		fmt.Println("one")
	} else if i == 2 {
		fmt.Println("two")
	} else if i == 3 {
		fmt.Println("three")
	} else {
		fmt.Println("your i is not in if case")
	}
}

func No1() {
	for i := 1; i <= 100; i++ {
		if i%3 != 0 {
			fmt.Printf("i = %d, หาร 3 ไม่ลงตัว\n", i)
		} else {
			fmt.Printf("i = %d, หาร 3 ลงตัว\n", i)
		}
	}
}

func No2() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	min := x[0]
	max := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println("min =", min, "max =", max)
}

func No3(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		for _, c := range s {
			if c == '9' {
				count++
			}
		}
	}
	return count
}

func No4(s string) string {
	result := ""
	for _, c := range s {
		if c != ' ' {
			result += string(c)
		}
	}
	return result
}

func No5() {
	data := []map[string]string{
		{
			"name":    "Ms. Marry Jane (Age 22)",
			"address": "Bangkok",
		},
		{
			"name":    "Mr. John Doe (Age 25)",
			"address": "Chiang Mai",
		},
		{
			"name":    "Ms. Jane Doe (Age 24)",
			"address": "Phuket",
		},
	}
	for i, v := range data {
		fmt.Printf("%d. name: %s, Address: %s\n", i+1, v["name"], v["Address"])
	}
}

type Company struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

func No6() {
	companies := []Company{
		{
			Name:    "ABC Company",
			Address: "Bangkok",
			Tel:     "02-1234567",
		},
		{
			Name:    "XYZ Company",
			Address: "Chiang Mai",
			Tel:     "02-7654321",
		},
		{
			Name:    "123 Company",
			Address: "Phuket",
			Tel:     "02-4567890",
		},
	}
	for i, v := range companies {
		fmt.Printf("%d. name: %s, Address: %s, Tel: %s\n", i+1, v.Name, v.Address, v.Tel)
	}
}

func Power(num int, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= num
	}
	return result
}

func NoExtra() {
	for i := 0; i <= 6; i++ {
		stars := ""
		for j := 0; j < i; j++ {
			stars += "*"
		}
		fmt.Println(stars)
	}
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.InetRoutes(app)
	// No0()
	// No1()
	// Power(20, 2)
	No2()
	fmt.Println("จำนวนเลข 9 (1-1000):", No3(1000))
	fmt.Println(No4("AW SOME GO!"))
	No5()
	No6()
	NoExtra()
	app.Listen(":3000")
}
