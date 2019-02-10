var play_m1ch3l_jingle = function() {
    var track_id = Math.floor(Math.random() * 145);
    var url = "https://w.soundcloud.com/player/?url=api.soundcloud.com/playlists/37668512&show_artwork=false&show_comments=false&show_playcount=false&show_user=false&sharing=false&download=false&liking=false&buying=false&start_track="+track_id+"&auto_play=true";
    var iframe = $('<iframe src="'+url+'" id="sc-iframe"/>').hide();
    $('body').append(iframe);

    //$.getScript('https://w.soundcloud.com/player/api.js',
    //function(data, textStatus, jqxhr) {
    // var sc_widget = SC.widget('sc-iframe');
    //});
};

var getUrlVars = function() {
    var vars = {};
    var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(m,key,value) {
        vars[key] = value;
    });
    return vars;
}

var get_username = function(force) {
    if (!force && $.cookie('username')) {
        return $.cookie('username');
    }
    var username = prompt('C\'est quoi ton joli petit nom ?');
    $.cookie('username', username, {'path': '/'});
    $('.username_data').text(username);
    return username;
};

var calc_update_score = function(what, score) {
    score = parseInt(score);
    what = encodeURIComponent(what);
    username = encodeURIComponent(get_username());
    if (score < 0) {
        return;
    }
    var url = '/scorz/inc/' + username + '/' + what + '/' + score;
    $.ajax(url);
};

var easter_egg_callbacks = [];
easter_egg_callbacks.push(function() {
    $('body').addClass('barrel_roll');
    setTimeout(function() {
        $('body').removeClass('barrel_roll')
    }, 4000);
});

// http://www.had2know.com/technology/color-contrast-calculator-web-design.html
var brightness_difference = function(a, b) {
    return Math.abs(
        brightness(a.r, a.g, a.b) -
        brightness(b.r, b.g, b.b)
    );
};

var hue_difference = function(a, b) {
    return Math.abs(a.r - b.r) + Math.abs(a.g - b.g) + Math.abs(a.b - b.b);
};

var check_contrast = function(a, b) {
    return (brightness_difference(a, b) > 125 && hue_difference(a, b) > 500);
};

// http://stackoverflow.com/questions/9733288/how-to-programmatically-calculate-the-contrast-ratio-between-two-colors
var brightness = function(r, g, b) {
    return (299*r + 587*g + 114*b) / 1000;
};

// http://stackoverflow.com/questions/9733288/how-to-programmatically-calculate-the-contrast-ratio-between-two-colors
//    (luminanace(255, 255, 255) + 0.05) / (luminanace(255, 255, 0) + 0.05); // 1.074 for yellow
//    (luminanace(255, 255, 255) + 0.05) / (luminanace(0, 0, 255) + 0.05); // 8.592 for blue
//    minimal recommended contrast ratio is 4.5 or 3 for larger font-sizes
var luminanace = function(r, g, b) {
    var a = [r,g,b].map(function(v) {
        v /= 255;
        return (v <= 0.03928) ?
            v / 12.92 :
            Math.pow( ((v+0.055)/1.055), 2.4 );
        });
    return a[0] * 0.2126 + a[1] * 0.7152 + a[2] * 0.0722;
};

// http://www.had2know.com/technology/color-contrast-calculator-web-design.html


