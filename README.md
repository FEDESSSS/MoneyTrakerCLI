# Money Traker CLI

A simple command-line tool for money tasks, written in Go.

This project is the solution for [This Project]([https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/expense-tracker))

## Features
- Add, update, delete 
- Download CSV file
- List all expense and viewing a specific month
- Persistent storage in `moneyList.json`


### Available Commands

- `add --description <expense_description> --amount <amount>`: Add a new expense with a specified description and amount
- `list`: List all expense
- `summary --month <month number>`: List expense a specific month
- `delete <expense_id>`: Delete a expense by its ID
- `csv`: Form a csv file
- `summary`: Viewing the total cost
## Examples

Add a new expense with description and amount:
```sh
MoneyTrackerCLI add --description "Coffee" --amount 3
```

List all recorded expenses:
```sh
MoneyTrackerCLI list
```

Show expenses for specific month:
```sh
MoneyTrackerCLI summary --month 5
```

Delete expense by ID:
```sh
MoneyTrackerCLI delete 2
```

Export all expenses to CSV file:
```sh
MoneyTrackerCLI csv
```

Show total spending:
```sh
MoneyTrackerCLI summary
```

## License
FEDESSSS
