# -*- coding: utf-8 -*-

import sys

from recettator import Recettator


def recettator_cli():
    seed = None
    if len(sys.argv) > 1:
        seed = sys.argv[1]

    recettator = Recettator(seed=seed)

    title = recettator.title
    print('#{} - {}'.format(recettator.seed, title))
    print((len(title) + 4 + len(str(recettator.seed))) * '=')
    print('')

    print(recettator.people)
    print('')

    # for k, v in recettator.infos.items():
    #     print('{}: {}'.format(k, v))
    # print('')

    print('Ingredients')
    print('-----------')
    for ingredient in recettator.ingredients:
        print('- {}'.format(ingredient))
    print('')

    print('How-to')
    print('-------')
    for step in recettator.steps:
        print('- {}'.format(step))

    # print('Debug')
    # print('-----')
    # for k, v in recettator.amount.items():
    #     print('{} amount: {}'.format(k, v))

if __name__ == '__main__':
    recettator_cli()
