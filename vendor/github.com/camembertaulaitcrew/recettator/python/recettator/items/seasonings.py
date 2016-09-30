# -*- coding: utf-8 -*-

from .item import GenderizedItem


class Seasoning(GenderizedItem):
    kind = 'seasoning'

    @property
    def people(self):
        return 0.1

    def str_in_ingredients_list(self):
        parts = []
        if self._picked['value']:
            parts.append(self._picked['value'])
        if self._picked['unite']:
            parts.append(self._picked['unite'])
        parts.append(self.name)
        return parts

    def pick_some(self):
        value = None
        unite = None

        value = self._random.randrange(1, 31) / 10.0

        unite_base = 'litre'
        if value < 1:
            unite_base = 'centilitre'
            value *= 100

        unite = self._genderize(
            {'{} de'.format(unite_base): {'value': 1, '1st_voyel': False}},
            {'{} d\''.format(unite_base): {'value': 1, '1st_voyel': True}},
            {'{}s de'.format(unite_base): {'1st_voyel': False}},
            {'{}s d\''.format(unite_base): {'1st_voyel': True}},
            value=value,
        )

        if value and value == int(value):
            value = int(value)

        self._picked['value'] = value
        self._picked['unite'] = unite

    @property
    def steps(self):
        steps = []

        step = self._genderize(
            {'remuez {} dans un pot en terre cuite': {}},
            {'versez doucement {} et melangez suffisement': {}},
            shuffle=True,
        )
        step = step.format(self.name_with_prefix)
        steps.append(step)

        return steps


class Tisane(Seasoning):
    name = 'tisane'
    gender = 'female'
    quantity = 'single'


class ExtraitDeFleurDOranger(Seasoning):
    name = 'extrait de fleur d\'oranger'
    gender = 'male'
    quantity = 'single'


class Viandox(Seasoning):
    name = 'viandox'
    gender = 'male'
    quantity = 'single'


class BiereDeNoel(Seasoning):
    name = 'biere de noel'
    gender = 'female'
    quantity = 'single'


class VinRouge(Seasoning):
    name = 'vin rouge'
    gender = 'male'
    quantity = 'single'


class VinBlanc(Seasoning):
    name = 'vin blanc'
    gender = 'male'
    quantity = 'single'


class HuileDArachide(Seasoning):
    name = 'huile d\'arachide'
    gender = 'female'
    quantity = 'single'


class SauceDHuitre(Seasoning):
    name = 'sauce d\'huitre'
    gender = 'female'
    quantity = 'single'


class CremeFraiche(Seasoning):
    name = 'creme fraiche'
    gender = 'female'
    quantity = 'single'


class Creme(Seasoning):
    name = 'creme'
    gender = 'female'
    quantity = 'single'


class LiqueurDeRaisin(Seasoning):
    name = 'liqueur de raisin'
    gender = 'female'
    quantity = 'single'


class GrandMarnier(Seasoning):
    name = 'grand marnier'
    gender = 'male'
    quantity = 'single'


class Lait(Seasoning):
    name = 'lait'
    gender = 'male'
    quantity = 'single'


class LaitFermente(Seasoning):
    name = 'lait fermente'
    gender = 'male'
    quantity = 'single'


class HuileDOlive(Seasoning):
    name = 'huile d\'olive'
    gender = 'female'
    quantity = 'single'


class VinaigreDeRiz(Seasoning):
    name = 'vinaigre de riz'
    gender = 'male'
    quantity = 'single'


class VinaigreDeCidre(Seasoning):
    name = 'vinaigre de cidre'
    gender = 'male'
    quantity = 'single'


class VinaigreDeVin(Seasoning):
    name = 'vinaigre de vin'
    gender = 'male'
    quantity = 'single'


class JusDeCitron(Seasoning):
    name = 'jus de citron'
    gender = 'male'
    quantity = 'single'


def all_items():
    return (
        Tisane, ExtraitDeFleurDOranger, Viandox, BiereDeNoel, VinRouge,
        VinBlanc, HuileDArachide, SauceDHuitre, CremeFraiche, Creme,
        LiqueurDeRaisin, GrandMarnier, Lait, LaitFermente, HuileDOlive,
        VinaigreDeRiz, VinaigreDeCidre, VinaigreDeVin, JusDeCitron,
    )
