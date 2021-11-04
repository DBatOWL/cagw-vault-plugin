/*
 * Copyright (c) 2019 Entrust Datacard Corporation.
 * All rights reserved.
 */

package main

import (
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathIssue(b *backend) *framework.Path {
	ret := &framework.Path{
		Pattern: "issue/" + framework.GenericNameRegex("roleName") + "/?$",

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation:   &framework.PathOperation{Callback: b.opReadIssue},
			logical.ListOperation:   &framework.PathOperation{Callback: b.opListIssue},
			logical.UpdateOperation: &framework.PathOperation{Callback: b.opWriteIssue},
		},

		HelpSynopsis:    "Certificate Enrollment",
		HelpDescription: "Enroll for certificate.",
		Fields:          addIssueAndSignCommonFields(map[string]*framework.FieldSchema{}),
	}

	return ret
}
