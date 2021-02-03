// gionice is a port of ionice to Go

package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/jessevdk/go-flags"
	"github.com/xyproto/ionice"
)

const versionString = "gionice 1.0.0"

const usageString = `Usage:
 gionice [options] -p <pid>...
 gionice [options] -P <pgid>...
 gionice [options] -u <uid>...
 gionice [options] <command>

Show or change the I/O-scheduling class and priority of a process.

Options:
 -c, --class <class>    name or number of scheduling class,
                          0: none, 1: realtime, 2: best-effort, 3: idle
 -n, --classdata <num>  priority (0..7) in the specified scheduling class,
                          only for the realtime and best-effort classes
 -p, --pid <pid>...     act on these already running processes
 -P, --pgid <pgrp>...   act on already running processes in these groups
 -t, --ignore           ignore failures
 -u, --uid <uid>...     act on already running processes owned by these users

 -h, --help             display this help
 -V, --version          display version

For more details see gionice(1).`

//func oldmain() {
//var (
//data = 4
//set         = 0
//ioclass = ionice.IOPRIO_CLASS_BE
//c           = 0
//which       = 0
//who         = 0
//invalid_msg = ""
// )

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
	       	// convert string to int32 or fail
	           data = strtos32_or_err(optarg, "invalid class data argument");
	           set |= 1;
	           break;
	       case 'c':
	           if (isdigit(*optarg)) {
					// convert string to int32 or fail
	               ioclass = strtos32_or_err(optarg,
	                   "invalid class argument");
	           } else {
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
			   // convert string to int32 or fail
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
	           fmt.Println(versionString);
	           exit(EXIT_SUCCESS);
	       case 'h':
	           usage();
	       default:
	           fprintf(stderr, "Try '%s --help' for more information.\n", "gionice");
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
	        // gionice without options, print the current ioprio
	       ioprio_print(0, IOPRIO_WHO_PROCESS);
	   } else if (!set && who) {
	        // gionice -p|-P|-u ID [ID ...]
	       ioprio_print(which, who);
	       while argv[optind] != 0 {
	           which = strtos32_or_err(argv[optind], invalid_msg)
	           ioprio_print(which, who)
	       	optind++
	       }
	   } else if (set && who) {
	        // gionice -c CLASS -p|-P|-u ID [ID ...]
	       ioprio_setid(which, ioclass, data, who);
	       while argv[optind] != 0 {
	           which = strtos32_or_err(argv[optind], invalid_msg);
	           ioprio_setid(which, ioclass, data, who);
	           optind++
	       }
	   } else if (argv[optind]) {
	        // gionice [-c CLASS] COMMAND
	       ioprio_setid(0, ioclass, data, IOPRIO_WHO_PROCESS);
	       execvp(argv[optind], &argv[optind]);
	       static int EX_EXEC_FAILED = 126; // Program located, but not usable
	       static int EX_EXEC_ENOENT = 127; // Could not find program to exec
	       err(errno == ENOENT ? EX_EXEC_ENOENT : EX_EXEC_FAILED, "failed to execute %s", argv[optind]);

	   } else {
	       warnx("bad usage");
	       log.Fatalln("Try 'gionice --help' for more information.")
	   }*/

// experimental code follows:

//tolerant := false

// ion [-c CLASS] COMMAND
//ionice.SetIDPri(0, ioclass, data, ionice.IOPRIO_WHO_PROCESS, tolerant)

//execvp(argv[optind], &argv[optind])
//const EX_EXEC_FAILED = 126 // Program located, but not usable
//const EX_EXEC_ENOENT = 127 // Could not find program to exec
//log.Fatalf("failed to execute %s", commandString)
//err(errno == ENOENT ? EX_EXEC_ENOENT : EX_EXEC_FAILED, "failed to execute %s", argv[optind]);
//}

type Options struct {
	Class     string `short:"c" long:"class" description:"name or number of scheduling class, 0: none, 1: realtime, 2: best-effort, 3: idle" choice:"0" choice:"1" choice:"2" choice:"3" choice:"none" choice:"realtime" choice:"best-effort" choice:"idle"`
	ClassData int    `short:"n" long:"classdata" description:"priority (0..7) in the specified scheduling class, only for the realtime and best-effort classes" choice:"0" choice:"1" choice:"2" choice:"3" choice:"4" choice:"5" choice:"6" choice:"7" choice:"8" choice:"9"`
	PID       int    `short:"p" long:"pid" description:"act on these already running processes" value-name:"PID"`
	PGID      int    `short:"P" long:"pgid" description:"act on already running processes in these groups" value-name:"PGID"`
	Ignore    bool   `short:"t" long:"ignore" description:"ignore failures"`
	UID       int    `short:"u" long:"uid" description:"act on already running processes owned by these users" value-name:"UID"`
	Help      bool   `short:"h" long:"help" description:"display this help"`
	Version   bool   `short:"V" long:"version" description:"display version"`
	Args      struct {
		Command []string
	} //`positional-args:"yes" required:"yes"`
}

