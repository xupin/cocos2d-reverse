/// <reference path='./index.d.ts'/>

Java.perform(function () {
  function intercept() {
    const libSoName = "libcocos2dlua.so";
    const funcName = "_Z13xxtea_decryptPhjS_jPj";
    const ptr = Module.findExportByName(libSoName, funcName);
    if (!ptr) {
      setTimeout(intercept, 500);
      return;
    }
    Interceptor.attach(ptr, {
      onEnter: function (args) {
        console.log("----------------BEGIN----------------");
        console.log("Key:");
        console.log(args[2].readCString(16));
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
