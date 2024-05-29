package main

import (
	"log"
	"regexp"
	"strconv"
	"github.com/LucasxTS/GoBackEndProject/src/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	app.Post("/verify", VerifyHandler) 
		
	
	log.Fatal(app.Listen(":8080"))
}

func VerifyHandler(c *fiber.Ctx) error {
	var scoreModel model.ScoreModel
		if err := c.BodyParser(&scoreModel); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		regex := regexp.MustCompile(`^(\d+)[xX](\d+)$`)
		match := regex.FindStringSubmatch(scoreModel.Score)
		if match == nil {
			return c.Status(fiber.StatusBadRequest).SendString("Regex Error")
		}

		
		score1, err := StringToInt(match[1])
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
			
		}

		score2, err := StringToInt(match[2])
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
			
		}
		
		total := calculateCombinations(score1) * calculateCombinations(score2)

		combinationModel := model.CombinationModel{Combination: total}
		return c.JSON(combinationModel)
}

func StringToInt(s string) (int, error)  {
	return strconv.Atoi(s)
}


func calculateCombinations(score int) int {
    if score < 0 {
        return 0
    }
    ways := make([]int, score+1)
    ways[0] = 1

    for i := 3; i <= score; i++ {
        ways[i] += ways[i-3]
    }
    for i := 6; i <= score; i++ {
        ways[i] += ways[i-6]
    }
    for i := 7; i <= score; i++ {
        ways[i] += ways[i-7]
    }
    for i := 8; i <= score; i++ {
        ways[i] += ways[i-8]
    }

    return ways[score]
}
