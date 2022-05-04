package main

import (
	"archive/zip"
	"bytes"
	"consterr"
	"errors"
	"fmt"
	"os"
	"reflect"
)

func CalcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("(" + fmt.Sprintf("%d", numerator) + "/" + fmt.Sprintf("%d", denominator) + ") cannot divide by 0")
	}
	quotient := numerator / denominator
	mod := numerator % denominator
	return quotient, mod, nil

}
func BasicErrorHandling() {
	quotient, modulus, err := CalcRemainderAndMod(1, 0)
	fmt.Println("quotient:", quotient, "modulus:", modulus, "error:", err)
}

func QuotientRemainderDivisorByErrorsNewFunc(numerator, denominator int) (int, int, int, error) {
	if denominator == 0 {
		return 0, 0, 0, errors.New("cannot divide by 0")
	}
	quotient := numerator / denominator
	remainder := numerator % denominator

	return quotient, remainder, denominator, nil
}

func QuotientRemainderDivisorByErrorf(numerator, denominator int) (int, int, int, error) {
	if denominator == 0 {
		return 0, 0, 0, fmt.Errorf("cannot divide by 0")
	}
	quotient := numerator / denominator
	remainder := numerator % denominator

	return quotient, remainder, denominator, nil
}

func UsingStringsForSimpleErrors() {
	fmt.Println(QuotientRemainderDivisorByErrorf(20, 0))
	fmt.Println(QuotientRemainderDivisorByErrorf(3, 2))
	fmt.Println(QuotientRemainderDivisorByErrorsNewFunc(20, 0))
	fmt.Println(QuotientRemainderDivisorByErrorsNewFunc(5, 4))
}

// to signal the processing cannot continue due to a problem
// of the **** CURRENT STATE **** (basically a state management)
func SentinelErrors() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat { // " 'Err'Format is the sentinel error "
		fmt.Println(fmt.Errorf("zip file erorr"))
	}

	const (
		ErrFoo = consterr.Sentinel("foo error")
		ErrBar = consterr.Sentinel("bar error")
	)
	fmt.Println("consterr 1:", ErrFoo)
	fmt.Println("consterr 2:", ErrBar)
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	message string
}

func (se StatusErr) Error() string {
	return se.message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := Login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := GetData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			message: fmt.Sprintf("%s file not found", file),
		}
	}
	return data, nil
}

func GetData(file string) ([]byte, error) {
	if err := fileChecker(file); err != nil {
		return nil, errors.New("file not found")
	}
	return []byte{255, 255}, nil
}
func Login(uid, pwd string) error {
	if len(uid) == 0 {
		return errors.New("login failed")
	} else {
		return nil
	}
}

//custom defined error types
func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

// idiomatic custom error types
func GenerateErrorIdiomaticOne(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}

	return nil
}

func GenerateErrorIdiomaticTwo(flag bool) error {
	var genErr error // always declare standard error type
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}

	return genErr
}

func ErrorsAreValues() {
	fmt.Println(LoginAndGetData("", "", ""))
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil) // this is not nil because 'error' is an interface. An interface is not nil until both the type and the value of the interface are nil

	// fix and idiomatic approach to write custom error types
	err = GenerateErrorIdiomaticOne(true)
	fmt.Println("GenerateErrorIdiomaticOne(true):", err != nil)
	err = GenerateErrorIdiomaticOne(false)
	fmt.Println("GenerateErrorIdiomaticOne(false):", err != nil)
	err = GenerateErrorIdiomaticTwo(true)
	fmt.Println("GenerateErrorIdiomaticTwo(true):", err != nil)
	err = GenerateErrorIdiomaticTwo(false)
	fmt.Println("GenerateErrorIdiomaticTwo(false):", err != nil)
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in filechecker: %w", err)
	}
	f.Close()
	return nil
}

//using Unwrap in custom defined error type
type StatusErrForUnwrap struct {
	Status  Status
	message string
	err     error
}

func (seu StatusErrForUnwrap) Error() string {
	return seu.message
}
func (seu StatusErrForUnwrap) Unwrap() error {
	return seu.err
}
func LoginAndGetDataForUnwrap(uid, pwd, file string) ([]byte, error) {
	err := Login(uid, pwd)
	if err != nil {
		return nil, StatusErrForUnwrap{
			Status:  InvalidLogin,
			message: fmt.Sprintf("invalid credentials for user %s", uid),
			err:     err,
		}
	}
	data, err := GetData(file)
	if err != nil {
		return nil, StatusErrForUnwrap{
			Status:  NotFound,
			message: fmt.Sprintf("file %s not found", file),
			err:     err,
		}
	}
	return data, err
}
func WrappingErrors() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
	data, err := LoginAndGetDataForUnwrap("", "", "not_here.txt")
	fmt.Println(data, err)
	data, err = LoginAndGetDataForUnwrap("test", "test", "not_here.txt")
	fmt.Println(data, err)
}

