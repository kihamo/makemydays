Cross compiler
==============

Переменные окружения Go
-----------------------
    $ go env
    GOARCH="386"
    GOBIN=""
    GOCHAR="8"
    GOEXE=""
    GOHOSTARCH="386"
    GOHOSTOS="linux"
    GOOS="linux"
    GOPATH=""
    GORACE=""
    GOROOT="/usr/lib/go"
    GOTOOLDIR="/usr/lib/go/pkg/tool/linux_386"
    TERM="dumb"
    CC="gcc"
    GOGCCFLAGS="-g -O2 -fPIC -m32 -pthread"
    CXX="g++"
    CGO_ENABLED="1"
    
Компиляция
---------------------------------------
    go build -gccgoflags "-m32" -compiler gccgo hello.go
    
* **-m32** - для target архитектуры i386
* **-m64** - для target архитектуры amd64

Если host = target, то параметр **-gccgoflags "-m32"** необходимо убрать.

Если добавить параметр **-n** то будет показан план выполнения команд, без фактической компиляции.

    $ go build -n -gccgoflags "-m32" -compiler gccgo hello.go
    
    #
    # command-line-arguments
    #
    
    mkdir -p $WORK/command-line-arguments/_obj/
    cd /home/kihamo/go
    gccgo -I $WORK -c -g -m64 -fgo-relative-import-path=_/home/kihamo/go -o $WORK/command-line-arguments/_obj/main.o -m32 ./hello.go
    ar cru $WORK/libcommand-line-arguments.a $WORK/command-line-arguments/_obj/main.o
    cd .
    gccgo -o hello $WORK/command-line-arguments/_obj/main.o -Wl,-( -m64 -Wl,-) -m32
    
Если на target машине не установлена **libgo**, то ее необходимо "вшить" в исполняемый файл, например так:

    go build -gccgoflags "-m32 -static-libgo" -compiler gccgo hello.go
    
Инструменты для отладки
-----------------------
Посмотреть информацию об исполняемом файле можно утилитой **file**

    $ file hello 
    hello: ELF 32-bit LSB  executable, Intel 80386, version 1 (SYSV), dynamically linked (uses shared libs), for GNU/Linux 2.6.24, BuildID[sha1]=01ba9b5c63d74e90441c756f6a9c05fb3dcffd5c, not stripped
    
Динамические библиотеки, от которых зависит исполняемый файл можно посмотреть командой **ldd** или **readelf**

    $ readelf -d hello                                                                                                                         
    
    Динамический раздел со смещением 0x10ab90 содержит 27 элементов:
      Тег        Тип                          Имя/Знач
     0x00000003 (PLTGOT)                     0x8153c9c
     0x00000002 (PLTRELSZ)                   1824 (байт)
     0x00000017 (JMPREL)                     0x8049dd4
     0x00000014 (PLTREL)                     REL
     0x00000011 (REL)                        0x8049dc4
     0x00000012 (RELSZ)                      16 (байт)
     0x00000013 (RELENT)                     8 (байт)
     0x00000015 (DEBUG)                      0x0
     0x00000006 (SYMTAB)                     0x80481ac
     0x0000000b (SYMENT)                     16 (байт)
     0x00000005 (STRTAB)                     0x804903c
     0x0000000a (STRSZ)                      2542 (байт)
     0x6ffffef5 (GNU_HASH)                   0x8049a2c
     0x00000001 (NEEDED)                     Совм. исп. библиотека: [libpthread.so.0]
     0x00000001 (NEEDED)                     Совм. исп. библиотека: [libgcc_s.so.1]
     0x00000001 (NEEDED)                     Совм. исп. библиотека: [libc.so.6]
     0x00000001 (NEEDED)                     Совм. исп. библиотека: [ld-linux.so.2]
     0x0000000c (INIT)                       0x804a4f4
     0x0000000d (FINI)                       0x80c7f2c
     0x0000001a (FINI_ARRAY)                 0x8165f98
     0x0000001c (FINI_ARRAYSZ)               4 (байт)
     0x00000019 (INIT_ARRAY)                 0x8165f9c
     0x0000001b (INIT_ARRAYSZ)               12 (байт)
     0x6ffffff0 (VERSYM)                     0x8049a50
     0x6ffffffe (VERNEED)                    0x8049c24
     0x6fffffff (VERNEEDNUM)                 4
     0x00000000 (NULL)                       0x0
     
Список всех библиотек, установленных в системе можно посмотреть командой

    ldconfig -p
     
Особенности компиляции SQLite
-----------------------------
* [https://github.com/mattn/go-sqlite3/pull/123]
* [https://github.com/mattn/go-sqlite3/issues/20]

Проблема в следующем

    This problem does not occur with the gc compiler because it runs a totally different sequence of compiler commands, with different output filenames (for example the cgo-generated files are compiled into _go_.6 rather than sqlite3.o).
    
    To work around the problem, rename sqlite3.{h,c} to sqlite3-orig.{h,c}, and modify the #includes in sqlite3.go, backup.go, and sqlite3ext.h.
    
Ubuntu
======
Установить дополнительный пакет **gccgo-multilib** без него невозможна компиляция при разных host и target системах.
    
Openwrt
=======

- Применить патч для toolchain
- Собрать toolchain
- Убедиться, что gccgo скомпилировался. В папке **staging_dir/toolchain-mips_34kc_gcc-4.8-linaro_eglibc-2.19/bin** должен присутствовать исполняемый файл **mips-openwrt-linux-gnu-gccgo**

Запускаем сборку ( *не забыть проверить версии библиотек, которые учавствуют в путях* )

    $ ./mips-openwrt-linux-gnu-gccgo -Wl,--verbose,-R,/data/openwrt/barrier-breaker/staging_dir/toolchain-mips_34kc_gcc-4.8-linaro_eglibc-2.19/lib/gcc/mips-openwrt-linux-gnu/4.8.3 -L/data/openwrt/barrier-breaker/staging_dir/toolchain-mips_34kc_gcc-4.8-linaro_eglibc-2.19/lib -static-libgo /data/go/hello.go -o /data/go/hello-mips

Сборка Go
-----------
Toolchain с gccgo собирается только под **eglibc** ( потому что **uClibc** не предоставляет make/get/setcontext() функции [https://groups.google.com/forum/#!topic/golang-nuts/fAElwJu-QUM]). При сборке ее надо выбрать явно, так как по-умолчанию стоит **uClibc**

    $ make menuconfig
    
    [*] Advanced configuration options (for developers)  --->
        [*] Toolchain Options  --->
                C Library implementation (Use uClibc)  --->
                    (X) Use eglibc
                    
Далее необходимо включить сам **gccgo** и библиотеку **libgo** в сборку

    $ make menuconfig
    
    [*] Advanced configuration options (for developers)  --->
        [*] Toolchain Options  --->
            [*]   Build/install gccgo compiler?
            
    Base system  --->
        <*> libgo................................................. Go support library