package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

var tolerant bool

func ioprio_set(which, who int, ioprio uint) (uint, error) {
	r1, r2, errNo := syscall.Syscall(syscall.SYS_IOPRIO_SET, uintptr(which), uintptr(who), uintptr(ioprio))
	fmt.Println("FROM SYSCALL (a):", r1)
	fmt.Println("FROM SYSCALL (b):", r2)
	fmt.Println("FROM SYSCALL (err):", errNo)
	var err error
	if errNo != 0 {
		err = errNo
	}
	// TODO: r1 or r2?
	return uint(r1), err
}

func ioprio_get(which, who int) (uint, error) {
	r1, r2, errNo := syscall.Syscall(syscall.SYS_IOPRIO_GET, uintptr(which), uintptr(who), uintptr(0))
	fmt.Println("FROM SYSCALL (a):", r1)
	fmt.Println("FROM SYSCALL (b):", r2)
	fmt.Println("FROM SYSCALL (err):", errNo)
	var err error
	if errNo != 0 {
		err = errNo
	}
	// TODO: r1 or r2?
	return uint(r1), err
}

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

func parse_ioclass(str string) int {
	for i := 0; i < len(to_prio); i++ {
		if strings.ToLower(str) == strings.ToLower(to_prio[i]) {
			return i
		}
	}
	return -1
}

