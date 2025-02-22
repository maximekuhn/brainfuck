document.body.addEventListener("htmx:responseError", function(event) {
    if (event.detail.xhr.status >= 500) {
        alert("Internal server error. Try again");
        return;
    }
    if (event.detail.xhr.status >= 400) {
        alert("Invalid input or invalid brainfuck code");
    }
});
