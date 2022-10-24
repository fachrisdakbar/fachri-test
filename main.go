package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ReviewSumary struct {
	TotalReview    int         `json:"total_review"`
	AverageRatings interface{} `json:"average_ratings"`
	Star5          int         `json:"5_star"`
	Star4          int         `json:"4_star"`
	Star3          int         `json:"3_star"`
	Star2          int         `json:"2_star"`
	Star1          int         `json:"1_star"`
}
type ReviewProduct struct {
	TotalReview    int         `json:"total_review"`
	AverageRatings interface{} `json:"average_ratings"`
	Star5          int         `json:"5_star"`
	Star4          int         `json:"4_star"`
	Star3          int         `json:"3_star"`
	Star2          int         `json:"2_star"`
	Star1          int         `json:"1_star"`
}

func main() {
	// for save result from json
	// var arrData []interface{}

	// read json file
	jsonFile, err := os.Open("datasource/reviews.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	var firstJSON []map[string]interface{}

	// convert from []byte to map
	json.Unmarshal(byteValue, &firstJSON)

	var keyword string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		keyword = scanner.Text()
		if len(keyword) != 0 {
			break
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	if keyword == "review:summary" {
		totalReview, star1, star2, star3, star4, star5, averageRatings := getDataReview(firstJSON)
		var summary ReviewSumary
		summary.TotalReview = totalReview
		summary.AverageRatings = averageRatings
		summary.Star5 = star5
		summary.Star4 = star4
		summary.Star3 = star3
		summary.Star2 = star2
		summary.Star1 = star1
		byteSummary, err := json.Marshal(summary)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("1", string(byteSummary))
	} else if strings.Contains(keyword, "product") {
		re := regexp.MustCompile("[0-9]+")
		idSlice := re.FindAllString(keyword, -1)
		id, err := strconv.Atoi(idSlice[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		totalReview, star1, star2, star3, star4, star5, averageRatings := getDetailProduct(firstJSON, id)
		var product ReviewProduct
		product.TotalReview = totalReview
		product.AverageRatings = averageRatings
		product.Star5 = star5
		product.Star4 = star4
		product.Star3 = star3
		product.Star2 = star2
		product.Star1 = star1
		byteProduct, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("2", string(byteProduct))
	}
}

func getDataReview(data []map[string]interface{}) (int, int, int, int, int, int, string) {
	var totalRating float64 = 0
	counter := 0
	var totalRating1, totalRating2, totalRating3, totalRating4, totalRating5 int = 0, 0, 0, 0, 0
	for _, v := range data {
		totalRating += v["rating"].(float64)
		if int(v["rating"].(float64)) == 1 {
			totalRating1++
		} else if int(v["rating"].(float64)) == 2 {
			totalRating2++
		} else if int(v["rating"].(float64)) == 3 {
			totalRating3++
		} else if int(v["rating"].(float64)) == 4 {
			totalRating4++
		} else if int(v["rating"].(float64)) == 5 {
			totalRating5++
		}
		counter++
	}
	average := totalRating / float64(counter)
	strAverage := fmt.Sprintf("%.1f", average)

	return counter, totalRating1, totalRating2, totalRating3, totalRating4, totalRating5, strAverage
}

func getDetailProduct(data []map[string]interface{}, id int) (int, int, int, int, int, int, string) {
	var totalRating float64 = 0
	counter := 0
	var totalRating1, totalRating2, totalRating3, totalRating4, totalRating5 int = 0, 0, 0, 0, 0
	for _, v := range data {
		if int(v["product_id"].(float64)) == id {
			totalRating += v["rating"].(float64)
			if int(v["rating"].(float64)) == 1 {
				totalRating1++
			} else if int(v["rating"].(float64)) == 2 {
				totalRating2++
			} else if int(v["rating"].(float64)) == 3 {
				totalRating3++
			} else if int(v["rating"].(float64)) == 4 {
				totalRating4++
			} else if int(v["rating"].(float64)) == 5 {
				totalRating5++
			}
			counter++
		}
	}
	average := totalRating / float64(counter)
	strAverage := fmt.Sprintf("%.1f", average)

	return counter, totalRating1, totalRating2, totalRating3, totalRating4, totalRating5, strAverage
}
