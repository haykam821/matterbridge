// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GiphyRatingType undocumented
type GiphyRatingType int

const (
	// GiphyRatingTypeVStrict undocumented
	GiphyRatingTypeVStrict GiphyRatingType = 0
	// GiphyRatingTypeVModerate undocumented
	GiphyRatingTypeVModerate GiphyRatingType = 1
	// GiphyRatingTypeVUnknownFutureValue undocumented
	GiphyRatingTypeVUnknownFutureValue GiphyRatingType = 2
)

// GiphyRatingTypePStrict returns a pointer to GiphyRatingTypeVStrict
func GiphyRatingTypePStrict() *GiphyRatingType {
	v := GiphyRatingTypeVStrict
	return &v
}

// GiphyRatingTypePModerate returns a pointer to GiphyRatingTypeVModerate
func GiphyRatingTypePModerate() *GiphyRatingType {
	v := GiphyRatingTypeVModerate
	return &v
}

// GiphyRatingTypePUnknownFutureValue returns a pointer to GiphyRatingTypeVUnknownFutureValue
func GiphyRatingTypePUnknownFutureValue() *GiphyRatingType {
	v := GiphyRatingTypeVUnknownFutureValue
	return &v
}
