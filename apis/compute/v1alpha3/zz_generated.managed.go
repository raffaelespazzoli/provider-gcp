/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha3

import (
	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GetBindingPhase of this GKECluster.
func (mg *GKECluster) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this GKECluster.
func (mg *GKECluster) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this GKECluster.
func (mg *GKECluster) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this GKECluster.
func (mg *GKECluster) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetReclaimPolicy of this GKECluster.
func (mg *GKECluster) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this GKECluster.
func (mg *GKECluster) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this GKECluster.
func (mg *GKECluster) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this GKECluster.
func (mg *GKECluster) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this GKECluster.
func (mg *GKECluster) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this GKECluster.
func (mg *GKECluster) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetReclaimPolicy of this GKECluster.
func (mg *GKECluster) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this GKECluster.
func (mg *GKECluster) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this GlobalAddress.
func (mg *GlobalAddress) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this GlobalAddress.
func (mg *GlobalAddress) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this GlobalAddress.
func (mg *GlobalAddress) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this GlobalAddress.
func (mg *GlobalAddress) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetReclaimPolicy of this GlobalAddress.
func (mg *GlobalAddress) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this GlobalAddress.
func (mg *GlobalAddress) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this GlobalAddress.
func (mg *GlobalAddress) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this GlobalAddress.
func (mg *GlobalAddress) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this GlobalAddress.
func (mg *GlobalAddress) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this GlobalAddress.
func (mg *GlobalAddress) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetReclaimPolicy of this GlobalAddress.
func (mg *GlobalAddress) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this GlobalAddress.
func (mg *GlobalAddress) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this Network.
func (mg *Network) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this Network.
func (mg *Network) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this Network.
func (mg *Network) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this Network.
func (mg *Network) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetReclaimPolicy of this Network.
func (mg *Network) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this Network.
func (mg *Network) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this Network.
func (mg *Network) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this Network.
func (mg *Network) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this Network.
func (mg *Network) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this Network.
func (mg *Network) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetReclaimPolicy of this Network.
func (mg *Network) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this Network.
func (mg *Network) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this Subnetwork.
func (mg *Subnetwork) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this Subnetwork.
func (mg *Subnetwork) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this Subnetwork.
func (mg *Subnetwork) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this Subnetwork.
func (mg *Subnetwork) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetReclaimPolicy of this Subnetwork.
func (mg *Subnetwork) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this Subnetwork.
func (mg *Subnetwork) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this Subnetwork.
func (mg *Subnetwork) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this Subnetwork.
func (mg *Subnetwork) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this Subnetwork.
func (mg *Subnetwork) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this Subnetwork.
func (mg *Subnetwork) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetReclaimPolicy of this Subnetwork.
func (mg *Subnetwork) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this Subnetwork.
func (mg *Subnetwork) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}