var truck_move = function()
{
    $('#truck').animate({left: '+=80%'}, 9000);
}

var build_blocks = function(blocks_info)
{
    console.log("HELLo entered");
    blocks_info = blocks_info.split("}");
    var parent = $(".jumbotron");
    var div = $("<div></div>").addClass('row');
    for (var i = 0; i < 8; i++)
    {
        var object = blocks_info[i].substr(1);
        object += '}';
        var object = JSON.parse(object);
        var inner_div = $("<div></div>").addClass('col-sm-3');
        var extra_div = $("<div></div>").addClass('col-sm-1');
        var id_text = "myPopUp" + i;
        var id_text2 = "myPopup" + i;
        inner_div.addClass('popup');
        var span = $("<span></span>").attr('id',id_text);
        span.text(object["blockid"]);
        span.attr('hidden', 'true');
        span.addClass('popuptext');
        var txt = "";
        var p = $("<p></p>");

        inner_div.attr('onmouseenter','myFunction(\'' + id_text + '\',\'' + id_text2 + '\')'); inner_div.attr('onmouseleave', 'myFunction(\'' + id_text +  '\',\'' + id_text2 + '\')');
        txt = object["packageid "];
        p.text(txt);
        inner_div.append(span);
        inner_div.append(p);

        //console.log(inner_div);
        div.append(inner_div);
        div.append(extra_div);
    }
    parent.append(div);
    console.log(parent);
}

var fill_blocks = function()
{
    $.get("http://127.0.0.1:9040/getAllBlocks", function(data, status)
    {
        build_blocks(data);
        add_pitstops(data);
        add_checks(data);
    });
    return;
}

function myFunction(id,id2)
{
        var popup = document.getElementById(id);
        popup.classList.toggle("show");

        var popup2 = document.getElementById(id2);

        popup2.classList.toggle("show");
}

var add_pitstops = function(blocks_info)
{
    blocks_info = blocks_info.split("}");
    var parent = $(".pitstops");
    for (var i = 0; i < 8; i++)
    {
        var object = blocks_info[i].substr(1);
        object += '}';
        var object = JSON.parse(object);
        var div = $("<div></div>").addClass('pitstop');
        div.addClass('circle');
        var id_text = 'pitstop' + i;
        div.attr('id',id_text);
        div.css('position', 'absolute');
        div.css('margin-left', i*210);
        var id_text = "myPopup" + i;
        var id_text2 = "myPopUp" + i
        div.addClass('popup');
        var span = $("<span></span>").attr('id',id_text);
        span.text(object["lastupdated"]);
        span.attr('hidden', 'true');
        span.addClass('popuptext');

        div.attr('onmouseenter','myFunction(\'' + id_text + '\',\'' + id_text2 + '\')');
        div.attr('onmouseleave', 'myFunction(\'' + id_text +  '\',\'' + id_text2 + '\')');

        txt = object["packageid "];
        //console.log(txt);
        div.append(span);

        parent.append(div);
    }
}

var add_checks = function(blocks_info)
{
    blocks_info = blocks_info.split("}");
    var parent = $(".checks");
    for (var i = 0; i < 8; i++)
    {
        var object = blocks_info[i].substr(1);
        object += '}';
        var object = JSON.parse(object);
        var div = $("<div></div>").addClass('check_div');
        var id_text = 'check' + i;
        div.attr('id',id_text);
        div.css('position', 'absolute');
        if(object['packagestatus '] == 'ACTIVE')
        {
            div.css('margin-left', i*200);
            div.html('<svg width="150" height="150">' +
            '<path id="check' + i + 'active" d="M10,30 l30,50 l95,-70" />' +
            '</svg>');
        }
        else if(object['packagestatus '] == 'ALERT')
        {
            div.css('margin-left', i*210);
            div.html('<svg width="140" height="150">' +
            '<path id="check'+i+'alert" d="M10,-10 l0,90 l," />' +
            '</svg>'
            + '<svg width="140" height="150" position="relative" style = "margin-left: -50%">' +
            '<path id="check'+i+'alert" d="M10,90 l0,20 l," />' +
            '</svg>');
        }
        else
        {
            div.css('margin-left', i*200);
            div.html('<svg width="150" height="150">' +
            '<path id="check' + i + 'active" d="M10,30 l30,50 l95,-70" />' +
            '</svg>');
            /*div.html('<svg width="150" height="75   ">' +
            '<path id="check_inactive" d="M10,-15 l35,75 l55,-100" />' +
            '</svg>' + '<svg width="150" height="150" position="relative" style = "margin-left: -49.25%">' +
            '<path id="check_inactive" d="M10,150 l30,-85 l35, 100" />' +
            '</svg>');
*/        }
        parent.append(div);
        console.log("HIOOORAH");
    }
}


/*navbar*/
/*
$(window).scroll(function() {
    if($(this).scrollTop()>5) {
        $( ".navbar-me" ).addClass("fixed-me");
    } else {
        $( ".navbar-me" ).removeClass("fixed-me");
    }
});
*/
