package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

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

		if context.Endpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_ADDR=%s", context.Endpoint))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_ADDR")
		}

		if context.Token != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_TOKEN=%s", context.Token))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_TOKEN")
		}

		if context.Namespace != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_NAMESPACE=%s", context.Namespace))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_NAMESPACE")
		}

		if context.CaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CACERT=%s", context.CaCert))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CACERT")
		}

		if context.Cert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_CERT=%s", context.Cert))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CLIENT_CERT")
		}

		if context.CaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CAPATH=%s", context.CaPath))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CAPATH")
		}

		if context.Key != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_KEY=%s", context.Key))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CLIENT_KEY")
		}

		if context.Format != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_FORMAT=%s", context.Format))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_FORMAT")
		}

		if context.SkipVerify != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_SKIP_VERIFY=%s", context.SkipVerify))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_SKIP_VERIFY")
		}

		if context.ClientTimeout != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_TIMEOUT=%s", context.ClientTimeout))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CLIENT_TIMEOUT")
		}

		if context.ClusterAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLUSTER_ADDR=%s", context.ClusterAddr))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CLUSTER_ADDR")
		}

		if context.License != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LICENSE=%s", context.License))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_LICENSE")
		}

		if context.LicensePath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LICENSE_PATH=%s", context.LicensePath))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_LICENSE_PATH")
		}

		if context.LogLevel != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LOG_LEVEL=%s", context.LogLevel))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_LOG_LEVEL")
		}

		if context.MaxRetries != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_MAX_RETRIES=%s", context.MaxRetries))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_MAX_RETRIES")
		}

		if context.RedirectAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_REDIRECT_ADDR=%s", context.RedirectAddr))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_REDIRECT_ADDR")
		}

		if context.TlsServerName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_TLS_SERVER_NAME=%s", context.TlsServerName))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_TLS_SERVER_NAME")
		}

		if context.CliNoColour != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLI_NO_COLOR=%s", context.CliNoColour))
		} else if context.NoColor {
			exportCommands = append(exportCommands, "export VAULT_CLI_NO_COLOR=true")
		} else {
			exportCommands = append(exportCommands, "unset VAULT_CLI_NO_COLOR")
		}

		if context.RateLimit != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_RATE_LIMIT=%s", context.RateLimit))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_RATE_LIMIT")
		}

		if context.SvrLookup != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_SVR_LOOKUP=%s", context.SvrLookup))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_SVR_LOOKUP")
		}

		if context.Mfa != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_MFA=%s", context.Mfa))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_MFA")
		}

		if context.HttpProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_HTTP_PROXY=%s", context.HttpProxy))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_HTTP_PROXY")
		}

		if context.HttpsProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export HTTPS_PROXY=%s", context.HttpsProxy))
		} else {
			exportCommands = append(exportCommands, "unset HTTPS_PROXY")
		}

		if context.DisableRedirects != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_DISABLE_REDIRECTS=%s", context.DisableRedirects))
		} else {
			exportCommands = append(exportCommands, "unset VAULT_DISABLE_REDIRECTS")
		}

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

		if context.NomadEndpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_ADDR=%s", context.NomadEndpoint))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_ADDR")
		}

		if context.NomadToken != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_TOKEN=%s", context.NomadToken))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_TOKEN")
		}

		if context.NomadNamespace != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_NAMESPACE=%s", context.NomadNamespace))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_NAMESPACE")
		}

		if context.NomadCaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CACERT=%s", context.NomadCaCert))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_CACERT")
		}

		if context.NomadCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CLIENT_CERT=%s", context.NomadCert))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_CLIENT_CERT")
		}

		if context.NomadCaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CAPATH=%s", context.NomadCaPath))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_CAPATH")
		}

		if context.NomadKey != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CLIENT_KEY=%s", context.NomadKey))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_CLIENT_KEY")
		}

		if context.NomadRegion != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_REGION=%s", context.NomadRegion))
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_REGION")
		}

		if context.NomadHttpProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export HTTPS_PROXY=%s", context.NomadHttpProxy))
			exportCommands = append(exportCommands, fmt.Sprintf("export HTTP_PROXY=%s", context.NomadHttpProxy))
		} else {
			exportCommands = append(exportCommands, "unset HTTPS_PROXY")
			exportCommands = append(exportCommands, "unset HTTP_PROXY")
		}

		if context.NomadSkipVerify {
			exportCommands = append(exportCommands, "export NOMAD_SKIP_VERIFY=true")
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_SKIP_VERIFY")
		}

		if context.NomadDisableCliHints {
			exportCommands = append(exportCommands, "export NOMAD_CLI_SHOW_HINTS=false")
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_CLI_SHOW_HINTS")
		}

		if context.NomadNoColor {
			exportCommands = append(exportCommands, "export NOMAD_CLI_NO_COLOR=true")
		} else {
			exportCommands = append(exportCommands, "unset NOMAD_CLI_NO_COLOR")
		}

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

		if context.ConsulEndpoint != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_HTTP_ADDR=%s", context.ConsulEndpoint))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_HTTP_ADDR")
		}

		if context.ConsulToken != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_HTTP_TOKEN=%s", context.ConsulToken))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_HTTP_TOKEN")
		}

		if context.ConsulTokenFile != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_HTTP_TOKEN_FILE=%s", context.ConsulTokenFile))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_HTTP_TOKEN_FILE")
		}

		if context.ConsulCaCert != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CACERT=%s", context.ConsulCaCert))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_CACERT")
		}

		if context.ConsulCert != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CLIENT_CERT=%s", context.ConsulCert))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_CLIENT_CERT")
		}

		if context.ConsulCaPath != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CAPATH=%s", context.ConsulCaPath))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_CAPATH")
		}

		if context.ConsulKey != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_CLIENT_KEY=%s", context.ConsulKey))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_CLIENT_KEY")
		}

		if context.ConsulNamespace != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export CONSUL_NAMESPACE=%s", context.ConsulNamespace))
		} else {
			exportCommandStr = append(exportCommandStr, "unset CONSUL_NAMESPACE")
		}

		if context.ConsulHttpProxy != "" {
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export HTTPS_PROXY=%s", context.ConsulHttpProxy))
			exportCommandStr = append(exportCommandStr, fmt.Sprintf("export HTTP_PROXY=%s", context.ConsulHttpProxy))
		} else {
			exportCommandStr = append(exportCommandStr, "unset HTTPS_PROXY")
			exportCommandStr = append(exportCommandStr, "unset HTTP_PROXY")
		}

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

		if context.Endpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_ADDR=%s", context.Endpoint))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_ADDR")
		}

		// NOTE: users can use either a "BOUNARY_TOKEN" (for a file on disk) or
		// a "BOUNDARY_TOKEN_NAME" (for platform-specific OS credential store).

		if context.Token != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TOKEN=%s", context.Token))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_TOKEN")
		}

		if context.TokenName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TOKEN_NAME=%s", context.TokenName))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_TOKEN_NAME")
		}

		if context.CaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CACERT=%s", context.CaCert))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CACERT")
		}

		if context.Cert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLIENT_CERT=%s", context.Cert))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CLIENT_CERT")
		}

		if context.CaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CAPATH=%s", context.CaPath))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CAPATH")
		}

		if context.Key != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLIENT_KEY=%s", context.Key))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CLIENT_KEY")
		}

		if context.TlsInsecure != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TLS_INSECURE=%s", context.TlsInsecure))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_TLS_INSECURE")
		}

		if context.TlsServerName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TLS_SERVER_NAME=%s", context.TlsServerName))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_TLS_SERVER_NAME")
		}

		if context.RecoveryConfig != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_RECOVERY_CONFIG=%s", context.RecoveryConfig))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_RECOVERY_CONFIG")
		}

		if context.ConnectAuthZToken != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_AUTHZ_TOKEN=%s", context.ConnectAuthZToken))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CONNECT_AUTHZ_TOKEN")
		}

		if context.ConnectExec != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_EXEC=%s", context.ConnectExec))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CONNECT_EXEC")
		}

		if context.ConnectListenAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_ADDR=%s", context.ConnectListenAddr))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CONNECT_LISTEN_ADDR")
		}

		if context.ConnectListenPort != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_PORT=%s", context.ConnectListenPort))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CONNECT_LISTEN_PORT")
		}

		if context.ConnectTargetScopeId != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_ID=%s", context.ConnectTargetScopeId))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CONNECT_TARGET_SCOPE_ID")
		}

		if context.ConnectTargetScopeName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_NAME=%s", context.ConnectTargetScopeName))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CONNECT_TARGET_SCOPE_NAME")
		}

		if context.AuthMethodId != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_AUTH_METHOD_ID=%s", context.AuthMethodId))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_AUTH_METHOD_ID")
		}

		if context.LogLevel != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_LOG_LEVEL=%s", context.LogLevel))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_LOG_LEVEL")
		}

		if context.Format != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLI_FORMAT=%s", context.Format))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_CLI_FORMAT")
		}

		if context.ScopeId != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_SCOPE_ID=%s", context.ScopeId))
		} else {
			exportCommands = append(exportCommands, "unset BOUNDARY_SCOPE_ID")
		}

		if context.HttpProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export HTTPS_PROXY=%s", context.HttpProxy))
		} else {
			exportCommands = append(exportCommands, "unset HTTPS_PROXY")
		}

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

		commandStr := strings.Join(exportCommands, "; ")
		fmt.Println(commandStr)
	},
}
