package RSL_Log

func msgExample() {
	if AutoMsg == true {
		//Statements.
	}
}

func msgInfoInitLogger(err string) {
	if AutoMsg == true {
		PrintError("Unable to init Logger! Error:%s", err)
	}
}
