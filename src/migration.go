package src

//AutoMigration update the database schema
func AutoMigration() {
	db := dbConnect()
	defer db.Close()

	db.AutoMigrate(user{})
	db.AutoMigrate(debt{})

	db.Model(&debt{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
