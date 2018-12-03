cmd_/usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o := gcc -Wp,-MD,/usr/local/go/src/yulong-hids/syscall_hook/.syscall_stub_hook.o.d  -nostdinc -isystem /usr/lib/gcc/x86_64-redhat-linux/4.4.7/include -Iinclude  -I/usr/src/kernels/2.6.32-431.el6.x86_64/include/uapi -I/usr/src/kernels/2.6.32-431.el6.x86_64/arch/x86/include -include /usr/src/kernels/2.6.32-431.el6.x86_64/include/linux/kconfig.h -D__KERNEL__ -D__ASSEMBLY__ -m64 -DCONFIG_AS_CFI=1 -DCONFIG_AS_CFI_SIGNAL_FRAME=1 -DCONFIG_AS_CFI_SECTIONS=1 -DCONFIG_AS_AVX=1  -gdwarf-2      -DMODULE -c -o /usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o /usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.S

deps_/usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o := \
  /usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.S \
  /usr/src/kernels/2.6.32-431.el6.x86_64/include/linux/kconfig.h \
    $(wildcard include/config/h.h) \
    $(wildcard include/config/.h) \
    $(wildcard include/config/booger.h) \
    $(wildcard include/config/foo.h) \

/usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o: $(deps_/usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o)

$(deps_/usr/local/go/src/yulong-hids/syscall_hook/syscall_stub_hook.o):
