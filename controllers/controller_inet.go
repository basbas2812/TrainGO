package controllers

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello World!!!")
}

func InputName(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name)
	log.Println(p.Pass)
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamName(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryName(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is " + a
	return c.JSON(str)
}

func ValidateTest(c *fiber.Ctx) error {

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet
			DogID: v.DogID, //112
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	type ResultData struct {
		Data  []m.DogsRes `json:"data"`
		Name  string      `json:"name"`
		Count int         `json:"count"`
	}
	r := ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

func Factorial(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("number"))
	if err != nil {
		return c.Status(400).SendString("invalid number")
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return c.Status(200).SendString(strconv.Itoa(fact))
}

func Ascii(c *fiber.Ctx) error {
	taxID := c.Query("tax_id")
	if taxID == "" {
		return c.Status(400).SendString("tax_id is required")
	}
	var result []string
	for _, char := range taxID {
		result = append(result, strconv.Itoa(int(char)))
	}
	return c.Status(200).SendString(strings.Join(result, " "))
}

func Register(c *fiber.Ctx) error {
	reg := new(m.Register)
	if err := c.BodyParser(&reg); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ข้อมูลไม่ถูกต้อง"})
	}

	patterns := map[string]string{
		"email":        `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
		"username":     `^[a-zA-Z0-9_-]+$`,
		"password":     `^.{6,20}$`,
		"lineid":       `^[a-zA-Z0-9_-]+$`,
		"phoneid":      `^[0-9]+$`,
		"businesstype": `^[a-zA-Z0-9_-]+$`,
		"websitename":  `^[a-z0-9\-]{2,30}$`,
	}

	values := map[string]string{
		"email":        reg.Email,
		"username":     reg.Username,
		"password":     reg.Password,
		"lineid":       reg.LineID,
		"phoneid":      reg.PhoneID,
		"businesstype": reg.BusinessType,
		"websitename":  reg.WebsiteName,
	}

	for field, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, values[field])
		if !matched {
			return c.Status(400).JSON(fiber.Map{
				"message": "ใส่ข้อมูลผิดพลาด",
				"field":   field,
			})
		}
	}

	return c.Status(201).JSON(fiber.Map{"message": "สมัครสมาชิกสำเร็จ"})
}
