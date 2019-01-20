$(document).ready(function() {
    var url = 'http://paint-mmo.sbrk.org/?username='+encodeURIComponent(get_username());
    var iframe = $('<iframe />');
    iframe.attr({
        src: url,
        frameborder: 0,
        width: 840,
        height: 450,
    });
    $('#paint_frame').append(iframe);
});
