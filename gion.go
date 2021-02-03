// Converted from ionice in util-linux (GPL2 licensed) to a Go package
package gion

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

// SetPri is "ioprio_set" in ion
func SetPri(which, who int, ioprio uint) (uint, error) {
	r1, _, errNo := syscall.Syscall(syscall.SYS_IOPRIO_SET, uintptr(which), uintptr(who), uintptr(ioprio))
	//fmt.Println("FROM SYSCALL (a):", r1)
	//fmt.Println("FROM SYSCALL (b):", r2)
	//fmt.Println("FROM SYSCALL (err):", errNo)
	var err error
	if errNo != 0 {
		err = errNo
	}
	// TODO: r1 or r2?
	return uint(r1), err
}

// Pri is "ioprio_get" in ion
func Pri(which, who int) (uint, error) {
	r1, _, errNo := syscall.Syscall(syscall.SYS_IOPRIO_GET, uintptr(which), uintptr(who), uintptr(0))
	//fmt.Println("FROM SYSCALL (a):", r1)
	//fmt.Println("FROM SYSCALL (b):", r2)
	//fmt.Println("FROM SYSCALL (err):", errNo)
	var err error
	if errNo != 0 {
		err = errNo
	}
	// TODO: r1 or r2?
	return uint(r1), err
}

func IOPRIO_PRIO_MASK() uint {
	return (uint(1) << IOPRIO_CLASS_SHIFT) - 1
}

func IOPRIO_PRIO_CLASS(mask uint) uint {
	return mask >> IOPRIO_CLASS_SHIFT
}

func IOPRIO_PRIO_DATA(mask uint) uint {
	return mask & IOPRIO_PRIO_MASK()
}

func IOPRIO_PRIO_VALUE(classn, data uint) uint {
	return ((classn << IOPRIO_CLASS_SHIFT) | data)
}

var to_prio map[int]string = map[int]string{
	IOPRIO_CLASS_NONE: "none",
	IOPRIO_CLASS_RT:   "realtime",
	IOPRIO_CLASS_BE:   "best-effort",
	IOPRIO_CLASS_IDLE: "idle",
}

func Parse(str string) int {
	for i := 0; i < len(to_prio); i++ {
		if strings.ToLower(str) == strings.ToLower(to_prio[i]) {
			return i
		}
	}
	return -1
}

func Print(pid, who int) {
	ioprio, err := Pri(who, pid)
	if err != nil {
		log.Fatalln("ioprio_get failed", err)
	}
	ioclass := IOPRIO_PRIO_CLASS(ioprio)
	name := "unknown"
	if ioclass >= 0 && ioclass < uint(len(to_prio)) {
		name = to_prio[int(ioclass)]
	}
	if ioclass != IOPRIO_CLASS_IDLE {
		fmt.Printf("%s: prio %d\n", name, IOPRIO_PRIO_DATA(ioprio))
	} else {
		fmt.Printf("%s\n", name)
	}
}

// SetIDPri is "ioprio_setid" in ion
func SetIDPri(which, ioclass, data, who int, tolerant bool) {
	_, err := SetPri(who, which, IOPRIO_PRIO_VALUE(uint(ioclass), uint(data)))
	if err != nil && !tolerant {
		log.Fatalln("ioprio_set failed", err)
	}
}
