# water-intake-tracker
A command-line application that tracks daily water intake, calculates totals, and checks if a daily goal is met. Optionally, explore data persistence using a file or a simple database.

## Core Features

- Log Water Intake

    Users can log the amount of water consumed with timestamps.
    Example: "Logged 500ml at 10:00 AM."

- Calculate Total Intake

    Sum up the total water consumed for the current day.

- Set and Track Daily Goals

    Users can set a daily water intake goal.
    The app tracks progress and notifies the user when the goal is reached or exceeded.

- Display Progress

    Show the user how much water they’ve consumed so far and how close they are to their goal.

- Data Persistence (Optional)

    Save water intake logs to a file or database for later use.
    Load and display previous entries (optional).

## Day 1: Define Scope and Setup

    Goals:
        Define the features for the tracker (e.g., log intake, calculate totals, check goals).
        Plan how the user interacts with the app (e.g., menu-driven CLI or command-line arguments).
        Create a project folder structure.

    Tasks:
        Brainstorm and document the project features.
        Initialize a Git repository and create a README.md with the project description.
        Set up a basic main.go file.

## Day 2: Implement Basic Functionality

    Goals:
        Create a CLI for logging water intake and displaying total intake.

    Tasks:
        Write a function to log water intake with timestamps.
        Write a function to calculate the total water consumed for the day.
        Test these functionalities with mock data.

## Day 3: Add Goal-Tracking

    Goals:
        Implement a feature to set a daily water intake goal.
        Notify the user when the goal is reached or exceeded.

    Tasks:
        Create a function to set and retrieve the daily goal.
        Integrate goal tracking into the total calculation logic.
        Add user-friendly messages to indicate progress toward the goal.

## Day 4: Data Persistence (Optional)

    Goals:
        Store water intake logs persistently in a file or lightweight database.

    Tasks:
        Use a .txt or .json file to save logs.
        Explore SQLite for a simple database solution (optional).
        Add functions to read/write logs to/from the storage.

## Day 5: Refactor and Polish

    Goals:
        Improve code structure, handle edge cases, and test thoroughly.

    Tasks:
        Refactor repetitive logic into reusable functions.
        Add error handling (e.g., invalid input, missing logs).
        Test edge cases like no intake logged or goal not set.

## Day 6: Documentation and Demo

    Goals:
        Document how to use the app and create a demo for social media.

    Tasks:
        Update the README.md with usage instructions and examples.
        Record a demo of the app in action (e.g., showing the CLI usage).
        Take a screenshot or photo for social media.

