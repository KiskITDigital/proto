// Code generated by ogen, DO NOT EDIT.

package api

// setDefaults set default value of fields.
func (s *V1QuestionnaireOrganizationIDPostReq) setDefaults() {
	{
		val := bool(false)
		s.IsCompleted = val
	}
}

// setDefaults set default value of fields.
func (s *V1QuestionnaireOrganizationIDStatusGetOKData) setDefaults() {
	{
		val := bool(false)
		s.IsCompleted = val
	}
}

// setDefaults set default value of fields.
func (s *V1TendersPostReq) setDefaults() {
	{
		val := bool(false)
		s.IsDraft.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *V1TendersTenderIDPutReq) setDefaults() {
	{
		val := bool(false)
		s.IsDraft.SetTo(val)
	}
}
