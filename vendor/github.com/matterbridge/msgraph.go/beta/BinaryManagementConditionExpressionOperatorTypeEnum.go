// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BinaryManagementConditionExpressionOperatorType undocumented
type BinaryManagementConditionExpressionOperatorType int

const (
	// BinaryManagementConditionExpressionOperatorTypeVOr undocumented
	BinaryManagementConditionExpressionOperatorTypeVOr BinaryManagementConditionExpressionOperatorType = 0
	// BinaryManagementConditionExpressionOperatorTypeVAnd undocumented
	BinaryManagementConditionExpressionOperatorTypeVAnd BinaryManagementConditionExpressionOperatorType = 1
)

// BinaryManagementConditionExpressionOperatorTypePOr returns a pointer to BinaryManagementConditionExpressionOperatorTypeVOr
func BinaryManagementConditionExpressionOperatorTypePOr() *BinaryManagementConditionExpressionOperatorType {
	v := BinaryManagementConditionExpressionOperatorTypeVOr
	return &v
}

// BinaryManagementConditionExpressionOperatorTypePAnd returns a pointer to BinaryManagementConditionExpressionOperatorTypeVAnd
func BinaryManagementConditionExpressionOperatorTypePAnd() *BinaryManagementConditionExpressionOperatorType {
	v := BinaryManagementConditionExpressionOperatorTypeVAnd
	return &v
}
