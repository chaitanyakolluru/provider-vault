/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/upbound/provider-vault/apis/v1beta1"
)

const (
	// provider config
	// source: https://registry.terraform.io/providers/hashicorp/vault/latest/docs

	// required origin URL of the Vault server
	keyAddress = "address"

	// optional provider arguments
	// remove options that we do not yet want to implement
	keyAddAddressToEnv      = "add_address_to_env"
	keyToken                = "token"
	keyTokenName            = "token_name"
	keyCaCertFile           = "ca_cert_file"
	keyCaCertDir            = "ca_cert_dir"
	keyAuthLoginUserpass    = "auth_login_userpath"
	keyAuthLoginAWS         = "auth_login_aws"
	keyAuthLoginCert        = "auth_login_cert"
	keyAuthLoginGCP         = "auth_login_gcp"
	keyAuthLoginKerberos    = "auth_login_kerberos"
	keyAuthLoginRadius      = "auth_login_radius"
	keyAuthLoginOCI         = "auth_login_oci"
	keyAuthLoginOIDC        = "auth_login_oidc"
	keyAuthLoginJWT         = "auth_login_jwt"
	keyAuthLoginAzure       = "auth_login_azure"
	keyAuthLogin            = "auth_login"
	keyClientAuth           = "client_auth"
	keySkipTlsVerify        = "skip_tls_verify"
	keyTlsServerName        = "tls_server_name"
	keySkipChildToken       = "skip_child_token"
	keyMaxLeaseTtlSeconds   = "max_lease_ttl_seconds"
	keyMaxRetries           = "max_retries"
	keyMaxRetriesCcc        = "max_retries_ccc"
	keyNamespace            = "namespace"
	keySkipGetVaultVersion  = "skip_get_vault_version"
	keyVaultVersionOverride = "vault_version_override"
	keyHeaders              = "headers"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal vault credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		/*ps.Configuration = map[string]any{
			"username": creds["username"],
			"password": creds["password"],
		}*/
		// set provider configuration
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyAddress]; ok {
			ps.Configuration[keyAddress] = v
		}
		if v, ok := creds[keyAddAddressToEnv]; ok {
			ps.Configuration[keyAddAddressToEnv] = v
		}
		if v, ok := creds[keyToken]; ok {
			ps.Configuration[keyToken] = v
		}
		if v, ok := creds[keyTokenName]; ok {
			ps.Configuration[keyTokenName] = v
		}
		if v, ok := creds[keyCaCertFile]; ok {
			ps.Configuration[keyCaCertFile] = v
		}
		if v, ok := creds[keyCaCertDir]; ok {
			ps.Configuration[keyCaCertDir] = v
		}
		if v, ok := creds[keyAuthLoginUserpass]; ok {
			ps.Configuration[keyAuthLoginUserpass] = v
		}
		if v, ok := creds[keyAuthLoginAWS]; ok {
			ps.Configuration[keyAuthLoginAWS] = v
		}
		if v, ok := creds[keyAuthLoginCert]; ok {
			ps.Configuration[keyAuthLoginCert] = v
		}
		if v, ok := creds[keyAuthLoginGCP]; ok {
			ps.Configuration[keyAuthLoginGCP] = v
		}
		if v, ok := creds[keyAuthLoginKerberos]; ok {
			ps.Configuration[keyAuthLoginKerberos] = v
		}
		if v, ok := creds[keyAuthLoginRadius]; ok {
			ps.Configuration[keyAuthLoginRadius] = v
		}
		if v, ok := creds[keyAuthLoginOCI]; ok {
			ps.Configuration[keyAuthLoginOCI] = v
		}
		if v, ok := creds[keyAuthLoginOIDC]; ok {
			ps.Configuration[keyAuthLoginOIDC] = v
		}
		if v, ok := creds[keyAuthLoginJWT]; ok {
			ps.Configuration[keyAuthLoginJWT] = v
		}
		if v, ok := creds[keyAuthLoginAzure]; ok {
			ps.Configuration[keyAuthLoginAzure] = v
		}
		if v, ok := creds[keyAuthLogin]; ok {
			ps.Configuration[keyAuthLogin] = v
		}
		if v, ok := creds[keyClientAuth]; ok {
			ps.Configuration[keyClientAuth] = v
		}
		if v, ok := creds[keySkipTlsVerify]; ok {
			ps.Configuration[keySkipTlsVerify] = v
		}
		if v, ok := creds[keyTlsServerName]; ok {
			ps.Configuration[keyTlsServerName] = v
		}
		if v, ok := creds[keySkipChildToken]; ok {
			ps.Configuration[keySkipChildToken] = v
		}
		if v, ok := creds[keyMaxLeaseTtlSeconds]; ok {
			ps.Configuration[keyMaxLeaseTtlSeconds] = v
		}
		if v, ok := creds[keyMaxRetries]; ok {
			ps.Configuration[keyMaxRetries] = v
		}
		if v, ok := creds[keyMaxRetriesCcc]; ok {
			ps.Configuration[keyMaxRetriesCcc] = v
		}
		if v, ok := creds[keyNamespace]; ok {
			ps.Configuration[keyNamespace] = v
		}
		if v, ok := creds[keySkipGetVaultVersion]; ok {
			ps.Configuration[keySkipGetVaultVersion] = v
		}
		if v, ok := creds[keyVaultVersionOverride]; ok {
			ps.Configuration[keyVaultVersionOverride] = v
		}
		if v, ok := creds[keyHeaders]; ok {
			ps.Configuration[keyHeaders] = v
		}
		return ps, nil
	}
}