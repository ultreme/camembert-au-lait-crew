var invert = function (obj) {
    var new_obj = {};
    for (var prop in obj) {
        if(obj.hasOwnProperty(prop)) {
            new_obj[obj[prop]] = prop;
        }
    }
    return new_obj;
};
var yougo_french_attrs_old = {
    'a': '¡',
    'b': '¢',
    'c': '£',
    'd': '¤',
    'e': '¥',
    'f': '¦',
    'g': '§',
    'h': '¨',
    'i': 'ª',
    'j': '«',
    'k': '¬',
    'l': '®',
    'm': '¯',
    'n': '°',
    'o': '©',
    'p': '±',
    'q': '²',
    'r': '³',
    's': '´',
    't': 'µ',
    'u': '¶',
    'v': '·',
    'w': '¸',
    'x': '¹',
    'y': 'º',
    'z': '»',
    'A': '¼',
    'B': '½',
    'C': '¾',
    'D': '¿',
    'E': 'À',
    'F': 'Á',
    'G': 'Â',
    'H': 'Ã',
    'I': 'Ä',
    'J': 'Å',
    'K': 'Æ',
    'L': 'Ç',
    'M': 'È',
    'N': 'É',
    'O': 'Ê',
    'P': 'Ë',
    'Q': 'Ì',
    'R': 'Í',
    'S': 'Î',
    'T': 'Ï',
    'U': 'Ð',
    'V': 'Ñ',
    'W': 'Ò',
    'X': 'Ó',
    'Y': 'Ô',
    'Z': 'Õ',
    '0': 'Ö',
    '1': '×',
    '2': 'Ø',
    '3': 'Ù',
    '4': 'Ú',
    '5': 'Û',
    '6': 'Ü',
    '7': 'Ý',
    '8': 'Þ',
    '9': 'ß'
};
var yougo_french_attrs = {
    'a': '™',
    'b': '€',
    'c': '‰',
    'd': '…',
    'e': '•',
    'f': '‡',
    'g': '†',
    'h': '„',
    'i': '”',
    'j': '“',
    'k': '‚',
    'l': '’',
    'm': '‘',
    'n': '—',
    'o': '–',
    'p': 'Œ',
    'q': 'œ',
    'r': 'Š',
    's': 'š',
    't': 'Ÿ',
    'u': 'ƒ',
    'v': 'ÿ',
    'w': 'þ',
    'x': 'ý',
    'y': 'ü',
    'z': 'ø',
    'A': '÷',
    'B': 'ö',
    'C': 'ñ',
    'D': 'ï',
    'E': 'ë',
    'F': 'ç',
    'G': 'æ',
    'H': 'å',
    'I': 'ß',
    'J': 'Þ',
    'K': 'Ý',
    'L': 'Ü',
    'M': 'Ø',
    'N': '×',
    'O': 'Ñ',
    'P': 'Ð',
    'Q': 'Ï',
    'R': 'Ë',
    'S': 'Ç',
    'T': 'Æ',
    'U': 'Å',
    'V': '¿',
    'W': '¾',
    'X': '½',
    'Y': '¼',
    'Z': '»',
    '0': 'º',
    '1': '¹',
    '2': '¸',
    '3': '·',
    '4': 'µ',
    '5': '¶',
    '6': '´',
    '7': '³',
    '8': '²',
    '9': '±'
};
var french_yougo_attrs = invert(yougo_french_attrs);

var french_scrabble_attrs = {
    'a': 'JAMBON',
    'b': 'VOITURE',
    'c': 'CENDRILLER',
    'd': 'CACHOU',
    'e': 'OREILLER',
    'f': 'MALE',
    'g': 'FEMME',
    'h': 'ISTANBUL',
    'i': 'POITRINE',
    'j': 'DENT',
    'k': 'PEIGNOIR',
    'l': 'FRIGO',
    'm': 'TELEVISION',
    'n': 'CASQUE',
    'o': 'CHAMBRE',
    'p': 'GIRAFE',
    'q': 'UKULELE',
    'r': 'MIROIR',
    's': 'ASCENCEUR',
    't': 'ROBINET',
    'u': 'LAMPADAIRE',
    'v': 'WEBCAM',
    'w': 'ECRAN',
    'x': 'SOURIS',
    'y': 'CLAVIER',
    'z': 'MONOPOLY'
};
var scrabble_french_attrs = invert(french_scrabble_attrs);

function translate_french_yougo(text) {
    var newtext = '';
    for (var i = 0; i < text.length; i++) {
        var ch = text[i];
        if (ch in yougo_french_attrs) {
            newtext += yougo_french_attrs[ch];
        } else {
            newtext += ch;
        }
    }
    return newtext;
}
function translate_yougo_french(text) {
    var newtext = '';
    for (var i = 0; i < text.length; i++) {
        var ch = text[i];
        if (ch in french_yougo_attrs) {
            newtext += french_yougo_attrs[ch];
        } else {
            newtext += ch;
        }
    }
    return newtext;
}

function translate_french_scrabble(text) {
    var parts = text.split("");
    var newtext = '';
    for (var i in parts) {
        var ch = parts[i];
        if (ch in french_scrabble_attrs) {
            newtext += '[' + french_scrabble_attrs[ch] + ']';
        } else {
            newtext += ch;
        }
    }
    return newtext;
}

function translate_scrabble_french(text) {
    var replaceCallback = function(match) {
        match = match.substring(1, match.length - 1);
        if (match in scrabble_french_attrs) {
            return scrabble_french_attrs[match];
        } else {
            return match;
        }
    };
    return text.replace(/\[([^\]]*)\]/g, replaceCallback);
}

function translate_french_verlant(text) {
    return text.split("").reverse().join("");
}

function translate_verlant_french(text) {
    return text.split("").reverse().join("");
}

function cryptage() {
    var from = 'french';
    var to = 'yougo';
    var func, text;
    var cestparti = function() {
        from = $('#langfrom input:checked').val();
        to = $('#langto input:checked').val();
        if (from == to) {
            if (from == 'french') {
                $('#langto2').attr('checked', 'checked');
                to = 'yougo';
            } else {
                $('#langto1').attr('checked', 'checked');
                to = 'french';
            }
        }
        $('#langto input[disabled]').prop('disabled', false);
        $('#langto input[value='+from+']').attr('disabled', 'disabled');


        var text = $('textarea#source').val();
        console.log(text, from, to);
        if (from != 'french') {
            func = 'translate_' + from + '_french';
            text = window[func](text);
        }

        if (to != 'french') {
            func = 'translate_french_' + to;
            text = window[func](text);
        }
        $('textarea#dest').val(text);
    };
    //$('#cryptage').click(cestparti);
    $('button').click(cestparti);
    $('input').click(cestparti).change(cestparti);
    $('textarea').focus(cestparti).blur(cestparti);
    $('#cryptage').text('CRYPTAGE').prop('disabled', false);
    var latest = null;
    function cron() {
        var newone = $('#source').val() + $('#langto1').val();
        if (latest != newone) {
            cestparti();
            latest = newone;
        }
    }
    setInterval(cron, 500);
}
$(document).ready(function() {cryptage();});
