package validator

func BadBillRecordType(v *int) bool {
	if *v != 0 && *v != 1 && *v != 2 {
		return true
	}
	return false
}

func BadTimeDuration(startTime, endTime *int64) bool {
	return (startTime == nil) != (endTime == nil)
}
