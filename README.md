# TTracker
Simple CLI to track your work hours

## Configuration
Most configuration is handled by the ```config.json```.
|Option|Description|Default Value|
|--|--|--|
|InitialOvertime|Amount of overtime when starting to work with ttracker. (In minutes) |0|
|DefaultWorkingHours|Expected work hours per day. |8|
|BreakTime|Usual total break time for each day. (In minutes)|60|

## Usage
```
# Starting a work day
./ttracker in

# Ending a work day
./ttracker out

# Taking some time off
./ttracker take <Time in Minutes>

# Show summary of currently available overtime minutes
./ttracker summary
```

## Build
### Dependencies
* Golang 1.18

### Instructions
```shell
git clone https://github.com/FritzJo/ttracker.git
cd ttracker
go build *.go
```

## Features
- [x] Clock-in / Clock-out
- [x] CSV based storage
- [x] One file for each year
- [x] Summary of hours worked overtime
- [x] Taking time off
- [x] Configuration of default work hours
- [ ] Colored output
- [ ] CSV validation
- [ ] Code quality and error handling

## FAQ
### Does this tool handle different working hours for individual days?
* No
### How Can I update an older record?
* The recommended way is to simply edit the csv with a text editor
