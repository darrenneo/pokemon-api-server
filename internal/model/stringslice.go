package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// This is just a wrapper around []string which implements the sql.Scanner interface
// This is so that we can use the StringSlice type as []string in our models
type StringSlice []string

// Scan scan value into Jsonb, implements sql.Scanner interface
// This is so that the data from our database can be converted into a StringSlice type
// This is called when we query data from the database
func (ss *StringSlice) Scan(value interface{}) error {
	// We check the type assertion here that it is indeed a jsonb field
	tmp, ok := value.([]byte)
	if !ok {
		return errors.New("type not []byte")
	}
	// Then we unmarshal the data into the StringSlice type
	return json.Unmarshal(tmp, ss)
}

// We need to define the GormDataType for our StringSlice type
// This tells the GORM what database type to use when working with the table
func (ss StringSlice) GormDataType() string {
	// There's a few values that's supposed to be here, JSON is just one of them
	// These are the "database types" that GORM supports
	return "JSON"
}

// Value return json value, implement driver.Valuer interface
// This is called when we insert data into the database
func (ss StringSlice) Value() (driver.Value, error) {
	// We check if the StringSlice is empty
	if len(ss) == 0 {
		return nil, nil
	}
	// This is the reverse of the Scan function
	// This converts back to the []byte type to send back to DB
	return json.Marshal(ss)
}

// String returns the string representation of the StringSlice type
// This is just a convenient feature for you to use when need to convert to string
func (ss *StringSlice) String() string {
	// We marshal the StringSlice type into the []byte
	str, _ := json.Marshal(*ss)
	// Then we convert the []byte into a string
	return string(str)
}
