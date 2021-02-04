// Package ionice contains code that has been ported from util-linux/ionice (GPL2 licensed)
package ionice

import (
	"fmt"
	"log"
	"strings"
	"syscall"
)

const (
	IOPRIO_CLASS_NONE = 0
	IOPRIO_CLASS_RT   = 1
	IOPRIO_CLASS_BE   = 2
	IOPRIO_CLASS_IDLE = 3

	IOPRIO_WHO_PROCESS = 1
	IOPRIO_WHO_PGRP    = 2
	IOPRIO_WHO_USER    = 3

	IOPRIO_CLASS_SHIFT = 13
)

type IOPrioClass int

// SetPri sets the IO priority for the given "which" (process, pgrp or user) and "who" (the ID),
// using the given io priority number.
func SetPri(which, who int, ioprio uint) (uint, error) {
	r1, _, errNo := syscall.Syscall(syscall.SYS_IOPRIO_SET, uintptr(which), uintptr(who), uintptr(ioprio))
	var err error
	if errNo != 0 {
		err = errNo
	}
	return uint(r1), err
}

// Pri returns the IO priority for the given "which" (process, pgrp or user) and "who" (the ID).
func Pri(which, who int) (uint, error) {
	r1, _, errNo := syscall.Syscall(syscall.SYS_IOPRIO_GET, uintptr(which), uintptr(who), uintptr(0))
	var err error
	if errNo != 0 {
		err = errNo
	}
	return uint(r1), err
}

func IOPrioMask() uint {
	return (uint(1) << IOPRIO_CLASS_SHIFT) - 1
}

func IOPrioPrioClass(mask uint) IOPrioClass {
	return IOPrioClass(mask >> IOPRIO_CLASS_SHIFT)
}

func IOPrioData(mask uint) uint {
	return mask & IOPrioMask()
}

func IOPrioValue(classn, data uint) uint {
	return ((classn << IOPRIO_CLASS_SHIFT) | data)
}

var to_prio = map[IOPrioClass]string{
	IOPRIO_CLASS_NONE: "none",
	IOPRIO_CLASS_RT:   "realtime",
	IOPRIO_CLASS_BE:   "best-effort",
	IOPRIO_CLASS_IDLE: "idle",
}

// Parse converts a string containing either:
// "none", "realtime", best-effort" or "idle", to a corresponding IOPRIO_CLASS.
// will also handle "0", "1", "2" or "3"
// The parsing is case-insensitive, so "REALTIME" or "rEaLtImE" is also fine.
func Parse(ioprio string) (IOPrioClass, error) {
	switch strings.ToLower(ioprio) {
	case "0", "none":
		return IOPRIO_CLASS_NONE, nil
	case "1", "realtime":
		return IOPRIO_CLASS_RT, nil
	case "2", "best-effort":
		return IOPRIO_CLASS_BE, nil
	case "3", "idle":
		return IOPRIO_CLASS_IDLE, nil
	}
	return 0, fmt.Errorf("could not parse %s as an IOPRIO_CLASS constant", ioprio)
}

// Print outputs the IO nice status for the given PID and "who"
func Print(pid, who int) {
	ioprio, err := Pri(who, pid)
	if err != nil {
		log.Fatalln("ioprio_get failed", err)
	}
	ioclass := IOPrioPrioClass(ioprio)
	name := "unknown"
	to_prio_len := IOPrioClass(len(to_prio))
	if ioclass >= 0 && ioclass < to_prio_len {
		name = to_prio[ioclass]
	}
	if ioclass != IOPRIO_CLASS_IDLE {
		fmt.Printf("%s: prio %d\n", name, IOPrioData(ioprio))
	} else {
		fmt.Printf("%s\n", name)
	}
}

func SetIDPri(which int, ioclass IOPrioClass, data, who int, tolerant bool) {
	_, err := SetPri(who, which, IOPrioValue(uint(ioclass), uint(data)))
	if err != nil && !tolerant {
		log.Fatalln("ioprio_set failed", err)
	}
}
