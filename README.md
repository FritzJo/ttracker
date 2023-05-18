# TTracker

Simple CLI to track your work hours. Written in Go with zero external dependencies and a local, CSV-based storage.

## Configuration

Most configuration is handled by the ```config.json```.

| Option              | Description                                                          | Default Value |
|---------------------|----------------------------------------------------------------------|---------------|
| InitialOvertime     | Amount of overtime when starting to work with ttracker. (In minutes) | 0             |
| DefaultWorkingHours | Expected work hours per day.                                         | 8             |
| BreakTime           | Usual total break time for each day. (In minutes)                    | 60            |
| StorageLocation     | Output directory for the created CSV files.                          | .             |

## Usage

In its most basic form, simply use the ```in``` and ```out``` commands to start and end your work day.
Everything else is handled in the background by ttracker. 
The log of all working hours will be stored in a CSV file, named after the current year.
An example file can be found in [```docs/2023_data.csv```](docs/2023_data.csv)

```
# Starting a work day
./ttracker in <Optional start time (hh:mm)>

# Ending a work day
./ttracker out <Optional end time (hh:mm)>
```

The full list of supported commands is listed below.

| Command  | Description                                                                                  | Example         |
|----------|----------------------------------------------------------------------------------------------|-----------------|
| in       | Start work day. Optionally provide a start time, if it differs from the current system time. | ```in 7:30```   |
| out      | End work day. Optionally provide a end time, if it differs from the current system time.     | ```out 17:30``` |
| take     | Taking some time off. Amount of minutes are input via the second parameter.                  | ```take 240```  |
| summary  | Show summary of currently available overtime minutes.                                        | ```summary```   |
| status   | Show change in overtime if work would end now and the clock-in time.                         | ```status```    |
| validate | Validate the currently stored records. Prints invalid records.                               | ```validate```  |

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

## Features / Roadmap

- [x] Clock-in / Clock-out
- [x] CSV based storage
- [x] One file for each year
- [x] Summary of hours worked overtime
- [x] Taking time off
- [x] Configuration of default work hours
- [x] CSV validation
- [ ] Code quality and error handling
- [ ] Map time taken off to a specific day

## FAQ

### Does this tool handle different working hours for individual days?

* No

### How can I update an older record?

* The recommended way is to simply edit the csv with a text editor
