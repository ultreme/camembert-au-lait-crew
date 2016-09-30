#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import re

from setuptools import setup, find_packages
from recettator import __version__

MODULE_NAME = 'recettator'

DEPENDENCIES = []
TEST_DEPENDENCIES = []

setup(
    name=MODULE_NAME,
    version=__version__,
    description='Recettator',
    long_description='Generateur de recettes de cuisine',
    author='CALC',
    author_email='contact@camembertaulaitcrew.biz',
    url='http://camembertaulaitcrew.biz',
    packages=find_packages(),
    install_requires=DEPENDENCIES,
    tests_require=DEPENDENCIES + TEST_DEPENDENCIES,
    dependency_links=[],
    test_suite=MODULE_NAME + '.tests',
    classifiers=[
        # As from http://pypi.python.org/pypi?%3Aaction=list_classifiers
        'Development Status :: 2 - Pre-Alpha',
        'License :: Other/Proprietary License',
        'Operating System :: POSIX',
        'Operating System :: MacOS',
        'Operating System :: Unix',
        'Programming Language :: Python',
        'Topic :: Internet',
    ],
    extras_require={},
    entry_points={
        'console_scripts': [
            'recettator-cli = recettator.bin:recettator_cli',
        ],
    },
)
