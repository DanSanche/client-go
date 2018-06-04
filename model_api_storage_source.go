/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// StorageSource describes the location of the source in an archive file in Google Cloud Storage.
type ApiStorageSource struct {
	// Google Cloud Storage bucket containing source (see [Bucket Name Requirements] (https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	Bucket string `json:"bucket,omitempty"`

	// Google Cloud Storage object containing source.
	Object string `json:"object,omitempty"`

	// Google Cloud Storage generation for the object.
	Generation string `json:"generation,omitempty"`
}
