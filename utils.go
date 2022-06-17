package protonmail

func iHasFlag(i, flag int) bool           { return i&flag == flag }
func iHasAtLeastOneFlag(i, flag int) bool { return i&flag > 0 }
func iIsFlag(i, flag int) bool            { return i == flag }
func iHasNoneOfFlag(i, flag int) bool     { return !iHasAtLeastOneFlag(i, flag) }
