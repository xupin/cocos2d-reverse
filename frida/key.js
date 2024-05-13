/// <reference path='./index.d.ts'/>

Java.perform(function () {
	function intercept() {
		const libName = "libcocos2dlua.so";
		const funcName = "_Z13xxtea_decryptPhjS_jPj";
		const ptr = Module.findExportByName(libName, funcName);
		if (!ptr) {
			setTimeout(intercept, 500);
			return;
		}
		Interceptor.attach(ptr, {
			onEnter: function (args) {
				console.log("----------------BEGIN----------------");
				console.log("Key:");
				console.log(Memory.readUtf8String(args[2]));
			},
			onLeave: function (retval) {
				console.log("Encrypted:");
				console.log(retval.readByteArray(16));
				console.log("-----------------END-----------------");
			},
		});
	}
	intercept();
});
