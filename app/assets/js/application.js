
require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");
require("jquery-ujs/src/rails.js");

const option= {
  enableTime: false,
  dateFormat: "Y-m-d"
}

const pickers = document.querySelectorAll(".date-picker")

pickers.forEach(element => {
  flatpickr(element, option)

})



$(".check_box").change(function(){
  if($(this).prop('checked') == true){
    $("form").submit();
    console.log("aqui es")
  }
});