func ioprio_print(pid, who int) {
	ioprio, err := ioprio_get(who, pid)
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

func ioprio_setid(which, ioclass, data, who int) {
	_, err := ioprio_set(who, which, IOPRIO_PRIO_VALUE(uint(ioclass), uint(data)))
	if err != nil && !tolerant {
		log.Fatalln("ioprio_set failed", err)
	}
}

func usage() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(" ion [options] -p <pid>...")
	fmt.Println(" ion [options] -P <pgid>...")
	fmt.Println(" ion [options] -u <uid>...")
	fmt.Println(" ion [options] <command>")
	fmt.Println()
	fmt.Println("Show or change the I/O-scheduling class and priority of a process.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println(" -c, --class <class>    name or number of scheduling class,")
	fmt.Println("                          0: none, 1: realtime, 2: best-effort, 3: idle")
	fmt.Println(" -n, --classdata <num>  priority (0..7) in the specified scheduling class,")
	fmt.Println("                          only for the realtime and best-effort classes")
	fmt.Println(" -p, --pid <pid>...     act on these already running processes")
	fmt.Println(" -P, --pgid <pgrp>...   act on already running processes in these groups")
	fmt.Println(" -t, --ignore           ignore failures")
	fmt.Println(" -u, --uid <uid>...     act on already running processes owned by these users")
	fmt.Println()
	fmt.Printf("%-24s%s\n", " -h, --help", "display this help")
	fmt.Printf("%-24s%s\n", " -V, --version", "display version")
	os.Exit(0)
}

func main() {
	var (
		data = 4
		//set         = 0
		ioclass = IOPRIO_CLASS_BE
		//c           = 0
		//which       = 0
		//who         = 0
		//invalid_msg = ""
	)

	/*    static const struct option longopts[] = {
	          { "classdata", required_argument, NULL, 'n' },
	          { "class", required_argument, NULL, 'c' },
	          { "help", no_argument, NULL, 'h' },
	          { "ignore", no_argument, NULL, 't' },
	          { "pid", required_argument, NULL, 'p' },
	          { "pgid", required_argument, NULL, 'P' },
	          { "uid", required_argument, NULL, 'u' },
	          { "version", no_argument, NULL, 'V' },
	          { NULL, 0, NULL, 0 }
	      };
	*/

	//defer close_stdout();

	/*
	   while ((c = getopt_long(argc, argv, "+n:c:p:P:u:tVh", longopts, NULL)) != EOF)
	       switch (c) {
	       case 'n':
	           data = strtos32_or_err(optarg, "invalid class data argument");
	           set |= 1;
	           break;
	       case 'c':
	           if (isdigit(*optarg))
	               ioclass = strtos32_or_err(optarg,
	                   "invalid class argument");
	           else {
	               ioclass = parse_ioclass(optarg);
	               if (ioclass < 0)
	                   errx(EXIT_FAILURE,
	                       "unknown scheduling class: '%s'",
	                       optarg);
	           }
	           set |= 2;
	           break;
	       case 'p':
	           if (who)
	               errx(EXIT_FAILURE,
	                   "can handle only one of pid, pgid or uid at once");
	           invalid_msg = "invalid PID argument";
	           which = strtos32_or_err(optarg, invalid_msg);
	           who = IOPRIO_WHO_PROCESS;
	           break;
	       case 'P':
	           if (who)
	               errx(EXIT_FAILURE,
	                   "can handle only one of pid, pgid or uid at once");
	           invalid_msg = "invalid PGID argument";
	           which = strtos32_or_err(optarg, invalid_msg);
	           who = IOPRIO_WHO_PGRP;
	           break;
	       case 'u':
	           if (who)
	               errx(EXIT_FAILURE,
	                   "can handle only one of pid, pgid or uid at once");
	           invalid_msg = "invalid UID argument";
	           which = strtos32_or_err(optarg, invalid_msg);
	           who = IOPRIO_WHO_USER;
	           break;
	       case 't':
	           tolerant = true;
	           break;
	       case 'V':
	           printf("%s %s\n", "ion", "1.0.0");
	           exit(EXIT_SUCCESS);
	       case 'h':
	           usage();
	       default:
	           fprintf(stderr, "Try '%s --help' for more information.\n", "ion");
	           exit(EXIT_FAILURE);
	       }
	   switch (ioclass) {
	   case IOPRIO_CLASS_NONE:
	       if ((set & 1) && !tolerant) {
	           warnx("ignoring given class data for none class");
	       }
	       data = 0;
	       break;
	   case IOPRIO_CLASS_RT:
	   case IOPRIO_CLASS_BE:
	       break;
	   case IOPRIO_CLASS_IDLE:
	       if ((set & 1) && !tolerant) {
	           warnx("ignoring given class data for idle class");
	       }
	       data = 7;
	       break;
	   default:
	       if (!tolerant) {
	           warnx("unknown prio class %d", ioclass);
	       }
	       break;
	   }
	   if (!set && !which && optind == argc) {
	        // ion without options, print the current ioprio
	       ioprio_print(0, IOPRIO_WHO_PROCESS);
	   } else if (!set && who) {
	        // ion -p|-P|-u ID [ID ...]
	       ioprio_print(which, who);
	       while argv[optind] != 0 {
	           which = strtos32_or_err(argv[optind], invalid_msg)
	           ioprio_print(which, who)
	       	optind++
	       }
	   } else if (set && who) {
	        // ion -c CLASS -p|-P|-u ID [ID ...]
	       ioprio_setid(which, ioclass, data, who);
	       while argv[optind] != 0 {
	           which = strtos32_or_err(argv[optind], invalid_msg);
	           ioprio_setid(which, ioclass, data, who);
	           optind++
	       }
	   } else if (argv[optind]) {
	        // ion [-c CLASS] COMMAND
	       ioprio_setid(0, ioclass, data, IOPRIO_WHO_PROCESS);
	       execvp(argv[optind], &argv[optind]);
	       static int EX_EXEC_FAILED = 126; // Program located, but not usable
	       static int EX_EXEC_ENOENT = 127; // Could not find program to exec
	       err(errno == ENOENT ? EX_EXEC_ENOENT : EX_EXEC_FAILED, "failed to execute %s", argv[optind]);

	   } else {
	       warnx("bad usage");
	       log.Fatalln("Try 'ion --help' for more information.")
	   }*/

	// ion [-c CLASS] COMMAND
	ioprio_setid(0, ioclass, data, IOPRIO_WHO_PROCESS)

	var argv0 string = "/usr/bin/ls"
	var argv []string = []string{"/usr/bin/ls"}
	var envv []string = []string{}

	err := syscall.Exec(argv0, argv, envv)
	if err != nil {
		log.Fatalf("failed to execute %s", argv0)
	}

	//execvp(argv[optind], &argv[optind])
	//const EX_EXEC_FAILED = 126 // Program located, but not usable
	//const EX_EXEC_ENOENT = 127 // Could not find program to exec
	//log.Fatalf("failed to execute %s", commandString)
	//err(errno == ENOENT ? EX_EXEC_ENOENT : EX_EXEC_FAILED, "failed to execute %s", argv[optind]);

}
