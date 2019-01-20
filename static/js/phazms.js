$(document).ready(function () {
    $("#form-add-phazms").submit(function () {
        var username = get_username();
        if (!username) {
            return false;
        }
        $("#phazms-login").attr('value', username);
        return true;
    });
});
