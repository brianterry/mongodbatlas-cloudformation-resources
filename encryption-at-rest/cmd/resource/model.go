package resource

import "github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'generate' command.

Updates to this type are made my editing the schema file
and executing the 'generate' command
*/

// Model is autogenerated from the json schema
type Model struct {
	AwsKms    AwsKms           `json:"AwsKms,omitempty"`
	ApiKeys   ApiKeyDefinition `json:"ApiKeys,omitempty"`
	ProjectId *encoding.String `json:"ProjectId,omitempty"`
}

// AwsKms is autogenerated from the json schema
type AwsKms struct {
	AccessKeyID         *encoding.String
	CustomerMasterKeyID *encoding.String
	Enabled             *encoding.Bool
	Region              *encoding.String
	SecretAccessKey     *encoding.String
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PublicKey  *encoding.String
	PrivateKey *encoding.String
}
