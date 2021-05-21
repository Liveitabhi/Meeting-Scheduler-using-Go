package main

import (
	"fmt"
	"os"
	"bufio"
	// "encoding/json"
	"io/ioutil"
	// "bytes"
	"log"
	"net/http"
	"net/url"
	"strings"
	// "strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	// read faculty number
	fmt.Println("What is your faculty number?")
	id, _ := reader.ReadString('\n')

	// read the work to be done
	state := -1

	// while loop
	for state != 0 { 
		fmt.Println("What do you want to do?\nBlock/Unblock Calendar - 1\nSchedule/Unschedule Meeting - 2\nGet Schedule for Day - 3\nGet Meeting Summary for Month - 4\nExit - 0")
		_, err:= fmt.Scanln(&state)

		if err != nil{
			fmt.Println(err)

		} else if state == 1 {

			//////// BLOCK CALENDAR ////////

			// read date
			fmt.Println("Enter the date (format DD/MM)")
			date, _ := reader.ReadString('\n')

			// read time
			fmt.Println("Enter the time (format HH)")
			time, _ := reader.ReadString('\n')

			// read request type
			fmt.Println("Block - 1\nUnblock - 2")
			flag, _ := reader.ReadString('\n')

			// send request
			data := url.Values{
		        "date": {date},
		        "time": {time},
		        "owner": {id},
		        "type": {flag},
		    }
		    resp, err := http.PostForm("http://localhost:8080/block/", data)

		    if err != nil {
		        log.Fatal(err)
		    }
		    
		    // read response data
			res, _ := ioutil.ReadAll( resp.Body )
			resp.Body.Close()
			fmt.Printf( "%s\n", res)


		} else if state == 2 {
			//////// SCHEDULE MEETING ////////

			// read date
			fmt.Println("Enter the date (format DD/MM)")
			date, _ := reader.ReadString('\n')

			// read time
			fmt.Println("Enter the time (format HH)")
			time, _ := reader.ReadString('\n')

			// read request type
			fmt.Println("Schedule - 1\nUnschedule - 2")
			flag, _ := reader.ReadString('\n')
			flag = strings.TrimRight(flag,"\n\r")			

			data := url.Values{}
			if flag=="1" {
				// post request
				
				// read meeting title
				fmt.Println("Enter the meeting title")
				title, _ := reader.ReadString('\n')

				// read faculty group
				fmt.Println("Enter the faculty group (space separated faculty numbers)")
				grp, _ := reader.ReadString('\n')

				data = url.Values{
			        "date": {date},
			        "time": {time},
			        "owner": {id},
			        "type": {flag},
			        "title": {title},
			        "grp": {grp},
			    }
			} else {
				// delete request

				data = url.Values{
			        "date": {date},
			        "time": {time},
			        "owner": {id},
			        "type": {flag},
			    }
			}
			
			// send request
			resp, err := http.PostForm("http://localhost:8080/schedule/", data)

		    if err != nil {
		        log.Fatal(err)
		    }
		    
		    // read response data
			res, _ := ioutil.ReadAll( resp.Body )
			resp.Body.Close()
			fmt.Printf( "%s\n", res)

		} else if state == 3 {
			
			//////// GET SCHEDULE ////////

			// read date
			fmt.Println("Enter the date (format DD/MM)")
			date, _ := reader.ReadString('\n')

			// send request
			data := url.Values{
		        "date": {date},
		        "owner": {id},
		    }
		    resp, err := http.PostForm("http://localhost:8080/get/", data)

		    if err != nil {
		        log.Fatal(err)
		    }
		    
		    // read response data
			res, _ := ioutil.ReadAll( resp.Body )
			resp.Body.Close()
			fmt.Printf( "%s\n", res)

		}  else if state == 4 {

			//////// HOD getSUMMARY ////////

			// check if HOD
			id = strings.TrimRight(id,"\n\r")
			if id != "10" {
				fmt.Println("You are not the HOD. Not allowed !!")
				continue
			}

			// read month
			fmt.Println("Enter the month in digits (1-12)")
			month, _ := reader.ReadString('\n')

			// send request
			data := url.Values{
		        "month": {month},
		        "owner": {id},
		    }
		    resp, err := http.PostForm("http://localhost:8080/summary/", data)

		    if err != nil {
		        log.Fatal(err)
		    }
		    
		    // read response data
			res, _ := ioutil.ReadAll( resp.Body )
			resp.Body.Close()
			fmt.Printf( "%s\n", res)


		} else {
			// wrong input
			fmt.Println("Your input is wrong. Please try again")
		}

	}


}