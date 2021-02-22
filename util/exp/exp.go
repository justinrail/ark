package exp

//CheckError error handler global
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
