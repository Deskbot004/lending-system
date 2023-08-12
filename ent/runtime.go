// Code generated by ent, DO NOT EDIT.

package ent

import (
	"lending-system/ent/game"
	"lending-system/ent/lending"
	"lending-system/ent/schema"
	"lending-system/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	gameFields := schema.Game{}.Fields()
	_ = gameFields
	// gameDescName is the schema descriptor for name field.
	gameDescName := gameFields[0].Descriptor()
	// game.DefaultName holds the default value on creation for the name field.
	game.DefaultName = gameDescName.Default.(string)
	// gameDescType is the schema descriptor for type field.
	gameDescType := gameFields[1].Descriptor()
	// game.DefaultType holds the default value on creation for the type field.
	game.DefaultType = gameDescType.Default.(string)
	lendingFields := schema.Lending{}.Fields()
	_ = lendingFields
	// lendingDescDate is the schema descriptor for date field.
	lendingDescDate := lendingFields[0].Descriptor()
	// lending.DefaultDate holds the default value on creation for the date field.
	lending.DefaultDate = lendingDescDate.Default.(string)
	// lendingDescNotes is the schema descriptor for notes field.
	lendingDescNotes := lendingFields[1].Descriptor()
	// lending.DefaultNotes holds the default value on creation for the notes field.
	lending.DefaultNotes = lendingDescNotes.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
