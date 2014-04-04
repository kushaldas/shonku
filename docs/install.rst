Installation
=============

Install golang
---------------

Download golang from `here <http://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz>`_ , extract go directory
under your home directory.

::

    $ mkdir ~/gocode

Now write the following lines in your ~/.bashrc file.
::

    export PATH=$PATH:~/go/bin
    export GOPATH=~/gocode/
    export GOROOT=~/go/

and then ::

    $ source ~/.bashrc

Install the dependencies
-------------------------

After golang installation, get the dependent libraries. 
::

    $ go get github.com/russross/blackfriday
    $ go get github.com/gorilla/feeds
    $ go get code.google.com/p/go-sqlite/go1/sqlite3

Building the source
===================

::

    $ make

This should create a binary called `shonku.bin`.
