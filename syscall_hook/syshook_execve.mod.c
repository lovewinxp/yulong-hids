#include <linux/module.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

MODULE_INFO(vermagic, VERMAGIC_STRING);

struct module __this_module
__attribute__((section(".gnu.linkonce.this_module"))) = {
 .name = KBUILD_MODNAME,
 .init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
 .exit = cleanup_module,
#endif
 .arch = MODULE_ARCH_INIT,
};

static const struct modversion_info ____versions[]
__used
__attribute__((section("__versions"))) = {
	{ 0x14522340, "module_layout" },
	{ 0x9853b40d, "d_path" },
	{ 0x4f1939c7, "per_cpu__current_task" },
	{ 0x5a34a45c, "__kmalloc" },
	{ 0xd691cba2, "malloc_sizes" },
	{ 0x8ce3169d, "netlink_kernel_create" },
	{ 0xde0bdcff, "memset" },
	{ 0xd9a9bb30, "getname" },
	{ 0xea147363, "printk" },
	{ 0xd4defbf4, "netlink_kernel_release" },
	{ 0xa1c76e0a, "_cond_resched" },
	{ 0xb4390f9a, "mcount" },
	{ 0x1e6d26a8, "strstr" },
	{ 0xa77d88f6, "strnlen_user" },
	{ 0x1c740bd6, "init_net" },
	{ 0x8b6c553c, "fput" },
	{ 0x2b4c47df, "putname" },
	{ 0x25421969, "__alloc_skb" },
	{ 0x312919, "netlink_broadcast" },
	{ 0x3d75cbcf, "kfree_skb" },
	{ 0x6d334118, "__get_user_8" },
	{ 0x27f96468, "pv_cpu_ops" },
	{ 0x2044fa9e, "kmem_cache_alloc_trace" },
	{ 0xba497f13, "loops_per_jiffy" },
	{ 0x37a0cba, "kfree" },
	{ 0x9edbecae, "snprintf" },
	{ 0x207b7e2c, "skb_put" },
	{ 0x3302b500, "copy_from_user" },
	{ 0x64159ec7, "open_exec" },
	{ 0xdcb0349b, "sys_close" },
};

static const char __module_depends[]
__used
__attribute__((section(".modinfo"))) =
"depends=";


MODULE_INFO(srcversion, "FFA468F76CF0EB5AD2ED9F4");

static const struct rheldata _rheldata __used
__attribute__((section(".rheldata"))) = {
	.rhel_major = 6,
	.rhel_minor = 5,
};