/*
	errors.Is - when we are looking for a specific instance or value
	errors.As - when we are looking for a specific type
*/
func fileCheckerWithIsToUnwrap(file string) error {
	f, err := os.Open(file)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(fmt.Errorf("in fileCheckerWithIsToUnwrap: %w", err))
		return err
	}
	f.Close()
	return nil
}

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}
func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}
func GenerateErrorCustom() error {
	return MyErr{
		Codes: []int{404},
	}
}

type ResourceErr struct {
	Resource string
	Codes    []int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("codes: %v and Resource: %s", re.Codes, re.Resource)
}
func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := len(other.Codes) == 0
		matchResource := re.Resource == other.Resource
		matchCodes := reflect.DeepEqual(re.Codes, other.Codes)

		return (matchResource && matchCodes) ||
			(matchResource && ignoreCode) ||
			(matchCodes && ignoreResource)
	}
	return false
}
func (re ResourceErr) DigitalResource() string {
	return re.Resource
}
func (re ResourceErr) ErrorCodes() []int {
	return re.Codes
}

func GenerateResourceError() error {
	return ResourceErr{
		Resource: "Database",
		Codes:    []int{401},
	}
}

func IsAndAsInErrorHandling() {
	// errros.Is
	err := fileCheckerWithIsToUnwrap("not_here.txt")
	fmt.Println(err)
	err = GenerateErrorCustom()
	ErrNotFound := MyErr{
		Codes: []int{404},
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println(err)
	}

	err = GenerateResourceError()
	unauthorizedDBAccess := ResourceErr{Resource: "Database", Codes: []int{401}}
	if errors.Is(err, unauthorizedDBAccess) {
		fmt.Println(fmt.Sprintf("Resource: '%s' is unauthorized to access (%v)", unauthorizedDBAccess.Resource, unauthorizedDBAccess.Codes))
	}

	// errors.As
	var dBAccess401 ResourceErr
	if errors.As(err, &dBAccess401) {
		fmt.Println(dBAccess401.Resource, dBAccess401.Codes)
	}

	var dbAccessor interface {
		DigitalResource() string
		ErrorCodes() []int
	}
	if errors.As(err, &dbAccessor) {
		fmt.Println(dbAccessor.DigitalResource())
		fmt.Println(dbAccessor.ErrorCodes())
	}
}

// Multiple errors within a same function
// before wrapping with defer
func HandleMultipleErrorBeforeDefer(val1, val2 int) (string, error) {
	if val1 < val2 {
		return "", fmt.Errorf("in HandleMultipleErrorBeforeDefer: %w", errors.New("val1 less than val2"))
	}
	if val1 == val2 {
		return "", fmt.Errorf("in HandleMultipleErrorBeforeDefer: %w", errors.New("val1 equals val2"))
	}
	if val2 == 0 {
		return "", fmt.Errorf("in HandleMultipleErrorBeforeDefer: %w", errors.New("val2 is 0"))
	}
	result := val1 / val2
	return fmt.Sprint(result), nil
}

// after wrapping with defer
func HandleMultipleErrorAfterDefer(val1, val2 int) (outcome string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in HandleMultipleErrorAfterDefer: %w", err)
		}
		if outcome == "failed validation" {
			outcome = "deferred failed validation"
		}
	}() // this extra parenthesis is to call this anonymous function

	if val1 < val2 {
		return "failed validation", errors.New("val1 less than val2")
	}
	if val1 == val2 {
		return "failed validation", errors.New("val1 equals val2")
	}
	if val2 == 0 {
		return "failed validation", errors.New("val2 is 0")
	}
	result := val1 / val2

	return fmt.Sprint(result), nil
}

func WrappingErrorsWithDefer() {
	fmt.Println(HandleMultipleErrorBeforeDefer(1, 2))
	fmt.Println(HandleMultipleErrorBeforeDefer(2, 2))
	fmt.Println(HandleMultipleErrorBeforeDefer(1, 0))
	fmt.Println(HandleMultipleErrorBeforeDefer(3, 2))

	fmt.Println(HandleMultipleErrorAfterDefer(1, 2))
	fmt.Println(HandleMultipleErrorAfterDefer(2, 2))
	fmt.Println(HandleMultipleErrorAfterDefer(1, 0))
	fmt.Println(HandleMultipleErrorAfterDefer(3, 2))
}

func doPanic(msg string) {
	panic(msg)
}
func div60(i int) {
	defer func() {
		if v := recover(); v != nil { // recover() catches the panic
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
func PanicAndRecover() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}

	doPanic(os.Args[0])
}

func GetStackTraceFromError() {
	err := fileChecker("not_here.txt")
	fmt.Printf("%+v\n", err)
}
