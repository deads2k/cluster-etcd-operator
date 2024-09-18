// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AWSCSIDriverConfigSpecApplyConfiguration represents an declarative configuration of the AWSCSIDriverConfigSpec type for use
// with apply.
type AWSCSIDriverConfigSpecApplyConfiguration struct {
	KMSKeyARN        *string                                `json:"kmsKeyARN,omitempty"`
	EFSVolumeMetrics *AWSEFSVolumeMetricsApplyConfiguration `json:"efsVolumeMetrics,omitempty"`
}

// AWSCSIDriverConfigSpecApplyConfiguration constructs an declarative configuration of the AWSCSIDriverConfigSpec type for use with
// apply.
func AWSCSIDriverConfigSpec() *AWSCSIDriverConfigSpecApplyConfiguration {
	return &AWSCSIDriverConfigSpecApplyConfiguration{}
}

// WithKMSKeyARN sets the KMSKeyARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KMSKeyARN field is set to the value of the last call.
func (b *AWSCSIDriverConfigSpecApplyConfiguration) WithKMSKeyARN(value string) *AWSCSIDriverConfigSpecApplyConfiguration {
	b.KMSKeyARN = &value
	return b
}

// WithEFSVolumeMetrics sets the EFSVolumeMetrics field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EFSVolumeMetrics field is set to the value of the last call.
func (b *AWSCSIDriverConfigSpecApplyConfiguration) WithEFSVolumeMetrics(value *AWSEFSVolumeMetricsApplyConfiguration) *AWSCSIDriverConfigSpecApplyConfiguration {
	b.EFSVolumeMetrics = value
	return b
}
