package cmd

import (
	"errors"
	"fmt"
	"github.com/devops-rob/target-cli/pkg/targetdir"
	"os"
	"reflect"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

//var cfgFile = ""

var version string = "dev"

// Config struct containing different product profiles
type Config struct {
	Vault     map[string]*Vault     `json:"vault,omitempty" mapstructure:"vault"`
	Consul    map[string]*Consul    `json:"consul,omitempty" mapstructure:"consul"`
	Nomad     map[string]*Nomad     `json:"nomad,omitempty" mapstructure:"nomad"`
	Boundary  map[string]*Boundary  `json:"boundary,omitempty" mapstructure:"boundary"`
	Terraform map[string]*Terraform `json:"terraform,omitempty" mapstructure:"terraform"`
	//Default  map[string]*Default  `json:"default,omitempty" mapstructure:"default"`
}

type Terraform struct {
	Vars      map[string]string `json:"vars,omitempty" mapstructure:"vars"`
	HttpProxy string            `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
}

type Boundary struct {
	Endpoint               string `json:"endpoint,omitempty" mapstructure:"endpoint"`
	Token                  string `json:"token,omitempty" mapstructure:"token"`
	TokenName              string `json:"token_name,omitempty" mapstructure:"token_name"`
	CaPath                 string `json:"ca_path,omitempty" mapstructure:"ca_path"`
	CaCert                 string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	Cert                   string `json:"cert,omitempty" mapstructure:"cert"`
	Key                    string `json:"key,omitempty" mapstructure:"key"`
	TlsInsecure            string `json:"tls_insecure,omitempty" mapstructure:"tls_insecure"`
	TlsServerName          string `json:"tls_server_name,omitempty" mapstructure:"tls_server_name"`
	RecoveryConfig         string `json:"recovery_config,omitempty" mapstructure:"recovery_config"`
	ConnectAuthZToken      string `json:"connect_auth_z_token,omitempty" mapstructure:"connect_auth_z_token"`
	ConnectExec            string `json:"connect_exec,omitempty" mapstructure:"connect_exec"`
	ConnectListenAddr      string `json:"connect_listen_addr,omitempty" mapstructure:"connect_listen_addr"`
	ConnectListenPort      string `json:"connect_listen_port,omitempty" mapstructure:"connect_listen_port"`
	ConnectTargetScopeId   string `json:"connect_target_scope_id,omitempty" mapstructure:"connect_target_scope_id"`
	ConnectTargetScopeName string `json:"connect_target_scope_name,omitempty" mapstructure:"connect_target_scope_name"`
	AuthMethodId           string `json:"auth_method_id,omitempty" mapstructure:"auth_method_id"`
	LogLevel               string `json:"log_level,omitempty" mapstructure:"log_level"`
	Format                 string `json:"format,omitempty" mapstructure:"format"`
	ScopeId                string `json:"scope_id,omitempty" mapstructure:"scope_id"`
	HttpProxy              string `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
}

// Vault struct with flag parameters
type Vault struct {
	Endpoint         string `json:"endpoint,omitempty" mapstructure:"endpoint"`
	Token            string `json:"token,omitempty" mapstructure:"token"`
	CaPath           string `json:"ca_path,omitempty" mapstructure:"ca_path"`
	CaCert           string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	Cert             string `json:"cert,omitempty" mapstructure:"cert"`
	Key              string `json:"key,omitempty" mapstructure:"key"`
	Format           string `json:"format,omitempty" mapstructure:"format"`
	Namespace        string `json:"namespace,omitempty" mapstructure:"namespace"`
	SkipVerify       string `json:"skip_verify,omitempty" mapstructure:"skip_verify"`
	ClientTimeout    string `json:"client_timeout,omitempty" mapstructure:"client_timeout"`
	ClusterAddr      string `json:"cluster_addr,omitempty" mapstructure:"cluster_addr"`
	License          string `json:"license,omitempty" mapstructure:"license"`
	LicensePath      string `json:"license_path,omitempty" mapstructure:"license_path"`
	LogLevel         string `json:"log_level,omitempty" mapstructure:"log_level"`
	MaxRetries       string `json:"max_retries,omitempty" mapstructure:"max_retries"`
	RedirectAddr     string `json:"redirect_addr,omitempty" mapstructure:"redirect_addr"`
	TlsServerName    string `json:"tls_server_name,omitempty" mapstructure:"tls_server_name"`
	CliNoColour      string `json:"cli_no_colour,omitempty" mapstructure:"cli_no_colour"`
	RateLimit        string `json:"rate_limit,omitempty" mapstructure:"rate_limit"`
	SvrLookup        string `json:"svr_lookup,omitempty" mapstructure:"svr_lookup"`
	Mfa              string `json:"mfa,omitempty" mapstructure:"mfa"`
	HttpProxy        string `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
	HttpsProxy       string `json:"https_proxy,omitempty" mapstructure:"https_proxy"`
	DisableRedirects string `json:"disable_redirects,omitempty" mapstructure:"disable_redirects"`
	NoColor          bool   `json:"no_color,omitempty" mapstructure:"no_color"`
}

var (
	c *Config
)

// Consul struct with flag parameters
type Consul struct {
	ConsulEndpoint  string `json:"endpoint" mapstructure:"endpoint"`
	ConsulToken     string `json:"token,omitempty" mapstructure:"token"`
	ConsulCaPath    string `json:"ca_path,omitempty" mapstructure:"ca_path"`
	ConsulCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	ConsulCert      string `json:"cert,omitempty" mapstructure:"cert"`
	ConsulKey       string `json:"key,omitempty" mapstructure:"key"`
	ConsulTokenFile string `json:"token_file,omitempty" mapstructure:"token_file"`
	ConsulNamespace string `json:"namespace,omitempty" mapstructure:"namespace"`
	ConsulHttpProxy string `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
}

// Nomad struct with flag parameters
type Nomad struct {
	NomadEndpoint  string `json:"endpoint" mapstructure:"endpoint"`
	NomadToken     string `json:"token,omitempty" mapstructure:"token"`
	NomadCaPath    string `json:"ca_path,omitempty" mapstructure:"ca_path"`
	NomadCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	NomadCert      string `json:"cert,omitempty" mapstructure:"cert"`
	NomadKey       string `json:"key,omitempty" mapstructure:"key"`
	NomadRegion    string `json:"region,omitempty" mapstructure:"region"`
	NomadNamespace  string `json:"namespace,omitempty" mapstructure:"namespace"`
	NomadHttpProxy  string `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
	NomadSkipVerify bool   `json:"skip_verify,omitempty" mapstructure:"skip_verify"`
	NomadDisableCliHints bool `json:"disable_cli_hints,omitempty" mapstructure:"disable_cli_hints"`
	NomadNoColor bool `json:"no_color,omitempty" mapstructure:"no_color"`
}

// Default struct with default profiles
//type Default struct {
//	VaultProfile     string `json:"vault_profile,omitempty" mapstracture:"vault_profile"`
//	NomadProfile     string `json:"nomad_profile,omitempty" mapstracture:"nomad_profile"`
//	ConsulProfile    string `json:"consul_profile,omitempty" mapstracture:"consul_profile"`
//	BoundaryfProfile string `json:"boundary_profile,omitempty" mapstracture:"boundary_profile"`
//}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "target",
	Short: "A context switcher CLI tool for Hashicorp Tools",
	Long: `Target allows a user to configure and switch between different contexts for their Vault, Nomad, Consul and Boundary targets by setting tool specific environment variables.
	
A context contains connection details for a given target.
Example: 
	a vault-dev context could point to 
	https://example-dev-vault.com:8200 with a vault token value is s.jidjibndiyuqepjepwo`,
	ValidArgs: []string{
		"vault",
		"nomad",
		"consul",
		"boundary",
		"terraform",
		"config",
		"version",
		"select",
	},
	Args:    cobra.OnlyValidArgs,
	Version: version,
}

var selectCmd = &cobra.Command{
	Use:     "select [profile]",
	Short:   "select a context profile for all tools",
	Long:    `select a context profile for all tools that have the profile defined.`,
	Example: `target select dev`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a profile name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		globalProxies := make(map[string]string)

		// Collect proxies from all tools
		if c.Vault[profile] != nil {
			context := c.Vault[profile]
			if context.HttpsProxy != "" {
				globalProxies["HTTPS_PROXY"] = context.HttpsProxy
			}
		}

		if c.Boundary[profile] != nil {
			context := c.Boundary[profile]
			if context.HttpProxy != "" {
				globalProxies["HTTPS_PROXY"] = context.HttpProxy
			}
		}

		if c.Nomad[profile] != nil {
			context := c.Nomad[profile]
			if context.NomadHttpProxy != "" {
				globalProxies["HTTPS_PROXY"] = context.NomadHttpProxy
				globalProxies["HTTP_PROXY"] = context.NomadHttpProxy
			}
		}

		if c.Consul[profile] != nil {
			context := c.Consul[profile]
			if context.ConsulHttpProxy != "" {
				globalProxies["HTTPS_PROXY"] = context.ConsulHttpProxy
				globalProxies["HTTP_PROXY"] = context.ConsulHttpProxy
			}
		}

		// Print global proxies first
		for key, value := range globalProxies {
			fmt.Printf("export %s=%s\n", key, value)
		}

		// Vault
		if c.Vault[profile] != nil {
			context := c.Vault[profile]
			exportCommands := []string{}

			if context.Endpoint != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_ADDR=%s", context.Endpoint))
			}

			if context.Token != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_TOKEN=%s", context.Token))
			}

			if context.Namespace != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_NAMESPACE=%s", context.Namespace))
			}

			if context.CaCert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CACERT=%s", context.CaCert))
			}

			if context.Cert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_CERT=%s", context.Cert))
			}

			if context.CaPath != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CAPATH=%s", context.CaPath))
			}

			if context.Key != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_KEY=%s", context.Key))
			}

			if context.Format != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_FORMAT=%s", context.Format))
			}

			if context.SkipVerify != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_SKIP_VERIFY=%s", context.SkipVerify))
			}

			if context.ClientTimeout != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLIENT_TIMEOUT=%s", context.ClientTimeout))
			}

			if context.ClusterAddr != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLUSTER_ADDR=%s", context.ClusterAddr))
			}

			if context.License != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LICENSE=%s", context.License))
			}

			if context.LicensePath != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LICENSE_PATH=%s", context.LicensePath))
			}

			if context.LogLevel != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_LOG_LEVEL=%s", context.LogLevel))
			}

			if context.MaxRetries != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_MAX_RETRIES=%s", context.MaxRetries))
			}

			if context.RedirectAddr != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_REDIRECT_ADDR=%s", context.RedirectAddr))
			}

			if context.TlsServerName != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_TLS_SERVER_NAME=%s", context.TlsServerName))
			}

			if context.CliNoColour != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_CLI_NO_COLOR=%s", context.CliNoColour))
			}

			if context.NoColor {
				exportCommands = append(exportCommands, "export VAULT_CLI_NO_COLOR=true")
			}

			if context.RateLimit != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_RATE_LIMIT=%s", context.RateLimit))
			}

			if context.SvrLookup != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_SRV_LOOKUP=%s", context.SvrLookup))
			}

			if context.Mfa != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_MFA=%s", context.Mfa))
			}

			if context.HttpProxy != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_HTTP_PROXY=%s", context.HttpProxy))
			}

			if context.DisableRedirects != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export VAULT_DISABLE_REDIRECTS=%s", context.DisableRedirects))
			}

			for _, cmd := range exportCommands {
				fmt.Println(cmd)
			}
		}

		// Boundary
		if c.Boundary[profile] != nil {
			context := c.Boundary[profile]
			exportCommands := []string{}

			if context.Endpoint != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_ADDR=%s", context.Endpoint))
			}

			if context.Token != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TOKEN=%s", context.Token))
			}

			if context.TokenName != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TOKEN_NAME=%s", context.TokenName))
			}

			if context.CaPath != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CAPATH=%s", context.CaPath))
			}

			if context.CaCert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CACERT=%s", context.CaCert))
			}

			if context.Cert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLIENT_CERT=%s", context.Cert))
			}

			if context.Key != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLIENT_KEY=%s", context.Key))
			}

			if context.TlsInsecure != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TLS_INSECURE=%s", context.TlsInsecure))
			}

			if context.TlsServerName != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_TLS_SERVER_NAME=%s", context.TlsServerName))
			}

			if context.RecoveryConfig != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_RECOVERY_CONFIG=%s", context.RecoveryConfig))
			}

			if context.ConnectAuthZToken != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_AUTHZ_TOKEN=%s", context.ConnectAuthZToken))
			}

			if context.ConnectExec != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_EXEC=%s", context.ConnectExec))
			}

			if context.ConnectListenAddr != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_ADDR=%s", context.ConnectListenAddr))
			}

			if context.ConnectListenPort != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_LISTEN_PORT=%s", context.ConnectListenPort))
			}

			if context.ConnectTargetScopeId != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_ID=%s", context.ConnectTargetScopeId))
			}

			if context.ConnectTargetScopeName != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CONNECT_TARGET_SCOPE_NAME=%s", context.ConnectTargetScopeName))
			}

			if context.AuthMethodId != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_AUTH_METHOD_ID=%s", context.AuthMethodId))
			}

			if context.LogLevel != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_LOG_LEVEL=%s", context.LogLevel))
			}

			if context.Format != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_CLI_FORMAT=%s", context.Format))
			}

			if context.ScopeId != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export BOUNDARY_SCOPE_ID=%s", context.ScopeId))
			}

			for _, cmd := range exportCommands {
				fmt.Println(cmd)
			}
		}

		// Nomad
		if c.Nomad[profile] != nil {
			context := c.Nomad[profile]
			exportCommands := []string{}

			if context.NomadEndpoint != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_ADDR=%s", context.NomadEndpoint))
			}

			if context.NomadToken != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_TOKEN=%s", context.NomadToken))
			}

			if context.NomadNamespace != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_NAMESPACE=%s", context.NomadNamespace))
			}

			if context.NomadCaCert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CACERT=%s", context.NomadCaCert))
			}

			if context.NomadCert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CLIENT_CERT=%s", context.NomadCert))
			}

			if context.NomadCaPath != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CAPATH=%s", context.NomadCaPath))
			}

			if context.NomadKey != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_CLIENT_KEY=%s", context.NomadKey))
			}

			if context.NomadRegion != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export NOMAD_REGION=%s", context.NomadRegion))
			}

			if context.NomadSkipVerify {
				exportCommands = append(exportCommands, "export NOMAD_SKIP_VERIFY=true")
			}

			if context.NomadDisableCliHints {
				exportCommands = append(exportCommands, "export NOMAD_CLI_SHOW_HINTS=false")
			}

			if context.NomadNoColor {
				exportCommands = append(exportCommands, "export NOMAD_CLI_NO_COLOR=true")
			}

			for _, cmd := range exportCommands {
				fmt.Println(cmd)
			}
		}

		// Consul
		if c.Consul[profile] != nil {
			context := c.Consul[profile]
			exportCommands := []string{}

			if context.ConsulEndpoint != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_HTTP_ADDR=%s", context.ConsulEndpoint))
			}

			if context.ConsulToken != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_HTTP_TOKEN=%s", context.ConsulToken))
			}

			if context.ConsulTokenFile != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_HTTP_TOKEN_FILE=%s", context.ConsulTokenFile))
			}

			if context.ConsulNamespace != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_NAMESPACE=%s", context.ConsulNamespace))
			}

			if context.ConsulCaCert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_CACERT=%s", context.ConsulCaCert))
			}

			if context.ConsulCert != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_CLIENT_CERT=%s", context.ConsulCert))
			}

			if context.ConsulCaPath != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_CAPATH=%s", context.ConsulCaPath))
			}

			if context.ConsulKey != "" {
				exportCommands = append(exportCommands, fmt.Sprintf("export CONSUL_CLIENT_KEY=%s", context.ConsulKey))
			}

			for _, cmd := range exportCommands {
				fmt.Println(cmd)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	targetdir.TargetHomeCreate()
}

