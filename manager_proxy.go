package protonmail

// AllowProxy allows the client manager to switch clients over to a proxy if need be.
func (m *manager) AllowProxy() {
	if m.proxyDialer != nil {
		m.proxyDialer.AllowProxy()
	}
}

// DisallowProxy prevents the client manager from switching clients over to a proxy if need be.
func (m *manager) DisallowProxy() {
	if m.proxyDialer != nil {
		m.proxyDialer.DisallowProxy()
	}
}
