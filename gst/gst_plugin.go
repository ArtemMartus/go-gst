package gst

/*
#cgo pkg-config: gstreamer-1.0
#cgo CFLAGS: -Wno-deprecated-declarations -g -Wall
#include <gst/gst.h>
#include "gst.go.h"
*/
import "C"
import "github.com/gotk3/gotk3/glib"

// Plugin is a go representation of a GstPlugin.
type Plugin struct{ *Object }

// Instance returns the underlying GstPlugin instance.
func (p *Plugin) Instance() *C.GstPlugin { return C.toGstPlugin(p.unsafe()) }

func wrapPlugin(obj *glib.Object) *Plugin {
	return &Plugin{wrapObject(obj)}
}

// Description returns the description for this plugin.
func (p *Plugin) Description() string {
	ret := C.gst_plugin_get_description((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}

// Filename returns the filename for this plugin.
func (p *Plugin) Filename() string {
	ret := C.gst_plugin_get_filename((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}

// Version returns the version for this plugin.
func (p *Plugin) Version() string {
	ret := C.gst_plugin_get_version((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}

// License returns the license for this plugin.
func (p *Plugin) License() string {
	ret := C.gst_plugin_get_license((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}

// Source returns the source module for this plugin.
func (p *Plugin) Source() string {
	ret := C.gst_plugin_get_source((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}

// Package returns the binary package for this plugin.
func (p *Plugin) Package() string {
	ret := C.gst_plugin_get_package((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}

// Origin returns the origin URL for this plugin.
func (p *Plugin) Origin() string {
	ret := C.gst_plugin_get_origin((*C.GstPlugin)(p.Instance()))
	if ret == nil {
		return ""
	}
	return C.GoString(ret)
}
