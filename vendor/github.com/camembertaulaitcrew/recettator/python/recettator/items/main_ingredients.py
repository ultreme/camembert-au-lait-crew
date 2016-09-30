# -*- coding: utf-8 -*-

from ..utils import parts_to_string

from .item import GenderizedItem


class MainIngredient(GenderizedItem):
    kind = 'main_ingredient'
    method = None

    @property
    def people(self):
        return 0.7

    @property
    def steps(self):
        steps = []

        if self.method:
            steps += self.method.steps

        step = self._genderize(
            {'decoupez {} en fines petits tranches': {}},
            {'deposez {} juste au dessus': {}},
            shuffle=True,
        )
        step = step.format(self.name_with_prefix)
        steps.append(step)

        return steps

    def str_in_ingredients_list(self):
        parts = []
        if self._picked['value']:
            parts.append(self._picked['value'])
        if self._picked['unite']:
            parts.append(self._picked['unite'])
        parts.append(self.name)
        return parts

    def str_in_title(self, left):
        parts = []

        if self._random.randrange(2):
            parts.append(self._genderize(
                {'aux': {'quantity': 'multiple'}},
                {'a l\'': {'1st_voyel': True}},
                {'au': {'gender': 'male', '1st_voyel': False}},
                {'a la': {'gender': 'female'}},
            ))
        else:
            parts.append(left._genderize(
                {'assorti': {'gender': 'male', 'quantity': 'single'}},
                {'assortie': {'gender': 'female', 'quantity': 'single'}},
                {'assortis': {'gender': 'male', 'quantity': 'multiple'}},
                {'assorties': {'gender': 'female', 'quantity': 'multiple'}},
            ))
            parts.append(self._genderize(
                {'de': {'1st_voyel': False}},
                {'d\'': {'1st_voyel': True}},
            ))
        parts.append(self.name)

        if self.method:
            parts += self.method.str_in_title(self)

        return parts

    def pick_some(self):
        value = None
        unite = None

        if self._random.randrange(20):
            self.method = self._db.pick_random(kind='ingredient_method',
                                               parent=self)

        rand = self._random.randrange(3)
        if rand == 0:
            value = self._random.randrange(1, 51) * 10
            unite = self._genderize(
                {'gramme de': {'value': 1, '1st_voyel': False}},
                {'gramme d\'': {'value': 1, '1st_voyel': True}},
                {'grammes de': {'1st_voyel': False}},
                {'grammes d\'': {'1st_voyel': True}},
                value=value,
            )

        elif rand == 1:
            value = self._random.randrange(1, 6) + 1
            unite = self._genderize(
                {'tranche de': {'value': 1, '1st_voyel': False}},
                {'tranche d\'': {'value': 1, '1st_voyel': True}},
                {'tranches de': {'1st_voyel': False}},
                {'tranches d\'': {'1st_voyel': True}},
                value=value,
            )

        elif rand == 2:
            value = None
            options = [
                {'un bon gros': {'gender': 'male', 'quantity': 'single'}},
                {'une assez grosse': {'gender': 'female',
                                      'quantity': 'single'}},
                {'plusieurs gros': {'gender': 'male', 'quantity': 'multiple'}},
            ]

            for beginning in ('une quantite suffisante', 'pas mal',
                              'quelques morceaux', 'un bon paquet', 'beaucoup'):
                for ending, constraints in {
                        'de': {'1st_voyel': False},
                        'd\'': {'1st_voyel': True},
                }.items():
                    key = '{} {}'.format(beginning, ending)
                    option = {}
                    option[key] = constraints
                    options.append(option)

            unite = self._genderize(*options, shuffle=True)

        if value and value == int(value):
            value = int(value)

        self._picked['value'] = value
        self._picked['unite'] = unite


class FoieGras(MainIngredient):
    name = 'foie gras'
    gender = 'male'
    quantity = 'single'


class FoieDOie(MainIngredient):
    name = 'foie d\'oie'
    gender = 'male'
    quantity = 'single'


class Lievre(MainIngredient):
    name = 'lievre'
    gender = 'male'
    quantity = 'single'


class Jambon(MainIngredient):
    name = 'jambon'
    gender = 'male'
    quantity = 'single'


class Poulet(MainIngredient):
    name = 'poulet'
    gender = 'male'
    quantity = 'single'


class Surimi(MainIngredient):
    name = 'surimi'
    gender = 'male'
    quantity = 'single'


class Requin(MainIngredient):
    name = 'requin'
    gender = 'male'
    quantity = 'single'


class Cheval(MainIngredient):
    name = 'cheval'
    gender = 'male'
    quantity = 'single'


class Veau(MainIngredient):
    name = 'veau'
    gender = 'male'
    quantity = 'single'


class Lotte(MainIngredient):
    name = 'lotte'
    gender = 'female'
    quantity = 'single'


class Oie(MainIngredient):
    name = 'oie'
    gender = 'female'
    quantity = 'single'


class Carpe(MainIngredient):
    name = 'carpe'
    gender = 'female'
    quantity = 'single'


class Dinde(MainIngredient):
    name = 'dinde'
    gender = 'female'
    quantity = 'single'


class Autruche(MainIngredient):
    name = 'autruche'
    gender = 'female'
    quantity = 'single'


class Lardons(MainIngredient):
    name = 'lardons'
    gender = 'male'
    quantity = 'multiple'


def all_items():
    return (
        FoieGras, FoieDOie, Lievre, Jambon, Poulet, Surimi, Requin, Cheval,
        Veau, Lotte, Oie, Carpe, Dinde, Autruche, Lardons,
    )
