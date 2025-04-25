package envconst

const (
	// tools name.
	AgnosticName   = "tf"
	AtmosName      = "atmos"
	TenvName       = "tenv"
	TerraformName  = "terraform"
	TerragruntName = "terragrunt"
	TofuName       = "tofu"

	// full project name.
	OpentofuName = "opentofu"

	// hidden proxy sub command.
	CallSubCmd = "call"

	BasicErrorExitCode = 1
	// exit code used by tenv proxy to signal a failure before the proxied command call.
	EarlyErrorExitCode = 42
)
