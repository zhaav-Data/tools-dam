# Dam
Dam (هنگام) Formula for Persian Date Calculations

## How to run?

`install go version 1.16`

`go run main.go`

## How to Use?

Below are the main API Endpoints:


### 1. Current Date

```
curl --location --request GET 'localhost:8080/today'

"1401-مهر-21"
```

### 2. Current Time

```
curl --location --request GET 'localhost:8080/now'

"18:58:59"
```


### 3. Current UNIX Time

```
curl --location --request GET 'localhost:8080/unix'

"1665675010"
```


### 4. Current Date (Custom Format)

```
curl --location --request POST 'localhost:8080/today' \
--header 'Content-Type: application/json' \
--data-raw '{"Format_Time": "yyyy/MM/dd E hh:mm:ss a"}'

"1401/07/21 پنج‌شنبه 07:01:42 ب.ظ"
```

List of applicable time format symbols:

|  Time Symbol          |         Value                               |
|--------------|------------------------------------------------------|
| yyyy, yyy, y | year (e.g. 1394)                                     |
| yy           | 2-digits representation of year (e.g. 94)            |
| MMM          | the Persian name of month (e.g. فروردین)             |
| MMI          | the Dari name of month (e.g. حمل)                    |
| MM           | 2-digits representation of month (e.g. 01)           |
| M            | month (e.g. 1)                                       |
| rw           | remaining weeks of year                              |
| w            | week of year                                         |
| W            | week of month                                        |
| RD           | remaining days of year                               |
| D            | day of year                                          |
| rd           | remaining days of month                              |
| dd           | 2-digits representation of day (e.g. 01)             |
| d            | day (e.g. 1)                                         |
| E            | the Persian name of weekday (e.g. شنبه)              |
| e            | the Persian short name of weekday (e.g. ش)           |
| A            | the Persian name of 12-Hour marker (e.g. قبل از ظهر) |
| a            | the Persian short name of 12-Hour marker (e.g. ق.ظ)  |
| HH           | 2-digits representation of hour [00-23]              |
| H            | hour [0-23]                                          |
| kk           | 2-digits representation of hour [01-24]              |
| k            | hour [1-24]                                          |
| hh           | 2-digits representation of hour [01-12]              |
| h            | hour [1-12]                                          |
| KK           | 2-digits representation of hour [00-11]              |
| K            | hour [0-11]                                          |
| mm           | 2-digits representation of minute [00-59]            |
| m            | minute [0-59]                                        |
| ss           | 2-digits representation of seconds [00-59]           |
| s            | seconds [0-59]                                       |
| ns           | nanoseconds                                          |
| S            | 3-digits representation of milliseconds (e.g. 001)   |
| z            | the name of location                                 |
| Z            | zone offset (e.g. +03:30)                            |


usefull time format phrases:

- "yyyy-MM-dd"
- "yy-MM-dd e"
- "hh:mm:ss a"



### 5. Format My Unix

Takes `unix time stamp` and desired `time format` as an argument and gives back the equivalent shamsi Date(and time) based on requested format:

```
curl --location --request POST 'localhost:8080/format_my_unix' \
--header 'Content-Type: application/json' \
--data-raw '{"Unix_Time": "1664121321", "Format_Time": "yyyy-MM-dd E hh:mm ss"}'

"1401-07-03 یک‌شنبه 07:25 21"
```


### 6. Convert Gregorian DateTime stamp tp Persian Equivalent  

Takes gregorian `datetime stamp` and its related gregorian `datetime format`, plus the desired persian datetime format, so called (`PFormat_Time`)
as an argument and gives back the equivalent Shamsi Date(and time) based on the requested format:

```
curl --location --request POST 'localhost:8080/gor2per' \
--header 'Content-Type: text/plain' \
--data-raw '{"Gor_Time": "2022/10/22 04:35", "GFormat_Time":"YYYY/MM/dd HH:mm","PFormat_Time": "yyyy-MM-dd E HH:mm"}'

"1401-07-30 شنبه 04:35"
```


### 6. Convert Persian Datetime Stamp to Gregorian Equivalent  

Takes persian `datetime stamp` and its related persian `datetime format`, plus the desired gregorian datetime format, so called (`GFormat_Time`)
as an argument and gives back the equivalent Gregorian Date(and time) based on the requested format:

```
curl --location --request POST 'localhost:8080/per2gor' \
--header 'Content-Type: application/json' \
--data-raw '{"Per_Time": "1401/07/30 09:54:30", "PFormat_Time":"YYYY/MM/dd HH:mm:ss","GFormat_Time": "yyyy/MM/dd"}'

"2022-10-22T09:54:30+03:30"
```

Or alternatively you can only pass Date, as it's shown below:


```
curl --location --request POST 'localhost:8080/per2gor' \
--header 'Content-Type: application/json' \
--data-raw '{"Per_Time": "1401/07/30", "PFormat_Time":"YYYY/MM/dd","GFormat_Time": "yyyy/MM/dd"}'

"2022-10-22T00:00:00+03:30"
```