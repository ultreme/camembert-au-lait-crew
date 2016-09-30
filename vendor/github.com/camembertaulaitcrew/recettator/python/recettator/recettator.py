#!/usr/bin/env python

from collections import OrderedDict
from math import ceil
from random import Random, randrange
import os
import pkg_resources
import sys

from .items import all_items
from .utils import parts_to_string, pick_random


class Recettator:
    """ Recettator class. """

    def __init__(self, seed=None):
        self._data = None
        self._items = []
        self._amounts = {}

        # random
        self._random = Random()
        if not seed:
            seed = randrange(10000)
        self.seed = seed
        self._random.seed(seed)

        # db
        self._db = all_items(seed=self._random.randrange(1000))

    @property
    def is_valid(self):
        if not self._amounts:
            return False
        return self._amounts['main_ingredient'] + \
            self._amounts['secondary_ingredient'] > 1

    @property
    def items(self):
        if not len(self._items):
            while not self.is_valid:
                self._amounts = OrderedDict([
                    ('recette', 1),
                    ('main_ingredient', self._random.randrange(4) - 1),
                    ('secondary_ingredient', self._random.randrange(7) - 1),
                    ('seasoning', self._random.randrange(5) - 1),
                    ('method', int(self._random.randrange(10) < 4)),
                ])

            for k, v in self._amounts.items():
                for i in xrange(max(v, 0)):
                    item = self._db.pick_random(kind=k)
                    if item:
                        self._items.append(item)

        return self._items

    @property
    def steps(self):
        steps = []
        for item in self.items:
            steps += item.steps
        steps.append(pick_random([
            'Faites cuire 20 minutes au four, thermostat 7',
            'Laissez reposer pendant une a deux heures environ',
            'Decopuez le tout en morceaux assez copieux et deposez-les dans de '
            'petits ramequins',
            None,
        ], self._random))
        steps.append(pick_random([
            'rassemblez tous les ingredients dans un grand plat et consommez '
            'vite !',
            '... et bon appetit !',
            None,
        ], self._random))
        steps = [
            step.capitalize()
            for step in steps
            if step
        ]
        return steps

    @property
    def title(self):
        title = []
        left = None
        for item in self.items:
            title += item.str_in_title(left)
            left = item
        title = parts_to_string(title)
        title = title.capitalize()
        return title

    @property
    def ingredients(self):
        ingredients = []
        for item in self.items:
            ingredient = item.str_in_ingredients_list()
            if ingredient and len(ingredient):
                ingredient = parts_to_string(ingredient)
                ingredient = ingredient.capitalize()
                ingredients.append(ingredient)
        return ingredients

    @property
    def _people(self):
        people = 0
        for item in self.items:
            people += item.people
        return int(ceil(max(people, 1)))

    @property
    def people(self):
        people = self._people
        parts = ['Pour']
        if self._random.randrange(100) < 20:
            parts.append('environ')
        parts.append(people)
        if self._random.randrange(100) < 20:
            parts.append('a')
            parts.append(people + self._random.randrange(1, 4))
        parts.append('personne(s)')
        return ' '.join([str(part) for part in parts])
