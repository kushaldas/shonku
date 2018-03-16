Installation
=============

Install golang
---------------

Download golang from `here <https://golang.org/doc/install?download=go1.8.3.linux-amd64.tar.gz>`_ , extract go directory
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
    $ go get golang.org/x/net/html

Get the latest Shonku code
===========================

Use git to clone the repository
::

  $ git clone https://github.com/kushaldas/shonku.git

Building the source
===================

::

    $ make

This should create a binary called `shonku.bin`.

Rebuilding bindata for default theme
=====================================

In case you make any changes to the default theme, you want those changes inside
the binary file also. For that issue the following command before building the
binary.

::

	$ go-bindata assets/... templates/

.. note::
	Remember to install go-bindata from `here <https://github.com/jteeuwen/go-bindata>`_.
