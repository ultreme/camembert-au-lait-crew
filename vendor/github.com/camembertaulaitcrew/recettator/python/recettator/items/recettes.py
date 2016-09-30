# -*- coding: utf-8 -*-

from .item import GenderizedItem


class Recette(GenderizedItem):
    kind = 'recette'
    raw_steps = []

    def str_in_title(self, left=None):
        if left:
            raise NotImplementedError()
        return [self.name]

    @property
    def steps(self):
        steps = []
        for step in self.raw_steps:
            steps.append(step)
            # FIXME: format string
        return steps


class Tranches(Recette):
    name = 'tranches'
    gender = 'female'
    quantity = 'multiple'


class Galettes(Recette):
    name = 'galettes'
    gender = 'female'
    quantity = 'multiple'


class Lasagnes(Recette):
    name = 'lasagnes'
    gender = 'female'
    quantity = 'multiple'


class Chips(Recette):
    name = 'chips'
    gender = 'female'
    quantity = 'multiple'


class Cereales(Recette):
    name = 'cereales'
    gender = 'female'
    quantity = 'multiple'


class Escalopes(Recette):
    name = 'escalopes'
    gender = 'female'
    quantity = 'multiple'


class Endives(Recette):
    name = 'endives'
    gender = 'female'
    quantity = 'multiple'


class Pates(Recette):
    name = 'pates'
    gender = 'female'
    quantity = 'multiple'


class Patates(Recette):
    name = 'patates'
    gender = 'female'
    quantity = 'multiple'


class Truffe(Recette):
    name = 'truffe'
    gender = 'female'
    quantity = 'single'


class Mousse(Recette):
    name = 'mousse'
    gender = 'female'
    quantity = 'single'


class Buche(Recette):
    name = 'buche'
    gender = 'female'
    quantity = 'single'


class Puree(Recette):
    name = 'puree'
    gender = 'female'
    quantity = 'single'


class Ratatouille(Recette):
    name = 'ratatouille'
    gender = 'female'
    quantity = 'single'
    raw_steps = [
        'faites une ratatouille'
    ]


class Soupe(Recette):
    name = 'soupe'
    gender = 'female'
    quantity = 'single'


class PetitsGateaux(Recette):
    name = 'petits gateaux'
    gender = 'male'
    quantity = 'multiple'


class Rochers(Recette):
    name = 'rochers'
    gender = 'male'
    quantity = 'multiple'


class Champignons(Recette):
    name = 'champignons'
    gender = 'male'
    quantity = 'multiple'


class Parfait(Recette):
    name = 'parfait'
    gender = 'male'
    quantity = 'single'


class Civet(Recette):
    name = 'civet'
    gender = 'male'
    quantity = 'single'


class Gateau(Recette):
    name = 'gateau'
    gender = 'male'
    quantity = 'single'


class Gratin(Recette):
    name = 'gratin'
    gender = 'male'
    quantity = 'single'


class Kebab(Recette):
    name = 'kebab'
    gender = 'male'
    quantity = 'single'


class Rouleau(Recette):
    name = 'rouleau'
    gender = 'male'
    quantity = 'single'


def all_items():
    return (
        Tranches, Galettes, Lasagnes, Chips, Cereales, Escalopes, Endives,
        Pates, Patates, Truffe, Mousse, Buche, Puree, Ratatouille, Soupe,
        PetitsGateaux, Rochers, Champignons, Parfait, Civet, Gateau, Gratin,
        Kebab, Rouleau,
    )
