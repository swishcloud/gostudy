function blockUi() {
    $("#mask").removeClass("d-none")
}

function unBlockUi() {
    $("#mask").addClass("d-none")
}
$(document).ajaxStart(blockUi).ajaxComplete(unBlockUi)