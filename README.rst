masa-snap
=========

This is a sample project to demonstrate how to use go to generate yaml
from a :code:`struct`.  It's overly simple since I have no idea what goes
into a snap project, and Masahiro's actual requirements will be much more
complex than I want to deal with right now.

Usage
-----
There are currently two commands:

  * :code:`generate` to generate new yaml
  * :code:`edit` to edit existing yaml

Help is available via the :code:`--help` flag.  When generating yaml, a
missing description will be read via stdin.

.. code-block::

    $ ./masa-snap generate -n "masa" -s "a sample snap" <<EOF
    > This is a longer description.
    > It can be multiple lines.
    > Hello, Masahiro.
    >
    > Goodbye Masahiro.
    > EOF
    name: masa
    version: 1.0.0
    summary: a sample snap
    description: |
        This is a longer description.
        It can be multiple lines.
        Hello, Masahiro.

        Goodbye Masahiro.

The layout section is also supported:

.. code-block::

    $ ./masa-snap generate -n "masa" -s "a sample snap" \
        -f '/etc/hosts:$SNAP_DATA/etc/hosts' \
        -b '/etc/apache/hosts.d:$SNAP/etc/apache/hosts.d'  <<EOF
    This is a longer description.
    It can be multiple lines.
    Hello, Masahiro.

    Goodbye Masahiro.
    EOF
    name: masa
    version: 1.0.0
    summary: a sample snap
    description: |
        This is a longer description.
        It can be multiple lines.
        Hello, Masahiro.

        Goodbye Masahiro.
    layout:
        /etc/apache/hosts.d:
            bind: $SNAP/etc/apache/hosts.d
        /etc/hosts:
            bind-file: $SNAP_DATA/etc/hosts


When editing yaml, the same flags are supported and all are optional:

.. code-block::

    $ ./masa-snap generate -n "masa" -s "a sample snap" -d "a long description" |
      ./masa-snap edit --version "1.2.3"
    name: masa
    version: 1.2.3
    summary: a sample snap
    description: a long description
