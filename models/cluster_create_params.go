// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterCreateParams cluster create params
//
// swagger:model cluster-create-params
type ClusterCreateParams struct {

	// Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.
	BaseDNSDomain string `json:"base_dns_domain,omitempty"`

	// IP address block from which Pod IPs are allocated This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$
	ClusterNetworkCidr *string `json:"cluster_network_cidr,omitempty"`

	// The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.
	// Maximum: 32
	// Minimum: 1
	ClusterNetworkHostPrefix int64 `json:"cluster_network_host_prefix,omitempty"`

	// Virtual IP used for cluster ingress traffic.
	// Pattern: ^(([0-9]{1,3}\.){3}[0-9]{1,3})?$
	IngressVip string `json:"ingress_vip,omitempty"`

	// Name of the OpenShift cluster.
	// Required: true
	Name *string `json:"name"`

	// Version of the OpenShift cluster.
	// Required: true
	// Enum: [4.5]
	OpenshiftVersion *string `json:"openshift_version"`

	// The pull secret that obtained from the Pull Secret page on the Red Hat OpenShift Cluster Manager site.
	PullSecret string `json:"pull_secret,omitempty"`

	// The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$
	ServiceNetworkCidr *string `json:"service_network_cidr,omitempty"`

	// SSH public key for debugging OpenShift nodes.
	SSHPublicKey string `json:"ssh_public_key,omitempty"`
}

// Validate validates this cluster create params
func (m *ClusterCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkHostPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenshiftVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("cluster_network_cidr", "body", string(*m.ClusterNetworkCidr), `^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkHostPrefix(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkHostPrefix) { // not required
		return nil
	}

	if err := validate.MinimumInt("cluster_network_host_prefix", "body", int64(m.ClusterNetworkHostPrefix), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cluster_network_host_prefix", "body", int64(m.ClusterNetworkHostPrefix), 32, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateIngressVip(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressVip) { // not required
		return nil
	}

	if err := validate.Pattern("ingress_vip", "body", string(m.IngressVip), `^(([0-9]{1,3}\.){3}[0-9]{1,3})?$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var clusterCreateParamsTypeOpenshiftVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["4.5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterCreateParamsTypeOpenshiftVersionPropEnum = append(clusterCreateParamsTypeOpenshiftVersionPropEnum, v)
	}
}

const (

	// ClusterCreateParamsOpenshiftVersionNr45 captures enum value "4.5"
	ClusterCreateParamsOpenshiftVersionNr45 string = "4.5"
)

// prop value enum
func (m *ClusterCreateParams) validateOpenshiftVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterCreateParamsTypeOpenshiftVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterCreateParams) validateOpenshiftVersion(formats strfmt.Registry) error {

	if err := validate.Required("openshift_version", "body", m.OpenshiftVersion); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpenshiftVersionEnum("openshift_version", "body", *m.OpenshiftVersion); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateServiceNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("service_network_cidr", "body", string(*m.ServiceNetworkCidr), `^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreateParams) UnmarshalBinary(b []byte) error {
	var res ClusterCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
