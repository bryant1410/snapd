// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

import (
	"github.com/snapcore/snapd/interfaces"
)

const autopilotIntrospectionPlugAppArmor = `
# Description: Allows an application to be introspected and export its ui
# status over DBus

#include <abstractions/dbus-session-strict>

dbus (receive)
    bus=session
    path=/com/canonical/Autopilot/**
    interface=org.freedesktop.DBus.Introspectable
    member=Introspect
    peer=(label=unconfined),
dbus (receive)
    bus=session
    path=/com/canonical/Autopilot/Introspection
    interface=com.canonical.Autopilot.Introspection
    member=GetVersion
    peer=(label=unconfined),
dbus (receive)
    bus=session
    path=/com/canonical/Autopilot/Introspection
    interface=com.canonical.Autopilot.Introspection
    member=GetState
    peer=(label=unconfined),
`
const autopilotIntrospectionPlugSecComp = `
# Description: Allows an application to be introspected and export its ui
# status over DBus.
recvmsg
sendmsg
sendto
`

// NewAutopilotIntrospectionInterface returns a new "autopilot-introspection"
// interface.
func NewAutopilotIntrospectionInterface() interfaces.Interface {
	return &commonInterface{
		name: "autopilot-introspection",
		connectedPlugAppArmor: autopilotIntrospectionPlugAppArmor,
		connectedPlugSecComp:  autopilotIntrospectionPlugSecComp,
		reservedForOS:         true,
	}
}

func init() {
	registerIface(NewAutopilotIntrospectionInterface())
}