var colours = {
    "aliceblue":"#f0f8ff","antiquewhite":"#faebd7","aqua":"#00ffff","aquamarine":"#7fffd4","azure":"#f0ffff",
    "beige":"#f5f5dc","bisque":"#ffe4c4","black":"#000000","blanchedalmond":"#ffebcd","blue":"#0000ff","blueviolet":"#8a2be2","brown":"#a52a2a","burlywood":"#deb887",
    "cadetblue":"#5f9ea0","chartreuse":"#7fff00","chocolate":"#d2691e","coral":"#ff7f50","cornflowerblue":"#6495ed","cornsilk":"#fff8dc","crimson":"#dc143c","cyan":"#00ffff",
    "darkblue":"#00008b","darkcyan":"#008b8b","darkgoldenrod":"#b8860b","darkgray":"#a9a9a9","darkgreen":"#006400","darkkhaki":"#bdb76b","darkmagenta":"#8b008b","darkolivegreen":"#556b2f",
    "darkorange":"#ff8c00","darkorchid":"#9932cc","darkred":"#8b0000","darksalmon":"#e9967a","darkseagreen":"#8fbc8f","darkslateblue":"#483d8b","darkslategray":"#2f4f4f","darkturquoise":"#00ced1",
    "darkviolet":"#9400d3","deeppink":"#ff1493","deepskyblue":"#00bfff","dimgray":"#696969","dodgerblue":"#1e90ff",
    "firebrick":"#b22222","floralwhite":"#fffaf0","forestgreen":"#228b22","fuchsia":"#ff00ff",
    "gainsboro":"#dcdcdc","ghostwhite":"#f8f8ff","gold":"#ffd700","goldenrod":"#daa520","gray":"#808080","green":"#008000","greenyellow":"#adff2f",
    "honeydew":"#f0fff0","hotpink":"#ff69b4",
    "indianred ":"#cd5c5c","indigo":"#4b0082","ivory":"#fffff0","khaki":"#f0e68c",
    "lavender":"#e6e6fa","lavenderblush":"#fff0f5","lawngreen":"#7cfc00","lemonchiffon":"#fffacd","lightblue":"#add8e6","lightcoral":"#f08080","lightcyan":"#e0ffff","lightgoldenrodyellow":"#fafad2",
    "lightgrey":"#d3d3d3","lightgreen":"#90ee90","lightpink":"#ffb6c1","lightsalmon":"#ffa07a","lightseagreen":"#20b2aa","lightskyblue":"#87cefa","lightslategray":"#778899","lightsteelblue":"#b0c4de",
    "lightyellow":"#ffffe0","lime":"#00ff00","limegreen":"#32cd32","linen":"#faf0e6",
    "magenta":"#ff00ff","maroon":"#800000","mediumaquamarine":"#66cdaa","mediumblue":"#0000cd","mediumorchid":"#ba55d3","mediumpurple":"#9370d8","mediumseagreen":"#3cb371","mediumslateblue":"#7b68ee",
    "mediumspringgreen":"#00fa9a","mediumturquoise":"#48d1cc","mediumvioletred":"#c71585","midnightblue":"#191970","mintcream":"#f5fffa","mistyrose":"#ffe4e1","moccasin":"#ffe4b5",
    "navajowhite":"#ffdead","navy":"#000080",
    "oldlace":"#fdf5e6","olive":"#808000","olivedrab":"#6b8e23","orange":"#ffa500","orangered":"#ff4500","orchid":"#da70d6",
    "palegoldenrod":"#eee8aa","palegreen":"#98fb98","paleturquoise":"#afeeee","palevioletred":"#d87093","papayawhip":"#ffefd5","peachpuff":"#ffdab9","peru":"#cd853f","pink":"#ffc0cb","plum":"#dda0dd","powderblue":"#b0e0e6","purple":"#800080",
    "red":"#ff0000","rosybrown":"#bc8f8f","royalblue":"#4169e1",
    "saddlebrown":"#8b4513","salmon":"#fa8072","sandybrown":"#f4a460","seagreen":"#2e8b57","seashell":"#fff5ee","sienna":"#a0522d","silver":"#c0c0c0","skyblue":"#87ceeb","slateblue":"#6a5acd","slategray":"#708090","snow":"#fffafa","springgreen":"#00ff7f","steelblue":"#4682b4",
    "tan":"#d2b48c","teal":"#008080","thistle":"#d8bfd8","tomato":"#ff6347","turquoise":"#40e0d0",
    "violet":"#ee82ee",
    "wheat":"#f5deb3","white":"#ffffff","whitesmoke":"#f5f5f5",
    "yellow":"#ffff00","yellowgreen":"#9acd32"
};


// http://stackoverflow.com/questions/5623838/rgb-to-hex-and-hex-to-rgb
var hexToRgb = function(hex) {
    // Expand shorthand form (e.g. "03F") to full form (e.g. "0033FF")
    var shorthandRegex = /^#?([a-f\d])([a-f\d])([a-f\d])$/i;
    hex = hex.replace(shorthandRegex, function(m, r, g, b) {
        return r + r + g + g + b + b;
    });

    var result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
    return result ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16)
    } : null;
};

