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

It also supports the layout section:

.. code-block::

    $ ./masa-snap -n "masa" -s "a sample snap" -f '/etc/hosts:$SNAP_DATA/etc/hosts' -b '/etc/apache/hosts.d:$SNAP/etc/apache/hosts.d' <<EOF
    This is a longer description.
    It can be multiple lines.
    Hello, Masahiro.

    Goodbye, Masahiro.
    EOF
    name: masa
    version: 1.0.0
    summary: a sample snap
    description: |
        This is a longer description.
        It can be multiple lines.
        Hello, Masahiro.

        Goodbye, Masahiro.
    layout:
        /etc/apache/hosts.d:
            bind: $SNAP/etc/apache/hosts.d
        /etc/hosts:
            bind-file: $SNAP_DATA/etc/hosts
