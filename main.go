package main

import (
	tickets "desafio-go-bases/internal/tickets"
	"fmt"
)

func main() {
	// Calcular a quantidade e a média de pessoas que viajam para um país determinado
	destination := "Brazil"
	count, average, err := tickets.GetTotalTickets(destination)
	//tratar o erro se houver
	if err != nil {
		fmt.Printf("Error calculating the number and average of people traveling to the destination %s: %v\n", destination, err)
	} else {
		fmt.Printf("Number of people traveling to the destination %s: %d, and average: %.2f\n", destination, count, average)
	}

	// Calcular a quantidade de pessoas que viajam em cada período do dia
	time := "morning"
	countTime, err := tickets.GetCountByPeriod(time)
	//tratar o erro se houver
	if err != nil {
		fmt.Printf("Error when calculating the number of people traveling in the period %s: %v\n", time, err)
	} else {
		fmt.Printf("Number of people traveling during the period %s: %d\n", time, countTime)
	}
}
