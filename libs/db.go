package libs

var UserDBConfig = DatabaseConfig{
	Host:     "localhost", // Replace with your user database details
	Port:     "3306",
	Username: "root",
	Password: "",
	Database: "mst",
}

var TodoDBConfig = DatabaseConfig{
	Host:     "localhost", // Replace with your to-do database details
	Port:     "3306",
	Username: "root", // Replace with actual credentials
	Password: "",
	Database: "todo",
}
