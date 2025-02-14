//go:build !windows && !linux

package uspfilter

// Reset firewall to the default state
func (m *Manager) Reset() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.outgoingRules = make(map[string]RuleSet)
	m.incomingRules = make(map[string]RuleSet)

	return nil
}

// AllowNetbird allows netbird interface traffic
func (m *Manager) AllowNetbird() error {
	return nil
}
