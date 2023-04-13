package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// tipo que representa cada campo do arquivo
type Ticket struct {
	ID          string
	Name        string
	Email       string
	Destination string
	FlightTime  string
	Price       float64
}

func GetTotalTickets(destination string) (int, float64, error) {
	// ler o arquivo
	fileContent, err := os.ReadFile("tickets.csv")
	// verificar se houve erro na leitura
	if err != nil {
		return 0, 0, err
	}

	// converter o conteúdo do arquivo em uma lista de linhas
	fileLines := strings.Split(string(fileContent), "\n")

	// inicializar a contagem de bilhetes para o destino especificado
	totalTicketsDestination := 0

	// inicializar a contagem e a soma dos preços para o destino especificado
	countDestination := 0
	sumPriceDestination := 0.0

	// iterar sobre cada linha
	for _, line := range fileLines {
		// ignorar linhas vazias
		if len(line) == 0 {
			continue
		}

		// converter a linha em uma lista de campos
		fields := strings.Split(line, ",")

		// analisar os dados
		ticket := Ticket{
			ID:          fields[0],
			Name:        fields[1],
			Email:       fields[2],
			Destination: fields[3],
			FlightTime:  fields[4],
		}

		// verificar se existe preco a ser convertido
		price, err := strconv.ParseFloat(fields[5], 64)
		// tratar o erro acima
		if err != nil {
			return 0, 0, err
		}
		// atribuir o preco convertido ao tipo
		ticket.Price = price

		// incrementar a contagem se o destino for o desejado
		if ticket.Destination == destination {
			totalTicketsDestination++
			countDestination++
			sumPriceDestination += ticket.Price
		}
	}

	// verificar se houve bilhetes vendidos para o destino
	if countDestination == 0 {
		return 0, 0, fmt.Errorf("no tickets sold for the destination %s", destination)
	} else {
		// calcular a média dos preços
		average := sumPriceDestination / float64(countDestination)

		return totalTicketsDestination, average, nil
	}
}

func GetCountByPeriod(time string) (int, error) {
	// ler o arquivo
	fileContent, err := os.ReadFile("tickets.csv")
	// verificar se houve erro na leitura
	if err != nil {
		return 0, err
	}

	// converter o conteúdo do arquivo em uma lista de linhas
	fileLines := strings.Split(string(fileContent), "\n")

	// inicializar a contagem de bilhetes para o período de tempo especificado
	count := 0

	// iterar sobre cada linha
	for _, line := range fileLines {
		// ignorar linhas vazias
		if len(line) == 0 {
			continue
		}

		// converter a linha em uma lista de campos
		fields := strings.Split(line, ",")

		// analisar os dados
		ticket := Ticket{
			ID:          fields[0],
			Name:        fields[1],
			Email:       fields[2],
			Destination: fields[3],
			FlightTime:  fields[4],
		}

		// converter a string em float
		price, err := strconv.ParseFloat(fields[5], 64)
		// tratar o erro acima
		if err != nil {
			return 0, err
		}
		// atribuir o preco convertido ao tipo
		ticket.Price = price

		// pegar apenas a hora
		hourString := strings.Split(ticket.FlightTime, ":")[0]
		// converter a string para inteiro
		hour, err := strconv.Atoi(hourString)
		if err != nil {
			return 0, err
		}

		// incrementar a contagem se o horário estiver dentro do período desejado
		switch time {
		case "dawn":
			if hour >= 0 && hour <= 6 {
				count++
			}
		case "morning":
			if hour >= 7 && hour <= 12 {
				count++
			}
		case "afternoon":
			if hour >= 13 && hour <= 19 {
				count++
			}
		case "night":
			if hour >= 20 && hour <= 23 {
				count++
			}
		}
	}

	// verificar se houve bilhetes vendidos para o periodo de tempo especificado
	if count == 0 {
		return 0, fmt.Errorf("no tickets sold for the time %s", time)
	}

	return count, nil
}
