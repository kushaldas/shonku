Usage
======

.. note:: We support only `Markdown <http://daringfireball.net/projects/markdown/>`_ format.
  You can learn it very fast and use any text editor to edit.


Creating new site
-------------------

First go to an empty directory and run the following command.
::

    $ ./shonku.bin -new_site

This will create the required files and directories for shonku to run.

Writing a new post
-------------------

To write a new blog post do the following command.
::

    $  ./shonku.bin -new
    Enter the title of the post: Hello World
    Your new post is ready at ./posts/hello-world.md

As the output shows your first blog post is ready. Make the changes as you
want in that file.

.. note:: Remember to keep a blank line at the end of each post or page.


Building your post
------------------

Just run the following command.
::

    $  ./shonku.bin
    {SITE AUTHOR SITE TITLE http://localhost/ Copyright 2014 yourdisqus author@email Description of the site URL for logo [{/pages/about-me.html About} {/categories/ Categories} {/archive.html Archive}]}
    Building post: ./posts/hello-world.md

You can check the output directory for the finished blog post.

Force rebuild of the whole site
--------------------------------

::

    $ ./shonku.bin -force

The above command will rebuild the whole site. You may want to use this command when
you make any change to your theme or configuration file.

Details of each post
---------------------

When you create a new post it will contain something similar to the details below

::

  <!--
  .. title: Hello World
  .. slug: hello-world
  .. date: 2014-05-19T12:15:41+05:30
  .. tags: Blog
  .. link:
  .. description:
  .. type: text
  -->

  Write your post here.

You can add more tags to the post, they are comma separated. This post format is
same of Nikola v7.x, that means it is interchangable between these two blog engines.


Individual author per post
---------------------------

We can have individual author for each post. Just add the following line in any of the 
post where you want a different author (do it in the comments as show in above).

::

  .. author: AUTHOR NAME




