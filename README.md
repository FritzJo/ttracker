# TTracker
## What is ttracker
Simple CLI to track your work hours

# Usage
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
- [ ] Configuration of default work hours
- [ ] Colored output
- [ ] CSV validation
- [ ] Code quality and error handling

## Process / Workflow
### Clock-In / Starting a work day
![alt text](doc/images/ttracker_process_clockin.drawio.png "Clock-In process")


### Clock-Out / Ending a work day
![alt text](doc/images/ttracker_process_clockout.drawio.png "Clock-Out process")
