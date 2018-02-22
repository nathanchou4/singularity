/*
Copyright (c) 2018, Sylabs, Inc. All rights reserved.
This software is licensed under a 3-clause BSD license.  Please
consult LICENSE file distributed with the sources of this project regarding
your rights to use or distribute this software.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shell called")
	},
}

func init() {
	singularityCmd.AddCommand(shellCmd)

    // ultimately all of this should go into a seperate function to be shared
    // between exec, run, and shell
	shellCmd.PersistentFlags().String("overlay", "o", "")
	shellCmd.PersistentFlags().String("shell", "s", "")
    shellCmd.PersistentFlags().BoolP("userns", "u", false, "")
    shellCmd.PersistentFlags().BoolP("readonly", "r", false, "")
    shellCmd.PersistentFlags().String("home", "H", "")
	shellCmd.PersistentFlags().String("workdir", "W", "")
    shellCmd.PersistentFlags().String("scratchdir", "S", "")
	shellCmd.PersistentFlags().String("app", "a", "")
	shellCmd.PersistentFlags().String("bind", "B", "")
    shellCmd.PersistentFlags().BoolP("contain", "c", false, "")
    shellCmd.PersistentFlags().BoolP("containall", "C", false, "")
    shellCmd.PersistentFlags().BoolP("cleanenv", "e", false, "")
    shellCmd.PersistentFlags().BoolP("pid", "p", false, "")
    shellCmd.PersistentFlags().BoolP("ipc", "i", false, "")
    shellCmd.PersistentFlags().BoolP("uts", "", false, "")
    shellCmd.PersistentFlags().BoolP("hostname", "", false, "")
	shellCmd.PersistentFlags().String("pwd", "p", "")
    shellCmd.PersistentFlags().BoolP("nv", "", false, "")
    shellCmd.PersistentFlags().BoolP("fakeroot", "f", false, "")
    shellCmd.PersistentFlags().BoolP("keep-privs", "", false, "")
    shellCmd.PersistentFlags().BoolP("no-privs", "", false, "")
    shellCmd.PersistentFlags().BoolP("add-caps", "", false, "")
    shellCmd.PersistentFlags().BoolP("drop-caps", "", false, "")
    shellCmd.PersistentFlags().BoolP("allow-setuid", "", false, "")

    shellCmd.SetHelpTemplate(`
USAGE: singularity [...] shell [shell options...] <container path>
Obtain an interactive shell (/bin/bash) within the container image.


SHELL OPTIONS:
    -a|--app            Include an installed container app in the environment
    -B|--bind <spec>    A user-bind path specification.  spec has the format
                        src[:dest[:opts]], where src and dest are outside and
                        inside paths.  If dest is not given, it is set equal
                        to src.  Mount options ('opts') may be specified as
                        'ro' (read-only) or 'rw' (read/write, which is the
                        default). This option can be called multiple times.
    -c|--contain        Use minimal /dev and empty other directories (e.g. /tmp
                        and /home/ubuntu) instead of sharing filesystems on your host
    -C|--containall     Contain not only file systems, but also PID and IPC
    -e|--cleanenv       Clean environment before running container
    -H|--home <spec>    A home directory specification.  spec can either be a
                        src path or src:dest pair.  src is the source path
                        of the home directory outside the container and dest
                        overrides the home directory within the container
    -i|--ipc            Run container in a new IPC namespace
    -j|--join           Join a running named instance of the given container 
                        image
    -n|--net            Run container in a new network namespace (loopback is
                        only network device active)
    --nv                Enable experimental Nvidia support
    -o|--overlay        Use a persistent overlayFS via a writable image
    -p|--pid            Run container in a new PID namespace (creates child)
    --pwd               Initial working directory for payload process inside
                        the container
    -S|--scratch <path> Include a scratch directory within the container that
                        is linked to a temporary dir (use -W to force location)
    -s|--shell <shell>  Path to program to use for interactive shell
    -u|--userns         Run container in a new user namespace (this allows
                        Singularity to run completely unprivileged on recent
                        kernels and doesn't support all features)
    -W|--workdir        Working directory to be used for /tmp, /var/tmp and
                        /home/ubuntu (if -c/--contain was also used)
    -r|--readonly       By default all Singularity containers are available as
                        writable if they contain a writable partition. This
                        option makes the file system accessible as read only.

CONTAINER FORMATS SUPPORTED:
    *.sqsh              SquashFS format.  Native to Singularity 2.4+
    *.img               This is the native Singularity image format for all
                        Singularity versions < 2.4.
    *.tar*              Tar archives are exploded to a temporary directory and
                        run within that directory (and cleaned up after). The
                        contents of the archive is a root file system with root
                        being in the current directory. Compression suffixes as
                        '.gz' and '.bz2' are supported.
    directory/          Container directories that contain a valid root file
                        system.
    instance://*        A local running instance of a container. (See the
                        instance command group.)
    shub://*            A container hosted on Singularity Hub
    docker://*          A container hosted on Docker Hub


EXAMPLES:

$ singularity shell /tmp/Debian.img
Singularity/Debian.img> pwd
/home/gmk/test
Singularity/Debian.img> exit

$ singularity shell -C /tmp/Debian.img
Singularity/Debian.img> pwd
/home/gmk
Singularity/Debian.img> ls -l
total 0
Singularity/Debian.img> exit

$ sudo singularity shell -w /tmp/Debian.img
$ sudo singularity shell --writable /tmp/Debian.img

$ singularity shell instance://my_instance 

$ singularity shell instance://my_instance
Singularity: Invoking an interactive shell within container...
Singularity container:~> ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
ubuntu       1     0  0 20:00 ?        00:00:00 /usr/local/bin/singularity/bin/sinit
ubuntu       2     0  0 20:01 pts/8    00:00:00 /bin/bash --norc
ubuntu       3     2  0 20:02 pts/8    00:00:00 ps -ef

For additional help, please visit our public documentation pages which are
found at:

http://singularity.lbl.gov/
    `)

    shellCmd.SetUsageTemplate(`
USAGE: singularity [...] shell [shell options...] <container path>
    `)
}
