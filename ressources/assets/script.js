$(function(){
  $("#custom_search_group").hide();
  $("#custom_search").change(function() {
    if($(this).prop("checked")){
      $("#custom_search_group").show();
    } else {
      $("#custom_search_group").hide();
    }
  });

  $("#search-form").submit(function(e){
    if($("#search_lang").val().length == 0) {
      $("#search_alert").removeClass("hide");
      e.preventDefault();
    }
  });
});
