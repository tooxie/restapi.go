package errors

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Error struct {
	message string
	code    int
}

type Errors struct {
	// Server errors
	ErrorRetrievingBooks     int `code:"100" http:"500" message:"Error retrieving books"`
	ErrorSavingBook          int `code:"100" http:"500" message:"Error saving book"`
	ErrorUpdatingBook        int `code:"100" http:"500" message:"Error updating book"`
	InternalServerError      int `code:"100" http:"500" message:"Internal Server Error"`
	JwtTokenGenerationFailed int `code:"100" http:"500" message:"Token generation failed"`

	// Client errors
	BookNotFound       int `code:"200" http:"404" message:"Book not found"`
	EmailAlreadyInUse  int `code:"200" http:"400" message:"Email already in use"`
	Forbidden          int `code:"200" http:"403" message:"Forbidden"`
	InvalidCredentials int `code:"200" http:"400" message:"Invalid email and/or password"`
	InvalidHostHeader  int `code:"200" http:"400" message:"Invalid 'Host' header"`
	InvalidInput       int `code:"200" http:"400" message:"Invalid input provided"`
	JwtInvalidToken    int `code:"200" http:"400" message:"Invalid token"`
	JwtTokenExpired    int `code:"200" http:"400" message:"Token expired"`
	Unauthorized       int `code:"200" http:"401" message:"Unauthorized"`
}

func getHttpCode(field reflect.StructField) int {
	tag := field.Tag.Get("http")
	code, err := strconv.Atoi(tag)
	if err != nil {
		panic(err)
	}

	return code
}

func getErrorCode(field reflect.StructField) int {
	tag := field.Tag.Get("code")
	code, err := strconv.Atoi(tag)
	if err != nil {
		panic(err)
	}

	return code
}

func getErrorMessage(field reflect.StructField) string {
	tag := field.Tag.Get("message")
	return tag
}

func getError(name errorName, message string) (int, *gin.H) {
	errorsType := reflect.TypeOf(Errors{})
	_name := string(name)
	field, _ := errorsType.FieldByName(_name)

	_message := message
	if message == "" {
		_message = getErrorMessage(field)
	}
	hRet := gin.H{
		"code":    getErrorCode(field),
		"message": _message,
	}

	httpCode := getHttpCode(field)
	return httpCode, &hRet
}

func Raise(c *gin.Context, name errorName) {
	httpCode, errorStruct := getError(name, "")
	c.JSON(httpCode, gin.H{"error": *errorStruct})
}

func Abort(c *gin.Context, name errorName) {
	Raise(c, name)
	c.Abort()
}

func RaiseWithMessage(c *gin.Context, name errorName, message string) {
	httpCode, errorStruct := getError(name, message)
	c.JSON(httpCode, gin.H{"error": *errorStruct})
}

func AbortWithMessage(c *gin.Context, name errorName, message string) {
	RaiseWithMessage(c, name, message)
	c.Abort()
}
