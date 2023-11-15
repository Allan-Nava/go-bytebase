package bytebase

/*
Query parameters
Parameter	Required?	Description	Example
environment	Required	The environment name for your schema review policy. Case sensitive	Dev
statement	Required	The SQL statement.	SELECT * FROM `table`
host	Optional	The instance host. Available values: MySQL, PostgreSQL, TiDB.	127.0.0.1
port	Optional	The instance port. Available values: MySQL, PostgreSQL, TiDB.	3306
databaseName	Optional	The database name in the instance.	DB Name
databaseType	Optional	The database type. Required if port, host, and databaseName are not specified. Available values: MySQL, PostgreSQL, TiDB.	MySQL
*/
type sqlAdviseRequest struct {
	Environment string `json:"environment"`
	Parameter 	string `json:"parameter"`
	Statement 	string `json:"statement"` //
}