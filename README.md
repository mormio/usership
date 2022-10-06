# Usership software in Go
Many of the things we consume we need but seldom use. Usership describes a mode of possession centered around sharing fewer objects within a group, rather than individual ownership of these items. The benefits of usership are cost efficiency and reduced consumption. <br><br>

# How to use this (mostly notes to self)
1. Establish the MySQL server. I did this locally:
   1. Download [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/) (note the SQL code in this repo is for v8.0.30)
   2. Install MySQL. For instance, in Mac OS shell: `brew install mysql` and then `brew services start mysql`
   3. You can now log into the DBMS with `mysql -u root` if you don't have a password, or `mysql -u root -p` if you do. 
2. If you're starting afresh, instantiate the tables:
   1. `source /usership/data-access/create-tables.sql` 
3. Import dependencies. The following commands import the dependencies in the go.mod file:
   1. `cd usership`
   2. `go get .`
4. While still in the data-access folder, run go.main to connect to your database according to your user and password (if you have one):
   1. `export DBUSER=root`
   2. `export DBPASS=`
   3. `go run .`
   4. The command interface should return a message: "Connected!"
5. Launch the interface to manage and use your usership group:
   1. This is coming 

# To do
- [x] Investigate data connections. mysql ~~or just dataframes in the meantime?~~
- [x] Establish database structure 
- [ ] Write functions for accessing database. Remaining functions:
  - [ ] Return items by matching name
  - [ ] For a given item ID, return its current user
  - [ ] Add an item
  - [ ] Delete an item 
  - [ ] Add a user
  - [ ] Delete a user
  - [ ] Change the current user of an item
  - [ ] Change the contact info of a user
- [ ] Write interface. 
  - [ ] https://earthly.dev/blog/tui-app-with-go/ <br><br>

# Requirements
* Installation of [Go](https://go.dev/doc/install)
* Installation of [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)


  
