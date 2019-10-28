package errors


// duplicate of the generate_function.go file simply nicer to be able to say errors.Check(something)
func Check(e error) {
	if e != nil {
		panic(e)
	}
}