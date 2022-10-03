// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// APIServerServingCertsApplyConfiguration represents an declarative configuration of the APIServerServingCerts type for use
// with apply.
type APIServerServingCertsApplyConfiguration struct {
	NamedCertificates []APIServerNamedServingCertApplyConfiguration `json:"namedCertificates,omitempty"`
}

// APIServerServingCertsApplyConfiguration constructs an declarative configuration of the APIServerServingCerts type for use with
// apply.
func APIServerServingCerts() *APIServerServingCertsApplyConfiguration {
	return &APIServerServingCertsApplyConfiguration{}
}

// WithNamedCertificates adds the given value to the NamedCertificates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NamedCertificates field.
func (b *APIServerServingCertsApplyConfiguration) WithNamedCertificates(values ...*APIServerNamedServingCertApplyConfiguration) *APIServerServingCertsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNamedCertificates")
		}
		b.NamedCertificates = append(b.NamedCertificates, *values[i])
	}
	return b
}
