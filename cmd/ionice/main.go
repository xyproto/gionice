package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/xyproto/gion"
)

func usage() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(" ionice [options] -p <pid>...")
	fmt.Println(" ionice [options] -P <pgid>...")
	fmt.Println(" ionice [options] -u <uid>...")
	fmt.Println(" ionice [options] <command>")
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
		ioclass = gion.IOPRIO_CLASS_BE
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
		           printf("%s %s\n", "ionice", "1.0.0");
		           exit(EXIT_SUCCESS);
		       case 'h':
		           usage();
		       default:
		           fprintf(stderr, "Try '%s --help' for more information.\n", "ionice");
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
		        // ionice without options, print the current ioprio
		       ioprio_print(0, IOPRIO_WHO_PROCESS);
		   } else if (!set && who) {
		        // ionice -p|-P|-u ID [ID ...]
		       ioprio_print(which, who);
		       while argv[optind] != 0 {
		           which = strtos32_or_err(argv[optind], invalid_msg)
		           ioprio_print(which, who)
		       	optind++
		       }
		   } else if (set && who) {
		        // ionice -c CLASS -p|-P|-u ID [ID ...]
		       ioprio_setid(which, ioclass, data, who);
		       while argv[optind] != 0 {
		           which = strtos32_or_err(argv[optind], invalid_msg);
		           ioprio_setid(which, ioclass, data, who);
		           optind++
		       }
		   } else if (argv[optind]) {
		        // ionice [-c CLASS] COMMAND
		       ioprio_setid(0, ioclass, data, IOPRIO_WHO_PROCESS);
		       execvp(argv[optind], &argv[optind]);
		       static int EX_EXEC_FAILED = 126; // Program located, but not usable
		       static int EX_EXEC_ENOENT = 127; // Could not find program to exec
		       err(errno == ENOENT ? EX_EXEC_ENOENT : EX_EXEC_FAILED, "failed to execute %s", argv[optind]);

		   } else {
		       warnx("bad usage");
		       log.Fatalln("Try 'ionice --help' for more information.")
		   }*/

	// experimental code follows:

	tolerant := false

	// ion [-c CLASS] COMMAND
	gion.SetIDPri(0, ioclass, data, gion.IOPRIO_WHO_PROCESS, tolerant)

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