func init() {
	cobra.OnInitialize(initConfig)
	targetdir.TargetHomeCreate()

	rootCmd.AddCommand(vaultCmd)
	rootCmd.AddCommand(nomadCmd)
	rootCmd.AddCommand(consulCmd)
	rootCmd.AddCommand(configlCmd)
	rootCmd.AddCommand(boundaryCmd)
	rootCmd.AddCommand(terraformCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(selectCmd)

}

// sliceOfMapsToMapHookFunc merges a slice of maps to a map
func sliceOfMapsToMapHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if from.Kind() == reflect.Slice && from.Elem().Kind() == reflect.Map && (to.Kind() == reflect.Struct || to.Kind() == reflect.Map) {
			source, ok := data.([]map[string]interface{})
			if !ok {
				return data, nil
			}
			if len(source) == 0 {
				return data, nil
			}
			if len(source) == 1 {
				return source[0], nil
			}
			// flatten the slice into one map
			convert := make(map[string]interface{})
			for _, mapItem := range source {
				for key, value := range mapItem {
					convert[key] = value
				}
			}
			return convert, nil
		}
		return data, nil
	}
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath("$HOME/.target")
	viper.SetConfigName("profiles")
	viper.SetConfigType("json")

	// Attempt to read the config file to see if it exists
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found, use default configuration
			c = &Config{}
		}
	}

	// Config file found and successfully loaded
	configOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		sliceOfMapsToMapHookFunc(),
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))
	if err := viper.Unmarshal(&c, configOption); err != nil {
		fmt.Println("Error unmarshaling config:", err)
		os.Exit(1)
	}

	if c.Vault == nil {
		c.Vault = map[string]*Vault{}
	}
	if c.Nomad == nil {
		c.Nomad = map[string]*Nomad{}
	}
	if c.Consul == nil {
		c.Consul = map[string]*Consul{}
	}
	if c.Boundary == nil {
		c.Boundary = map[string]*Boundary{}
	}
	if c.Terraform == nil {
		c.Terraform = map[string]*Terraform{}
	}

	// Automatically bind environment variables
	viper.AutomaticEnv()

}
