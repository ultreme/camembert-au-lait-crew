# -*- coding: utf-8 -*-

import random
import re

from ..utils import parts_to_string

from .main_ingredients import all_items as main_ingredients_items
from .secondary_ingredients import all_items as secondary_ingredients_items
from .seasonings import all_items as seasonings_items
from .methods import all_items as methods_items
from .recettes import all_items as recettes_items
from .ingredient_methods import all_items as ingredient_methods_items


class ItemGroup(object):
    def __init__(self, items, shuffle=True, seed=None):
        self._random = random.Random()
        if seed:
            self._random.seed(seed)

        self.availables = items
        if shuffle:
            self._random.shuffle(self.availables)

    def pick_random(self, recycle_item=False, parent=None, **kwargs):
        for item in self.availables:
            found = True
            for k, v in kwargs.items():
                if not hasattr(item, k):
                    found = False
                    break
                item_value = getattr(item, k)
                if not item_value:
                    break

                # Prepare for matching
                if v == 'any':
                    v = '.*'
                if item_value == 'any':
                    item_value = '.*'
                if v[0] != '^':
                    v = '^{}$'.format(v)
                if item_value[0] != '^':
                    item_value = '^{}$'.format(item_value)

                if item_value != v and \
                   not re.match(item_value, v) and \
                   not re.match(v, item_value):
                    found = False
                    break
            if found:
                self.availables.remove(item)
                if recycle_item:
                    self.availables.append(item)
                seed = self._random.random()
                return item(self, parent=parent, seed=seed)
        return None


def all_items(seed=None):
    items = []
    items += main_ingredients_items()
    items += secondary_ingredients_items()
    items += seasonings_items()
    items += methods_items()
    items += recettes_items()
    items += ingredient_methods_items()
    return ItemGroup(items, seed=seed)
