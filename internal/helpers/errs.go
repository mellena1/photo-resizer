package helpers

// PanicIfErr panic if there is an error
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
