package main

import (
	"flag"
	"fmt"
	"time"
)

// go run main.go -fd 201710302300 -td 201710312301
// /home/suwit/go/src/github.com/suwit471661124/softnix-getfile
func main() {
	p := fmt.Println

	fromDateTime := flag.String("dateFrom", "", "a string. Such a 2017-10-01")
	toDateTime := flag.String("dateTo", "", "a string. Such a 2017-10-02")
	pathSource := flag.String("path", "", "a string path data source. Such a /home/softnixlogger/xxxx")
	port := flag.Int("port", 514, "an int")
	//sources := flag. ("source", "", "an array")
	logFile := flag.String("log", "/var/log/softnix/query_time.log", "a string path")

	//wordPtr := flag.String("word", "foo", "a string")
	//numbPtr := flag.Int("n", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")

	flag.Parse()
	p("fromDateTime:", *fromDateTime)
	p("toDateTime:", *toDateTime)
	p("pathSource:", *pathSource)
	p("logFile:", *logFile)
	p("port:", *port)
	//p("sources:", *sources)
	p("tail:", flag.Args())
	p()

	//fmt.Println(diffDays2(2012, 10, 10, 2012, 10, 10))

	diffDays2()
}

func diffDays2() {
	p := fmt.Println

	//now := time.Now().Add(-24 * time.Hour)
	toDate := time.Date(2017, 05, 29, 10, 15, 58, 000000, time.UTC)
	p(toDate)
	fromDate := time.Date(2017, 05, 29, 10, 10, 58, 00000, time.UTC)
	p(fromDate)
	p(fromDate.Year())
	p(fromDate.Month())
	p(fromDate.Day())
	p(fromDate.Hour())
	p("Minute:", fromDate.Minute())
	//p("Second:", fromDate.Second())
	//p(fromDate.Nanosecond())
	p(fromDate.Location())
	p(fromDate.Weekday())
	p(fromDate.Before(toDate))
	p(fromDate.After(toDate))
	p(fromDate.Equal(toDate))

	p()
	diff := toDate.Sub(fromDate)

	p(diff)
	p("Hours:", diff.Hours())
	p("Minutes:", diff.Minutes())
	//p("Seconds:", diff.Seconds())
	//p(diff.Nanoseconds())
	p(fromDate.Add(diff))
	p(fromDate.Add(-diff))
}

func diffDays(
	yearFrom int, monthFrom time.Month, dayFrom int,
	yearTo int, monthTo time.Month, dayTo int,
) int {
	dateFrom := time.Date(
		yearFrom, monthFrom, dayFrom,
		10, 19, 0, 0, time.UTC,
	)
	dateTo := time.Date(
		yearTo, monthTo, dayTo,
		10, 20, 10, 0, time.UTC,
	)
	//fmt.Println(int(dateTo.Sub(dateFrom) / 60))

	//fmt.Println((60 * time.Minute))

	fmt.Println(dateTo.Sub(dateFrom))

	return int(dateFrom.Sub(dateTo) / (60 * time.Minute))
}

func setArgs(interval string, datefrom int, dateto int) {
	/*
				case "m": // Number of full months
					$months_difference = floor($difference / 2678400);
					while (mktime(date("H", $datefrom), date("i", $datefrom), date("s", $datefrom), date("n", $datefrom)+($months_difference), date("j", $dateto), date("Y", $datefrom)) < $dateto) {
						$months_difference++;
					}
					$months_difference--;
					$datediff = $months_difference;
					break;

					$datefrom = strtotime($datefrom, 0);
		            $dateto = strtotime($dateto, 0);
	*/

	//datefrom = strtotime(datefrom, 0);

	//$difference = $dateto - $datefrom;
	switch interval {
	case "m":
		//months_difference := floor($difference / 2678400);
		fmt.Println("from date")
	case "-td":
		fmt.Println("to date")
	default:
		fmt.Println("It's default")
	}

}

/*
//setFormatTime('2017-10-11')
func setFormatTime() {
	for i := 0; i < count; i++ {

	}
}
*/
