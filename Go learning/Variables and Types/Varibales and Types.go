package main

import "fmt"

func main() {
	//first Exercise
	//fmt.Println(2235 * 1231)

	//Second exercise
	const earthGravity = 9.80665
	fmt.Println(earthGravity)

	//Exercise three
	var numOfFlavors int = 57
	var flavorScale float32 = 5.8
	fmt.Println(numOfFlavors, flavorScale)

	//Exercise four
	var somehthing string = "something"
	var s string = "another thing"
	fmt.Println(somehthing + s)

	//Exercise five
	x := 23
	y := (x*9)/5 + 354
	fmt.Println(x, " degree Cesius is ", y, " degree Farenheit")

	//Exercise six
	r, w := "something", 10
	fmt.Println(r, w)

	//Exercise seven
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "writeen by", writer, "drawn by", artist, "published by", publisher, "published in", year, "has", pageNumber, "of pages and a score of", grade)
}
