# -*- coding: utf-8 -*-

from .item import GenderizedItem


class SecondaryIngredient(GenderizedItem):
    kind = 'secondary_ingredient'
    is_spreadable = False
    is_powder = False
    is_uncountable = False
    is_by_piece = False
    is_spice = False
    is_citrus = False

    @property
    def people(self):
        return 0.4

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
        if left.kind in ('main_ingredient', 'secondary_ingredient'):
            if self._random.randrange(10) < 5:
                parts.append('et')
        parts.append(self._genderize(
            {'aux': {'quantity': 'multiple'}},
            {'a l\'': {'1st_voyel': True}},
            {'au': {'gender': 'male', '1st_voyel': False}},
            {'a la': {'gender': 'female'}},
        ))
        parts.append(self.name)
        return parts

    @property
    def steps(self):
        steps = []

        step = self._genderize(
            {'rechauffez {} a feu doux': {}},
            {'placez {} au bain-marie quelques minutes': {}},
            {'selon votre gout, vous pouvez voiler {} d\'un fond de sucre': {}},
            {'ajoutez {} par dessus': {}},
            {'faites cuire {} dans un wok': {}},
            {'faites chauffer {} et pensez a vanner pendant le '
             'refroidissement': {}},
            shuffle=True,
        )
        step = step.format(self.name_with_prefix)
        steps.append(step)

        return steps

    def pick_some(self):
        value = None
        unite = None

        if self.is_uncountable:
            unite = self._genderize(
                {'des': {'quantity': 'plural'}},
                {'de l\'': {'1st_voyel': True}},
                {'du': {'gender': 'male'}},
                {'de la': {'gender': 'female'}},
                value=value,
            )

        elif self.is_powder:
            value = self._random.randrange(1, 51) * 10
            unite = self._genderize(
                {'gramme de': {'value': 1, '1st_voyel': False}},
                {'gramme d\'': {'value': 1, '1st_voyel': True}},
                {'grammes de': {'1st_voyel': False}},
                {'grammes d\'': {'1st_voyel': True}},
                value=value,
            )

        elif self.is_by_piece:
            if self.quantity == 'single':
                value = 1
            else:
                value = self._random.randrange(1, 21)

        elif self.is_spice:
            options = []
            for beginning in ('une poignee', 'une dosette', 'un verre',
                              'une pincee'):
                for ending, constraints in {
                        'de': {'1st_voyel': False},
                        'd\'': {'1st_voyel': True},
                }.items():
                    key = '{} {}'.format(beginning, ending)
                    option = {}
                    option[key] = constraints
                    options.append(option)

            unite = self._genderize(*options, shuffle=True)

        elif self.is_spreadable:
            options = []
            for beginning in ('une noix', 'un morceau', 'une dose',
                              'une cuillere a cafe'):
                for ending, constraints in {
                        'de': {'1st_voyel': False},
                        'd\'': {'1st_voyel': True},
                }.items():
                    key = '{} {}'.format(beginning, ending)
                    option = {}
                    option[key] = constraints
                    options.append(option)

            unite = self._genderize(*options, shuffle=True)

        elif self.is_citrus:
            options = []
            for beginning in ('un zeste', 'un quartier', 'une pelure',
                              'de la pulpe'):
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

    @property
    def attrs(self):
        attrs = super(SecondaryIngredient, self).attrs
        attrs['is_spreadable'] = self.is_spreadable
        attrs['is_powder'] = self.is_powder
        attrs['is_uncountable'] = self.is_uncountable
        attrs['is_by_piece'] = self.is_by_piece
        attrs['is_spice'] = self.is_spice
        attrs['is_citrus'] = self.is_citrus
        return attrs


class Noisettes(SecondaryIngredient):
    name = 'noisettes'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class Amandes(SecondaryIngredient):
    name = 'amandes'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class Noix(SecondaryIngredient):
    name = 'noix'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class PommesDeTerre(SecondaryIngredient):
    name = 'pommes de terre'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class Dattes(SecondaryIngredient):
    name = 'dattes'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class GrainesDePavot(SecondaryIngredient):
    name = 'graines de pavot'
    gender = 'female'
    quantity = 'multiple'
    is_powder = True


class Epices(SecondaryIngredient):
    name = 'epices'
    gender = 'female'
    quantity = 'multiple'
    is_spice = True


class Tomates(SecondaryIngredient):
    name = 'tomates'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class GoussesDeVanille(SecondaryIngredient):
    name = 'gousses de vanille'
    gender = 'female'
    quantity = 'multiple'
    is_by_piece = True


class Canelle(SecondaryIngredient):
    name = 'canelle'
    gender = 'female'
    quantity = 'single'
    is_powder = True



class NoixDeCoco(SecondaryIngredient):
    name = 'noix de coco'
    gender = 'female'
    is_by_piece = True


class Mascarpone(SecondaryIngredient):
    name = 'mascarpone'
    gender = 'female'
    quantity = 'single'
    is_spreadable = True


