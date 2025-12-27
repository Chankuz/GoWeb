package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"webserver/util"
)

func FormHanlder(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	username := r.FormValue("user")

	fmt.Fprintf(w, "Hello %s\n", username)

	weight := r.FormValue("weight")
	height := r.FormValue("height")
	ticket := r.FormValue("ticket")

	// price, err := strconv.Atoi(ticketquantity)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(price)
	fmt.Fprintf(w, "weight: %s\n", weight)
	fmt.Fprintf(w, "height: %s\n", height)
	fmt.Fprintf(w, "ticket: %s\n", ticket)

	weightFloat, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid weight: %v", err)
		return
	}
	heightFloat, err := strconv.ParseFloat(height, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid height: %v", err)
		return
	}
	bmi := util.BMICalculator(weightFloat, heightFloat)

	fmt.Fprintf(w, "BMI: %f\n", bmi)
	// fmt.Fprintf(w, "Ticket Price: %s\n", ticketprice)
	// fmt.Fprintf(w, "Ticket Quantity: %s\n", ticketquantity)

}
