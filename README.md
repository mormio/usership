# Usership software in Go
Many of the things we consume we need but seldom use. Usership describes a mode of possession centered around sharing fewer objects within a group, rather than individual ownership of these items. The benefits of usership are cost efficiency and reduced consumption. <br><br>

# How to use this
1. Establish the database. I did this locally:
   1. Download [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)
   2. Install MySQL. For instance, in Mac OS shell: `brew install mysql` and then `brew services start mysql`
   3. You can now log into the DBMS with `mysql -u root` if you don't have a password, or `mysql -u root -p` if you do. 
2. If you're starting afresh, instantiate the tables:
   1. `cd /usership/data-access/create-tables.sql` 
3. Launch the interface to manage and use your usership group:
   1. This is coming 

# To do
- [x] Investigate data connections. mysql ~~or just dataframes in the meantime?~~
- [ ] Establish database structure 
- [ ] Write functions for accessing database. 
- [ ] Write interface. <br><br>

# Requirements
* Installation of [Go](https://go.dev/doc/install)
* Installation of [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)


  
