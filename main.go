package main

import (
	"bufio"
	"errors"
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

	amount, err := parseInput(input)
	if err != nil {
		fmt.Println("Error:", err)
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
	if len(waterLogs) == 0 {
		fmt.Println("No water intake logs available today.")
		return
	}

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

	goal, err := parseInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dailyGoal = goal
	fmt.Printf("Your new daily goal is set to %d ml.\n", dailyGoal)
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
	reader := bufio.NewReader(os.Stdin)
	for {
		displayMenu()
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
			showHelp()
		case "5":
			fmt.Println("Exiting... Stay hydrated!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option (1-5).")
		}
	}
}
