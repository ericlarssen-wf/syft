package pkg

import (
	"github.com/anchore/packageurl-go"
	"github.com/anchore/syft/syft/linux"
)

type PubMetadata struct {
	Name        string   `mapstructure:"name" json:"name"`
	Version     string   `mapstructure:"version" json:"version"`
	Description string   `mapstructure:"description" json:"description,omitempty"`
	Authors     []string `mapstructure:"authors" json:"authors,omitempty"`
	Homepage    string   `mapstructure:"homepage" json:"homepage,omitempty"`
	Repository  string   `mapstructure:"repository" json:"repository,omitempty"`
	HostedURL   string   `mapstructure:"hosted_url" json:"hosted_url,omitempty"`
	VcsURL      string   `mapstructure:"vcs_url" json:"vcs_url,omitempty"`
}

func (m PubMetadata) PackageURL(_ *linux.Release) string {
	var qualifiers packageurl.Qualifiers

	if m.HostedURL != "" {
		qualifiers = append(qualifiers, packageurl.Qualifier{
			Key:   "hosted_url",
			Value: m.HostedURL,
		})
	} else if m.VcsURL != "" { // Default to using Hosted if somehow both are provided
		qualifiers = append(qualifiers, packageurl.Qualifier{
			Key:   "vcs_url",
			Value: m.VcsURL,
		})
	}

	return packageurl.NewPackageURL(
		packageurl.TypePub,
		"",
		m.Name,
		m.Version,
		qualifiers,
		"",
	).ToString()
}
