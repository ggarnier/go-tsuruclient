/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type ServiceInstanceBoundUnit struct {
	AppName string `json:"app_name,omitempty"`

	Id string `json:"id,omitempty"`

	Ip string `json:"ip,omitempty"`
}