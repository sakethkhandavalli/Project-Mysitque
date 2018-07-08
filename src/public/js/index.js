var myFunction = function(id) {
      var popup = document.getElementById(id);
      popup.classList.toggle("show");
}

var truck_move = function()
{
    $('#truck').animate({left: '+=45%'}, 5000);
}

var build_table = function(table_info)
{
    console.log("I entered");
    table_info = table_info.split("}");
    var parent = $(".table");
    var tbody = $("tbody");
    for(var i=0 ; i<6; i++)
    {
      var object = table_info[i].substr(1);
      object += '}';
      var object = JSON.parse(object);
      console.log(object);
      var row = $("<tr></tr>");
      var key,j=0;
      for(key in object)
      {
        var cell = $("<td></td>");
        if(j==0)
        {
          var anch = $("<a></a>").attr("href", "/trackingPage");
          anch.text(object[key]);
          cell.append(anch);
        }
        else {
          if(object[key]=='')
          {
            object[key]='null';
          }
          if(object[key]==null)
          {
            object[key]='null';
          }
          cell.text(object[key]);
          if(object[key]=='ACTIVE')
          {
            cell.attr("style", "color:#00cc66");
          }
          if(object[key]=='ALERT')
          {
            cell.attr("style", "color:#ED4340");
          }
          if(object[key]=='INACTIVE')
          {
            cell.attr("style", "color:#E3A21A");
          }
        }
        row.append(cell);
        j+=1;
      }
      tbody.append(row);
    }
    parent.append(tbody);
}

var build_blocks = function(blocks_info)
{
    console.log("HELLo entered");
    blocks_info = blocks_info.split("}");
    var parent = $(".pokemon");
    var div = $("<div></div>").addClass('row');
    for (var i = 0; i < 8; i++)
    {
        var object = blocks_info[i].substr(1);
        object += '}';
        var object = JSON.parse(object);
        //console.log(object);
        var inner_div = $("<div></div>").addClass("col-sm-3");
        var space_div = $("<div></div>").addClass("col-sm-1");
        inner_div.addClass("popup");
        var span_id = "myPopup" + i;
        var span = $("<span></span>").attr("id", span_id);
        span.text(object["blockid"]);
        span.addClass("popuptext");
        inner_div.attr("id", "block" + i);
        inner_div.attr("onmouseenter", "myFunction('" + span_id + "')");
        inner_div.attr("onmouseleave", "myFunction('" + span_id + "')");

        var txt = "";
        var p = $("<h4></h4>");
        txt = object["packageid "];
        //console.log(txt);
        p.text(txt);

        inner_div.append(span);
        inner_div.append(p);
        //console.log(inner_div);
        div.append(inner_div);
        div.append(space_div);
    }
    parent.append(div);
    console.log(parent);
}

var fill_blocks = function()
{
        $.get("http://127.0.0.1:9040/getAllBlocks", function(data, status){
        build_blocks(data);
    });
    return;
}

var fill_table = function()
{
        $.get("http://127.0.0.1:9040/getAllPackages", function(data, status){
        build_table(data);
    });
    return;
}
