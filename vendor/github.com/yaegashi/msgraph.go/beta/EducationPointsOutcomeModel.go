// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationPointsOutcome undocumented
type EducationPointsOutcome struct {
	// EducationOutcome is the base model of EducationPointsOutcome
	EducationOutcome
	// Points undocumented
	Points *EducationAssignmentPointsGrade `json:"points,omitempty"`
	// PublishedPoints undocumented
	PublishedPoints *EducationAssignmentPointsGrade `json:"publishedPoints,omitempty"`
}
