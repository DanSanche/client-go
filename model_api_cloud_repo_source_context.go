/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A CloudRepoSourceContext denotes a particular revision in a Google Cloud Source Repo.
type ApiCloudRepoSourceContext struct {
	// The ID of the repo.
	RepoId *ApiRepoId `json:"repo_id,omitempty"`

	// A revision ID.
	RevisionId string `json:"revision_id,omitempty"`

	// An alias, which may be a branch or tag.
	AliasContext *ApiAliasContext `json:"alias_context,omitempty"`
}
