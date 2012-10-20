from __future__ import with_statement
from fabric.api import *
from fabric.contrib.console import confirm
from fabric.context_managers import cd, lcd


def compile_less():
	local('lessc --compress static/less/main.less > static/css/main.css')
