package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	ptime "github.com/yaa110/go-persian-calendar"
	jodaTime "github.com/vjeantet/jodaTime"
	"strconv"
	"strings"
	// "reflect"
	// "log"
    // "math/big"
	// "log"
	// "time"
)

func Split(r rune) bool {
    return r == ':' || r == '/' || r == ' ' || r == '-'
}

// Get The Current Time
func get_clock(c *gin.Context) {
	rez_, this_err := get_clock_func()
	if this_err != nil {
		panic(this_err)
	}
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_clock_func() (func_result string, err error) {

	pt := ptime.Now()
	func_result = fmt.Sprintf("%v:%v:%v", pt.Hour(), pt.Minute(), pt.Second())
	return

}


// Get Today's Date
func get_today(c *gin.Context) {
	rez_, this_err := get_today_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_today_func() (func_result string, err error) {

	pt := ptime.Now()
	func_result = fmt.Sprintf("%v-%v-%v", pt.Year(), pt.Month(), pt.Day())
	return

}

// Get UNIX Time
func get_unix(c *gin.Context) {
	rez_, this_err := get_unix_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_unix_func() (func_result string, err error) {

	pt := ptime.Now()
	func_result = fmt.Sprintf("%v", pt.Unix())
	return

}

// Week of Month
func get_week_of_month(c *gin.Context) {
	rez_, this_err := get_week_of_month_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_week_of_month_func() (func_result string, err error) {

	pt := ptime.Now()

	func_result = fmt.Sprintf("%v", pt.MonthWeek())
	return

}

// Get Week of Year
func get_week_of_year(c *gin.Context) {
	rez_, this_err := get_week_of_year_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_week_of_year_func() (func_result string, err error) {

	pt := ptime.Now()

	func_result = fmt.Sprintf("%v", pt.YearWeek())
	return

}


//  Get Remaining Weeks
func get_remaining_weeks(c *gin.Context) {
	rez_, this_err := get_remaining_weeks_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_remaining_weeks_func() (func_result string, err error) {

	pt := ptime.Now()

	func_result = fmt.Sprintf("%v", pt.RYearWeek())
	return
}

func get_remaining_month(c *gin.Context) {
	rez_, this_err := get_remaining_month_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_remaining_month_func() (func_result string, err error) {

	pt := ptime.Now()

	func_result = fmt.Sprintf("%v", pt.RMonthDay())
	return
}




//  Format my unix timestamp
type Bazak_Date struct {
    Unix_Time     string  `json:"unix_time"`
    Format_Time   string  `json:"format_time"`
}

var bazak_dates = []Bazak_Date{
    // {Unix_Time: "1664121321", Format_Time: "yyyy/MM/dd E "},
    // {Unix_Time: "1664121321", Format_Time: "yyyy/MM/dd E hh:mm:ss a"},
}

func format_my_unix(c *gin.Context) {
	var bazak_date Bazak_Date

	if err := c.ShouldBindJSON(&bazak_date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// bazak_dates = append(bazak_dates, bazak_date)

	// printing the catched values
	// fmt.Println(bazak_date.Unix_Time)
	// fmt.Println(bazak_date.Format_Time)

	// Converting the fetched String Unix Time into proper data type (Int64)
	i, err := strconv.ParseInt(bazak_date.Unix_Time, 10, 64)
	if err != nil {
		panic(err)
	}
	pt := ptime.Unix(i, 0)
	// fmt.Println(pt.Format(bazak_date.Format_Time))
	func_result := fmt.Sprintf("%v", pt.Format(bazak_date.Format_Time))
	// fmt.Println(pt.Format("yyyy/MM/dd E hh:mm:ss a"))
	// fmt.Printf("Hello, %v with type %s!\n", i, reflect.TypeOf(i))
	
	c.JSON(http.StatusCreated, func_result)
}




//  Format Today
type Bazak_Today struct {
    Format_Time   string  `json:"format_time"`
}

func format_today(c *gin.Context) {
	var bazak_now Bazak_Today

	if err := c.ShouldBindJSON(&bazak_now); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pt := ptime.Now()
	func_result := fmt.Sprintf("%v", pt.Format(bazak_now.Format_Time))
	c.JSON(http.StatusCreated, func_result)
}


//  Convert Gregorian DateTime stamp tp Persian Equivalent  
type Bazak_GDay struct {
	Gor_Time     string  `json:"gor_time"`
    GFormat_Time   string  `json:"Gformat_time"`
	PFormat_Time   string  `json:"Pformat_time"`
}

func convert_gor2per(c *gin.Context) {
	var bazak_now Bazak_GDay

	if err := c.ShouldBindJSON(&bazak_now); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	layout := bazak_now.GFormat_Time
	input := bazak_now.Gor_Time
	output := bazak_now.PFormat_Time
	// tmp_date , error := time.Parse(layout, input)
	tmp_date, error := jodaTime.Parse(layout, input)
	if error != nil {
		fmt.Println(error)
		return
	}

	pt := ptime.New(tmp_date)
	func_result := fmt.Sprintf("%v", pt.Format(output))
	
	c.JSON(http.StatusCreated, func_result)
}




//  Convert Gregorian DateTime stamp tp Persian Equivalent  
type Bazak_PDay struct {
	Per_Time     string  `json:"Per_time"`
    PFormat_Time   string  `json:"PFormat_time"`
	GFormat_Time   string  `json:"GFormat_time"`
}

func convert_per2gor(c *gin.Context) {
	var bazak_now Bazak_PDay

	if err := c.ShouldBindJSON(&bazak_now); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := bazak_now.Per_Time
	layout := bazak_now.PFormat_Time
	// output := bazak_now.GFormat_Time

	// splitting the whole string into array
    a := strings.FieldsFunc(input, Split)
    fmt.Println(a)

	// Converting the Datetime elements to Integer
    var a_int = []int{}
    for _, i := range a {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        a_int = append(a_int, j)
    }
	
	pt_tmp_ := ptime.Now()
	var a1_val = pt_tmp_.Month()

	switch a1_ := a_int[1]; {
	case a1_ == 1:
		a1_val = ptime.Farvardin
	case a1_ == 2:
		a1_val = ptime.Ordibehesht
	case a1_ == 3:
		a1_val = ptime.Khordad
	case a1_ == 4:
		a1_val = ptime.Tir
	case a1_ == 5:
		a1_val = ptime.Mordad
	case a1_ == 6:
		a1_val = ptime.Shahrivar
	case a1_ == 7:
		a1_val = ptime.Mehr
	case a1_ == 8:
		a1_val = ptime.Aban
	case a1_ == 9:
		a1_val = ptime.Azar
	case a1_ == 10:
		a1_val = ptime.Dey
	case a1_ == 11:
		a1_val = ptime.Bahman
	case a1_ == 12:
		a1_val = ptime.Esfand
	default:
		break;
	}
	
		
    if len(layout) > 12 {
        var pt ptime.Time = ptime.Date(a_int[0], a1_val, a_int[2], a_int[3], a_int[4], a_int[5], 0, ptime.Iran())
		t_gor := pt.Time()
		// func_result := fmt.Sprintf("%v", t_gor.Format(output))
		c.JSON(http.StatusCreated, t_gor)
    } else {
        var pt ptime.Time = ptime.Date(a_int[0], a1_val, a_int[2],0, 0, 0, 0, ptime.Iran())
		t_gor := pt.Time()
		// func_result := fmt.Sprintf("%v", t_gor.Format(output))
		c.JSON(http.StatusCreated, t_gor)
    }

}



// misc. functions, added by Dear Hamid


func get_Tomorrow(c *gin.Context) {
	rez_, this_err := get_Tomorrow_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_Tomorrow_func() (func_result string, err error) {

	pt := ptime.Now().Tomorrow()
	func_result = fmt.Sprintf("%v-%v-%v", pt.Year(), pt.Month(), pt.Day())
	return

}

func get_Yesterday(c *gin.Context) {
	rez_, this_err := get_Yesterday_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_Yesterday_func() (func_result string, err error) {

	pt := ptime.Now().Yesterday()
	func_result = fmt.Sprintf("%v-%v-%v", pt.Year(), pt.Month(), pt.Day())
	return

}

func get_weekday(c *gin.Context) {
	rez_, this_err := get_weekday_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_weekday_func() (func_result string, err error) {

	pt := ptime.Now().Weekday()
	func_result = fmt.Sprintf("%v", pt.String())
	return

}

func get_daytime(c *gin.Context) {
	rez_, this_err := get_daytime_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_daytime_func() (func_result string, err error) {

	pt := ptime.Now().BeginningOfMonth().DayTime()
	func_result = fmt.Sprintf("%v", pt)
	return

}

func get_month(c *gin.Context) {
	rez_, this_err :=  get_month_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_month_func() (func_result string, err error) {

	pt := ptime.Now().Month()
	func_result = fmt.Sprintf("%v", pt.String())
	return

}


func get_zone(c *gin.Context) {
	rez_, this_err := get_zone_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_zone_func() (func_result string, err error) {

	pt := ptime.Now().ZoneOffset()

	func_result = fmt.Sprintf("%v",pt)
	return
}


func get_FirstMonthDay(c *gin.Context) {
	rez_, this_err := get_FirstMonthDay_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_FirstMonthDay_func() (func_result string, err error) {

	pt := ptime.Now().FirstMonthDay()

	func_result = fmt.Sprintf("%v",pt)
	return
}

func get_FirstWeekDay(c *gin.Context) {
	rez_, this_err := get_FirstWeekDay_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_FirstWeekDay_func() (func_result string, err error) {

	pt := ptime.Now().FirstWeekDay()
	func_result = fmt.Sprintf("%v",pt)
	return
}


func get_FirstYearDay(c *gin.Context) {
	rez_, this_err := get_FirstYearDay_func()
	if this_err != nil {
		panic(this_err)
	}
	// fmt.Println(rez_)
	// fmt.Printf("AFTER_CALLING: i has value: %v and type: %T\n", rez_, rez_)
	c.IndentedJSON(http.StatusOK, rez_)
}

func get_FirstYearDay_func() (func_result string, err error) {

	pt := ptime.Now().FirstYearDay()
	func_result = fmt.Sprintf("%v",pt)
	return
}







func main() {

	router := gin.Default()

	// Time based APIs
	router.GET("/now", get_clock)

	// Get Todays date
	router.GET("/today", get_today)
	router.POST("/today", format_today)

	// Unix Time Based APIs
	router.GET("/unix", get_unix)
	router.POST("/format_my_unix", format_my_unix)

	// Informative things...
	router.GET("/week_of_month", get_week_of_month)
	router.GET("/week_of_year", get_week_of_year)
	router.GET("/remaining_weeks", get_remaining_weeks)


	// Data Convertion Functions...

	// gor to persian
	router.POST("/gor2per", convert_gor2per)
	// persian to gor
	router.POST("/per2gor", convert_per2gor)



	// misc.
	router.GET("/Tomorrow", get_Tomorrow)
	router.GET("/Yesterday", get_Yesterday)
	router.GET("/zone", get_zone)
	router.GET("/weekday", get_weekday)
	router.GET("/daytime", get_daytime)
	router.GET("/month", get_month)
	router.GET("/FirstMonthDay", get_FirstMonthDay)
	router.GET("/FirstWeekDay", get_FirstWeekDay)
	router.GET("/FirstYearDay", get_FirstYearDay)
	

	router.Run("0.0.0.0:8020")

}

