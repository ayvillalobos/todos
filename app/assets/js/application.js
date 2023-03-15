
require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");
require("jquery-ujs/src/rails.js");





window.watchCheckboxes = function() {
  $(".item-form input[type=checkbox]").on("change", (e) => {
    let $e = $(e.target);
    let $form = $e.closest("form");
    $form.submit();
  });
};

$(() => {
  watchCheckboxes();
});