/*

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
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// ProductBaseWebHooker is the real validator
// +kubebuilder:object:generate=false
type ProductBaseWebHooker interface {
	Default(*ProductBase)
	ValidateCreate(*ProductBase) error
	ValidateUpdate(old runtime.Object, new *ProductBase) error
	ValidateDelete(*ProductBase) error
}

var webHooker ProductBaseWebHooker

// SetWebHooker set the real validator
func SetWebHooker(v ProductBaseWebHooker) {
	webHooker = v
}

// +kubebuilder:webhook:path=/mutate-productbase,mutating=true,failurePolicy=fail,groups=product.alauda.io,resources=productbases,verbs=create;update,versions=v1alpha1,name=mproductbase.kb.io,admissionReviewVersions={v1alpha1},sideEffects=none

var _ webhook.Defaulter = &ProductBase{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ProductBase) Default() {
	if webHooker != nil {
		webHooker.Default(r)
	}
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-productbase,mutating=false,failurePolicy=fail,groups=product.alauda.io,resources=productbases,versions=v1alpha1,name=vproductbase.kb.io,admissionReviewVersions={v1alpha1},sideEffects=none

var _ webhook.Validator = &ProductBase{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ProductBase) ValidateCreate() error {
	if webHooker != nil {
		return webHooker.ValidateCreate(r)
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ProductBase) ValidateUpdate(old runtime.Object) error {
	if webHooker != nil {
		return webHooker.ValidateUpdate(old, r)
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ProductBase) ValidateDelete() error {
	if webHooker != nil {
		return webHooker.ValidateDelete(r)
	}
	return nil
}
