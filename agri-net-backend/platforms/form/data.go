package form

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ImageExtensions list of valid image extensions
var ImageExtensions = []string{"jpeg", "png", "jpg", "gif", "btmp"}

// PhoneRX represents phone number maching pattern
var PhoneRX = regexp.MustCompile("(^\\+[0-9]{2}|^\\+[0-9]{2}\\(0\\)|^\\(\\+[0-9]{2}\\)\\(0\\)|^00[0-9]{2}|^0)([0-9]{9}$|[0-9\\-\\s]{10}$)")

// EmailRX represents email address maching pattern
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Input represents form input values and validations
type Input struct {
	Values  url.Values
	VErrors ValidationErrors
	CSRF    string
}

// MinLength checks if a given minium length is satisfied
func (inVal *Input) MinLength(field string, d int) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		inVal.VErrors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}
}

// IsImage function checking whether the file is an image or not
func IsImage(filepath string) bool {
	extension := GetExtension(filepath)
	extension = strings.ToLower(extension)
	for _, e := range ImageExtensions {
		if e == extension {
			return true
		}
	}
	return false
}

// GetExtension function to return the extension of the File Input FileName
func GetExtension(Filename string) string {
	fileSlice := strings.Split(Filename, ".")
	if len(fileSlice) >= 1 {
		return fileSlice[len(fileSlice)-1]
	}
	return ""
}

// GetFormFile returning the file and it's header and error while opening the file
// if the file doesnt exist Error message will be added to the Input Validation Errors List
func (inVal *Input) GetFormFile(request *http.Request, filename string) (multipart.File, *multipart.FileHeader, error) {
	file, header, err := request.FormFile(filename)
	if err != nil || file == nil || header == nil || !(IsImage(header.Filename)) {
		if err != nil {
			inVal.VErrors.Add(filename, fmt.Sprintf("Invalid File Value %s ", filename))
		} else if !(IsImage(header.Filename)) {
			inVal.VErrors.Add(filename, fmt.Sprintf("Invalid File Type %s ", GetExtension(filename)))
			file.Close()
			return nil, nil, errors.New(" Invalid File Error ")
		} else {
			inVal.VErrors.Add(filename, fmt.Sprintf("Invalid File Value %s ", filename))
			file.Close()
			return nil, nil, errors.New(" Invalid File Error ")
		}
	}
	return file, header, nil
}

// Required checks if list of provided form input fields have values
func (inVal *Input) Required(fields ...string) {
	for _, f := range fields {
		value := inVal.Values.Get(f)
		if value == "" {
			inVal.VErrors.Add(f, "This field is required field")
		}
	}
}

// MatchesPattern checks if a given input form field matchs a given pattern
func (inVal *Input) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		inVal.VErrors.Add(field, "The value entered is invalid")
	}
}

// MatchesPattern PATTERN , REGULAR EXPRESION
func MatchesPattern(value string, pattern *regexp.Regexp) bool {
	if value == "" {
		return false
	}
	if !pattern.MatchString(value) {
		return false
	}
	return true
}

// MatchesPattern checks if a given input form field matchs a given pattern
func (inVal *Input) ParseBoolean(field string) (res bool) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	res, era := strconv.ParseBool(value)
	if era != nil {
		inVal.VErrors.Add(field, fmt.Sprintf("Parse Error : variable %s must be type Boolean ", field))
	}
	return res
}

// PasswordMatches checks if Password and Confirm Password fields match
func (inVal *Input) PasswordMatches(password string, confPassword string) {
	pwd := inVal.Values.Get(password)
	confPwd := inVal.Values.Get(confPassword)
	if pwd == "" || confPwd == "" {
		return
	}
	if pwd != confPwd {
		inVal.VErrors.Add(password, "The Password and Confim Password values did not match")
		inVal.VErrors.Add(confPassword, "The Password and Confim Password values did not match")
	}
}

// Valid checks if any form input validation has failed or not
func (inVal *Input) Valid() bool {
	return len(inVal.VErrors) == 0
}
