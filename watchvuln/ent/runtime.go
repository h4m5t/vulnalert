// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/zema1/watchvuln/ent/schema"
	"github.com/zema1/watchvuln/ent/vulninformation"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	vulninformationFields := schema.VulnInformation{}.Fields()
	_ = vulninformationFields
	// vulninformationDescTitle is the schema descriptor for title field.
	vulninformationDescTitle := vulninformationFields[1].Descriptor()
	// vulninformation.DefaultTitle holds the default value on creation for the title field.
	vulninformation.DefaultTitle = vulninformationDescTitle.Default.(string)
	// vulninformationDescDescription is the schema descriptor for description field.
	vulninformationDescDescription := vulninformationFields[2].Descriptor()
	// vulninformation.DefaultDescription holds the default value on creation for the description field.
	vulninformation.DefaultDescription = vulninformationDescDescription.Default.(string)
	// vulninformationDescSeverity is the schema descriptor for severity field.
	vulninformationDescSeverity := vulninformationFields[3].Descriptor()
	// vulninformation.DefaultSeverity holds the default value on creation for the severity field.
	vulninformation.DefaultSeverity = vulninformationDescSeverity.Default.(string)
	// vulninformationDescCve is the schema descriptor for cve field.
	vulninformationDescCve := vulninformationFields[4].Descriptor()
	// vulninformation.DefaultCve holds the default value on creation for the cve field.
	vulninformation.DefaultCve = vulninformationDescCve.Default.(string)
	// vulninformationDescDisclosure is the schema descriptor for disclosure field.
	vulninformationDescDisclosure := vulninformationFields[5].Descriptor()
	// vulninformation.DefaultDisclosure holds the default value on creation for the disclosure field.
	vulninformation.DefaultDisclosure = vulninformationDescDisclosure.Default.(string)
	// vulninformationDescSolutions is the schema descriptor for solutions field.
	vulninformationDescSolutions := vulninformationFields[6].Descriptor()
	// vulninformation.DefaultSolutions holds the default value on creation for the solutions field.
	vulninformation.DefaultSolutions = vulninformationDescSolutions.Default.(string)
	// vulninformationDescFrom is the schema descriptor for from field.
	vulninformationDescFrom := vulninformationFields[10].Descriptor()
	// vulninformation.DefaultFrom holds the default value on creation for the from field.
	vulninformation.DefaultFrom = vulninformationDescFrom.Default.(string)
	// vulninformationDescPushed is the schema descriptor for pushed field.
	vulninformationDescPushed := vulninformationFields[11].Descriptor()
	// vulninformation.DefaultPushed holds the default value on creation for the pushed field.
	vulninformation.DefaultPushed = vulninformationDescPushed.Default.(bool)
	// vulninformationDescCreateTime is the schema descriptor for create_time field.
	vulninformationDescCreateTime := vulninformationFields[12].Descriptor()
	// vulninformation.DefaultCreateTime holds the default value on creation for the create_time field.
	vulninformation.DefaultCreateTime = vulninformationDescCreateTime.Default.(func() time.Time)
	// vulninformationDescUpdateTime is the schema descriptor for update_time field.
	vulninformationDescUpdateTime := vulninformationFields[13].Descriptor()
	// vulninformation.DefaultUpdateTime holds the default value on creation for the update_time field.
	vulninformation.DefaultUpdateTime = vulninformationDescUpdateTime.Default.(func() time.Time)
	// vulninformation.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	vulninformation.UpdateDefaultUpdateTime = vulninformationDescUpdateTime.UpdateDefault.(func() time.Time)
}
