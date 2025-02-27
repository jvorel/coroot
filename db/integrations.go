package db

import (
	"fmt"
	"github.com/coroot/coroot/timeseries"
)

type IntegrationType string

const (
	IntegrationTypePrometheus IntegrationType = "prometheus"
	IntegrationTypePyroscope  IntegrationType = "pyroscope"
	IntegrationTypeSlack      IntegrationType = "slack"
	IntegrationTypePagerduty  IntegrationType = "pagerduty"
	IntegrationTypeTeams      IntegrationType = "teams"
	IntegrationTypeOpsgenie   IntegrationType = "opsgenie"
)

type Integrations struct {
	BaseUrl string `json:"base_url"`

	Slack     *IntegrationSlack     `json:"slack,omitempty"`
	Pagerduty *IntegrationPagerduty `json:"pagerduty,omitempty"`
	Teams     *IntegrationTeams     `json:"teams,omitempty"`
	Opsgenie  *IntegrationOpsgenie  `json:"opsgenie,omitempty"`

	Pyroscope *IntegrationPyroscope `json:"pyroscope,omitempty"`
}

type IntegrationInfo struct {
	Type        IntegrationType
	Configured  bool
	Incidents   bool
	Deployments bool
	Title       string
	Details     string
}

func (integrations Integrations) GetInfo() []IntegrationInfo {
	var res []IntegrationInfo

	i := IntegrationInfo{Type: IntegrationTypeSlack, Title: "Slack"}
	if cfg := integrations.Slack; cfg != nil {
		i.Configured = true
		i.Incidents = cfg.Incidents
		i.Deployments = cfg.Deployments
		i.Details = fmt.Sprintf("channel: #%s", cfg.DefaultChannel)
	}
	res = append(res, i)

	i = IntegrationInfo{Type: IntegrationTypeTeams, Title: "MS Teams"}
	if cfg := integrations.Teams; cfg != nil {
		i.Configured = true
		i.Incidents = cfg.Incidents
		i.Deployments = cfg.Deployments
	}
	res = append(res, i)

	i = IntegrationInfo{Type: IntegrationTypePagerduty, Title: "Pagerduty"}
	if cfg := integrations.Pagerduty; cfg != nil {
		i.Configured = true
		i.Incidents = cfg.Incidents
	}
	res = append(res, i)

	i = IntegrationInfo{Type: IntegrationTypeOpsgenie, Title: "Opsgenie"}
	if cfg := integrations.Opsgenie; cfg != nil {
		i.Configured = true
		i.Incidents = cfg.Incidents
		region := "US"
		if cfg.EUInstance {
			region = "EU"
		}
		i.Details = fmt.Sprintf("region: %s", region)
	}
	res = append(res, i)

	return res
}

type IntegrationsPrometheus struct {
	Url             string              `json:"url"`
	RefreshInterval timeseries.Duration `json:"refresh_interval"`
	TlsSkipVerify   bool                `json:"tls_skip_verify"`
	BasicAuth       *BasicAuth          `json:"basic_auth"`
	ExtraSelector   string              `json:"extra_selector"`
}

type IntegrationPyroscope struct {
	Url           string     `json:"url"`
	TlsSkipVerify bool       `json:"tls_skip_verify"`
	BasicAuth     *BasicAuth `json:"basic_auth,omitempty"`
	ApiKey        string     `json:"api_key"`
}

type IntegrationSlack struct {
	Token          string `json:"token"`
	DefaultChannel string `json:"default_channel"`
	Enabled        bool   `json:"enabled"` // deprecated: use Incidents and Deployments
	Incidents      bool   `json:"incidents"`
	Deployments    bool   `json:"deployments"`
}

type IntegrationTeams struct {
	WebhookUrl  string `json:"webhook_url"`
	Incidents   bool   `json:"incidents"`
	Deployments bool   `json:"deployments"`
}

type IntegrationPagerduty struct {
	IntegrationKey string `json:"integration_key"`
	Incidents      bool   `json:"incidents"`
}

type IntegrationOpsgenie struct {
	ApiKey     string `json:"api_key"`
	EUInstance bool   `json:"eu_instance"`
	Incidents  bool   `json:"incidents"`
}

type BasicAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (db *DB) SaveIntegrationsBaseUrl(id ProjectId, baseUrl string) error {
	p, err := db.GetProject(id)
	if err != nil {
		return err
	}
	p.Settings.Integrations.BaseUrl = baseUrl
	return db.saveProjectSettings(p)
}
