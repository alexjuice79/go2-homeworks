package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var num1 int
var num2 int

func GenerateNumFir(w http.ResponseWriter, r *http.Request) {
	num1 = rand.Intn(100)
	fmt.Fprintf(w, strconv.Itoa(num1))
}

func GenerateNumSec(w http.ResponseWriter, r *http.Request) {
	num2 = rand.Intn(100)
	fmt.Fprintf(w, strconv.Itoa(num2))
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Available commands:</h1><h2>\"<a href=\"./info\">/info</a>\" - выводит данное сообщение<br>\"<a href=\"./first\">/first</a>\" - выводит случайное число<br>\"<a href=\"./second\">/second</a>\" - выводит случайное число<br>\"<a href=\"./add\">/add</a>\" - выводит сумму двух случайных чисел<br>\"<a href=\"./sub\">/sub</a>\" - выводит разность двух случайных чисел<br>\"<a href=\"./mul\">/mul</a>\" - выводит произведение двух случайных чисел<br>\"<a href=\"./div\">/div</a>\" - выводит деление двух случайных чисел</h2>")
}

func GetSumRandNums(w http.ResponseWriter, r *http.Request) {
	output := strconv.Itoa(num1) + "+" + strconv.Itoa(num2) + "=" + strconv.Itoa(num1+num2)

	fmt.Fprintf(w, output)
}

func GetSubRandNums(w http.ResponseWriter, r *http.Request) {
	output := strconv.Itoa(num1) + "-" + strconv.Itoa(num2) + "=" + strconv.Itoa(num1-num2)

	fmt.Fprintf(w, output)
}

func GetMulRandNums(w http.ResponseWriter, r *http.Request) {
	output := strconv.Itoa(num1) + "*" + strconv.Itoa(num2) + "=" + strconv.Itoa(num1*num2)

	fmt.Fprintf(w, output)
}

func GetDivRandNums(w http.ResponseWriter, r *http.Request) {
	var output string

	if num2 == 0 {
		output = strconv.Itoa(num1) + "/" + strconv.Itoa(num2) + "~ На ноль делить нельзя!"
	} else {
		output = strconv.Itoa(num1) + "/" + strconv.Itoa(num2) + "~" + strconv.Itoa(num1/num2)
	}

	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/info", GetInfo)
	http.HandleFunc("/first", GenerateNumFir)
	http.HandleFunc("/second", GenerateNumSec)
	http.HandleFunc("/add", GetSumRandNums)
	http.HandleFunc("/sub", GetSubRandNums)
	http.HandleFunc("/mul", GetMulRandNums)
	http.HandleFunc("/div", GetDivRandNums)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
