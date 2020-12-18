package commonpb

// ConvertToBool when value == yes return true
func (i YesNo) ConvertToBool() bool {
	if i == YesNo_YES {
		return true
	}
	return false
}
