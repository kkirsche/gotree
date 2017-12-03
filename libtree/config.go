package libtree

// Config contains the different configuration modes
type Config struct {
	Mode string
}

// Validate is used to validate different configuration options
func (c *Config) Validate() bool {
	if !c.ValidateMode() {
		return false
	}

	return true
}

// ValidateMode is used to validate the mode is one that we support
func (c *Config) ValidateMode() bool {
	directoryModes := []string{
		"d",
		"di",
		"dir",
		"dire",
		"direc",
		"direct",
		"directo",
		"director",
		"directory",
	}

	if stringInSlice(c.Mode, directoryModes) {
		c.Mode = "directory"
	}

	fileModes := []string{
		"f",
		"fi",
		"fil",
		"file",
	}

	if stringInSlice(c.Mode, fileModes) {
		c.Mode = "file"
	}

	modes := []string{"directory", "file"}
	if stringInSlice(c.Mode, modes) {
		return true
	}

	return false
}
