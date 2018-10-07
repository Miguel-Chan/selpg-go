# selpg-go 说明

Selpg-go 是使用 Go 语言编写的一个[ Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) , 它在命令行中的使用方法如下：

`./selpg-go -sstart_page -eend_page [ -f | -llines_per_page ][ -ddest ] [ in_filename ]`

其中， `-s` 后接的是要打印的页面范围的起始页号码（从1开始计算），`-e` 后接要打印的范围的最后一页页码，这两个是必须指定的选项。

后面的可选选项中，`-f` 和`-l` 选项互斥，不可同时指定，`-f`选项表示程序从输入中会将ASCII 换页字符（十进制数值为 12，用“\f”表示）作为换页的标记；`-l` 选项后接每页的行数，表示程序将把输入中每个指定行数的区间视为一页（若未指定则默认每72行一页）。

`-d` 选项指定打印的输出目的地，若未指定则默认输出至标准输出。指定后程序将会调用`lp` 命令，并将数据通过标准输入传入至程序。

最后一个`[in_filename]` 参数用于指定输入文件名，若未指定则程序将从程序的标准输入获得输入数据。

## 设计

selpg-go 使用了 cobra 库，创建了一个根 command，并使用 cobra 库中的 flags 来配置程序命令行选项，利用`MarkFlagRequired` 来令`-s ` 和`-e` 选项为强制选项。

在 selpg 模块中，`GetPage` 函数从指定的输入中读取数据，并将数据写入到传入的`io.Writer`中。若数据量未达到指定的输出要求量，将返回一个错误。

`Run` 函数将根据指定的选项将对应的输出用`io.Writer`传入`GetPage` 并将返回的错误值输出至程序 stderr 中。若程序有指定输出目的地，则会调用`lp` 命令并将输出的内容传入`lp` 命令, 同时使用`io.Pipe` 来将`lp`的 stderr 的输出输出到程序的 stderr 中。

## 测试

测试主要测试了`GetPage` 函数，测试了使用行数时正常运行，起始页 大于总页数与结束页大于总页数的情况，以及使用\f 为分页符的运行情况。测试代码位于 selpg/selpg_test.go 。