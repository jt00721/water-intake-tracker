package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type WaterEntry struct {
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type TrackerData struct {
	Goal int          `json:"goal"`
	Logs []WaterEntry `json:"logs"`
}

func logWaterIntake(data *TrackerData) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the amount of water (in ml): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	amount, err := parseInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	entry := WaterEntry{
		Amount:    amount,
		Timestamp: time.Now(),
	}

	data.Logs = append(data.Logs, entry)
	fmt.Printf("Logged %d ml at %s\n", amount, entry.Timestamp.Format("2006-01-02 15:04:05"))
}

func viewDailyProgress(data TrackerData) {
	if len(data.Logs) == 0 {
		fmt.Println("No water intake logs available today.")
		return
	}

	today := time.Now().Format("2006-01-02")
	total := 0

	for _, entry := range data.Logs {
		if entry.Timestamp.Format("2006-01-02") == today {
			total += entry.Amount
		}
	}

	fmt.Printf("\nDaily Progress: %d/%d ml\n", total, data.Goal)
	if total >= data.Goal {
		fmt.Println("Congratulations! You've met your daily goal!")
	} else {
		fmt.Printf("You need %d more ml to reach your goal.\n", data.Goal-total)
	}
}

func setDailyGoal(data *TrackerData) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your new daily goal (in ml): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	goal, err := parseInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data.Goal = goal
	fmt.Printf("Your new daily goal is set to %d ml.\n", data.Goal)
}

func saveData(data TrackerData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	fmt.Println("Data saved successfully.")
	return nil
}

func loadData(filename string) (TrackerData, error) {
	var data TrackerData

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No existing data file found. Starting with default settings.")
			return TrackerData{Goal: 2000}, nil
		}
		return data, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return data, fmt.Errorf("failed to decode data: %v", err)
	}

	fmt.Println("Data loaded successfully.")
	return data, nil
}

func parseInput(input string) (int, error) {
	amount, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid input: please enter a number")
	}
	if amount <= 0 {
		return 0, errors.New("amount must be greater than zero")
	}
	return amount, nil
}

func displayMenu() {
	fmt.Println("\nWater Intake Tracker")
	fmt.Println("1. Log Water Intake")
	fmt.Println("2. View Daily Progress")
	fmt.Println("3. Set/Update Daily Goal")
	fmt.Println("4. Help")
	fmt.Println("5. Exit")
	fmt.Print("\nEnter your choice: ")
}

func showHelp() {
	fmt.Println("\nHelp Menu:")
	fmt.Println("1. Log Water Intake - Record the amount of water consumed.")
	fmt.Println("2. View Daily Progress - Check your progress toward the daily goal.")
	fmt.Println("3. Set/Update Daily Goal - Define or update your daily water intake target.")
	fmt.Println("4. Help - Display this help menu.")
	fmt.Println("5. Exit - Quit the application.")
}

func main() {
	const filename = "tracker.json"

	data, err := loadData(filename)
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		displayMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			logWaterIntake(&data)
		case "2":
			viewDailyProgress(data)
		case "3":
			setDailyGoal(&data)
		case "4":
			showHelp()
		case "5":
			err := saveData(data, filename)
			if err != nil {
				fmt.Printf("Error saving data: %v\n", err)
			}
			fmt.Println("Exiting... Stay hydrated!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option (1-5).")
		}
	}
}
