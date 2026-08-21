/*
 * Copyright 2021 kloeckner.i GmbH
 * Copyright 2023 DB-Operator Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1beta1

import (
	"k8s.io/apimachinery/pkg/types"
)

// NamespacedName mirrors Kubernetes' namespaced name type with the JSON tags required
// for CRD generation.
type NamespacedName struct {
	Namespace string `json:"Namespace"`
	Name      string `json:"Name"`
}

// ToKubernetesType converts our local type to the kubernetes API equivalent.
func (nn *NamespacedName) ToKubernetesType() types.NamespacedName {
	if nn == nil {
		return types.NamespacedName{}
	}

	return types.NamespacedName{
		Name:      nn.Name,
		Namespace: nn.Namespace,
	}
}

// Template defines one generated credential entry.
type Template struct {
	// Name is the data key written to the generated Secret or ConfigMap.
	Name string `json:"name"`
	// Template is a Go template evaluated with database connection values and template helper functions.
	Template string `json:"template"`
	// Secret writes the entry to a Secret when true and to a ConfigMap when false.
	// DbUser templates must set this field to true.
	Secret bool `json:"secret"`
}

type Templates []*Template

// CredentialsMetadata configures metadata on generated credential Secrets.
type CredentialsMetadata struct {
	// ExtraLabels is merged into the generated Secret's labels.
	// Values in this map replace existing values for the same keys.
	ExtraLabels map[string]string `json:"extraLabels,omitempty"`

	// ExtraAnnotations is merged into the generated Secret's annotations.
	// Values in this map replace existing values for the same keys.
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
}

// Credentials configures generated credential data and Secret metadata.
// TODO(@allanger): Field .spec.secretName should be moved here in the v1beta2 version
type Credentials struct {
	// Templates defines additional data entries for generated Secrets and ConfigMaps.
	Templates Templates `json:"templates,omitempty"`

	// Metadata configures labels and annotations on the generated credential Secret.
	Metadata *CredentialsMetadata `json:"metadata,omitempty"`
}
