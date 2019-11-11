# mealplan-restapi
Victor Obiahu 
November 7th 2019 

Web app for handling meal plans for various consumer markets
------
PROMPT
------
A company would like to sell meal plans.As part of the build to support meal plans, we will need a small webapp to allow internal staff to enter and edit meal plan information. Each plan will have a name, weekly cost, spring semester start and end dates, and fall semester start and end dates.  
Each plan will be associated with only one of the markets in which we operate and will contain a 'market' key which will store the name of one of four markets: Syracuse, Drexel, Colgate, or Maryland.  In addition, some of the schools where we operate have trimesters rather than semesters, so internal users will need to have the option to enter the start and end date for a third academic period.
------------
Requirements
------------
If a customer signs up for a meal plan, the id of the meal plan will be stored in the 'mealPlanId' key on the customer object. 
Customer objects are stored in a 'users' table in our MongoDB database.  

Assume that a user object has two keys: 'id' and 'mealPlanId'.  Importantly, if an internal user of the webapp deletes a meal plan, that operation will need to clear the 'mealPlan' key on all users tagged with the deleted meal plan and replace it with NULL.

Build a service/API in Go that uses a MongoDB table to accomplish all of the above.  
In addition, the API should allow for GET requests to retrieve all meal plans, all meal plans in a particular market, and a meal plan by id.  
The get all meal plans endpoint should return meal plans in alphabetical order by market name, then by weekly cost within the market. 
The service should also include a special GET to retrieve the meal plan for a given user.

--------------------
Guidelines for Linux 
--------------------
Install a textEditor of preference (In my case I used Atom) 
If running this application on a MAC OS..Hit Cmd + Shift + H and create a GO folder.
In the GO folder, create 3 sub folders  bin, pkg and src. Most of the code logic will be in src.
Used Gorilla MUX for web config and imported it into necessary files.
Set up MongoDB from Terminal using the following command: go get gopkg.in/mgo.v2
To install mongoDB Community Edition correctly via Linux on Mac OS using Homebrew run the following commands sequentially
brew tap mongodb/brew
brew install mongodb-community@4.2
mongod --config /usr/local/etc/mongod.conf --fork (To run mongo in the background include --fork)
brew services start mongodb-community@4.2
ps -ef | grep mongod (verifies MongoDB is running correctly)
Enter command : "mongo" 
Last enter command "show dbs" to show available dbs in Mongo

