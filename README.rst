masa-snap
=========

This is a sample project to demonstrate how to use go to generate yaml
from a :code:`struct`.  It's overly simple since I have no idea what goes
into a snap project, and Masahiro's actual requirements will be much more
complex than I want to deal with right now.

Usage
-----
If a description isn't provided, data from stdin will be used.

.. code-block::

    $ ./masa-snap -n "masa" -s "a sample snap" <<EOF
    > This is a longer description.
    > It can be multiple lines.
    > Hello, Masahiro.
    >
    > Goodbye, Masahiro.
    > EOF
    name: masa
    version: 1.0.0
    summary: a sample snap
    description: |
        This is a longer description.
        It can be multiple lines.
        Hello, Masahiro.

        Goodbye, Masahiro.

