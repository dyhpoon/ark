/*
Copyright 2018 the Heptio Ark contributors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/heptio/ark/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Backups returns a BackupInformer.
	Backups() BackupInformer
	// BackupStorageLocations returns a BackupStorageLocationInformer.
	BackupStorageLocations() BackupStorageLocationInformer
	// Configs returns a ConfigInformer.
	Configs() ConfigInformer
	// DeleteBackupRequests returns a DeleteBackupRequestInformer.
	DeleteBackupRequests() DeleteBackupRequestInformer
	// DownloadRequests returns a DownloadRequestInformer.
	DownloadRequests() DownloadRequestInformer
	// PodVolumeBackups returns a PodVolumeBackupInformer.
	PodVolumeBackups() PodVolumeBackupInformer
	// PodVolumeRestores returns a PodVolumeRestoreInformer.
	PodVolumeRestores() PodVolumeRestoreInformer
	// ResticRepositories returns a ResticRepositoryInformer.
	ResticRepositories() ResticRepositoryInformer
	// Restores returns a RestoreInformer.
	Restores() RestoreInformer
	// Schedules returns a ScheduleInformer.
	Schedules() ScheduleInformer
	// VolumeSnapshotLocations returns a VolumeSnapshotLocationInformer.
	VolumeSnapshotLocations() VolumeSnapshotLocationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Backups returns a BackupInformer.
func (v *version) Backups() BackupInformer {
	return &backupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupStorageLocations returns a BackupStorageLocationInformer.
func (v *version) BackupStorageLocations() BackupStorageLocationInformer {
	return &backupStorageLocationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Configs returns a ConfigInformer.
func (v *version) Configs() ConfigInformer {
	return &configInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DeleteBackupRequests returns a DeleteBackupRequestInformer.
func (v *version) DeleteBackupRequests() DeleteBackupRequestInformer {
	return &deleteBackupRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DownloadRequests returns a DownloadRequestInformer.
func (v *version) DownloadRequests() DownloadRequestInformer {
	return &downloadRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodVolumeBackups returns a PodVolumeBackupInformer.
func (v *version) PodVolumeBackups() PodVolumeBackupInformer {
	return &podVolumeBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodVolumeRestores returns a PodVolumeRestoreInformer.
func (v *version) PodVolumeRestores() PodVolumeRestoreInformer {
	return &podVolumeRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResticRepositories returns a ResticRepositoryInformer.
func (v *version) ResticRepositories() ResticRepositoryInformer {
	return &resticRepositoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Restores returns a RestoreInformer.
func (v *version) Restores() RestoreInformer {
	return &restoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Schedules returns a ScheduleInformer.
func (v *version) Schedules() ScheduleInformer {
	return &scheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeSnapshotLocations returns a VolumeSnapshotLocationInformer.
func (v *version) VolumeSnapshotLocations() VolumeSnapshotLocationInformer {
	return &volumeSnapshotLocationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
