package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	// cek apabila terjadi error dengan recover
	err := recover()
	if err != nil {
		errorrRollback := tx.Rollback()
		PanicIfError(errorrRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
