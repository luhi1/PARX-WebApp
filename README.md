# ParX
## Your new favorite tool for tracking student participation in school events!
Gif here!

## Installation
### Dependencies
* GO >= 1.19
* MySQL (or MariaDB)

### ⚠️ This project is intended to be run server-side, so setup involves creating a MySQL database and running a local instance of the webapp. Follow steps below carefully. ⚠️ 
### Quick Start (Most Users)
1. Navigate over to the releases tab and install ```createDatabase.sql``` and ```PARX-WebApp.exe```.
2. Execute the ```createDatabase.sql``` file to generate the database and populate it with sample data.
  * Option A: GUI using MySQL Workbench (Information derived from ![this guide.](https://www.geeksforgeeks.org/how-to-import-and-export-data-to-database-in-mysql-workbench/))
    * Open the application and log in with your username and password.
    * Navigate to the Server Administration tab
    * Click on Manage Import/Export
    * Click on Data Import/Restore (on the left side of the screen).
    * Select Import from Self-Contained File and input the path to createDatabase.sql.
    * Hit the start import button.
    * Finally, execute the script.
  * Option B (Command line)
    * If you have never used MySQL before, enter the installation utility and follow the steps:
      ```bash
      sudo mysql_secure_installation
    
    * Then, log in
      ```bash
      mysql -u {USERNAME} -p
    * Finally, to execute the createDatabase.sql file:
      ```mysql
      source full/path/to/createDatabase.sql
    * To re-run this script in the future (should not be neccesary), make sure you execute this before re-running:
      ```mysql
      drop schema `fbla`;
  2. Execute the binary ```PARX-WebApp.exe```.
  3. Open a web browser and navigate to localhost:8082.
  4. Enjoy!
      
### Advanced Setup (Compiling from Source)












![ParX icon](static/parxfull.png)
## Built By
* @givingdonation - Carlo Allietti (Dev 1)
* @luhi1 - Michael Borov (Dev 2)
* @ChadicalRadical - Chad Khan (Graphic Design)
