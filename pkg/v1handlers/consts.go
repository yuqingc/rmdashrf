package v1handlers

// BadRequestErrMsg is for general 400 status code
const BadRequestErrMsg = "invalid or too less query params"

// MountDir is hard coded for now. It will be configurable
// TODO: checks if mount dir exists. Create it if not
// TODO: mountpath must be an absolute path
const MountDir = "/home/matt/Projects/github.com/yuqingc/data"

// MaxListResults is the max for listing contents for a specified directory
const MaxListResults = 5000
