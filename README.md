# TTracker
## What is ttracker
Simple CLI to track your work hours

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
- [ ] One file for each year
- [ ] Summary of hours worked overtime

## Process / Workflow
### Clock-In / Starting a work day
![alt text](doc/images/ttracker_process_clockin.drawio.png "Clock-In process")


### Clock-Out / Ending a work day
![alt text](doc/images/ttracker_process_clockout.drawio.png "Clock-Out process")
