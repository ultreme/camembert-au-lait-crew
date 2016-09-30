# -*- coding: utf-8 -*-

import random


def parts_to_string(parts):
    return ' '.join([str(part) for part in parts]).replace("' ", "'")


def genderization(options, constraints):
    for option in options:
        checks = option.values()[0]
        matches = True
        for k, v in checks.items():
            if constraints[k] != v:
                matches = False
                break
        if matches:
            return option.keys()[0]


def pick_random(lst, random_=None):
    if not random_:
        random_ = random

    return lst[random_.randrange(len(lst))]
