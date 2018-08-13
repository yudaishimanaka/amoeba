$(document).ready(function () {
    fetch(window.location.protocol+'//'+window.location.host+'/api/containers')
        .then(function (response) {
            if(response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok.');
        })
        .then(function (json) {
            console.log(json);
            for(var i = 0; i < json['list'].length; i++) {
                $("#container-card").append(
                    "<div class='col s12 m4 l3'><div class='card blue-grey darken-2 white-text'><div class='card-content'><div class='row'><div class='col s3 m3 l3'></div><div class='col s6 m6 l6'><img class='responsive-img' src='assets/images/ubuntu-logo.png'></div><div class='col s3 m3 l3'></div></div><span class='card-title center'>test-container</span><div class='row'><div class='col s1 m1 l1'></div><div class='col s10 m10 l10 center'><p>163.138.193.242</p><p>状態: RUNNING</p></div><div class='col s1 m1 l1'></div></div><div class='row' style='margin-bottom: 0px;'><div class='col s1 m1 l1'></div><div class='col s10 m10 l10 center'><p><a class='white-text' href='/container-list/test-container'>コンテナの詳細</a></p></div><div class='col s1 m1 l1'></div></div></div></div></div>"
                );
            }
        })
});