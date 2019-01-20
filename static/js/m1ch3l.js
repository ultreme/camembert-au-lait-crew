$(document).ready(function () {
    var name = get_username();
    var lines = [{'date': new Date(), 'who': 'm1ch3l', 'what': "Salut, tu peux me parler si tu veux."}];
    var height = 42 / 2;
    var history = [];
    var history_offset = 0;

    var refresh_m1ch3l = function () {
        var result = '';
        for (var i in lines) {
            line = lines[i];
            date = '[' + line['date'].getHours() + ':' + line['date'].getMinutes() + ']';
            who = line['who'];
            result += date + ' ' + who + ': ' + line['what'] + '\n';
        }
        $("#m1ch3l-area").html(result);
    }

    var m1ch3l_send = function () {
        var what = $('#what-m1ch3l').val();
	history.unshift(what);
        $.ajax({
            url: '',
            method: 'POST',
            data: {'query': JSON.stringify({'Action': 'SAY', 'Login': name, 'Data': what})},
            dataType: 'json',
            success: function (data) {
                lines.push({'date': new Date(), 'who': name, 'what': what});
                while (lines.length > height) {
                    lines.shift();
                }
                refresh_m1ch3l();
            }
        });
        $('#what-m1ch3l').val("");
    }

    $("#what-m1ch3l").on("keydown", function (event) {
	// up
	if (event.keyCode == 38) {
	    ++history_offset;
	    if (history_offset >= history.length) {
		--history_offset;
	    }
	    if (history.length > 0) {
		$("#what-m1ch3l").val(history[history_offset]);
	    }
	}

	// down
	if (event.keyCode == 40) {
	    --history_offset;
	    if (history_offset < 0) {
		history_offset = 0;
	    }
	    if (history.length > 0) {
		$("#what-m1ch3l").val(history[history_offset]);
	    }
	}
    });
    
    $("#rien").submit(function () {
        m1ch3l_send();
        return false;
    });

    $("#talk-to-m1ch3l").click(function () {
        m1ch3l_send();
    });

    setInterval(function () {
        $.ajax({
            url: '',
            method: 'POST',
            data: {'query': JSON.stringify({'Action': 'POLL', 'Login': name, 'Data': ''})},
            dataType: 'json',
            success: function (data) {
                if (data.ReturnCode == "RC_OK") {
                    for (var i in data.Messages) {
                        lines.push({'date': new Date(), 'who': 'm1ch3l', 'what': data.Messages[i]});
                        while (lines.length > height) {
                            lines.shift();
                        }
                    }
                    refresh_m1ch3l();
                }
            }
        });
    }, 2000);

    refresh_m1ch3l();

    play_m1ch3l_jingle();

});