class ConfitureDOrangeAmeres(SecondaryIngredient):
    name = 'confiture d\'orange ameres'
    gender = 'female'
    quantity = 'single'
    is_spreadable = True


class Orange(SecondaryIngredient):
    name = 'orange'
    gender = 'female'
    quantity = 'single'
    is_citrus = True


class Pamplemousse(SecondaryIngredient):
    name = 'pamplemousse'
    gender = 'female'
    quantity = 'single'
    is_citrus = True


class Farine(SecondaryIngredient):
    name = 'farine'
    gender = 'female'
    quantity = 'single'
    is_powder = True


class Moutarde(SecondaryIngredient):
    name = 'moutarde'
    gender = 'female'
    quantity = 'single'
    is_uncountable = True


class Gui(SecondaryIngredient):
    name = 'gui'
    gender = 'male'
    quantity = 'single'
    is_uncountable = True


class Houx(SecondaryIngredient):
    name = 'houx'
    gender = 'male'
    quantity = 'single'
    is_uncountable = True


class Ble(SecondaryIngredient):
    name = 'ble'
    gender = 'male'
    quantity = 'single'
    is_powder = True


class Lierre(SecondaryIngredient):
    name = 'lierre'
    gender = 'male'
    quantity = 'single'
    is_uncountable = True


class Anis(SecondaryIngredient):
    name = 'anis'
    gender = 'male'
    quantity = 'single'
    is_powder = True


class Citron(SecondaryIngredient):
    name = 'citron'
    gender = 'male'
    quantity = 'single'
    is_citrus = True


class Mais(SecondaryIngredient):
    name = 'mais'
    gender = 'male'
    quantity = 'single'
    is_powder = True


class Beurre(SecondaryIngredient):
    name = 'beurre'
    gender = 'male'
    quantity = 'single'
    is_spreadable = True


class Sel(SecondaryIngredient):
    name = 'sel'
    gender = 'male'
    quantity = 'single'
    is_spice = True


class Riz(SecondaryIngredient):
    name = 'riz'
    gender = 'male'
    quantity = 'single'
    is_powder = True


class Cacao(SecondaryIngredient):
    name = 'cacao'
    gender = 'male'
    quantity = 'single'
    is_powder = True


class FromageRape(SecondaryIngredient):
    name = 'fromage rape'
    gender = 'male'
    quantity = 'single'
    is_powder = True


class CubeDeKubor(SecondaryIngredient):
    name = 'cube de Kubor'
    gender = 'male'
    quantity = 'single'
    is_by_piece = True


class Reblochon(SecondaryIngredient):
    name = 'reblochon'
    gender = 'male'
    quantity = 'single'
    is_spreadable = True


class FloconsDAvoine(SecondaryIngredient):
    name = 'flocons d\'avoine'
    gender = 'male'
    quantity = 'multiple'
    is_powder = True


class Fruits(SecondaryIngredient):
    name = 'fruits'
    gender = 'male'
    quantity = 'multiple'
    is_powder = True


class FruitsSeches(SecondaryIngredient):
    name = 'fruits seches'
    gender = 'male'
    quantity = 'multiple'
    is_powder = True


class ClousDeGirofle(SecondaryIngredient):
    name = 'clous de girofle'
    gender = 'male'
    quantity = 'multiple'
    is_spice = True


class PetitsPois(SecondaryIngredient):
    name = 'petits pois'
    gender = 'male'
    quantity = 'multiple'
    is_powder = True


class Oeufs(SecondaryIngredient):
    name = 'oeufs'
    gender = 'male'
    is_by_piece = True
    name_singular = 'oeuf'


class BlancsDOeuf(SecondaryIngredient):
    name = 'blancs d\'oeuf'
    gender = 'male'
    quantity = 'multiple'
    is_by_piece = True


class JaunesDOeuf(SecondaryIngredient):
    name = 'jaunes d\'oeuf'
    gender = 'male'
    quantity = 'multiple'
    is_by_piece = True


class MorceauxDeSucre(SecondaryIngredient):
    name = 'morceaux de sucre'
    gender = 'male'
    quantity = 'multiple'
    is_by_piece = True


class ChampignonsDeParis(SecondaryIngredient):
    name = 'champignons de Paris'
    gender = 'male'
    quantity = 'multiple'
    is_by_piece = True


def all_items():
    return (
        Noisettes, Amandes, Noix, PommesDeTerre, Dattes, GrainesDePavot,
        Epices, Tomates, GoussesDeVanille, Canelle, NoixDeCoco, Mascarpone,
        ConfitureDOrangeAmeres, Orange, Pamplemousse, Farine, Moutarde, Gui,
        Houx, Ble, Lierre, Anis, Citron, Mais, Beurre, Sel, Riz, Cacao,
        FromageRape, CubeDeKubor, Reblochon, FloconsDAvoine, Fruits,
        FruitsSeches, ClousDeGirofle, PetitsPois, Oeufs, BlancsDOeuf,
        JaunesDOeuf, MorceauxDeSucre, ChampignonsDeParis,
    )
