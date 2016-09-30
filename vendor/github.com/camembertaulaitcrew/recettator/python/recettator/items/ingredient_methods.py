# -*- coding: utf-8 -*-

from ..utils import parts_to_string

from .item import AddonItem


class IngredientMethod(AddonItem):
    kind = 'ingredient_method'
    raw_steps = []
    names = None

    @property
    def name(self):
        return self.names[0]

    def str_in_title(self, left):
        if left != self._parent:
            raise RuntimeError()

        options = {}
        options[self.names[0]] = {'quantity': 'single', 'gender': 'male'}
        options[self.names[1]] = {'quantity': 'single', 'gender': 'female'}
        options[self.names[2]] = {'quantity': 'multiple', 'gender': 'male'}
        options[self.names[3]] = {'quantity': 'multiple', 'gender': 'female'}
        options = [{k: v} for k, v in options.items()]
        return [self._parent._genderize(*options)]

    @property
    def steps(self):
        steps = []
        for step in self.raw_steps:
            step = step.format(self._parent.name_with_prefix)
            steps.append(step)
        return steps


class Glace(IngredientMethod):
    names = ['glace', 'glacee', 'glaces', 'glacees']
    raw_steps = [
        'mettez {} au refrigirateur quelques heures'
    ]


class Poele(IngredientMethod):
    names = ['poele', 'poelee', 'poeles', 'poelees']
    raw_steps = [
        'faites revenir {} dans une poele'
    ]


class Farci(IngredientMethod):
    names = ['farci', 'farcie', 'farcis', 'farcies']
    raw_steps = [
        'remplissez {} avec ce que vous voulez'
    ]
    # FIXME: try to replace "ce que vous voulez" with another item


class Roti(IngredientMethod):
    names = ['roti', 'rotie', 'rotis', 'roties']
    raw_steps = [
        'prechauffez le four pour y mettre {} par la suite'
    ]


class Chaud(IngredientMethod):
    names = ['chaud', 'chaude', 'chauds', 'chaudes']
    raw_steps = [
        'chauffez legerement {} au four'
    ]


class Decoupe(IngredientMethod):
    names = ['decoupe', 'decoupee', 'decoupes', 'decoupees']
    raw_steps = [
        'decoupez {} en tranches plutot epaisses'
    ]

class Grille(IngredientMethod):
    names = ['grille', 'grillee', 'grilles', 'grillees']
    raw_steps = [
        'mettez {} sur le grill'
    ]


class Battu(IngredientMethod):
    names = ['battu', 'battue', 'battus', 'battues']
    raw_steps = [
        'battez energiquement {} avec un fouet'
    ]


def all_items():
    return (
        Glace, Poele, Farci, Roti, Chaud, Decoupe, Grille, Battu,
    )
