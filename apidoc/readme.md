

1、安装nodejs


2、安装gitbook
  利用npm安装gitbook，安装步骤可见 Install GitBook。
npm install gitbook -g
npm install gitbook-cli -g

gitbook-cli 是构建book的工具，gitbook init时需要。

安装后的gitbook会在node安装bin目录下，所以要使用gitbook目录可以使用ln -s命令做好软连接。

3、gitbook使用
gitbook 的基本用法非常简单，基本上就只有两步：

使用 gitbook init 初始化书籍目录；
使用 gitbook serve 编译书籍；
mkdir book
cd book
gitbook init

README.md 和 SUMMARY.md 是两个必须文件，README.md 是对书籍的简单介绍：
# cat book/README.md 
# Introduction
 
This is my first book.

SUMMARY.md 是书籍的目录结构。内容如下：
# cat book/SUMMARY.md 
# SUMMARY
 
* [Chapter1](chapter1/README.md)
  * [Section1.1](chapter1/section1.1.md)
  * [Section1.2](chapter1/section1.2.md)
* [Chapter2](chapter2/README.md)

书籍目录结构创建完成以后，就可以使用 gitbook serve 来编译和预览书籍了。

gitbook serve 命令实际上会首先调用 gitbook build 编译书籍，然后启动web服务，监听在本地的 4000 端口。
gitbook 为我们创建了书籍目录结构后，就可以向其中添加内容了，文件的编写使用 markdown 语法，在文件修改过程中，每一次保存文件，gitbook serve 都会自动重新编译，所以可以实时通过浏览器来查看最新的书籍效果！
