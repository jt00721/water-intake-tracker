package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type WaterEntry struct {
	Amount    int
	Timestamp time.Time
}

var waterLogs []WaterEntry

func logWaterIntake() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the amount of water (in ml): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	amount, err := strconv.Atoi(input)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return
	}

	entry := WaterEntry{
		Amount:    amount,
		Timestamp: time.Now(),
	}
	waterLogs = append(waterLogs, entry)
	fmt.Printf("Logged %d ml at %s\n", amount, entry.Timestamp.Format("2006-01-02 15:04:05"))
}

var dailyGoal int = 2000 // Default goal

func viewDailyProgress() {
	today := time.Now().Format("2006-01-02")
	total := 0

	for _, entry := range waterLogs {
		if entry.Timestamp.Format("2006-01-02") == today {
			total += entry.Amount
		}
	}

	fmt.Printf("\nDaily Progress: %d/%d ml\n", total, dailyGoal)
	if total >= dailyGoal {
		fmt.Println("Congratulations! You've met your daily goal!")
	} else {
		fmt.Printf("You need %d more ml to reach your goal.\n", dailyGoal-total)
	}
}

func setDailyGoal() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your new daily goal (in ml): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	goal, err := strconv.Atoi(input)
	if err != nil || goal <= 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return
	}

	dailyGoal = goal
	fmt.Printf("Your new daily goal is set to %d ml.\n", dailyGoal)
}

// func setDailyGoal(goal int, filename string) error {
// 	return nil
// }

// func calculateTotalWater(filename string) (int, error) {
// 	// Get logged water intakes from file
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return 0, fmt.Errorf("could not open file: %v", err)
// 	}

// 	defer file.Close()

// 	currentDate := time.Now().Format("2006-01-02")

// 	totalWater := 0

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		parts := strings.Split(line, ",")
// 		if len(parts) != 2 {
// 			continue
// 		}

// 		timestamp := parts[0]
// 		amount, err := strconv.Atoi(parts[1])
// 		if err != nil {
// 			continue
// 		}

// 		formattedTimeStamp := timestamp[:10]
// 		if formattedTimeStamp == currentDate {
// 			totalWater += amount
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return 0, fmt.Errorf("error reading file: %v", err)
// 	}

// 	return totalWater, nil
// }

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nWater Intake Tracker")
		fmt.Println("\n1. Log Water Intake")
		fmt.Println("2. View Progress")
		fmt.Println("3. Set Goal")
		fmt.Println("4. Exit")
		fmt.Print("\nChoose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			logWaterIntake()
		case "2":
			viewDailyProgress()
		case "3":
			setDailyGoal()
		case "4":
			fmt.Println("Exiting. Stay hydrated!")
			return
		default:
			fmt.Println("Invalid choice, input the number for the option require. Please try again.")
		}
	}

	// fmt.Print("Enter the amount of water consumed (in ml): ")
	// _, err := fmt.Scan(&amount)
	// if err != nil || amount <= 0 {
	// 	fmt.Println("Invalid input. Please enter a positive number")
	// 	return
	// }

	// err = logWaterIntake(amount, logFilename)
	// if err != nil {
	// 	fmt.Printf("Error logging water intake: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Logged %d ml of water successfully!\n", amount)

	// totalWater, err := calculateTotalWater(logFilename)
	// if err != nil {
	// 	fmt.Printf("Error calculating total water intake: %v\n", err)
	// 	return
	// }

	// if totalWater >= 1000 {
	// 	fmt.Printf("You've consumed %.2f litres today", float64(totalWater)/1000)
	// } else {
	// 	fmt.Printf("You've consumed %d ml today", totalWater)
	// }
}
