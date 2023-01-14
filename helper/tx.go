package helper

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback().Error
		PanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit().Error
		PanicIfError(errCommit)
	}
}
