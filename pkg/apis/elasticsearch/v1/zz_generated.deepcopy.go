// +build !ignore_autogenerated

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	commonv1 "github.com/elastic/cloud-on-k8s/pkg/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RoleSource, len(*in))
		copy(*out, *in)
	}
	if in.FileRealm != nil {
		in, out := &in.FileRealm, &out.FileRealm
		*out = make([]FileRealmSource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangeBudget) DeepCopyInto(out *ChangeBudget) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(int32)
		**out = **in
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangeBudget.
func (in *ChangeBudget) DeepCopy() *ChangeBudget {
	if in == nil {
		return nil
	}
	out := new(ChangeBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSettings) DeepCopyInto(out *ClusterSettings) {
	*out = *in
	if in.InitialMasterNodes != nil {
		in, out := &in.InitialMasterNodes, &out.InitialMasterNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSettings.
func (in *ClusterSettings) DeepCopy() *ClusterSettings {
	if in == nil {
		return nil
	}
	out := new(ClusterSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elasticsearch) DeepCopyInto(out *Elasticsearch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	if in.AssocConfs != nil {
		in, out := &in.AssocConfs, &out.AssocConfs
		*out = make(map[types.NamespacedName]commonv1.AssociationConf, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elasticsearch.
func (in *Elasticsearch) DeepCopy() *Elasticsearch {
	if in == nil {
		return nil
	}
	out := new(Elasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Elasticsearch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchList) DeepCopyInto(out *ElasticsearchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Elasticsearch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchList.
func (in *ElasticsearchList) DeepCopy() *ElasticsearchList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSettings) DeepCopyInto(out *ElasticsearchSettings) {
	*out = *in
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(Node)
		(*in).DeepCopyInto(*out)
	}
	in.Cluster.DeepCopyInto(&out.Cluster)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSettings.
func (in *ElasticsearchSettings) DeepCopy() *ElasticsearchSettings {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	in.HTTP.DeepCopyInto(&out.HTTP)
	in.Transport.DeepCopyInto(&out.Transport)
	if in.NodeSets != nil {
		in, out := &in.NodeSets, &out.NodeSets
		*out = make([]NodeSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(commonv1.PodDisruptionBudgetTemplate)
		(*in).DeepCopyInto(*out)
	}
	in.Auth.DeepCopyInto(&out.Auth)
	if in.SecureSettings != nil {
		in, out := &in.SecureSettings, &out.SecureSettings
		*out = make([]commonv1.SecretSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemoteClusters != nil {
		in, out := &in.RemoteClusters, &out.RemoteClusters
		*out = make([]RemoteCluster, len(*in))
		copy(*out, *in)
	}
	out.Monitoring = in.Monitoring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchStatus) DeepCopyInto(out *ElasticsearchStatus) {
	*out = *in
	if in.ElasticsearchAssociationsStatus != nil {
		in, out := &in.ElasticsearchAssociationsStatus, &out.ElasticsearchAssociationsStatus
		*out = make(commonv1.AssociationStatusMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchStatus.
func (in *ElasticsearchStatus) DeepCopy() *ElasticsearchStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsMonitoringAssociation) DeepCopyInto(out *EsMonitoringAssociation) {
	*out = *in
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		*out = new(Elasticsearch)
		(*in).DeepCopyInto(*out)
	}
	out.ref = in.ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsMonitoringAssociation.
func (in *EsMonitoringAssociation) DeepCopy() *EsMonitoringAssociation {
	if in == nil {
		return nil
	}
	out := new(EsMonitoringAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileRealmSource) DeepCopyInto(out *FileRealmSource) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileRealmSource.
func (in *FileRealmSource) DeepCopy() *FileRealmSource {
	if in == nil {
		return nil
	}
	out := new(FileRealmSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsMonitoring) DeepCopyInto(out *LogsMonitoring) {
	*out = *in
	out.ElasticsearchRef = in.ElasticsearchRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsMonitoring.
func (in *LogsMonitoring) DeepCopy() *LogsMonitoring {
	if in == nil {
		return nil
	}
	out := new(LogsMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsMonitoring) DeepCopyInto(out *MetricsMonitoring) {
	*out = *in
	out.ElasticsearchRef = in.ElasticsearchRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsMonitoring.
func (in *MetricsMonitoring) DeepCopy() *MetricsMonitoring {
	if in == nil {
		return nil
	}
	out := new(MetricsMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	out.Metrics = in.Metrics
	out.Logs = in.Logs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = new(bool)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(bool)
		**out = **in
	}
	if in.Ingest != nil {
		in, out := &in.Ingest, &out.Ingest
		*out = new(bool)
		**out = **in
	}
	if in.ML != nil {
		in, out := &in.ML, &out.ML
		*out = new(bool)
		**out = **in
	}
	if in.Transform != nil {
		in, out := &in.Transform, &out.Transform
		*out = new(bool)
		**out = **in
	}
	if in.RemoteClusterClient != nil {
		in, out := &in.RemoteClusterClient, &out.RemoteClusterClient
		*out = new(bool)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VotingOnly != nil {
		in, out := &in.VotingOnly, &out.VotingOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSet) DeepCopyInto(out *NodeSet) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = (*in).DeepCopy()
	}
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSet.
func (in *NodeSet) DeepCopy() *NodeSet {
	if in == nil {
		return nil
	}
	out := new(NodeSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteCluster) DeepCopyInto(out *RemoteCluster) {
	*out = *in
	out.ElasticsearchRef = in.ElasticsearchRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteCluster.
func (in *RemoteCluster) DeepCopy() *RemoteCluster {
	if in == nil {
		return nil
	}
	out := new(RemoteCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleSource) DeepCopyInto(out *RoleSource) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleSource.
func (in *RoleSource) DeepCopy() *RoleSource {
	if in == nil {
		return nil
	}
	out := new(RoleSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportConfig) DeepCopyInto(out *TransportConfig) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportConfig.
func (in *TransportConfig) DeepCopy() *TransportConfig {
	if in == nil {
		return nil
	}
	out := new(TransportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportTLSOptions) DeepCopyInto(out *TransportTLSOptions) {
	*out = *in
	out.Certificate = in.Certificate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportTLSOptions.
func (in *TransportTLSOptions) DeepCopy() *TransportTLSOptions {
	if in == nil {
		return nil
	}
	out := new(TransportTLSOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStrategy) DeepCopyInto(out *UpdateStrategy) {
	*out = *in
	in.ChangeBudget.DeepCopyInto(&out.ChangeBudget)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStrategy.
func (in *UpdateStrategy) DeepCopy() *UpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(UpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZenDiscoveryStatus) DeepCopyInto(out *ZenDiscoveryStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZenDiscoveryStatus.
func (in *ZenDiscoveryStatus) DeepCopy() *ZenDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(ZenDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}
