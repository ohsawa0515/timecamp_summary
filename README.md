# timecamp_summary
Summarize Timecamp tasks for the specified time period

# Installation

```console
$ go get -u github.com/ohsawa0515/timecamp_summary
```

# Settings

```console
export TIMECAMP_TOKEN="xxxxxxxxxxxxxxxxxxxx"
export TIMECAMP_USER_ID="xxxxxxx"
```

# Usage

```console
$ timecamp_summary --help
Usage of timecamp_summary:
  -from string
        The beginning of the date(YYYY-MM-DD). The default is today.
  -to string
        End of the date(YYYY-MM-DD). The default is the same day as from option.
```

# Options

## -from

The beginning of the date(YYYY-MM-DD). The default is today.

## -to

End of the date(YYYY-MM-DD). The default is the same day as from option.

# Example

```console
# Today's summary report
$ timecamp_summary
Development, 4h41m41s
Paperwork, 1h45m49s
Meeting, 44m16s
Total: 7h11m46s

# Summary report for a specific day
$ timecamp_summary -from 2018-02-20
Operation check, 35m32s
Inquiry, 1h13m9s
Paperwork , 2h48m50s
Meeting, 1h1m4s
Project A, 3m5s
Project B, 28m14s
Total: 6h9m54s

# Summary report for a specific time period
$ timecamp_summary -from 2018-02-22 -to 2018-02-23
Project A, 4h47m25s
Meeting, 1h57m47s
Cost calculation, 1h43m51s
Paperwork, 45m56s
Project B, 1h22m31s
Operation check, 1h8m54s
Inquiry, 26m53s
Meeting, 3h9m57s
Total: 15h23m14s
```

# License

See [LICENSE](https://github.com/ohsawa0515/timecamp_summary/blob/master/LICENSE)

# Author

Shuichi Ohsawa (@ohsawa0515)