easter_egg_callbacks.push(function() {
    if (!window.cornify) {
        $.getScript("/components/cornify/cornify.js",
                    function(data, textStatus, jqxhr) {
                        window.cornify.pizzazz();
                        window.setTimeout(function() {
                            window.cornify.add();
                        }, 10000);
                    });
    }
});

easter_egg_callbacks.push(function() {
    if (!$.fn.raptorize) {
        $.getScript("/js/jquery.raptorize.1.0.js",
                    function(data, textStatus, jqxhr) {
                        var div = $('<div id="raptorize-button" />');
                        $('body').append(div);
                        div.raptorize();
                        div.click();
                    });
    } else {
        $('#raptorize-button').click();
    }
});

if (!$('body').hasClass('devel')) {
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
        (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                             m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
                            })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
    ga('create', 'UA-51632370-2', 'auto');
    ga('send', 'pageview');
}

// http://stackoverflow.com/questions/15191058/css-rotation-cross-browser-with-jquery-animate
$.fn.animateRotate = function(base, angle, duration, easing, complete) {
    var args = $.speed(duration, easing, complete);
    var step = args.step;
    return this.each(function(i, e) {
        args.step = function(now) {
            $.style(e, 'transform', 'rotate(' + now + 'deg)');
            if (step) {
                return step.apply(this, arguments);
            }
            return false;
        };
        $({deg: base}).animate({deg: angle}, args);
    });
};

var thumbnail_colors = ['red', 'yellow', 'green', 'blue', 'orange', 'purple', 'cyan', 'magenta', 'pink'];

var each_time = function() {
    var username = $.cookie('username');
    if (username) {
        $('.username_data').text(username);
    }

    $('#username_button').click(function() {
        get_username(true);
        return false;
    });

    $(".rotate-random").each(function () {
        var angle_min = 1;
        var angle_max = 5;
        var animate_multiplier = 2;
        var animate_speed_max = 4000;
        var animate_speed_min = 1500;

        var angle_diff = angle_max - angle_min;
        var strength = Math.random() * angle_diff * 2 - angle_diff;
        strength = strength > 0 ? strength + angle_min : strength - angle_min;
        var that = $(this);
        that.css('transform', 'rotate(' + strength + 'deg)');

        var speed = Math.floor(Math.random() * (animate_speed_max - animate_speed_min) + animate_speed_min);
        var angle = Math.floor(Math.random() * angle_diff * 2 - angle_diff) * animate_multiplier;
        angle = angle > 0 ? angle + angle_min * animate_multiplier : angle - angle_min * animate_multiplier;

        var animate = function(base, that, angle, speed) {
            that.animateRotate(base, angle, speed, 'swing', function() {
                that.animateRotate(angle, -angle, speed, 'swing', function() {
                    animate(-angle, that, angle, speed);
                });
            });
        };
        animate(0, that, angle, speed);
    });

    $('.iframe-auto-height').each(function() {
        var that = $(this);
        var resize = function() {
            that.height($(this).height());
        };
        that.ready(resize);
        setInterval(resize, 2000);
    });

    $(".thumbnail").each(function () {
        $(this).hover(
            function () {
                $(this).find(".caption").show();
            },
            function () {
                $(this).find(".caption").hide();
            }
        ).css('background-color', random_color(thumbnail_colors));
    });

    $(".blink").each(function () {
        var that = $(this);
        setInterval(function() {
            if (that.css("visibility") == "hidden") {
                that.css("visibility", "visible");
            } else {
                that.css("visibility", "hidden");
            }
        }, 100);
    });

    $(".bicool").each(function () {
        var current = $(this).attr("src");
        var next = $(this).data("alternate");

        $(this).hover(
            function () {
                $(this).attr("src", next);
            },
            function () {
                $(this).attr("src", current);
            }
        );
    });

    $('.refresh-page').click(function() {
        location.reload();
        return false;
    });


    var disco_colors = ['red', 'orange', 'yellow', 'blue', 'brown', 'white', 'gray', 'royalblue', 'pink', 'magenta', 'purple'];
    $('.disco').each(function() {
        var that = $(this);
        that.css('color', random_color(disco_colors));
        setInterval(function() {
            that.css('color', random_color(disco_colors));
        }, Math.floor(Math.random() * 800 + 800));
    });

    //$("img").unveil();
};

