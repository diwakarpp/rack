// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kms_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/session"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/kms"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleKMS_CancelKeyDeletion() {
	svc := kms.New(session.New())

	params := &kms.CancelKeyDeletionInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.CancelKeyDeletion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_CreateAlias() {
	svc := kms.New(session.New())

	params := &kms.CreateAliasInput{
		AliasName:   aws.String("AliasNameType"), // Required
		TargetKeyId: aws.String("KeyIdType"),     // Required
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_CreateGrant() {
	svc := kms.New(session.New())

	params := &kms.CreateGrantInput{
		GranteePrincipal: aws.String("PrincipalIdType"), // Required
		KeyId:            aws.String("KeyIdType"),       // Required
		Constraints: &kms.GrantConstraints{
			EncryptionContextEquals: map[string]*string{
				"Key": aws.String("EncryptionContextValue"), // Required
				// More values...
			},
			EncryptionContextSubset: map[string]*string{
				"Key": aws.String("EncryptionContextValue"), // Required
				// More values...
			},
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		Name: aws.String("GrantNameType"),
		Operations: []*string{
			aws.String("GrantOperation"), // Required
			// More values...
		},
		RetiringPrincipal: aws.String("PrincipalIdType"),
	}
	resp, err := svc.CreateGrant(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_CreateKey() {
	svc := kms.New(session.New())

	params := &kms.CreateKeyInput{
		Description: aws.String("DescriptionType"),
		KeyUsage:    aws.String("KeyUsageType"),
		Policy:      aws.String("PolicyType"),
	}
	resp, err := svc.CreateKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_Decrypt() {
	svc := kms.New(session.New())

	params := &kms.DecryptInput{
		CiphertextBlob: []byte("PAYLOAD"), // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.Decrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DeleteAlias() {
	svc := kms.New(session.New())

	params := &kms.DeleteAliasInput{
		AliasName: aws.String("AliasNameType"), // Required
	}
	resp, err := svc.DeleteAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DescribeKey() {
	svc := kms.New(session.New())

	params := &kms.DescribeKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DisableKey() {
	svc := kms.New(session.New())

	params := &kms.DisableKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DisableKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DisableKeyRotation() {
	svc := kms.New(session.New())

	params := &kms.DisableKeyRotationInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DisableKeyRotation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_EnableKey() {
	svc := kms.New(session.New())

	params := &kms.EnableKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.EnableKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_EnableKeyRotation() {
	svc := kms.New(session.New())

	params := &kms.EnableKeyRotationInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.EnableKeyRotation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_Encrypt() {
	svc := kms.New(session.New())

	params := &kms.EncryptInput{
		KeyId:     aws.String("KeyIdType"), // Required
		Plaintext: []byte("PAYLOAD"),       // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.Encrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GenerateDataKey() {
	svc := kms.New(session.New())

	params := &kms.GenerateDataKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		KeySpec:       aws.String("DataKeySpec"),
		NumberOfBytes: aws.Int64(1),
	}
	resp, err := svc.GenerateDataKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GenerateDataKeyWithoutPlaintext() {
	svc := kms.New(session.New())

	params := &kms.GenerateDataKeyWithoutPlaintextInput{
		KeyId: aws.String("KeyIdType"), // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		KeySpec:       aws.String("DataKeySpec"),
		NumberOfBytes: aws.Int64(1),
	}
	resp, err := svc.GenerateDataKeyWithoutPlaintext(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GenerateRandom() {
	svc := kms.New(session.New())

	params := &kms.GenerateRandomInput{
		NumberOfBytes: aws.Int64(1),
	}
	resp, err := svc.GenerateRandom(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GetKeyPolicy() {
	svc := kms.New(session.New())

	params := &kms.GetKeyPolicyInput{
		KeyId:      aws.String("KeyIdType"),      // Required
		PolicyName: aws.String("PolicyNameType"), // Required
	}
	resp, err := svc.GetKeyPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GetKeyRotationStatus() {
	svc := kms.New(session.New())

	params := &kms.GetKeyRotationStatusInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.GetKeyRotationStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListAliases() {
	svc := kms.New(session.New())

	params := &kms.ListAliasesInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListAliases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListGrants() {
	svc := kms.New(session.New())

	params := &kms.ListGrantsInput{
		KeyId:  aws.String("KeyIdType"), // Required
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListGrants(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListKeyPolicies() {
	svc := kms.New(session.New())

	params := &kms.ListKeyPoliciesInput{
		KeyId:  aws.String("KeyIdType"), // Required
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListKeyPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListKeys() {
	svc := kms.New(session.New())

	params := &kms.ListKeysInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListKeys(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListRetirableGrants() {
	svc := kms.New(session.New())

	params := &kms.ListRetirableGrantsInput{
		RetiringPrincipal: aws.String("PrincipalIdType"), // Required
		Limit:             aws.Int64(1),
		Marker:            aws.String("MarkerType"),
	}
	resp, err := svc.ListRetirableGrants(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_PutKeyPolicy() {
	svc := kms.New(session.New())

	params := &kms.PutKeyPolicyInput{
		KeyId:      aws.String("KeyIdType"),      // Required
		Policy:     aws.String("PolicyType"),     // Required
		PolicyName: aws.String("PolicyNameType"), // Required
	}
	resp, err := svc.PutKeyPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ReEncrypt() {
	svc := kms.New(session.New())

	params := &kms.ReEncryptInput{
		CiphertextBlob:   []byte("PAYLOAD"),       // Required
		DestinationKeyId: aws.String("KeyIdType"), // Required
		DestinationEncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		SourceEncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
	}
	resp, err := svc.ReEncrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_RetireGrant() {
	svc := kms.New(session.New())

	params := &kms.RetireGrantInput{
		GrantId:    aws.String("GrantIdType"),
		GrantToken: aws.String("GrantTokenType"),
		KeyId:      aws.String("KeyIdType"),
	}
	resp, err := svc.RetireGrant(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_RevokeGrant() {
	svc := kms.New(session.New())

	params := &kms.RevokeGrantInput{
		GrantId: aws.String("GrantIdType"), // Required
		KeyId:   aws.String("KeyIdType"),   // Required
	}
	resp, err := svc.RevokeGrant(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ScheduleKeyDeletion() {
	svc := kms.New(session.New())

	params := &kms.ScheduleKeyDeletionInput{
		KeyId:               aws.String("KeyIdType"), // Required
		PendingWindowInDays: aws.Int64(1),
	}
	resp, err := svc.ScheduleKeyDeletion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_UpdateAlias() {
	svc := kms.New(session.New())

	params := &kms.UpdateAliasInput{
		AliasName:   aws.String("AliasNameType"), // Required
		TargetKeyId: aws.String("KeyIdType"),     // Required
	}
	resp, err := svc.UpdateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_UpdateKeyDescription() {
	svc := kms.New(session.New())

	params := &kms.UpdateKeyDescriptionInput{
		Description: aws.String("DescriptionType"), // Required
		KeyId:       aws.String("KeyIdType"),       // Required
	}
	resp, err := svc.UpdateKeyDescription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}