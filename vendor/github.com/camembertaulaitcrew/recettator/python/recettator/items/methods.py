# -*- coding: utf-8 -*-

from .item import Item


class Method(Item):
    kind = 'method'
    gender = 'any'
    quantity = 'any'

    @property
    def attrs(self):
        attrs = super(Method, self).attrs
        attrs['gender'] = self.gender
        attrs['quantity'] = self.quantity
        return attrs

    def str_in_title(self, left):
        return [self.name]


class ALaJuive(Method):
    name = 'a la juive'


class ALaMexicaine(Method):
    name = 'a la mexicaine'


class MethodeTraditionnelle(Method):
    name = 'methode traditionnelle'


class ALAncienne(Method):
    name = 'a l\'ancienne'


class CommeALaMaison(Method):
    name = 'comme a la maison'


class RecetteOriginale(Method):
    name = 'recette originale'


class Perso(Method):
    name = 'perso'


class DuChef(Method):
    name = 'du chef'


class ALaProvencale(Method):
    name = 'a la provencale'


class RecetteDeMaGrandMere(Method):
    name = 'recette de ma grand-mere'


class Premiums(Method):
    name = 'premium\'s'


class VersionXXL(Method):
    name = 'version XXL'


class DeChezMaxims(Method):
    name = 'de chez Maxim\'s'


class FaconSouabe(Method):
    name = 'facon Souabe'


class SpecialPizzaiolo(Method):
    name = 'special pizzaiolo'
    gender = 'male'
    quantity = 'single'


class SpecialGrandesOccasions(Method):
    name = 'special grandes occasions'
    gender = 'male'
    quantity = 'single'


def all_items():
    return (
        ALaJuive, ALaMexicaine, MethodeTraditionnelle, ALAncienne,
        CommeALaMaison, RecetteOriginale, Perso, DuChef, ALaProvencale,
        RecetteDeMaGrandMere, Premiums, VersionXXL, DeChezMaxims, FaconSouabe,
        SpecialPizzaiolo, SpecialGrandesOccasions,
    )
