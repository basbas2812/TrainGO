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
	param := c.Params("number")
	if param == "" {
		return c.Status(400).SendString("number is required.")
	}
	n, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(400).SendString("invalid number.")
	}
	if n < 0 {
		return c.Status(400).SendString("number must not be negative.")
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
		return c.Status(400).SendString("ข้อมูลไม่ถูกต้อง")
	}

	validate := validator.New()
	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[A-Za-z0-9_-]{3,20}$`, fl.Field().String())
		return matched
	})
	validate.RegisterValidation("lineid", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[A-Za-z0-9_-]{4,20}$`, fl.Field().String())
		return matched
	})
	validate.RegisterValidation("business_allowed", func(fl validator.FieldLevel) bool {
		allowed := map[string]bool{
			"retail":  true,
			"service": true,
			"it":      true,
			"finance": true,
			"other":   true,
		}
		return allowed[fl.Field().String()]
	})
	validate.RegisterValidation("websitename", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		allowedSuffixes := []string{
			".sogoodweb.com",
			".sogoodweb.co.th",
			".sogoodweb.in.th",
		}
		for _, suffix := range allowedSuffixes {
			if strings.HasSuffix(value, suffix) {
				prefix := strings.TrimSuffix(value, suffix)
				matched, _ := regexp.MatchString(`^[a-z0-9\-]{2,30}$`, prefix)
				return matched
			}
		}
		return false
	})

	if err := validate.Struct(reg); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ใส่ข้อมูลผิดพลาด",
			"errors":  err.Error(),
		})
	}

	return c.Status(201).SendString("สมัครสมาชิกสำเร็จ")
}

func AddCompany(c *fiber.Ctx) error {
	company := new(m.Company)
	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).SendString("ข้อมูลไม่ถูกต้อง")
	}
	if err := database.DBConn.Create(&company).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(201).SendString("เพิ่มบริษัทสำเร็จ")
}

func GetAllCompanies(c *fiber.Ctx) error {
	companies := []m.Company{}
	if err := database.DBConn.Find(&companies).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(200).JSON(companies)
}

func GetIDCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	company := m.Company{}
	if err := database.DBConn.First(&company, id).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(200).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	company := m.Company{}
	if err := database.DBConn.First(&company, id).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).SendString("ข้อมูลไม่ถูกต้อง")
	}
	if err := database.DBConn.Save(&company).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(200).SendString("อัปเดตบริษัทสำเร็จ")
}

func DeleteCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	company := m.Company{}
	if err := database.DBConn.First(&company, id).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	falseVal := false
	database.DBConn.Model(&company).Update("is_active", &falseVal)
	if err := database.DBConn.Delete(&company).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(200).SendString("ลบบริษัทสำเร็จ")
}

func GetDeletedDogs(c *fiber.Ctx) error {
	dogs := []m.Dogs{}
	if err := database.DBConn.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(200).JSON(dogs)
}

func GetDogIDFiftytoHundred(c *fiber.Ctx) error {
	dogs := []m.Dogs{}
	if err := database.DBConn.Find(&dogs, "dog_id BETWEEN 50 AND 100").Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.Status(200).JSON(dogs)
}

func GetDogIDSumColor(c *fiber.Ctx) error {
	dogs := []m.Dogs{}
	if err := database.DBConn.Find(&dogs).Error; err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}

	var dbName string
	database.DBConn.Raw("SELECT DATABASE()").Scan(&dbName)

	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNocolor := 0

	var data []m.DogColor
	for _, dog := range dogs {
		color := "no color"
		if dog.DogID >= 10 && dog.DogID <= 50 {
			color = "red"
			sumRed++
		} else if dog.DogID >= 100 && dog.DogID <= 150 {
			color = "green"
			sumGreen++
		} else if dog.DogID >= 200 && dog.DogID <= 250 {
			color = "pink"
			sumPink++
		} else {
			sumNocolor++
		}
		data = append(data, m.DogColor{Name: dog.Name, DogID: dog.DogID, Color: color})
	}

	return c.Status(200).JSON(fiber.Map{
		"database":    dbName,
		"count":       len(dogs),
		"data":        data,
		"sum_red":     sumRed,
		"sum_green":   sumGreen,
		"sum_pink":    sumPink,
		"sum_nocolor": sumNocolor,
	})

}
