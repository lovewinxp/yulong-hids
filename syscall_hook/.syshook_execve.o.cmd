cmd_/usr/local/go/src/yulong-hids/syscall_hook/syshook_execve.o := ld -m elf_x86_64   -r -o /usr/local/go/src/yulong-hids/syscall_hook/syshook_execve.o /usr/local/go/src/yulong-hids/syscall_hook/syscall_hook.o /usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o ; scripts/mod/modpost /usr/local/go/src/yulong-hids/syscall_hook/syshook_execve.o