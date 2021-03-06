package fakecgo

import _ "unsafe" //needed for go:linkname

// The following "fakes" all the necessary stuff for pretending we're using cgo, without actually doing that
// -> iscgo will be set to true and all functions that are then required by the runtime implemented
// This is necessary to get TLS working in the mainthread (cgo_init) and in all other threads (cgo_thread_start).
// If we leave this out, libc can't use TLS since go runtime overwrites it (printf with %f already needs that)

// The actual functions are implemented in assembly trampolines (trampoline_*.s) that call into carefully crafted golang (cgo.go)

//go:linkname _cgo_init _cgo_init
//go:linkname x_cgo_init_trampoline x_cgo_init_trampoline
var x_cgo_init_trampoline byte
var _cgo_init = &x_cgo_init_trampoline

//go:linkname x_cgo_thread_start_trampoline x_cgo_thread_start_trampoline
//go:linkname _cgo_thread_start _cgo_thread_start
var x_cgo_thread_start_trampoline byte
var _cgo_thread_start = &x_cgo_thread_start_trampoline

//go:linkname x_cgo_notify_runtime_init_done_trampoline x_cgo_notify_runtime_init_done_trampoline
//go:linkname _cgo_notify_runtime_init_done _cgo_notify_runtime_init_done
var x_cgo_notify_runtime_init_done_trampoline byte
var _cgo_notify_runtime_init_done = &x_cgo_notify_runtime_init_done_trampoline

//go:linkname x_cgo_setenv_trampoline x_cgo_setenv_trampoline
//go:linkname _cgo_setenv runtime._cgo_setenv
var x_cgo_setenv_trampoline byte
var _cgo_setenv = &x_cgo_setenv_trampoline

//go:linkname x_cgo_unsetenv_trampoline x_cgo_unsetenv_trampoline
//go:linkname _cgo_unsetenv runtime._cgo_unsetenv
var x_cgo_unsetenv_trampoline byte
var _cgo_unsetenv = &x_cgo_unsetenv_trampoline

// let's pretend we have cgo:
//go:linkname _iscgo runtime.iscgo
var _iscgo = true

// Now all the symbols we need to import from various libraries to implement the above functions:
// (just using one variable and taking the address in libcFuncs.go works with amd64 - but the two variable dance is needed for 386, where we get an unknown symbol relocation otherwise :/)

//go:linkname libc_pthread_attr_init_x libc_pthread_attr_init_x
var libc_pthread_attr_init_x libcFunc
var libc_pthread_attr_init = &libc_pthread_attr_init_x

//go:linkname libc_pthread_attr_getstacksize_x libc_pthread_attr_getstacksize_x
var libc_pthread_attr_getstacksize_x libcFunc
var libc_pthread_attr_getstacksize = &libc_pthread_attr_getstacksize_x

//go:linkname libc_pthread_attr_destroy_x libc_pthread_attr_destroy_x
var libc_pthread_attr_destroy_x libcFunc
var libc_pthread_attr_destroy = &libc_pthread_attr_destroy_x

//go:linkname libc_pthread_sigmask_x libc_pthread_sigmask_x
var libc_pthread_sigmask_x libcFunc
var libc_pthread_sigmask = &libc_pthread_sigmask_x

//go:linkname libc_pthread_create_x libc_pthread_create_x
var libc_pthread_create_x libcFunc
var libc_pthread_create = &libc_pthread_create_x

//go:linkname libc_pthread_detach_x libc_pthread_detach_x
var libc_pthread_detach_x libcFunc
var libc_pthread_detach = &libc_pthread_detach_x

//go:linkname libc_setenv_x libc_setenv_x
var libc_setenv_x libcFunc
var libc_setenv = &libc_setenv_x

//go:linkname libc_unsetenv_x libc_unsetenv_x
var libc_unsetenv_x libcFunc
var libc_unsetenv = &libc_unsetenv_x

//go:linkname libc_malloc_x libc_malloc_x
var libc_malloc_x libcFunc
var libc_malloc = &libc_malloc_x

//go:linkname libc_free_x libc_free_x
var libc_free_x libcFunc
var libc_free = &libc_free_x

//go:linkname libc_nanosleep_x libc_nanosleep_x
var libc_nanosleep_x libcFunc
var libc_nanosleep = &libc_nanosleep_x

//go:linkname libc_sigfillset_x libc_sigfillset_x
var libc_sigfillset_x libcFunc
var libc_sigfillset = &libc_sigfillset_x

//go:linkname libc_abort_x libc_abort_x
var libc_abort_x libcFunc
var libc_abort = &libc_abort_x

//go:linkname libc_dprintf_x libc_dprintf_x
var libc_dprintf_x libcFunc
var libc_dprintf = &libc_dprintf_x

//go:linkname libc_strerror_x libc_strerror_x
var libc_strerror_x libcFunc
var libc_strerror = &libc_strerror_x
