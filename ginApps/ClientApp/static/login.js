function validate() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("pass").value;
    
    if (username === "" || username === null) {
        intensify($("#userLabel"));
        alert("No puede dejar el campo usuario vacio.");
        return false;
    }
    
    if (password === "" || password === null) {
        intensify($("#passLabel"));
        alert("No puede dejar el campo clave vacio.");
        return false;
    }
    clicked()
}

function intensify(intense) {
    intense
    .addClass("animated shakeit")
    .delay(6000)
    .queue(function() {
           intense.removeClass("animated shakeit").dequeue();
           });
}

function clicked() {
    alert("Iniciando sesion");
}

var submit = document.getElementById("submit");
//submit.addEventListener("click", clicked);
submit.addEventListener("click", validate);
