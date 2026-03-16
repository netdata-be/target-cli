package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func appendEnvOrUnset(exportCommands *[]string, key, value string) {
	if value != "" {
		*exportCommands = append(*exportCommands, fmt.Sprintf("export %s=%s", key, value))
	} else {
		*exportCommands = append(*exportCommands, fmt.Sprintf("unset %s", key))
	}
}

func appendBoolEnvOrUnset(exportCommands *[]string, key string, value bool) {
	if value {
		*exportCommands = append(*exportCommands, fmt.Sprintf("export %s=true", key))
	} else {
		*exportCommands = append(*exportCommands, fmt.Sprintf("unset %s", key))
	}
}


// selectVaultCmd represents the switch command for Vault
var selectVaultCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target vault select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Vault[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Vault[args[0]]

		exportCommands := []string{}
		appendEnvOrUnset(&exportCommands, "VAULT_ADDR", context.Endpoint)
		appendEnvOrUnset(&exportCommands, "VAULT_TOKEN", context.Token)
		appendEnvOrUnset(&exportCommands, "VAULT_NAMESPACE", context.Namespace)
		appendEnvOrUnset(&exportCommands, "VAULT_CACERT", context.CaCert)
		appendEnvOrUnset(&exportCommands, "VAULT_CLIENT_CERT", context.Cert)
		appendEnvOrUnset(&exportCommands, "VAULT_CAPATH", context.CaPath)
		appendEnvOrUnset(&exportCommands, "VAULT_CLIENT_KEY", context.Key)
		appendEnvOrUnset(&exportCommands, "VAULT_FORMAT", context.Format)
		appendEnvOrUnset(&exportCommands, "VAULT_SKIP_VERIFY", context.SkipVerify)
		appendEnvOrUnset(&exportCommands, "VAULT_CLIENT_TIMEOUT", context.ClientTimeout)
		appendEnvOrUnset(&exportCommands, "VAULT_CLUSTER_ADDR", context.ClusterAddr)
		appendEnvOrUnset(&exportCommands, "VAULT_LICENSE", context.License)
		appendEnvOrUnset(&exportCommands, "VAULT_LICENSE_PATH", context.LicensePath)
		appendEnvOrUnset(&exportCommands, "VAULT_LOG_LEVEL", context.LogLevel)
		appendEnvOrUnset(&exportCommands, "VAULT_MAX_RETRIES", context.MaxRetries)
		appendEnvOrUnset(&exportCommands, "VAULT_REDIRECT_ADDR", context.RedirectAddr)
		appendEnvOrUnset(&exportCommands, "VAULT_TLS_SERVER_NAME", context.TlsServerName)
		if context.CliNoColour != "" {
			appendEnvOrUnset(&exportCommands, "VAULT_CLI_NO_COLOR", context.CliNoColour)
		} else {
			appendBoolEnvOrUnset(&exportCommands, "VAULT_CLI_NO_COLOR", context.NoColor)
		}
		appendEnvOrUnset(&exportCommands, "VAULT_RATE_LIMIT", context.RateLimit)
		appendEnvOrUnset(&exportCommands, "VAULT_SRV_LOOKUP", context.SvrLookup)
		appendEnvOrUnset(&exportCommands, "VAULT_MFA", context.Mfa)
		appendEnvOrUnset(&exportCommands, "VAULT_HTTP_PROXY", context.HttpProxy)
		appendEnvOrUnset(&exportCommands, "HTTPS_PROXY", context.HttpsProxy)
		appendEnvOrUnset(&exportCommands, "VAULT_DISABLE_REDIRECTS", context.DisableRedirects)
		exportCommands = append(exportCommands, fmt.Sprintf("export TARGET_VAULT_PROFILE=%s", profile))
		fmt.Println(strings.Join(exportCommands, "; "))
	},
}

var selectNomadCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target nomad select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Nomad[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Nomad[args[0]]

		exportCommands := []string{}
		appendEnvOrUnset(&exportCommands, "NOMAD_ADDR", context.NomadEndpoint)
		appendEnvOrUnset(&exportCommands, "NOMAD_TOKEN", context.NomadToken)
		appendEnvOrUnset(&exportCommands, "NOMAD_NAMESPACE", context.NomadNamespace)
		appendEnvOrUnset(&exportCommands, "NOMAD_CACERT", context.NomadCaCert)
		appendEnvOrUnset(&exportCommands, "NOMAD_CLIENT_CERT", context.NomadCert)
		appendEnvOrUnset(&exportCommands, "NOMAD_CAPATH", context.NomadCaPath)
		appendEnvOrUnset(&exportCommands, "NOMAD_CLIENT_KEY", context.NomadKey)
		appendEnvOrUnset(&exportCommands, "NOMAD_REGION", context.NomadRegion)
		appendEnvOrUnset(&exportCommands, "HTTPS_PROXY", context.NomadHttpProxy)
		appendEnvOrUnset(&exportCommands, "HTTP_PROXY", context.NomadHttpProxy)
		appendBoolEnvOrUnset(&exportCommands, "NOMAD_SKIP_VERIFY", context.NomadSkipVerify)
		if context.NomadDisableCliHints {
			appendEnvOrUnset(&exportCommands, "NOMAD_CLI_SHOW_HINTS", "false")
		} else {
			appendEnvOrUnset(&exportCommands, "NOMAD_CLI_SHOW_HINTS", "")
		}
		appendBoolEnvOrUnset(&exportCommands, "NOMAD_CLI_NO_COLOR", context.NomadNoColor)
		exportCommands = append(exportCommands, fmt.Sprintf("export TARGET_NOMAD_PROFILE=%s", profile))
		fmt.Println(strings.Join(exportCommands, "; "))
	},
}

var selectConsulCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target consul select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Consul[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Consul[args[0]]

		exportCommandStr := []string{}
		appendEnvOrUnset(&exportCommandStr, "CONSUL_HTTP_ADDR", context.ConsulEndpoint)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_HTTP_TOKEN", context.ConsulToken)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_HTTP_TOKEN_FILE", context.ConsulTokenFile)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_CACERT", context.ConsulCaCert)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_CLIENT_CERT", context.ConsulCert)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_CAPATH", context.ConsulCaPath)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_CLIENT_KEY", context.ConsulKey)
		appendEnvOrUnset(&exportCommandStr, "CONSUL_NAMESPACE", context.ConsulNamespace)
		appendEnvOrUnset(&exportCommandStr, "HTTPS_PROXY", context.ConsulHttpProxy)
		appendEnvOrUnset(&exportCommandStr, "HTTP_PROXY", context.ConsulHttpProxy)
		exportCommandStr = append(exportCommandStr, fmt.Sprintf("export TARGET_CONSUL_PROFILE=%s", profile))
		fmt.Println(strings.Join(exportCommandStr, "; "))
	},
}

var selectBoundaryCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target boundary select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Boundary[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Boundary[args[0]]

		exportCommands := []string{}
		appendEnvOrUnset(&exportCommands, "BOUNDARY_ADDR", context.Endpoint)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_TOKEN", context.Token)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_TOKEN_NAME", context.TokenName)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CACERT", context.CaCert)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CLIENT_CERT", context.Cert)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CAPATH", context.CaPath)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CLIENT_KEY", context.Key)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_TLS_INSECURE", context.TlsInsecure)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_TLS_SERVER_NAME", context.TlsServerName)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_RECOVERY_CONFIG", context.RecoveryConfig)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CONNECT_AUTHZ_TOKEN", context.ConnectAuthZToken)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CONNECT_EXEC", context.ConnectExec)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CONNECT_LISTEN_ADDR", context.ConnectListenAddr)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CONNECT_LISTEN_PORT", context.ConnectListenPort)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CONNECT_TARGET_SCOPE_ID", context.ConnectTargetScopeId)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CONNECT_TARGET_SCOPE_NAME", context.ConnectTargetScopeName)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_AUTH_METHOD_ID", context.AuthMethodId)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_LOG_LEVEL", context.LogLevel)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_CLI_FORMAT", context.Format)
		appendEnvOrUnset(&exportCommands, "BOUNDARY_SCOPE_ID", context.ScopeId)
		appendEnvOrUnset(&exportCommands, "HTTPS_PROXY", context.HttpProxy)
		exportCommands = append(exportCommands, fmt.Sprintf("export TARGET_BOUNDARY_PROFILE=%s", profile))
		fmt.Println(strings.Join(exportCommands, "; "))
	},
}

var selectTerraformCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `target terraform select example"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.Terraform[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.Terraform[args[0]]

		exportCommands := []string{}

		for k, v := range context.Vars {
			exportCommands = append(exportCommands, fmt.Sprintf("export TF_VAR_%s=%s", k, v))
		}

		if context.HttpProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export HTTPS_PROXY=%s", context.HttpProxy))
		}

		exportCommands = append(exportCommands, fmt.Sprintf("export TARGET_TERRAFORM_PROFILE=%s", profile))
		commandStr := strings.Join(exportCommands, "; ")
		fmt.Println(commandStr)
	},
}
