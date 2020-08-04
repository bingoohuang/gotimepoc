# gotimepoc

POC of go time in location

## Build

1. `GOOS=linux GOARCH=amd64 go install  -ldflags="-s -w" ./...`

## POC

```bash
[root@app01 vagrant]# mv /usr/bin/rigaga /usr/bin/rigaga-bak
[root@app01 vagrant]# cp /vagrant/gotimepoc /usr/bin/rigaga
[root@app01 vagrant]# /vagrant/gotimepoc
2020/08/04 13:04:48 env TZ:Asia/Shanghai
2020/08/04 13:04:48 time.Local:Asia/Shanghai
2020/08/04 13:04:48 Local Time:2020-08-04 13:04:48.384226731 +0800 CST m=+0.000089938
2020/08/04 13:04:48 UTC Time:2020-08-04 05:04:48.38423532 +0000 UTC
2020/08/04 13:04:48 Unix Seconds:1596517488
2020/08/04 13:04:48 UnixNano:1596517488384240829
UTC 2020/08/04 05:04:48 time.Local:Asia/Shanghai
UTC 2020/08/04 05:04:48 Local Time:2020-08-04 13:04:48.384244641 +0800 CST m=+0.000107841
UTC 2020/08/04 05:04:48 UTC Time:2020-08-04 05:04:48.384247826 +0000 UTC
UTC 2020/08/04 05:04:48 Unix Seconds:1596517488
UTC 2020/08/04 05:04:48 UnixNano:1596517488384251970
^C
[root@app01 vagrant]# service rigaga start
Starting the process rigaga [ OK ]
rigaga process was started [ OK ]
[root@app01 vagrant]# tail -f /var/log/rigaga/rigaga.log
2020/08/04 05:05:25 env TZ:
2020/08/04 05:05:25 time.Local:Local
2020/08/04 05:05:25 Local Time:2020-08-04 05:05:25.846327867 +0000 UTC m=+10.001344853
2020/08/04 05:05:25 UTC Time:2020-08-04 05:05:25.846336081 +0000 UTC
2020/08/04 05:05:25 Unix Seconds:1596517525
2020/08/04 05:05:25 UnixNano:1596517525846342122
UTC 2020/08/04 05:05:25 time.Local:Local
UTC 2020/08/04 05:05:25 Local Time:2020-08-04 05:05:25.846346397 +0000 UTC m=+10.001363371
UTC 2020/08/04 05:05:25 UTC Time:2020-08-04 05:05:25.846349571 +0000 UTC
UTC 2020/08/04 05:05:25 Unix Seconds:1596517525
UTC 2020/08/04 05:05:25 UnixNano:1596517525846354342
^C
[root@app01 vagrant]# ls -l /etc/localtime
-rw-r--r--. 1 root root 118 5月   5 2014 /etc/localtime
[root@app01 vagrant]# rm -f /etc/localtime
[root@app01 vagrant]# ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
[root@app01 vagrant]# ls -l /etc/localtime
lrwxrwxrwx 1 root root 33 8月   4 13:07 /etc/localtime -> /usr/share/zoneinfo/Asia/Shanghai
[root@app01 vagrant]# service rigaga restart
rigaga process is not running [ FAILED ]
Starting the process rigaga [ OK ]
rigaga process was started [ OK ]
[root@app01 vagrant]# tail -f /var/log/rigaga/rigaga.log
2020/08/04 13:08:02 env TZ:
2020/08/04 13:08:02 time.Local:Local
2020/08/04 13:08:02 Local Time:2020-08-04 13:08:02.408777281 +0800 CST m=+10.000302263
2020/08/04 13:08:02 UTC Time:2020-08-04 05:08:02.408785165 +0000 UTC
2020/08/04 13:08:02 Unix Seconds:1596517682
2020/08/04 13:08:02 UnixNano:1596517682408791511
UTC 2020/08/04 05:08:02 time.Local:Local
UTC 2020/08/04 05:08:02 Local Time:2020-08-04 13:08:02.408795984 +0800 CST m=+10.000320962
UTC 2020/08/04 05:08:02 UTC Time:2020-08-04 05:08:02.408799309 +0000 UTC
UTC 2020/08/04 05:08:02 Unix Seconds:1596517682
UTC 2020/08/04 05:08:02 UnixNano:159651768240880414
```

## 关于UTC和CST

[孙雪峰: 关于UTC和CST][1]

1. UTC 协调世界时，简称UTC, Coordinated Universal Time
1. CST可以为如下4个不同的时区的缩写：
    - 美国中部时间：Central Standard Time (USA) UT-6:00
    - 澳大利亚中部时间：Central Standard Time (Australia) UT+9:30
    - 中国标准时间：China Standard Time UT+8:00, 其实就是UTC+8，也是GMT+8小时
    - 古巴标准时间：Cuba Standard Time UT-4:00

[V__KING__:linux 关于GMT， CST， UTC的编程和设置][2]

1. GMT=UTC 
    - Greenwich Mean Time 格林尼治平均时
    - UTC和GMT都与英国伦敦的本地时相同，所以程序中UTC与GMT没什么不同。意思UTC=GMT是相等的
1. Unix时间戳：
    - 在计算机中看到的UTC时间都是从（1970年01月01日 0:00:00)开始计算秒数的。
    - 所看到的UTC时间那就是从1970年这个时间点起到具体时间共有多少秒。 这个秒数就是Unix时间戳。

[1]: https://blog.csdn.net/snowin1994/article/details/77530033
[2]: https://blog.csdn.net/V__KING__/article/details/79046976
[3]: https://www.cyberciti.biz/faq/centos-linux-6-7-changing-timezone-command-line/
[4]: https://www.vultr.com/docs/setup-timezone-and-ntp-on-centos-6

## Change the current timezone in CentOS 6 and older

[How To Change Timezone on a CentOS 6 and 7][3]

You need to use the ln command to set timezone on Centos 6. 

Type the following commands as root:

```bash
[root@app01 vagrant]# rm /etc/localtime
rm：是否删除普通文件 "/etc/localtime"？y
[root@app01 vagrant]# ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
[root@app01 vagrant]# date
2020年 08月 04日 星期二 12:36:31 CST
[root@app01 vagrant]# ls -l /etc/localtime
lrwxrwxrwx 1 root root 33 8月   4 12:36 /etc/localtime -> /usr/share/zoneinfo/Asia/Shanghai
```

## write the system time info into the hardware clock.

[Setup Timezone and NTP on CentOS 6][4]

`vi /etc/sysconfig/clock`

```
ZONE="Asia/Shanghai"
UTC=false
ARC=false
```

Write the system time into the hardware clock.

`hwclock --systohc --localtime`

Input `hwclock` to see the result.

![image](https://user-images.githubusercontent.com/1940588/89253930-8f696c00-d650-11ea-82fe-aad1882dceed.png)
