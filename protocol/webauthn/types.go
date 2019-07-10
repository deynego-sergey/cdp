// Code generated by cdpgen. DO NOT EDIT.

package webauthn

// AuthenticatorID
type AuthenticatorID string

// AuthenticatorProtocol
type AuthenticatorProtocol string

// AuthenticatorProtocol as enums.
const (
	AuthenticatorProtocolNotSet AuthenticatorProtocol = ""
	AuthenticatorProtocolU2F    AuthenticatorProtocol = "u2f"
	AuthenticatorProtocolCTAP2  AuthenticatorProtocol = "ctap2"
)

func (e AuthenticatorProtocol) Valid() bool {
	switch e {
	case "u2f", "ctap2":
		return true
	default:
		return false
	}
}

func (e AuthenticatorProtocol) String() string {
	return string(e)
}

// AuthenticatorTransport
type AuthenticatorTransport string

// AuthenticatorTransport as enums.
const (
	AuthenticatorTransportNotSet   AuthenticatorTransport = ""
	AuthenticatorTransportUSB      AuthenticatorTransport = "usb"
	AuthenticatorTransportNFC      AuthenticatorTransport = "nfc"
	AuthenticatorTransportBLE      AuthenticatorTransport = "ble"
	AuthenticatorTransportCable    AuthenticatorTransport = "cable"
	AuthenticatorTransportInternal AuthenticatorTransport = "internal"
)

func (e AuthenticatorTransport) Valid() bool {
	switch e {
	case "usb", "nfc", "ble", "cable", "internal":
		return true
	default:
		return false
	}
}

func (e AuthenticatorTransport) String() string {
	return string(e)
}

// VirtualAuthenticatorOptions
type VirtualAuthenticatorOptions struct {
	Protocol                    AuthenticatorProtocol  `json:"protocol"`                              // No description.
	Transport                   AuthenticatorTransport `json:"transport"`                             // No description.
	HasResidentKey              bool                   `json:"hasResidentKey"`                        // No description.
	HasUserVerification         bool                   `json:"hasUserVerification"`                   // No description.
	AutomaticPresenceSimulation *bool                  `json:"automaticPresenceSimulation,omitempty"` // If set to true, tests of user presence will succeed immediately. Otherwise, they will not be resolved. Defaults to true.
}

// Credential
type Credential struct {
	CredentialID string `json:"credentialId"` // No description.
	RPIDHash     string `json:"rpIdHash"`     // SHA-256 hash of the Relying Party ID the credential is scoped to. Must be 32 bytes long. See https://w3c.github.io/webauthn/#rpidhash
	PrivateKey   string `json:"privateKey"`   // The private key in PKCS#8 format.
	SignCount    int    `json:"signCount"`    // Signature counter. This is incremented by one for each successful assertion. See https://w3c.github.io/webauthn/#signature-counter
}
