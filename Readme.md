CS17B002 - ABHISHEK KUMAR - DISTRIBUTED SYSTEMS - ASSIGNMENT 2

----------To use the Meeting Scheduler,
1. Run server.go file
2. Run as many instances of client.go file as you want
3. Enter your inputs according to the task you want to perform


----------Points
1. Only Post and Delete requests are allowed in "Block calendar" & "Scchedule meeting" features. (No PUT/PATCH)
2. Inputs have to be provided in correct format as asked, otherwise user can face errors.
3. The scheduler is designed for the year 2021. It means user will provide only Date & Month for any task.
4. At the same date and time, multiple meetings with mutually exclusive list of faculty members can be scheduled.
5. Only the host of a meeting can delete/unschedule it, not the members.