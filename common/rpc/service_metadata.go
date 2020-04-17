package rpc

import "strings"

type ServiceMetadata struct {
	ServiceKey           string
	ServiceInterfaceName string
	Version              string
	Group                string
	Attachments          map[string][]string
	Attributes           map[string]interface{}
}

func (s *ServiceMetadata) BuildServiceKey() {
	s.ServiceKey = buildServiceKey(s.ServiceInterfaceName, s.Group, s.Version)
}

func (s *ServiceMetadata) GetAttachment(key string) string {
	if s.Attachments == nil {
		return ""
	}
	vs := s.Attachments[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func (s *ServiceMetadata) GetAttachments(key string) []string {
	if s.Attachments == nil || len(s.Attachments) == 0 {
		return []string{}
	}
	vs := s.Attachments[key]
	return vs
}

// Set sets the key to value. It replaces any existing
// values.
func (s *ServiceMetadata) SetAttachment(key, value string) {
	s.Attachments[key] = []string{value}
}

func (s *ServiceMetadata) AddAttachments(maps map[string][]string) {
	for k, v := range maps {
		s.Attachments[k] = v
	}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (s *ServiceMetadata) AddAttachment(key, value string) {
	s.Attachments[key] = append(s.Attachments[key], value)
}

// Del deletes the values associated with key.
func (s *ServiceMetadata) DelAttachment(key string) {
	delete(s.Attachments, key)
}

func buildServiceKey(path string, group string, version string) string {
	var b strings.Builder
	if len(group) > 0 {
		b.WriteString(group)
		b.WriteString("/")
	}
	b.WriteString(path)
	if len(version) > 0 {
		b.WriteString(":")
		b.WriteString(version)
	}
	return b.String()
}
