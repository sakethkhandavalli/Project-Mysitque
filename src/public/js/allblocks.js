var sobj={
  "None": 1
};

var geobj={
  "None": 1
};

var build_table = function(table_info,a,b)
{
    table_info = table_info.split("}");
    var parent = $("#table");
    var tbody = $("#tbody");
    tbody.html('');
    var and1=0,and2=0;
    if(a!='None')
    {
      and1=1;
    }
    if(b!='None')
    {
      and2=1;
    }
    for(var i=0 ; i<table_info.length; i++)
    {
      var object = table_info[i].substr(1);
      if(object!='\n')
      {
        object += '}';
        object = JSON.parse(object);
        var row = $("<tr></tr>");
        var key;
        flg1=0;flg2=0;
        for(key in object)
        {
          var cell = $("<td></td>");
          if(object[key]=='')
          {
            object[key]='null';
          }
          if(key.trim()=='packagestatus')
          {
            if(object[key]==a)
            {
              flg1=1;
            }
          }
          if(key.trim()=='geohash')
          {
            if(object[key]==b)
            {
              flg2=1;
            }
          }
          if(and1==1 && and2==1)
          {
            if(flg1==1 && flg2==1)
            {
              row=func(object);
              break;
            }
          }
          else if(and1==1)
          {
            if(flg1==1)
            {
              row=func(object);
              break;
            }
          }
          else if(and2==1){
            if(flg2==1)
            {
              row=func(object);
              break;
            }
          }
          else{
          cell.text(object[key]);
          row.append(cell);
        }
      }
    }
    tbody.append(row);
  }

  parent.append(tbody);
  return;
}

function func(object)
{
  var key;
  var row = $("<tr></tr>");
  for(key in object)
  {
    var cell = $("<td></td>");
    if(object[key]=='')
    {
      object[key]='null';
    }
    cell.text(object[key]);
    row.append(cell);
  }
  return row;
}

var filter = function()
{
    var a = document.getElementById('select_status').value;
    var b = document.getElementById('select_geo').value;
    $.get("http://127.0.0.1:9040/getAllBlocks", function(data, status){
        build_table(data,a,b);
    });
    return;
}

var fill_options = function()
{
  console.log('hi');
  var iter;
  for(iter in sobj)
  {
    var opt=$("<option></option>");
    opt.text(iter);
    opt.attr("value", iter);
    $("select#select_status").append(opt);
  }

  for(iter in geobj)
  {
    var opt=$("<option></option>");
    opt.text(iter);
    opt.attr("value", iter);
    $("select#select_geo").append(opt);
  }

  return;
}

var fill_search = function()
{
  console.log('hi');
  var a = document.getElementById('search').value;
  console.log(a);
  $.get("http://127.0.0.1:9040/getAllBlocks", function(data, status){
      table_info = data.split("}");
      var parent = $("#table");
      var tbody = $("#tbody");
      tbody.html('');
      var flag=0;
      for(var i=0 ; i<table_info.length; i++)
      {
        var object = table_info[i].substr(1);
        if(object!='\n')
        {
          object += '}';
          object = JSON.parse(object);
          var key;
          console.log(object['packageid ']);
          //console.log(a);
          if(object['packageid ']==a)
          {
            var row = $("<tr></tr>");
            for(key in object)
            {
              var cell = $("<td></td>");
              if(object[key]=='')
              {
                object[key]='null';
              }
              cell.text(object[key]);
              row.append(cell);
            }
            tbody.append(row);
            parent.append(tbody);
            flag=1;
            break;
          }
        }
      }
      if(flag==0)
      {
        alert("Package does not exist!");
        document.getElementById('search').value = '';
        location.reload();
      }
  });
  return;
}

var fill_allblocks = function()
{
        $.get("http://127.0.0.1:9040/getAllBlocks", function(data, status){
            table_info = data.split("}");
            for(var i=0 ; i<table_info.length; i++)
            {
                var object = table_info[i].substr(1);
                if(object!='\n')
                {
                  object += '}';
                  object = JSON.parse(object);
                  var key;
                  for(key in object)
                  {
                      if(key.trim()=='packagestatus')
                      {
                        if(sobj.hasOwnProperty(object[key]) == false)
                        {
                          sobj[object[key]]=1;
                        }
                      }
                      if(key.trim()=='geohash')
                      {
                        if(geobj.hasOwnProperty(object[key]) == false)
                        {
                          geobj[object[key]]=1;
                        }
                      }
                  }
                }
            }
            //console.log(sobj);
            //console.log(geobj);
            fill_options();
            build_table(data,'None','None');
          }
        );

    return;
}