var random_color = function(colors) {
    return colors[Math.floor(Math.random() * colors.length)];
};

$(document).ready(function () {
    $('.social-links a').tooltip();
    var easter_egg = new Konami(function() {
        var func = easter_egg_callbacks[
            Math.floor(Math.random() * easter_egg_callbacks.length)
        ];
        func();
    });

    each_time();
});

// http://www.onextrapixel.com/2011/10/17/animating-colors-using-css3-transitions-with-jquery-fallback/
$.each(["backgroundColor","borderBottomColor","borderLeftColor","borderRightColor","borderTopColor","borderColor","color","outlineColor"],function(b,a){$.fx.step[a]=function(c){if(!c.colorInit){c.start=getColor(c.elem,a);c.end=getRGB(c.end);c.colorInit=true}c.elem.style[a]="rgb("+Math.max(Math.min(parseInt((c.pos*(c.end[0]-c.start[0]))+c.start[0],10),255),0)+","+Math.max(Math.min(parseInt((c.pos*(c.end[1]-c.start[1]))+c.start[1],10),255),0)+","+Math.max(Math.min(parseInt((c.pos*(c.end[2]-c.start[2]))+c.start[2],10),255),0)+")"}});function getRGB(b){var a;if(b&&b.constructor==Array&&b.length==3){return b}if(a=/rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)/.exec(b)){return[parseInt(a[1],10),parseInt(a[2],10),parseInt(a[3],10)]}if(a=/rgb\(\s*([0-9]+(?:\.[0-9]+)?)\%\s*,\s*([0-9]+(?:\.[0-9]+)?)\%\s*,\s*([0-9]+(?:\.[0-9]+)?)\%\s*\)/.exec(b)){return[parseFloat(a[1])*2.55,parseFloat(a[2])*2.55,parseFloat(a[3])*2.55]}if(a=/#([a-fA-F0-9]{2})([a-fA-F0-9]{2})([a-fA-F0-9]{2})/.exec(b)){return[parseInt(a[1],16),parseInt(a[2],16),parseInt(a[3],16)]}if(a=/#([a-fA-F0-9])([a-fA-F0-9])([a-fA-F0-9])/.exec(b)){return[parseInt(a[1]+a[1],16),parseInt(a[2]+a[2],16),parseInt(a[3]+a[3],16)]}if(a=/rgba\(0, 0, 0, 0\)/.exec(b)){return colors.transparent}return colors[$.trim(b).toLowerCase()]}function getColor(c,a){var b;do{b=$.css(c,a);if(b!=""&&b!="transparent"||$.nodeName(c,"body")){break}a="backgroundColor"}while(c=c.parentNode);return getRGB(b)};

var switch_cool_style = function() {
    var timer = 500;
    var _html = $('html');
    var _body = $('body', _html);
    var last_attrs = {};
    var attrs = ['background-color', 'color'];
    for (var i = 0; i < attrs.length; i++) {
        var attr = attrs[i];
        last_attrs[attr] = _body.css(attr);
    }

    var last_background_color = _body.css('background-color');
    var last_style = null;
    for (var i = 0; i < styles.length; i++) {
        style = 'cool-style-' + styles[i];
        if (_html.hasClass(style)) {
            _html.removeClass(style);
            last_style = style;
        }
    }
    var new_style = last_style;
    while (new_style == last_style) {
        new_style = 'cool-style-' + styles[Math.floor(Math.random() * styles.length)];
    }

    _html.addClass(new_style);
    var new_attrs = {};
    for (var i = 0; i < attrs.length; i++) {
        var attr = attrs[i];
        new_attrs[attr] = _body.css(attr);
    }
    _html.removeClass(new_style);
    _body.css(last_attrs);
    _body.animate(new_attrs, timer);
    setTimeout(function() {
        _html.addClass(new_style);
        //_body.css({backgroundColor: null});
    }, timer + 100);
};



InstantClick.init(50);
InstantClick.on('change', function(isInitialLoad) {
    ga('send', 'pageview', location.pathname + location.search);
    styles = ['cachou', 'jambon', 'epinard', 'lasagne', 'haricot', 'sandwich'];

    switch_cool_style();
    each_time();
});