func main() {
	opts := &Options{}
	parser := flags.NewParser(opts, flags.PassAfterNonOption|flags.PassDoubleDash)
	args, err := parser.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	hasClass := parser.FindOptionByLongName("class").IsSet()
	hasClassData := parser.FindOptionByLongName("classdata").IsSet()
	hasPID := parser.FindOptionByLongName("pid").IsSet()
	hasPGID := parser.FindOptionByLongName("pgid").IsSet()
	hasUID := parser.FindOptionByLongName("uid").IsSet()

	var (
		data                         = 4
		set                          = 0
		ioclass  ionice.IOPRIO_CLASS = ionice.IOPRIO_CLASS_BE
		which                        = 0
		who                          = 0
		tolerant bool
	)

	switch {
	case hasClassData:
		set |= 1
	case hasClass:
		ioclass, err = ionice.Parse(opts.Class)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if ioclass < 0 {
			fmt.Fprintf(os.Stderr, "uknown scheduling class: '%s'\n", opts.Class)
		}
		set |= 2
	case hasPID:
		if who != 0 {
			fmt.Fprintln(os.Stderr, "can handle only one of pid, pgid or uid at once")
			os.Exit(1)

		}
		which = opts.PID
		who = ionice.IOPRIO_WHO_PROCESS
	case hasPGID:
		if who != 0 {
			fmt.Fprintln(os.Stderr, "can handle only one of pid, pgid or uid at once")
			os.Exit(1)
		}
		which = opts.PGID
		who = ionice.IOPRIO_WHO_PGRP
	case hasUID:
		if who != 0 {
			fmt.Fprintln(os.Stderr, "can handle only one of pid, pgid or uid at once")
			os.Exit(1)
		}
		which = opts.UID
		who = ionice.IOPRIO_WHO_USER
	case opts.Ignore:
		tolerant = true
	case opts.Version:
		fmt.Println(versionString)
		os.Exit(0)
	case opts.Help:
		fmt.Println(usageString)
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "Try 'gionice --help' for more information.\n")
		os.Exit(1)
	}
	switch ioclass {
	case ionice.IOPRIO_CLASS_NONE:
		if (set&1) != 0 && !tolerant {
			// warning
			fmt.Fprintln(os.Stderr, "ignoring given cass data for none class")
		}
		data = 0
	case ionice.IOPRIO_CLASS_RT, ionice.IOPRIO_CLASS_BE:
		break
	case ionice.IOPRIO_CLASS_IDLE:
		if (set&1) != 0 && !tolerant {
			// warning
			fmt.Fprintln(os.Stderr, "ignoring given class data for idle class")
		}
		data = 7
	default:
		if !tolerant {
			// warning
			fmt.Fprintf(os.Stderr, "unknown prio class %d\n", ioclass)
		}
	}

	if set == 0 && which == 0 && len(args) == 0 {
		// gionice without options, print the current ioprio
		ionice.Print(0, ionice.IOPRIO_WHO_PROCESS)
	} else if set == 0 && who != 0 {
		// gionice -p|-P|-u ID [ID ...]
		ionice.Print(which, who)
		for _, id := range args {
			if n, err := strconv.Atoi(id); err == nil { // success, arg is a number
				which = n
				ionice.Print(which, who)
			}
		}

	} else if set != 0 && who != 0 {
		// gionice -c CLASS -p|-P|-u ID [ID ...]
		ionice.SetIDPri(which, ioclass, data, who, tolerant)
		for _, id := range args {
			if n, err := strconv.Atoi(id); err == nil { // success, arg is a number
				which = n
				ionice.SetIDPri(which, ioclass, data, who, tolerant)
			}
		}
	} else if len(args) > 0 {
		// gionice [-c CLASS] COMMAND
		ionice.SetIDPri(0, ioclass, data, ionice.IOPRIO_WHO_PROCESS, tolerant)
		fmt.Println("TO IMPLEMENT: EXECUTE OTHER THINGS THAN JUST LS")
		var argv0 string = "/usr/bin/ls"
		var argv []string = []string{"/usr/bin/ls"}
		var envv []string = []string{}
		err := syscall.Exec(argv0, argv, envv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute %s\n", argv0)
			os.Exit(1)
		}
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, "bad usage")
		fmt.Fprintln(os.Stderr, "Try 'gionice --help' for more information.")
		os.Exit(1)
	}

	if hasClass {
		fmt.Printf("Class: %v %T\n", opts.Class, opts.Class)
	}

	fmt.Println("ClassData:", opts.ClassData)
	fmt.Println("PID:", opts.PID)
	fmt.Println("Args:", args)
}
