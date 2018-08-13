$(document).ready(function () {
    fetch(window.location.protocol+'//'+window.location.host+'/info/containers')
        .then(function (response) {
            return response.json();
        })
        .then(function (json) {
            console.log(json);
            for(var i = 0; i < json['list'].length; i++) {

            }
        })
});