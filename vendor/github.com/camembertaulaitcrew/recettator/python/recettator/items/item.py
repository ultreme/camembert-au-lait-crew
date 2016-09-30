# -*- coding: utf-8 -*-

from random import Random
from ..utils import parts_to_string, genderization


class Item(object):
    name = None
    kind = None

    def __init__(self, db, parent=None, seed=None):
        self._picked = {}
        self._db = db
        self._parent = parent
        self._random = Random()
        if seed:
            self._random.seed(seed)
        self.pick_some()

    def pick_some(self):
        pass

    def str_in_ingredients_list(self):
        return []

    def str_in_title(self, left):
        return []

    @property
    def attrs(self):
        attrs = {
            'kind': self.kind,
            'name': self.kind,
            '1st_voyel': self._first_voyel,
        }
        return attrs

    def __repr__(self):
        return "<{}:{}>".format(self.kind, type(self).__name__)
        """
        More detail:

        return "<{} {}>".format(
            type(self).__name__,
            ', '.join(['{}={}'.format(k, v) for k, v in self.attrs.items()])
        )
        """

    @property
    def _first_voyel(self):
        return self.name[0] in ('a', 'e', 'i', 'o', 'u', 'y')

    @property
    def people(self):
        return 0

    @property
    def steps(self):
        return []

    @property
    def name_prefix(self):
        return [self._genderize(
            {'les': {'quantity': 'multiple'}},
            {'l\'': {'1st_voyel': True}},
            {'le': {'gender': 'male'}},
            {'la': {'gender': 'female'}},
        )]

    @property
    def name_with_prefix(self):
        parts = []
        parts += self.name_prefix
        parts.append(self.name)
        return parts_to_string(parts)


class AddonItem(Item):
    pass


class GenderizedItem(Item):
    gender = 'any'
    quantity = 'any'

    @property
    def attrs(self):
        attrs = super(GenderizedItem, self).attrs
        attrs['gender'] = self.gender
        attrs['quantity'] = self.quantity
        return attrs

    def _genderize(self, *args, **kwargs):
        shuffle = 'shuffle' in kwargs and kwargs['shuffle']

        options = list(args)
        if shuffle:
            self._random.shuffle(options)

        constraints = self.attrs
        for k, v in kwargs.items():

            if k in ('shuffle',):
                continue

            constraints[k] = v

        return genderization(options, constraints)
