Usage
======

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

