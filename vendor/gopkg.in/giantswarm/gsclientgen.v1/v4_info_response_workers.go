/*
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * OpenAPI spec version: legacy
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package gsclientgen

// Information related to worker nodes
type V4InfoResponseWorkers struct {
	CountPerCluster V4InfoResponseWorkersCountPerCluster `json:"count_per_cluster,omitempty"`

	InstanceType V4InfoResponseWorkersInstanceType `json:"instance_type,omitempty"`

	VmSize V4InfoResponseWorkersVmSize `json:"vm_size,omitempty"`
}