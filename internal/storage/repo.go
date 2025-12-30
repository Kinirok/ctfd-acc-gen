package gormodel

import "gorm.io/gorm"

func AddUserToDB(db *gorm.DB, user Account) error {
	return db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func AddTeamToDB(db *gorm.DB, team Team) error {
	return db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&team)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
