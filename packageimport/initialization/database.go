package database

var database string

// fungsi `init` akan otomatis terpanggil ketika `import`
func init() {
	database = "Postgres"
}

func GetDatabase() string {
	return database
}
