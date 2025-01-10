package main

import (
	"fmt"
	"os"
	"time"
)

func logWaterIntake(amount int, filename string) error {
	timestamp := time.Now().Format(time.RFC3339)

	logEntry := fmt.Sprintf("%s,%d\n", timestamp, amount)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}

	defer file.Close()

	_, err = file.WriteString(logEntry)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}

func main() {
	filename := "water_log.csv"
	var amount int

	fmt.Print("Enter the amount of water consumed (in ml): ")
	_, err := fmt.Scan(&amount)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid input. Please enter a positive number")
		return
	}

	err = logWaterIntake(amount, filename)
	if err != nil {
		fmt.Printf("Error logging water intake: %v\n", err)
		return
	}

	fmt.Printf("Logged %d ml of water successfully!\n", amount)
}
