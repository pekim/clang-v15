package clang

// #include "./clang-c/BuildSystem.h"
// #include "./clang-c/FatalErrorHandler.h"
// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "unsafe"

// GetBuildSessionTimestamp return the timestamp for use with Clang's -fbuild-session-timestamp= option.
func GetBuildSessionTimestamp() uint64 {
	return uint64(C.clang_getBuildSessionTimestamp())
}

// Install_aborting_llvm_fatal_error_handler installs error handler that prints error message to stderr and calls abort(). Replaces currently installed error handler (if any).
func InstallAbortingFatalErrorHandler() {
	C.clang_install_aborting_llvm_fatal_error_handler()
}

// Uninstall_llvm_fatal_error_handler removes currently installed error handler (if any). If no error handler is intalled, the default strategy is to print error message to stderr and call exit(1).
func UninstallFatalErrorHandler() {
	C.clang_uninstall_llvm_fatal_error_handler()
}

// DefaultDiagnosticDisplayOptions retrieve the set of display options most similar to the
// default behavior of the clang compiler.
//
// Returns A set of display options suitable for use with \c
// clang_formatDiagnostic().
func DefaultDiagnosticDisplayOptions() uint32 {
	return uint32(C.clang_defaultDiagnosticDisplayOptions())
}

// GetDiagnosticCategoryName retrieve the name of a particular diagnostic category. This
// is now deprecated. Use clang_getDiagnosticCategoryText()
// instead.
//
// Parameter Category A diagnostic category number, as returned by
// clang_getDiagnosticCategory().
//
// Returns The name of the given diagnostic category.
func GetDiagnosticCategoryName(category uint32) string {
	o := cxstring{C.clang_getDiagnosticCategoryText(C.CXDiagnostic(uintptr(category)))}
	defer o.Dispose()

	return o.String()
}

// DefaultEditingTranslationUnitOptions returns the set of flags that is suitable for parsing a translation
// unit that is being edited.
//
// The set of flags returned provide options for clang_parseTranslationUnit()
// to indicate that the translation unit is likely to be reparsed many times,
// either explicitly (via clang_reparseTranslationUnit()) or implicitly
// (e.g., by code completion (clang_codeCompletionAt())). The returned flag
// set contains an unspecified set of optimizations (e.g., the precompiled
// preamble) geared toward improving the performance of these routines. The
// set of optimizations enabled may change from one version to the next.
func DefaultEditingTranslationUnitOptions() uint32 {
	return uint32(C.clang_defaultEditingTranslationUnitOptions())
}

// ConstructUSR_ObjCClass construct a USR for a specified Objective-C class.
func ConstructUSR_ObjCClass(className string) string {
	c_className := C.CString(className)
	defer C.free(unsafe.Pointer(c_className))

	o := cxstring{C.clang_constructUSR_ObjCClass(c_className)}
	defer o.Dispose()

	return o.String()
}

// ConstructUSR_ObjCCategory construct a USR for a specified Objective-C category.
func ConstructUSR_ObjCCategory(className string, categoryName string) string {
	c_className := C.CString(className)
	defer C.free(unsafe.Pointer(c_className))
	c_categoryName := C.CString(categoryName)
	defer C.free(unsafe.Pointer(c_categoryName))

	o := cxstring{C.clang_constructUSR_ObjCCategory(c_className, c_categoryName)}
	defer o.Dispose()

	return o.String()
}

// ConstructUSR_ObjCProtocol construct a USR for a specified Objective-C protocol.
func ConstructUSR_ObjCProtocol(protocolName string) string {
	c_protocolName := C.CString(protocolName)
	defer C.free(unsafe.Pointer(c_protocolName))

	o := cxstring{C.clang_constructUSR_ObjCProtocol(c_protocolName)}
	defer o.Dispose()

	return o.String()
}

// ConstructUSR_ObjCIvar construct a USR for a specified Objective-C instance variable and the USR for its containing class.
func ConstructUSR_ObjCIvar(name string, classUSR cxstring) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	o := cxstring{C.clang_constructUSR_ObjCIvar(c_name, classUSR.c)}
	defer o.Dispose()

	return o.String()
}

// ConstructUSR_ObjCMethod construct a USR for a specified Objective-C method and the USR for its containing class.
func ConstructUSR_ObjCMethod(name string, isInstanceMethod uint32, classUSR cxstring) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	o := cxstring{C.clang_constructUSR_ObjCMethod(c_name, C.uint(isInstanceMethod), classUSR.c)}
	defer o.Dispose()

	return o.String()
}

// ConstructUSR_ObjCProperty construct a USR for a specified Objective-C property and the USR for its containing class.
func ConstructUSR_ObjCProperty(property string, classUSR cxstring) string {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	o := cxstring{C.clang_constructUSR_ObjCProperty(c_property, classUSR.c)}
	defer o.Dispose()

	return o.String()
}

func EnableStackTraces() {
	C.clang_enableStackTraces()
}

// DefaultCodeCompleteOptions returns a default set of code-completion options that can be passed toclang_codeCompleteAt().
func DefaultCodeCompleteOptions() uint32 {
	return uint32(C.clang_defaultCodeCompleteOptions())
}

// GetClangVersion return a version string, suitable for showing to a user, but not intended to be parsed (the format is not guaranteed to be stable).
func GetClangVersion() string {
	o := cxstring{C.clang_getClangVersion()}
	defer o.Dispose()

	return o.String()
}

// ToggleCrashRecovery enable/disable crash recovery.
//
// Parameter isEnabled Flag to indicate if crash recovery is enabled. A non-zero
// value enables crash recovery, while 0 disables it.
func ToggleCrashRecovery(isEnabled uint32) {
	C.clang_toggleCrashRecovery(C.uint(isEnabled))
}
