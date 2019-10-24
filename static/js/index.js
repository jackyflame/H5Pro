$(document).ready(function(){

    $("p").click(function(){
        $(this).hide();
    });

    $("#test").click(function(){
        test($(this));
    });

});

function test(obj) {
    alert(obj.html()+" AddedObject run");
    AddedObject.ShareH5("test!!!!");
    AddedObject.openWebClient("test!!!!");
}