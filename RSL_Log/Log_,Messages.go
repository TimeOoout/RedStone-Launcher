package RSL_Log

func msgExample() {
	if AutoMsg == true {
		//Statements.
	}
}

func msgInfoInitLogger(err string) {
	if AutoMsg == true {
		PrintWarning("Unable to init Logger! Error:%s", err)
	}
}
