function blockUi() {
    $("#mask").removeClass("d-none")
}

function unBlockUi() {
    $("#mask").addClass("d-none")
}

$.ajaxStart = blockUi
$.ajaxComplete = unBlockUi