/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/provider-aws/apis/ec2/v1beta1"
)

// ResolveReferences of this Zone
func (mg *ResourceRecordSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.zoneID
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To:           reference.To{Managed: &HostedZone{}, List: &HostedZoneList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return err
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of a VPC provided for a HostedZone
func (mg *HostedZone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.VPCID
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCId),
		Reference:    mg.Spec.ForProvider.VPCIdRef,
		Selector:     mg.Spec.ForProvider.VPCIdSelector,
		To:           reference.To{Managed: &v1beta1.VPC{}, List: &v1beta1.VPCList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return err
	}

	mg.Spec.ForProvider.VPCId = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIdRef = rsp.ResolvedReference

	return nil
}
