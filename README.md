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
./ttracker in <Optional start time (hh:mm)>

# Ending a work day
./ttracker out <Optional end time (hh:mm)>

# Taking some time off
./ttracker take <Time in Minutes>

# Show summary of currently available overtime minutes
./ttracker summary

# Show change in overtime if work would end now
./ttracker status

# Validate the currently stored records
./ttracker validate
```
## Build
### Dependencies
* Golang (tested with >=1.18, but older versions should also work)
### Instructions
```shell
git clone https://github.com/FritzJo/ttracker.git
cd ttracker
go build *.go

# Or use the provided make file
# Build
make build

# Install
make install

# Uninstall
make uninstall

# Run code tests
make test
```
## Features
- [x] Clock-in / Clock-out
- [x] CSV based storage
- [x] One file for each year
- [x] Summary of hours worked overtime
- [x] Taking time off
- [x] Configuration of default work hours
- [x] CSV validation
- [ ] Code quality and error handling
## FAQ
### Does this tool handle different working hours for individual days?
* No
### How can I update an older record?
* The recommended way is to simply edit the csv with a text editor
