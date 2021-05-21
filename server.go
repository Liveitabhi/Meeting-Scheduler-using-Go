package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "strconv"
    "sort"
)

type slot struct {
	blockORmeeting byte
	title          string
	members        []string
	owner          string
}

// Structure to store the schedule

type m1 map[int][24][]slot
var m = make(map[int]m1)

func Intersection(a, b []string) (c bool, d string) {
      mp := make(map[string]bool)
      c = false

      for _, item := range a {
              mp[item] = true
      }

      for _, item := range b {
              if _, ok := mp[item]; ok {
                      c = true
                      d = item
                      return
              }
      }
      return
}

func common(a, b slot) (c bool, d string){
	c, d = Intersection(append(a.members, a.owner), append(b.members, b.owner))
	return
}


func blockHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the json object to get date, time and req type
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	split := strings.Split(r.Form.Get("date"), "/")
	
	month, err := strconv.Atoi(strings.TrimRight(split[1],"\n\r"))
	if err!=nil {fmt.Println(err)}

	date, err := strconv.Atoi(strings.TrimRight(split[0],"\n\r"))
	if err!=nil {fmt.Println(err)}

	time, err := strconv.Atoi(strings.TrimRight(r.Form.Get("time"),"\n\r"))
	if err!=nil {fmt.Println(err)}
    
    req := strings.TrimRight(r.Form.Get("type"),"\n\r")

	tmp1 := m[month]
	if _, found := m[month]; !found {
    		tmp1 = make(m1)
	}
	var tmp [24][]slot = tmp1[date]

	// Length of array
	l := len(tmp[time])
	arr := tmp[time]

    // Check request type
	switch req {

	case "1":	//POST

		fmt.Println("----------------POST req----------------")

	    // Store in the map data structure
		var s slot
		s.blockORmeeting = 'b'
		
		s.owner = strings.TrimRight(r.Form.Get("owner"),"\n\r")

		if l>0 {
		// Meetings exist in that slot
			for i := 0; i < l; i++ {
		        if ok, _ := common(arr[i], s); ok {
		        	
		        	// Slot not available
			    	fmt.Printf("SLOT NOT AVAILABLE for F%s Date-%d/%d Time-%d hrs\n\n", s.owner, date, month, time)
					fmt.Fprintf(w, "SLOT NOT AVAILABLE for F%s Date-%d/%d Time-%d hrs\n\n", s.owner, date, month, time)
					return
		        }
		    }
		}

		// Slot available for faculty
		if l==0 {
			tmp[time] = []slot{s}
		} else {
			tmp[time] = append(tmp[time], s)
		}
		tmp1[date] = tmp
		m[month] = tmp1

		fmt.Printf("CALENDAR BLOCKED!!! for F%s Date-%d/%d Time-%d hrs\n\n", s.owner, date, month, time)
		fmt.Fprintf(w, "CALENDAR BLOCKED!!! for F%s Date-%d/%d Time-%d hrs\n\n", s.owner, date, month, time)

	case "2":	//DELETE

		fmt.Println("----------------DELETE req----------------")

		// Delete from the map data structure
		owner := strings.TrimRight(r.Form.Get("owner"),"\n\r")

		if l>0 {
		// Meetings exist in that slot

			for i := 0; i < l; i++ {
		        if (arr[i].owner == owner && arr[i].blockORmeeting == 'b') {
		        	// Match found

		        	arr = append(arr[:i], arr[i+1:]...)
					tmp[time] = arr
					tmp1[date] = tmp
					m[month] = tmp1
		        	fmt.Printf("BLOCK DELETED!!! for F%s Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)
					fmt.Fprintf(w, "BLOCK DELETED!!! for F%s Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)
		        	return
		        }
		    }
		}

		// Slot available for faculty
		fmt.Printf("CALENDAR NOT BLOCKED for F%s Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)
		fmt.Fprintf(w, "CALENDAR NOT BLOCKED for F%s Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)

    default:
        fmt.Fprintf(w, "Sorry, only POST and DELETE methods are supported.")
    }
}


func scheduleHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the json object to get date, time and req type
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	split := strings.Split(r.Form.Get("date"), "/")
	
	month, err := strconv.Atoi(strings.TrimRight(split[1],"\n\r"))
	if err!=nil {fmt.Println(err)}

	date, err := strconv.Atoi(strings.TrimRight(split[0],"\n\r"))
	if err!=nil {fmt.Println(err)}

	time, err := strconv.Atoi(strings.TrimRight(r.Form.Get("time"),"\n\r"))
	if err!=nil {fmt.Println(err)}
    
    req := strings.TrimRight(r.Form.Get("type"),"\n\r")

	tmp1 := m[month]
	if _, found := m[month]; !found {
    		tmp1 = make(m1)
	}
	var tmp [24][]slot = tmp1[date]

	// Length of array
	l := len(tmp[time])
	arr := tmp[time]

    // Check request type
	switch req {

	case "1":	//POST

		fmt.Println("----------------POST req----------------")

		// Fetch meeting title
		title := strings.TrimRight(r.Form.Get("title"),"\n\r")

		// Fetch meeting group
		grp := strings.Split(strings.TrimRight(r.Form.Get("grp"),"\n\r"), " ")

	    // Store in the map data structure
		var s slot
		s.blockORmeeting = 'm'
		s.title = title
		s.members = grp
		s.owner = strings.TrimRight(r.Form.Get("owner"),"\n\r")

		if l>0 {
		// Meetings exist in that slot
			for i := 0; i < l; i++ {
		        if ok, id := common(arr[i], s); ok {
		        	
		        	// Slot not available
			    	fmt.Printf("SLOT NOT AVAILABLE for F%s Date-%d/%d Time-%d hrs\n\n", id, date, month, time)
					fmt.Fprintf(w, "SLOT NOT AVAILABLE for F%s Date-%d/%d Time-%d hrs\n\n", id, date, month, time)
					return
		        }
		    }
		}

		// Slot available for faculty
		if l==0 {
			tmp[time] = []slot{s}
		} else {
			tmp[time] = append(tmp[time], s)
		}
		tmp1[date] = tmp
		m[month] = tmp1

		fmt.Printf("MEETING SCHEDULED!!! by F%s Date-%d/%d Time-%d hrs\n\n", s.owner, date, month, time)
		fmt.Fprintf(w, "MEETING SCHEDULED!!! by F%s Date-%d/%d Time-%d hrs\n\n", s.owner, date, month, time)

	case "2":	//DELETE

		fmt.Println("----------------DELETE req----------------")

		// Delete from the map data structure
		owner := strings.TrimRight(r.Form.Get("owner"),"\n\r")


		if l>0 {
		// Meetings exist in that slot

			for i := 0; i < l; i++ {
		        if (arr[i].owner == owner && arr[i].blockORmeeting == 'm') {
		        	// Match found

		        	arr = append(arr[:i], arr[i+1:]...)
					tmp[time] = arr
					tmp1[date] = tmp
					m[month] = tmp1
		        	fmt.Printf("MEETING DELETED!!! for F%s Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)
					fmt.Fprintf(w, "MEETING DELETED!!! for F%s Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)
		        	return
		        }
		    }
		}

		// Slot available for faculty
		fmt.Printf("NO MEETING OWNED BY F%s on Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)
		fmt.Fprintf(w, "NO MEETING OWNED BY F%s on Date-%d/%d Time-%d hrs\n\n", owner, date, month, time)

    default:
        fmt.Fprintf(w, "Sorry, only POST and DELETE methods are supported.")
    }
}


func getHandler(w http.ResponseWriter, r *http.Request) {

	// GET
	fmt.Println("----------------GET req----------------")

	// Parse the json object to get date, time and owner
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	split := strings.Split(r.Form.Get("date"), "/")
	
	month, err := strconv.Atoi(strings.TrimRight(split[1],"\n\r"))
	if err!=nil {fmt.Println(err)}

	date, err := strconv.Atoi(strings.TrimRight(split[0],"\n\r"))
	if err!=nil {fmt.Println(err)}

	owner := strings.TrimRight(r.Form.Get("owner"),"\n\r")
	fmt.Fprintf(w, "\n-------------SCHEDULE for F%s Date-%d/%d :-\n\n", owner, date, month)

	tmp1 := m[month]
	if _, found := m[month]; !found {
			fmt.Fprintf(w, "No blocks or meetings on the specified date\n\n")
			fmt.Printf("SCHEDULE RETURNED!!! for F%s Date-%d/%d\n\n", owner, date, month)
    		return
	}

	var tmp [24][]slot = tmp1[date]

	// To check blocks/meetings exist or not
	flag := false

	// Check for all timeslots on that day
	for time :=0; time < 24; time++ {

		// Length of array
		l := len(tmp[time])
		arr := tmp[time]

		if l>0 {
		// Meetings exist in that slot

			for i := 0; i < l; i++ {

	        	if arr[i].blockORmeeting == 'b'{
	        	// A block in that slot

			        if owner == arr[i].owner {
			        // Faculty is the owner
			        	
			        	flag = true
	        			fmt.Fprintf(w, "Block at Time-%d hrs\n", time)
	        		}
	        	} else {
	        	// A  meeting in that slot

	        		if owner == arr[i].owner {
			        // Faculty is the owner
			        	
			        	flag = true
	        			fmt.Fprintf(w, "Meeting at Time-%d hrs by F%s\n", time, owner)

	        		} else {
	        		// Check for faculty in the members

	        			for _, value := range(arr[i].members) {
	        				if owner == value {
	        					// Faculty is a member

	        					flag = true
	        					fmt.Fprintf(w, "Meeting at Time-%d hrs by F%s\n", time, arr[i].owner)
	        					continue
	        				}
	        			}
	        		}
	        	}
		    }
		}
	}

	fmt.Printf("SCHEDULE RETURNED!!! for F%s Date-%d/%d\n\n", owner, date, month)
	
	// No meetings or blocks
	if flag==false {
		fmt.Fprintf(w, "No blocks or meetings on the specified date\n\n")
		return
	}

	fmt.Fprintf(w, "\n-------------END OF SCHEDULE\n\n")
	
}


func summaryHandler(w http.ResponseWriter, r *http.Request) {

	// GET
	fmt.Println("----------------GET req----------------")

	// Parse the json object to get date, time and owner
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	month, err := strconv.Atoi(strings.TrimRight(r.Form.Get("month"),"\n\r"))
	if err!=nil {fmt.Println(err)}

	// owner := strings.TrimRight(r.Form.Get("owner"),"\n\r")
	fmt.Fprintf(w, "\n-------------SUMMARY for Month %d :-\n\n", month)

	tmp1 := m[month]
	if _, found := m[month]; !found {
			fmt.Fprintf(w, "No meetings in the specified month\n\n")
    		return
	}

	// Get all dates for specified month with blocks or meetings in sorted order
	var dates []int
	for date := range tmp1 {
	  dates = append(dates, date)
	}
	sort.Ints(dates)

	// To check meetings exist or not
	flag := false

	for _, date := range dates {

		var tmp [24][]slot = tmp1[date]

		// Check for all timeslots on that day
		for time :=0; time < 24; time++ {

			// Length of array
			l := len(tmp[time])
			arr := tmp[time]

			if l>0 {
			// Meetings/Blocks exist in that slot

				for i := 0; i < l; i++ {

		        	if arr[i].blockORmeeting == 'm'{
		        	// A  meeting in that slot

		        		flag = true
		        		fmt.Fprintf(w, "Meeting on Date-%d/%d Time-%d hrs by F%s\n", date, month, time, arr[i].owner)
		        		fmt.Fprintf(w, "Meeting title - %s\n", arr[i].title)
		        		mem := ""
		        		for _, str := range arr[i].members {
		        			mem += "F" + str + " "
		        		}
		        		fmt.Fprintf(w, "Members - %s\n\n", mem)
		        	}
			    }
			}
		}
	}

	fmt.Printf("SUMMARY RETURNED!!! for Month-%d\n\n", month)

	// No meetings
	if flag==false {
		fmt.Fprintf(w, "No meetings in the specified month\n\n")
		return
	}

	fmt.Fprintf(w, "-------------END OF SUMMARY\n\n")
}



func main() {

    http.HandleFunc("/block/", blockHandler)
    http.HandleFunc("/schedule/", scheduleHandler)
    http.HandleFunc("/get/", getHandler)
    http.HandleFunc("/summary/", summaryHandler)

    fmt.Printf("Starting server at port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}